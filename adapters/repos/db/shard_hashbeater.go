//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/weaviate/weaviate/entities/storobj"
	"github.com/weaviate/weaviate/usecases/objects"
	"github.com/weaviate/weaviate/usecases/replica"
	"github.com/weaviate/weaviate/usecases/replica/hashtree"
)

func (s *Shard) initHashBeater() {
	go func() {
		t := time.NewTicker(50 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-s.hashBeaterCtx.Done():
				return
			case <-t.C:
				stats, err := s.hashBeat()
				if s.hashBeaterCtx.Err() != nil {
					return
				}
				if err != nil {
					s.index.logger.Printf("hashbeat at shard %s: %v", s.name, err)
					// TODO (jeroiraz): an exp-backoff delay may be convenient at this point
				}

				hosts := make([]string, len(stats.hostStats))
				localObjects := 0
				remoteObjects := 0
				objectsPropagated := 0
				var objectProgationTook time.Duration

				for i, stat := range stats.hostStats {
					hosts[i] = stat.host
					localObjects += stat.localObjects
					remoteObjects += stat.remoteObjects
					objectsPropagated += stat.objectsPropagated
					objectProgationTook += stat.objectProgationTook
				}

				if objectsPropagated == 0 {
					// no need to propagate any local object
					continue
				}

				s.index.logger.WithField("action", "async_replication_iteration").
					WithField("diffCalculationTook", stats.diffCalculationTook.String()).
					WithField("hosts", hosts).
					WithField("localObjects", localObjects).
					WithField("remoteObjects", remoteObjects).
					WithField("objectsPropagated", objectsPropagated).
					WithField("objectProgationTook", objectProgationTook.String()).
					Debug("async replication iteration successfully completed")
			}
		}
	}()
}

type hashBeatStats struct {
	diffCalculationTook time.Duration
	hostStats           []hashBeatHostStats
}

type hashBeatHostStats struct {
	host                string
	localObjects        int
	remoteObjects       int
	objectsPropagated   int
	objectProgationTook time.Duration
}

func (s *Shard) hashBeat() (stats hashBeatStats, err error) {
	// Note: a copy of the hashtree could be used if a more stable comparison is desired s.hashtree.Clone()
	// in such a case nodes may also need to have a stable copy of their corresponding hashtree.
	ht := s.hashtree

	diffCalculationStart := time.Now()

	// Note: any consistency level could be used
	replyCh, err := s.index.replicator.CollectShardDifferences(s.hashBeaterCtx, s.name, ht, replica.One, "")
	if err != nil {
		if errors.Is(err, hashtree.ErrNoMoreDifferences) {
			s.index.logger.Printf("shard fully replicated %s", s.name)

			return hashBeatStats{
				diffCalculationTook: time.Since(diffCalculationStart),
			}, nil
		}

		return stats, fmt.Errorf("collecting differences for shard %s: %w", s.name, err)
	}

	stats.diffCalculationTook = time.Since(diffCalculationStart)

	for r := range replyCh {
		objectProgationStart := time.Now()

		if r.Err != nil {
			if !errors.Is(r.Err, hashtree.ErrNoMoreDifferences) {
				s.index.logger.Printf("reading collected differences for shard %s: %v", s.name, r.Err)
			}
			continue
		}

		shardDiffReader := r.Value
		diffReader := shardDiffReader.DiffReader

		localObjects := 0
		remoteObjects := 0
		objectsPropagated := 0

		for {
			if s.hashBeaterCtx.Err() != nil {
				return stats, s.hashBeaterCtx.Err()
			}

			initialToken, finalToken, err := diffReader.Next()
			if errors.Is(err, hashtree.ErrNoMoreDifferences) {
				break
			}
			if err != nil {
				return stats, fmt.Errorf("difference reading: %w", err)
			}

			localObjs, remoteObjs, propagations, err := s.stepsTowardsShardConsistency(
				s.hashBeaterCtx,
				s.name,
				shardDiffReader.Host,
				initialToken,
				finalToken,
			)
			if err != nil {
				s.index.logger.Printf("solving differences for shard %s: %v", s.name, err)
				continue
			}

			localObjects += localObjs
			remoteObjects += remoteObjs
			objectsPropagated += propagations
		}

		stat := hashBeatHostStats{
			host:                shardDiffReader.Host,
			localObjects:        localObjects,
			remoteObjects:       remoteObjects,
			objectsPropagated:   objectsPropagated,
			objectProgationTook: time.Since(objectProgationStart),
		}

		stats.hostStats = append(stats.hostStats, stat)
	}

	return stats, nil
}

