// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListResponseVatType list response vat type
//
// swagger:model ListResponseVatType
type ListResponseVatType struct {

	// count
	Count int32 `json:"count,omitempty"`

	// from
	From int32 `json:"from,omitempty"`

	// [DEPRECATED] Indicates whether there are more values available. Note: The value is not exact
	FullResultSize int32 `json:"fullResultSize,omitempty"`

	// values
	Values []*VatType `json:"values"`

	// Used to know if the paginated list has changed.
	VersionDigest string `json:"versionDigest,omitempty"`
}

// Validate validates this list response vat type
func (m *ListResponseVatType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListResponseVatType) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListResponseVatType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListResponseVatType) UnmarshalBinary(b []byte) error {
	var res ListResponseVatType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
