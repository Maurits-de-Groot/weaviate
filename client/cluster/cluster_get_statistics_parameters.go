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

// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewClusterGetStatisticsParams creates a new ClusterGetStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterGetStatisticsParams() *ClusterGetStatisticsParams {
	return &ClusterGetStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterGetStatisticsParamsWithTimeout creates a new ClusterGetStatisticsParams object
// with the ability to set a timeout on a request.
func NewClusterGetStatisticsParamsWithTimeout(timeout time.Duration) *ClusterGetStatisticsParams {
	return &ClusterGetStatisticsParams{
		timeout: timeout,
	}
}

// NewClusterGetStatisticsParamsWithContext creates a new ClusterGetStatisticsParams object
// with the ability to set a context for a request.
func NewClusterGetStatisticsParamsWithContext(ctx context.Context) *ClusterGetStatisticsParams {
	return &ClusterGetStatisticsParams{
		Context: ctx,
	}
}

// NewClusterGetStatisticsParamsWithHTTPClient creates a new ClusterGetStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterGetStatisticsParamsWithHTTPClient(client *http.Client) *ClusterGetStatisticsParams {
	return &ClusterGetStatisticsParams{
		HTTPClient: client,
	}
}

/*
ClusterGetStatisticsParams contains all the parameters to send to the API endpoint

	for the cluster get statistics operation.

	Typically these are written to a http.Request.
*/
type ClusterGetStatisticsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster get statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGetStatisticsParams) WithDefaults() *ClusterGetStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster get statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGetStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster get statistics params
func (o *ClusterGetStatisticsParams) WithTimeout(timeout time.Duration) *ClusterGetStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster get statistics params
func (o *ClusterGetStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster get statistics params
func (o *ClusterGetStatisticsParams) WithContext(ctx context.Context) *ClusterGetStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster get statistics params
func (o *ClusterGetStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster get statistics params
func (o *ClusterGetStatisticsParams) WithHTTPClient(client *http.Client) *ClusterGetStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster get statistics params
func (o *ClusterGetStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterGetStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
