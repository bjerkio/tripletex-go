// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperListInteger response wrapper list integer
//
// swagger:model ResponseWrapperListInteger
type ResponseWrapperListInteger struct {

	// value
	Value []int32 `json:"value"`
}

// Validate validates this response wrapper list integer
func (m *ResponseWrapperListInteger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperListInteger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperListInteger) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperListInteger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
