// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CompanyAuthorityDTO company authority d t o
//
// swagger:model CompanyAuthorityDTO
type CompanyAuthorityDTO struct {

	// has company authority
	HasCompanyAuthority *bool `json:"hasCompanyAuthority,omitempty"`
}

// Validate validates this company authority d t o
func (m *CompanyAuthorityDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompanyAuthorityDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyAuthorityDTO) UnmarshalBinary(b []byte) error {
	var res CompanyAuthorityDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
