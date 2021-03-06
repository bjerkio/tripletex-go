// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransportType transport type
//
// swagger:model TransportType
type TransportType struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// code
	// Read Only: true
	// Max Length: 100
	Code string `json:"code,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is pick up
	// Read Only: true
	IsPickUp *bool `json:"isPickUp,omitempty"`

	// name key
	// Read Only: true
	// Max Length: 100
	NameKey string `json:"nameKey,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this transport type
func (m *TransportType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportType) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TransportType) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("code", "body", string(m.Code), 100); err != nil {
		return err
	}

	return nil
}

func (m *TransportType) validateNameKey(formats strfmt.Registry) error {

	if swag.IsZero(m.NameKey) { // not required
		return nil
	}

	if err := validate.MaxLength("nameKey", "body", string(m.NameKey), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportType) UnmarshalBinary(b []byte) error {
	var res TransportType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
