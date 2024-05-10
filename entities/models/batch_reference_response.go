// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BatchReferenceResponse batch reference response
//
// swagger:model BatchReferenceResponse
type BatchReferenceResponse struct {
	BatchReference

	// result
	Result *BatchReferenceResponseAO1Result `json:"result,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BatchReferenceResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BatchReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BatchReference = aO0

	// AO1
	var dataAO1 struct {
		Result *BatchReferenceResponseAO1Result `json:"result,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BatchReferenceResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BatchReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Result *BatchReferenceResponseAO1Result `json:"result,omitempty"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this batch reference response
func (m *BatchReferenceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchReference
	if err := m.BatchReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchReferenceResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this batch reference response based on the context it is used
func (m *BatchReferenceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchReference
	if err := m.BatchReference.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchReferenceResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {
		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchReferenceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchReferenceResponse) UnmarshalBinary(b []byte) error {
	var res BatchReferenceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BatchReferenceResponseAO1Result Results for this specific reference.
//
// swagger:model BatchReferenceResponseAO1Result
type BatchReferenceResponseAO1Result struct {

	// errors
	Errors *ErrorResponse `json:"errors,omitempty"`

	// status
	// Enum: [SUCCESS PENDING FAILED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this batch reference response a o1 result
func (m *BatchReferenceResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchReferenceResponseAO1Result) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

var batchReferenceResponseAO1ResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","PENDING","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		batchReferenceResponseAO1ResultTypeStatusPropEnum = append(batchReferenceResponseAO1ResultTypeStatusPropEnum, v)
	}
}

const (

	// BatchReferenceResponseAO1ResultStatusSUCCESS captures enum value "SUCCESS"
	BatchReferenceResponseAO1ResultStatusSUCCESS string = "SUCCESS"

	// BatchReferenceResponseAO1ResultStatusPENDING captures enum value "PENDING"
	BatchReferenceResponseAO1ResultStatusPENDING string = "PENDING"

	// BatchReferenceResponseAO1ResultStatusFAILED captures enum value "FAILED"
	BatchReferenceResponseAO1ResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *BatchReferenceResponseAO1Result) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, batchReferenceResponseAO1ResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BatchReferenceResponseAO1Result) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("result"+"."+"status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this batch reference response a o1 result based on the context it is used
func (m *BatchReferenceResponseAO1Result) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchReferenceResponseAO1Result) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if m.Errors != nil {
		if err := m.Errors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchReferenceResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchReferenceResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res BatchReferenceResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
