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

// DeliveryAddress delivery address
//
// swagger:model DeliveryAddress
type DeliveryAddress struct {

	// address line1
	// Max Length: 255
	// Min Length: 0
	AddressLine1 *string `json:"addressLine1,omitempty"`

	// address line2
	// Max Length: 255
	// Min Length: 0
	AddressLine2 *string `json:"addressLine2,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// city
	// Max Length: 100
	// Min Length: 0
	City *string `json:"city,omitempty"`

	// country
	Country *Country `json:"country,omitempty"`

	// employee
	Employee *Employee `json:"employee,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	// Max Length: 100
	Name string `json:"name,omitempty"`

	// postal code
	// Max Length: 100
	// Min Length: 0
	PostalCode *string `json:"postalCode,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this delivery address
func (m *DeliveryAddress) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEmployee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *DeliveryAddress) validateAddressLine1(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine1) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine1", "body", string(*m.AddressLine1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine1", "body", string(*m.AddressLine1), 255); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress) validateAddressLine2(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine2", "body", string(*m.AddressLine2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine2", "body", string(*m.AddressLine2), 255); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress) validateChanges(formats strfmt.Registry) error {

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

func (m *DeliveryAddress) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(*m.City), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(*m.City), 100); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress) validateCountry(formats strfmt.Registry) error {

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

func (m *DeliveryAddress) validateEmployee(formats strfmt.Registry) error {

	if swag.IsZero(m.Employee) { // not required
		return nil
	}

	if m.Employee != nil {
		if err := m.Employee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee")
			}
			return err
		}
	}

	return nil
}

func (m *DeliveryAddress) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryAddress) validatePostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PostalCode) { // not required
		return nil
	}

	if err := validate.MinLength("postalCode", "body", string(*m.PostalCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("postalCode", "body", string(*m.PostalCode), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryAddress) UnmarshalBinary(b []byte) error {
	var res DeliveryAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
