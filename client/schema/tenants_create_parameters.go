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

// NewTenantsCreateParams creates a new TenantsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenantsCreateParams() *TenantsCreateParams {
	return &TenantsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenantsCreateParamsWithTimeout creates a new TenantsCreateParams object
// with the ability to set a timeout on a request.
func NewTenantsCreateParamsWithTimeout(timeout time.Duration) *TenantsCreateParams {
	return &TenantsCreateParams{
		timeout: timeout,
	}
}

// NewTenantsCreateParamsWithContext creates a new TenantsCreateParams object
// with the ability to set a context for a request.
func NewTenantsCreateParamsWithContext(ctx context.Context) *TenantsCreateParams {
	return &TenantsCreateParams{
		Context: ctx,
	}
}

// NewTenantsCreateParamsWithHTTPClient creates a new TenantsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenantsCreateParamsWithHTTPClient(client *http.Client) *TenantsCreateParams {
	return &TenantsCreateParams{
		HTTPClient: client,
	}
}

/*
TenantsCreateParams contains all the parameters to send to the API endpoint

	for the tenants create operation.

	Typically these are written to a http.Request.
*/
type TenantsCreateParams struct {

	// Body.
	Body []*models.Tenant

	// ClassName.
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenants create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantsCreateParams) WithDefaults() *TenantsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenants create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenants create params
func (o *TenantsCreateParams) WithTimeout(timeout time.Duration) *TenantsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenants create params
func (o *TenantsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenants create params
func (o *TenantsCreateParams) WithContext(ctx context.Context) *TenantsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenants create params
func (o *TenantsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenants create params
func (o *TenantsCreateParams) WithHTTPClient(client *http.Client) *TenantsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenants create params
func (o *TenantsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tenants create params
func (o *TenantsCreateParams) WithBody(body []*models.Tenant) *TenantsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tenants create params
func (o *TenantsCreateParams) SetBody(body []*models.Tenant) {
	o.Body = body
}

// WithClassName adds the className to the tenants create params
func (o *TenantsCreateParams) WithClassName(className string) *TenantsCreateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the tenants create params
func (o *TenantsCreateParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *TenantsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
