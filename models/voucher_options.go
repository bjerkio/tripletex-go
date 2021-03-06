// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VoucherOptions voucher options
//
// swagger:model VoucherOptions
type VoucherOptions struct {

	// A data structure containing information about the delete operation.
	// Read Only: true
	Delete *Delete `json:"delete,omitempty"`
}

// Validate validates this voucher options
func (m *VoucherOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelete(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoucherOptions) validateDelete(formats strfmt.Registry) error {

	if swag.IsZero(m.Delete) { // not required
		return nil
	}

	if m.Delete != nil {
		if err := m.Delete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delete")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoucherOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoucherOptions) UnmarshalBinary(b []byte) error {
	var res VoucherOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
