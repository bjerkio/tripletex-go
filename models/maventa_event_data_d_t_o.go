// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MaventaEventDataDTO maventa event data d t o
//
// swagger:model MaventaEventDataDTO
type MaventaEventDataDTO struct {

	// destination
	// Required: true
	// Read Only: true
	Destination string `json:"destination"`

	// error message
	// Read Only: true
	ErrorMessage string `json:"error_message,omitempty"`

	// invoice id
	// Required: true
	// Read Only: true
	InvoiceID string `json:"invoice_id"`

	// invoice number
	// Required: true
	// Read Only: true
	InvoiceNumber string `json:"invoice_number"`

	// recipient bid
	// Required: true
	// Read Only: true
	RecipientBid string `json:"recipient_bid"`

	// recipient name
	// Required: true
	// Read Only: true
	RecipientName string `json:"recipient_name"`
}

// Validate validates this maventa event data d t o
func (m *MaventaEventDataDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientBid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaventaEventDataDTO) validateDestination(formats strfmt.Registry) error {

	if err := validate.RequiredString("destination", "body", string(m.Destination)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaEventDataDTO) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("invoice_id", "body", string(m.InvoiceID)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaEventDataDTO) validateInvoiceNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("invoice_number", "body", string(m.InvoiceNumber)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaEventDataDTO) validateRecipientBid(formats strfmt.Registry) error {

	if err := validate.RequiredString("recipient_bid", "body", string(m.RecipientBid)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaEventDataDTO) validateRecipientName(formats strfmt.Registry) error {

	if err := validate.RequiredString("recipient_name", "body", string(m.RecipientName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaventaEventDataDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaventaEventDataDTO) UnmarshalBinary(b []byte) error {
	var res MaventaEventDataDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
