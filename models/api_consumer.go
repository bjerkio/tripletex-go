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

// APIConsumer Api consumer
//
// swagger:model ApiConsumer
type APIConsumer struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// consumer name
	// Required: true
	// Max Length: 100
	// Min Length: 5
	ConsumerName *string `json:"consumerName"`

	// emails
	// Required: true
	// Max Length: 255
	// Format: email
	Emails *strfmt.Email `json:"emails"`

	// id
	ID int32 `json:"id,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this Api consumer
func (m *APIConsumer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIConsumer) validateChanges(formats strfmt.Registry) error {

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

func (m *APIConsumer) validateConsumerName(formats strfmt.Registry) error {

	if err := validate.Required("consumerName", "body", m.ConsumerName); err != nil {
		return err
	}

	if err := validate.MinLength("consumerName", "body", string(*m.ConsumerName), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("consumerName", "body", string(*m.ConsumerName), 100); err != nil {
		return err
	}

	return nil
}

func (m *APIConsumer) validateEmails(formats strfmt.Registry) error {

	if err := validate.Required("emails", "body", m.Emails); err != nil {
		return err
	}

	if err := validate.MaxLength("emails", "body", string(*m.Emails), 255); err != nil {
		return err
	}

	if err := validate.FormatOf("emails", "body", "email", m.Emails.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIConsumer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIConsumer) UnmarshalBinary(b []byte) error {
	var res APIConsumer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
