// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIValidationMessage Api validation message
//
// swagger:model ApiValidationMessage
type APIValidationMessage struct {

	// field
	Field string `json:"field,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// root Id
	RootID int32 `json:"rootId,omitempty"`
}

// Validate validates this Api validation message
func (m *APIValidationMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIValidationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIValidationMessage) UnmarshalBinary(b []byte) error {
	var res APIValidationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