func (s *Shard) stepsTowardsShardConsistency(ctx context.Context,
	shardName string, host string, initialToken, finalToken uint64,
) (localObjects, remoteObjects, propagations int, err error) {
	const limit = 100

	for localLastReadToken := initialToken; localLastReadToken < finalToken; {
		localDigests, newLocalLastReadToken, err := s.index.digestObjectsInTokenRange(ctx, shardName, localLastReadToken, finalToken, limit)
		if err != nil && !errors.Is(err, storobj.ErrLimitReached) {
			return localObjects, remoteObjects, propagations, err
		}

		localDigestsByUUID := make(map[string]replica.RepairResponse, len(localDigests))

		for _, d := range localDigests {
			// deleted objects are not propagated
			if !d.Deleted {
				localDigestsByUUID[d.ID] = d
			}
		}

		localObjects += len(localDigestsByUUID)

		if len(localDigestsByUUID) == 0 {
			// no more local objects need to be propagated in this iteration
			localLastReadToken = newLocalLastReadToken
			continue
		}

		remoteLastTokenRead := localLastReadToken

		remoteStaleUpdateTime := make(map[string]int64, len(localDigestsByUUID))

		// fetch digests from remote host in order to avoid sending unnecessary objects
		for remoteLastTokenRead < newLocalLastReadToken {
			remoteDigests, newRemoteLastTokenRead, err := s.index.replicator.DigestObjectsInTokenRange(ctx,
				shardName, host, remoteLastTokenRead, newLocalLastReadToken, limit)
			if err != nil && !strings.Contains(err.Error(), storobj.ErrLimitReached.Error()) {
				return localObjects, remoteObjects, propagations, err
			}

			if len(remoteDigests) == 0 {
				// no more objects in remote host
				break
			}

			for _, d := range remoteDigests {
				if len(localDigestsByUUID) == 0 {
					// no more local objects need to be propagated in this iteration
					break
				}

				if d.Deleted {
					// object was deleted in remote host
					delete(localDigestsByUUID, d.ID)
					continue
				}

				remoteObjects++

				localDigest, ok := localDigestsByUUID[d.ID]
				if ok {
					if localDigest.UpdateTime <= d.UpdateTime {
						// older or up to date objects are not propagated
						delete(localDigestsByUUID, d.ID)
					} else {
						// older object is subject to be overwriten
						remoteStaleUpdateTime[d.ID] = d.UpdateTime
					}
				}
			}

			remoteLastTokenRead = newRemoteLastTokenRead

			if len(localDigestsByUUID) == 0 {
				// no more local objects need to be propagated in this iteration
				break
			}
		}

		if len(localDigestsByUUID) == 0 {
			// no more local objects need to be propagated in this iteration
			localLastReadToken = newLocalLastReadToken
			continue
		}

		uuids := make([]strfmt.UUID, 0, len(localDigestsByUUID))
		for uuid := range localDigestsByUUID {
			uuids = append(uuids, strfmt.UUID(uuid))
		}

		replicaObjs, err := s.index.fetchObjects(ctx, shardName, uuids)
		if err != nil {
			return localObjects, remoteObjects, propagations, fmt.Errorf("fetching objects from shard %s: %w", shardName, err)
		}

		mergeObjs := make([]*objects.VObject, len(replicaObjs))

		for i, replicaObj := range replicaObjs {
			obj := &objects.VObject{
				LatestObject:    &replicaObj.Object.Object,
				Vector:          replicaObj.Object.Vector,
				StaleUpdateTime: remoteStaleUpdateTime[replicaObj.ID.String()],
			}
			mergeObjs[i] = obj
		}

		_, err = s.index.replicator.Overwrite(ctx, host, s.class.Class, shardName, mergeObjs)
		if err != nil {
			return localObjects, remoteObjects, propagations, fmt.Errorf("overwriting objects in shard %s at host %s: %w", shardName, host, err)
		}

		propagations += len(mergeObjs)
		localLastReadToken = newLocalLastReadToken
	}

	// Note: propagations == 0 means local shard is laying behind remote shard,
	// the local shard may receive recent objects when remote shard propagates them

	return localObjects, remoteObjects, propagations, nil
}

func (s *Shard) stopHashBeater() {
	s.hashBeaterCancelFunc()
}
