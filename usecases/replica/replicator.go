//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package replica

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/storobj"
	"github.com/semi-technologies/weaviate/usecases/objects"
	"github.com/semi-technologies/weaviate/usecases/sharding"
)

type opID int

const (
	opPutObject opID = iota + 1
	opMergeObject
	opDeleteObject

	opPutObjects = iota + 97
	opAddReferences
	opDeleteObjects
)

type shardingState interface {
	NodeName() string
	ShardingState(class string) *sharding.State
}

type nodeResolver interface {
	NodeHostname(nodeName string) (string, bool)
}

type Replicator struct {
	class          string
	stateGetter    shardingState
	client         Client
	resolver       nodeResolver
	requestCounter atomic.Uint64
	*Finder
}

func NewReplicator(className string,
	stateGetter shardingState, nodeResolver nodeResolver,
	client Client, rClient RClient,
) *Replicator {
	return &Replicator{
		class:       className,
		stateGetter: stateGetter,
		client:      client,
		resolver:    nodeResolver,
		Finder:      NewFinder(className, stateGetter, nodeResolver, rClient),
	}
}

func (r *Replicator) PutObject(ctx context.Context, shard string,
	obj *storobj.Object, cl ConsistencyLevel,
) error {
	coord := newCoordinator[SimpleResponse](r, shard, r.requestID(opPutObject))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.PutObject(ctx, host, r.class, shard, requestID, obj)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	replyCh, level, err := coord.Replicate2(ctx, cl, op, r.simpleCommit(shard))
	if err != nil {
		return err
	}
	return errorsFromSimpleResponses2(1, level, replyCh)[0]
}

func (r *Replicator) MergeObject(ctx context.Context, shard string,
	mergeDoc *objects.MergeDocument, cl ConsistencyLevel,
) error {
	coord := newCoordinator[SimpleResponse](r, shard, r.requestID(opMergeObject))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.MergeObject(ctx, host, r.class, shard, requestID, mergeDoc)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	replyCh, level, err := coord.Replicate2(ctx, cl, op, r.simpleCommit(shard))
	if err != nil {
		return err
	}
	return errorsFromSimpleResponses2(1, level, replyCh)[0]
}

func (r *Replicator) PutObjects(ctx context.Context, shard string,
	objs []*storobj.Object,
) []error {
	coord := newCoordinator[SimpleResponse](r, shard, r.requestID(opPutObjects))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.PutObjects(ctx, host, r.class, shard, requestID, objs)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	err := coord.Replicate(ctx, op, r.simpleCommit(shard))
	return errorsFromSimpleResponses(len(objs), coord.responses, err)
}

func (r *Replicator) DeleteObject(ctx context.Context, shard string,
	id strfmt.UUID, cl ConsistencyLevel,
) error {
	coord := newCoordinator[SimpleResponse](r, shard, r.requestID(opDeleteObject))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.DeleteObject(ctx, host, r.class, shard, requestID, id)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	replyCh, level, err := coord.Replicate2(ctx, cl, op, r.simpleCommit(shard))
	if err != nil {
		return err
	}
	return errorsFromSimpleResponses2(1, level, replyCh)[0]
}

func (r *Replicator) simpleCommit(shard string) commitOp[SimpleResponse] {
	return func(ctx context.Context, host, requestID string) (SimpleResponse, error) {
		resp := SimpleResponse{}
		err := r.client.Commit(ctx, host, r.class, shard, requestID, &resp)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			err = fmt.Errorf("%s: %w", host, err)
		}
		return resp, err
	}
}

func (r *Replicator) DeleteObjects(ctx context.Context, shard string,
	docIDs []uint64, dryRun bool,
) []objects.BatchSimpleObject {
	coord := newCoordinator[DeleteBatchResponse](r, shard, r.requestID(opDeleteObjects))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.DeleteObjects(
			ctx, host, r.class, shard, requestID, docIDs, dryRun)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	commit := func(ctx context.Context, host, requestID string) (DeleteBatchResponse, error) {
		resp := DeleteBatchResponse{}
		err := r.client.Commit(ctx, host, r.class, shard, requestID, &resp)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			err = fmt.Errorf("%q: %w", host, err)
		}
		return resp, err
	}

	err := coord.Replicate(ctx, op, commit)
	return resultsFromDeletionResponses(len(docIDs), coord.responses, err)
}

func (r *Replicator) AddReferences(ctx context.Context, shard string,
	refs []objects.BatchReference,
) []error {
	coord := newCoordinator[SimpleResponse](r, shard, r.requestID(opAddReferences))
	op := func(ctx context.Context, host, requestID string) error {
		resp, err := r.client.AddReferences(ctx, host, r.class, shard, requestID, refs)
		if err == nil {
			err = resp.FirstError()
		}
		if err != nil {
			return fmt.Errorf("%q: %w", host, err)
		}
		return nil
	}
	err := coord.Replicate(ctx, op, r.simpleCommit(shard))
	return errorsFromSimpleResponses(len(refs), coord.responses, err)
}

func errorsFromSimpleResponses(batchSize int, rs []SimpleResponse, defaultErr error) []error {
	errs := make([]error, batchSize)
	n := 0
	for _, resp := range rs {
		if len(resp.Errors) != batchSize {
			continue
		}
		n++
		for i, err := range resp.Errors {
			if !err.Empty() && errs[i] == nil {
				errs[i] = err.Clone()
			}
		}
	}
	if n == 0 || n != len(rs) {
		for i := range errs {
			if errs[i] == nil {
				errs[i] = defaultErr
			}
		}
	}
	return errs
}

func resultsFromDeletionResponses(batchSize int, rs []DeleteBatchResponse, defaultErr error) []objects.BatchSimpleObject {
	ret := make([]objects.BatchSimpleObject, batchSize)
	n := 0
	for _, resp := range rs {
		if len(resp.Batch) != batchSize {
			continue
		}
		n++
		for i, x := range resp.Batch {
			if !x.Error.Empty() && ret[i].Err == nil {
				ret[i].Err = x.Error.Clone()
			}
			if ret[i].UUID == "" && x.UUID != "" {
				ret[i].UUID = strfmt.UUID(x.UUID)
			}
		}
	}
	if n == 0 || n != len(rs) {
		for i := range ret {
			if ret[i].Err == nil {
				ret[i].Err = defaultErr
			}
		}
	}
	return ret
}

func (r *Replicator) requestID(op opID) string {
	return fmt.Sprintf("%s-%.2x-%x-%x",
		r.stateGetter.NodeName(),
		op,
		time.Now().UnixMilli(),
		r.requestCounter.Add(1))
}

type simpleResult[T any] struct {
	Response T
	Err      error
}

func errorsFromSimpleResponses2(batchSize int, level int, ch <-chan simpleResult[SimpleResponse]) []error {
	urs := make([]SimpleResponse, 0, level)
	var firstError error
	for x := range ch {
		if x.Err != nil {
			urs = append(urs, x.Response)
			if _, ok := x.Err.(*Error); !ok && firstError == nil {
				firstError = x.Err
			}
		} else {
			level--
			if level == 0 {
				return make([]error, batchSize)
			}
		}
	}
	return errorsFromSimpleResponses(batchSize, urs, firstError)
}

// type doAfter[T any] func (batchSize int, level int, ch <-chan simpleResult[T])
