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

// Country country
//
// swagger:model Country
type Country struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// id
	ID int32 `json:"id,omitempty"`

	// The ISO 3166-1 Alpha2 code of the country (2 letters). https://en.wikipedia.org/wiki/ISO_3166-1
	// Read Only: true
	IsoAlpha2Code string `json:"isoAlpha2Code,omitempty"`

	// The ISO 3166-1 Alpha3 code of the country (3 letters). https://en.wikipedia.org/wiki/ISO_3166-1
	// Read Only: true
	// Max Length: 3
	IsoAlpha3Code string `json:"isoAlpha3Code,omitempty"`

	// The ISO 3166-1 numeric code of the country (3 digits). https://en.wikipedia.org/wiki/ISO_3166-1
	// Read Only: true
	IsoNumericCode string `json:"isoNumericCode,omitempty"`

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this country
func (m *Country) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsoAlpha3Code(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Country) validateChanges(formats strfmt.Registry) error {

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

func (m *Country) validateIsoAlpha3Code(formats strfmt.Registry) error {

	if swag.IsZero(m.IsoAlpha3Code) { // not required
		return nil
	}

	if err := validate.MaxLength("isoAlpha3Code", "body", string(m.IsoAlpha3Code), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Country) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Country) UnmarshalBinary(b []byte) error {
	var res Country
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
