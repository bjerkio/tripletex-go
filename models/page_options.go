// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PageOptions page options
//
// swagger:model PageOptions
type PageOptions struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// The actual options as a JSON blob
	// Required: true
	Data map[string]interface{} `json:"data"`

	// id
	ID int32 `json:"id,omitempty"`

	// The lookup key for this PageOptions entry
	// Max Length: 255
	Key string `json:"key,omitempty"`

	// The type that `data` is
	// Enum: [IncomingInvoiceViewOptions]
	Type string `json:"type,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this page options
func (m *PageOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOptions) validateChanges(formats strfmt.Registry) error {

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

func (m *PageOptions) validateData(formats strfmt.Registry) error {

	for k := range m.Data {

		if err := validate.Required("data"+"."+k, "body", m.Data[k]); err != nil {
			return err
		}

		if err := validate.Required("data"+"."+k, "body", m.Data[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PageOptions) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.MaxLength("key", "body", string(m.Key), 255); err != nil {
		return err
	}

	return nil
}

var pageOptionsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IncomingInvoiceViewOptions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pageOptionsTypeTypePropEnum = append(pageOptionsTypeTypePropEnum, v)
	}
}

const (

	// PageOptionsTypeIncomingInvoiceViewOptions captures enum value "IncomingInvoiceViewOptions"
	PageOptionsTypeIncomingInvoiceViewOptions string = "IncomingInvoiceViewOptions"
)

// prop value enum
func (m *PageOptions) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pageOptionsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PageOptions) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageOptions) UnmarshalBinary(b []byte) error {
	var res PageOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
