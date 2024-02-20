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

package schema

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

	"github.com/weaviate/weaviate/entities/models"
)

// NewSchemaObjectsShardsUpdateParams creates a new SchemaObjectsShardsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchemaObjectsShardsUpdateParams() *SchemaObjectsShardsUpdateParams {
	return &SchemaObjectsShardsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsShardsUpdateParamsWithTimeout creates a new SchemaObjectsShardsUpdateParams object
// with the ability to set a timeout on a request.
func NewSchemaObjectsShardsUpdateParamsWithTimeout(timeout time.Duration) *SchemaObjectsShardsUpdateParams {
	return &SchemaObjectsShardsUpdateParams{
		timeout: timeout,
	}
}

// NewSchemaObjectsShardsUpdateParamsWithContext creates a new SchemaObjectsShardsUpdateParams object
// with the ability to set a context for a request.
func NewSchemaObjectsShardsUpdateParamsWithContext(ctx context.Context) *SchemaObjectsShardsUpdateParams {
	return &SchemaObjectsShardsUpdateParams{
		Context: ctx,
	}
}

// NewSchemaObjectsShardsUpdateParamsWithHTTPClient creates a new SchemaObjectsShardsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchemaObjectsShardsUpdateParamsWithHTTPClient(client *http.Client) *SchemaObjectsShardsUpdateParams {
	return &SchemaObjectsShardsUpdateParams{
		HTTPClient: client,
	}
}

/*
SchemaObjectsShardsUpdateParams contains all the parameters to send to the API endpoint

	for the schema objects shards update operation.

	Typically these are written to a http.Request.
*/
type SchemaObjectsShardsUpdateParams struct {

	// Body.
	Body *models.ShardStatus

	// ClassName.
	ClassName string

	// ShardName.
	ShardName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schema objects shards update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaObjectsShardsUpdateParams) WithDefaults() *SchemaObjectsShardsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schema objects shards update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaObjectsShardsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithTimeout(timeout time.Duration) *SchemaObjectsShardsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithContext(ctx context.Context) *SchemaObjectsShardsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithHTTPClient(client *http.Client) *SchemaObjectsShardsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithBody(body *models.ShardStatus) *SchemaObjectsShardsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetBody(body *models.ShardStatus) {
	o.Body = body
}

// WithClassName adds the className to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithClassName(className string) *SchemaObjectsShardsUpdateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetClassName(className string) {
	o.ClassName = className
}

// WithShardName adds the shardName to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) WithShardName(shardName string) *SchemaObjectsShardsUpdateParams {
	o.SetShardName(shardName)
	return o
}

// SetShardName adds the shardName to the schema objects shards update params
func (o *SchemaObjectsShardsUpdateParams) SetShardName(shardName string) {
	o.ShardName = shardName
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsShardsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param shardName
	if err := r.SetPathParam("shardName", o.ShardName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
