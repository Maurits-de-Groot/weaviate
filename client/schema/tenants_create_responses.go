// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// TenantsCreateReader is a Reader for the TenantsCreate structure.
type TenantsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTenantsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTenantsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTenantsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTenantsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenantsCreateOK creates a TenantsCreateOK with default headers values
func NewTenantsCreateOK() *TenantsCreateOK {
	return &TenantsCreateOK{}
}

/*
TenantsCreateOK describes a response with status code 200, with default header values.

Added new tenants to the specified class
*/
type TenantsCreateOK struct {
	Payload []*models.Tenant
}

// IsSuccess returns true when this tenants create o k response has a 2xx status code
func (o *TenantsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenants create o k response has a 3xx status code
func (o *TenantsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants create o k response has a 4xx status code
func (o *TenantsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenants create o k response has a 5xx status code
func (o *TenantsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants create o k response a status code equal to that given
func (o *TenantsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenants create o k response
func (o *TenantsCreateOK) Code() int {
	return 200
}

func (o *TenantsCreateOK) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateOK  %+v", 200, o.Payload)
}

func (o *TenantsCreateOK) String() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateOK  %+v", 200, o.Payload)
}

func (o *TenantsCreateOK) GetPayload() []*models.Tenant {
	return o.Payload
}

func (o *TenantsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsCreateUnauthorized creates a TenantsCreateUnauthorized with default headers values
func NewTenantsCreateUnauthorized() *TenantsCreateUnauthorized {
	return &TenantsCreateUnauthorized{}
}

/*
TenantsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type TenantsCreateUnauthorized struct {
}

// IsSuccess returns true when this tenants create unauthorized response has a 2xx status code
func (o *TenantsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants create unauthorized response has a 3xx status code
func (o *TenantsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants create unauthorized response has a 4xx status code
func (o *TenantsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants create unauthorized response has a 5xx status code
func (o *TenantsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants create unauthorized response a status code equal to that given
func (o *TenantsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the tenants create unauthorized response
func (o *TenantsCreateUnauthorized) Code() int {
	return 401
}

func (o *TenantsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateUnauthorized ", 401)
}

func (o *TenantsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateUnauthorized ", 401)
}

func (o *TenantsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenantsCreateForbidden creates a TenantsCreateForbidden with default headers values
func NewTenantsCreateForbidden() *TenantsCreateForbidden {
	return &TenantsCreateForbidden{}
}

/*
TenantsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TenantsCreateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants create forbidden response has a 2xx status code
func (o *TenantsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants create forbidden response has a 3xx status code
func (o *TenantsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants create forbidden response has a 4xx status code
func (o *TenantsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants create forbidden response has a 5xx status code
func (o *TenantsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants create forbidden response a status code equal to that given
func (o *TenantsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the tenants create forbidden response
func (o *TenantsCreateForbidden) Code() int {
	return 403
}

func (o *TenantsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateForbidden  %+v", 403, o.Payload)
}

func (o *TenantsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateForbidden  %+v", 403, o.Payload)
}

func (o *TenantsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsCreateUnprocessableEntity creates a TenantsCreateUnprocessableEntity with default headers values
func NewTenantsCreateUnprocessableEntity() *TenantsCreateUnprocessableEntity {
	return &TenantsCreateUnprocessableEntity{}
}

/*
TenantsCreateUnprocessableEntity describes a response with status code 422, with default header values.

Invalid Tenant class
*/
type TenantsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants create unprocessable entity response has a 2xx status code
func (o *TenantsCreateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants create unprocessable entity response has a 3xx status code
func (o *TenantsCreateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants create unprocessable entity response has a 4xx status code
func (o *TenantsCreateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants create unprocessable entity response has a 5xx status code
func (o *TenantsCreateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants create unprocessable entity response a status code equal to that given
func (o *TenantsCreateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the tenants create unprocessable entity response
func (o *TenantsCreateUnprocessableEntity) Code() int {
	return 422
}

func (o *TenantsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TenantsCreateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TenantsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsCreateInternalServerError creates a TenantsCreateInternalServerError with default headers values
func NewTenantsCreateInternalServerError() *TenantsCreateInternalServerError {
	return &TenantsCreateInternalServerError{}
}

/*
TenantsCreateInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type TenantsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants create internal server error response has a 2xx status code
func (o *TenantsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants create internal server error response has a 3xx status code
func (o *TenantsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants create internal server error response has a 4xx status code
func (o *TenantsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenants create internal server error response has a 5xx status code
func (o *TenantsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tenants create internal server error response a status code equal to that given
func (o *TenantsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the tenants create internal server error response
func (o *TenantsCreateInternalServerError) Code() int {
	return 500
}

func (o *TenantsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /schema/{className}/tenants][%d] tenantsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
