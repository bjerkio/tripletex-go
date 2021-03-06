// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoPayMessageDTO auto pay message d t o
//
// swagger:model AutoPayMessageDTO
type AutoPayMessageDTO struct {

	// message
	// Required: true
	// Read Only: true
	Message string `json:"message"`

	// message Id
	// Required: true
	// Read Only: true
	MessageID string `json:"messageId"`

	// message type
	// Required: true
	// Read Only: true
	MessageType string `json:"messageType"`
}

// Validate validates this auto pay message d t o
func (m *AutoPayMessageDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoPayMessageDTO) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *AutoPayMessageDTO) validateMessageID(formats strfmt.Registry) error {

	if err := validate.RequiredString("messageId", "body", string(m.MessageID)); err != nil {
		return err
	}

	return nil
}

func (m *AutoPayMessageDTO) validateMessageType(formats strfmt.Registry) error {

	if err := validate.RequiredString("messageType", "body", string(m.MessageType)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoPayMessageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoPayMessageDTO) UnmarshalBinary(b []byte) error {
	var res AutoPayMessageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
