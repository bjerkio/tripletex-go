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

// Voucher voucher
//
// swagger:model Voucher
type Voucher struct {

	// If the documentation for the voucher has been provided from an external source (e.g. another system via API or a user upload) then this is a reference to the document. This is always a PDF. Note that a voucher may have both a document, an attachment and an ediDocument.
	// Read Only: true
	Attachment *Document `json:"attachment,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// date
	// Required: true
	Date *string `json:"date"`

	// description
	// Required: true
	Description *string `json:"description"`

	// If the documentation for the voucher has been generated by the system (e.g. an invoice) then this is a reference to the generated document. This document is always a PDF. Note that a voucher may have both a document, an attachment and an ediDocument.
	// Read Only: true
	Document *Document `json:"document,omitempty"`

	// If the voucher is created based on a machine readable document (such as EHF or EFO/NELFO) then this is a reference to that document. Note that a voucher may have both a document, an attachment and an ediDocument.
	// Read Only: true
	EdiDocument *Document `json:"ediDocument,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// System generated number that cannot be changed.
	// Read Only: true
	// Minimum: 0
	Number int32 `json:"number,omitempty"`

	// postings
	// Required: true
	Postings []*Posting `json:"postings"`

	// reverse voucher
	// Read Only: true
	ReverseVoucher *Voucher `json:"reverseVoucher,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// Voucher type. Must not be of type 'Utgående faktura' ('Outgoing Invoice') on new vouchers, instead use voucherType=null or use the Invoice endpoint.
	VoucherType *VoucherType `json:"voucherType,omitempty"`

	// System generated number that cannot be changed.
	// Read Only: true
	// Minimum: 0
	Year int32 `json:"year,omitempty"`
}

// Validate validates this voucher
func (m *Voucher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdiDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReverseVoucher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoucherType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Voucher) validateAttachment(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachment) { // not required
		return nil
	}

	if m.Attachment != nil {
		if err := m.Attachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

func (m *Voucher) validateChanges(formats strfmt.Registry) error {

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

func (m *Voucher) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *Voucher) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Voucher) validateDocument(formats strfmt.Registry) error {

	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *Voucher) validateEdiDocument(formats strfmt.Registry) error {

	if swag.IsZero(m.EdiDocument) { // not required
		return nil
	}

	if m.EdiDocument != nil {
		if err := m.EdiDocument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ediDocument")
			}
			return err
		}
	}

	return nil
}

func (m *Voucher) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MinimumInt("number", "body", int64(m.Number), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Voucher) validatePostings(formats strfmt.Registry) error {

	if err := validate.Required("postings", "body", m.Postings); err != nil {
		return err
	}

	for i := 0; i < len(m.Postings); i++ {
		if swag.IsZero(m.Postings[i]) { // not required
			continue
		}

		if m.Postings[i] != nil {
			if err := m.Postings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Voucher) validateReverseVoucher(formats strfmt.Registry) error {

	if swag.IsZero(m.ReverseVoucher) { // not required
		return nil
	}

	if m.ReverseVoucher != nil {
		if err := m.ReverseVoucher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reverseVoucher")
			}
			return err
		}
	}

	return nil
}

func (m *Voucher) validateVoucherType(formats strfmt.Registry) error {

	if swag.IsZero(m.VoucherType) { // not required
		return nil
	}

	if m.VoucherType != nil {
		if err := m.VoucherType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voucherType")
			}
			return err
		}
	}

	return nil
}

func (m *Voucher) validateYear(formats strfmt.Registry) error {

	if swag.IsZero(m.Year) { // not required
		return nil
	}

	if err := validate.MinimumInt("year", "body", int64(m.Year), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Voucher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Voucher) UnmarshalBinary(b []byte) error {
	var res Voucher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
