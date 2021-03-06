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

// OccupationCode occupation code
//
// swagger:model OccupationCode
type OccupationCode struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// code
	// Max Length: 7
	Code string `json:"code,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name n o
	// Max Length: 100
	NameNO string `json:"nameNO,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this occupation code
func (m *OccupationCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNO(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OccupationCode) validateChanges(formats strfmt.Registry) error {

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

func (m *OccupationCode) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("code", "body", string(m.Code), 7); err != nil {
		return err
	}

	return nil
}

func (m *OccupationCode) validateNameNO(formats strfmt.Registry) error {

	if swag.IsZero(m.NameNO) { // not required
		return nil
	}

	if err := validate.MaxLength("nameNO", "body", string(m.NameNO), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OccupationCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OccupationCode) UnmarshalBinary(b []byte) error {
	var res OccupationCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
