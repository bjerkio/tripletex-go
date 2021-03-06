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

// SalesForceAddress sales force address
//
// swagger:model SalesForceAddress
type SalesForceAddress struct {

	// address line1
	// Read Only: true
	// Max Length: 255
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	// Read Only: true
	// Max Length: 255
	AddressLine2 string `json:"addressLine2,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// city
	// Read Only: true
	// Max Length: 100
	City string `json:"city,omitempty"`

	// country
	// Read Only: true
	Country *SalesForceCountry `json:"country,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// postal code
	// Read Only: true
	// Max Length: 100
	PostalCode string `json:"postalCode,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this sales force address
func (m *SalesForceAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLine2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SalesForceAddress) validateAddressLine1(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine1) { // not required
		return nil
	}

	if err := validate.MaxLength("addressLine1", "body", string(m.AddressLine1), 255); err != nil {
		return err
	}

	return nil
}

func (m *SalesForceAddress) validateAddressLine2(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	if err := validate.MaxLength("addressLine2", "body", string(m.AddressLine2), 255); err != nil {
		return err
	}

	return nil
}

func (m *SalesForceAddress) validateChanges(formats strfmt.Registry) error {

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

func (m *SalesForceAddress) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MaxLength("city", "body", string(m.City), 100); err != nil {
		return err
	}

	return nil
}

func (m *SalesForceAddress) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if m.Country != nil {
		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

func (m *SalesForceAddress) validatePostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PostalCode) { // not required
		return nil
	}

	if err := validate.MaxLength("postalCode", "body", string(m.PostalCode), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesForceAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesForceAddress) UnmarshalBinary(b []byte) error {
	var res SalesForceAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
