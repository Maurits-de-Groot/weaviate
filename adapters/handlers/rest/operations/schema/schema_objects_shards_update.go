// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsShardsUpdateHandlerFunc turns a function with the right signature into a schema objects shards update handler
type SchemaObjectsShardsUpdateHandlerFunc func(SchemaObjectsShardsUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SchemaObjectsShardsUpdateHandlerFunc) Handle(params SchemaObjectsShardsUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SchemaObjectsShardsUpdateHandler interface for that can handle valid schema objects shards update params
type SchemaObjectsShardsUpdateHandler interface {
	Handle(SchemaObjectsShardsUpdateParams, *models.Principal) middleware.Responder
}

// NewSchemaObjectsShardsUpdate creates a new http.Handler for the schema objects shards update operation
func NewSchemaObjectsShardsUpdate(ctx *middleware.Context, handler SchemaObjectsShardsUpdateHandler) *SchemaObjectsShardsUpdate {
	return &SchemaObjectsShardsUpdate{Context: ctx, Handler: handler}
}

/*
	SchemaObjectsShardsUpdate swagger:route PUT /schema/{className}/shards/{shardName} schema schemaObjectsShardsUpdate

Update shard status of an Object Class
*/
type SchemaObjectsShardsUpdate struct {
	Context *middleware.Context
	Handler SchemaObjectsShardsUpdateHandler
}

func (o *SchemaObjectsShardsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSchemaObjectsShardsUpdateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
