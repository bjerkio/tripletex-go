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

// PaymentTypeOut payment type out
//
// swagger:model PaymentTypeOut
type PaymentTypeOut struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// credit account
	// Required: true
	CreditAccount *Account `json:"creditAccount"`

	// description
	// Required: true
	// Max Length: 255
	Description *string `json:"description"`

	// id
	ID int32 `json:"id,omitempty"`

	// true if it should be a deduction from the wage. The module PROVISIONSALARY is required to both view and change this setting
	IsBruttoWageDeduction *bool `json:"isBruttoWageDeduction,omitempty"`

	// true if the payment type should be hidden from available payment types
	IsInactive *bool `json:"isInactive,omitempty"`

	// true if a separate voucher is required
	RequiresSeparateVoucher *bool `json:"requiresSeparateVoucher,omitempty"`

	// determines in which order the types should be listed. No 1 is listed first
	Sequence int32 `json:"sequence,omitempty"`

	// true if the payment type should be available in supplier invoices
	ShowIncomingInvoice *bool `json:"showIncomingInvoice,omitempty"`

	// true if the payment type should be available in vat returns
	ShowVatReturns *bool `json:"showVatReturns,omitempty"`

	// true if the payment type should be available in wage payments. The wage module is required to both view and change this setting
	ShowWagePayment *bool `json:"showWagePayment,omitempty"`

	// true if the payment type should be available in period transactionsThe wage module is required to both view and change this setting
	ShowWagePeriodTransaction *bool `json:"showWagePeriodTransaction,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this payment type out
func (m *PaymentTypeOut) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentTypeOut) validateChanges(formats strfmt.Registry) error {

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

func (m *PaymentTypeOut) validateCreditAccount(formats strfmt.Registry) error {

	if err := validate.Required("creditAccount", "body", m.CreditAccount); err != nil {
		return err
	}

	if m.CreditAccount != nil {
		if err := m.CreditAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditAccount")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentTypeOut) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentTypeOut) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentTypeOut) UnmarshalBinary(b []byte) error {
	var res PaymentTypeOut
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
