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

// PurchaseOrderIncomingInvoiceRelation purchase order incoming invoice relation
//
// swagger:model PurchaseOrderIncomingInvoiceRelation
type PurchaseOrderIncomingInvoiceRelation struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// id
	ID int32 `json:"id,omitempty"`

	// order out
	// Required: true
	OrderOut *PurchaseOrder `json:"orderOut"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// voucher
	// Required: true
	Voucher *Voucher `json:"voucher"`
}

// Validate validates this purchase order incoming invoice relation
func (m *PurchaseOrderIncomingInvoiceRelation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoucher(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseOrderIncomingInvoiceRelation) validateChanges(formats strfmt.Registry) error {

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

func (m *PurchaseOrderIncomingInvoiceRelation) validateOrderOut(formats strfmt.Registry) error {

	if err := validate.Required("orderOut", "body", m.OrderOut); err != nil {
		return err
	}

	if m.OrderOut != nil {
		if err := m.OrderOut.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderOut")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrderIncomingInvoiceRelation) validateVoucher(formats strfmt.Registry) error {

	if err := validate.Required("voucher", "body", m.Voucher); err != nil {
		return err
	}

	if m.Voucher != nil {
		if err := m.Voucher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voucher")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurchaseOrderIncomingInvoiceRelation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchaseOrderIncomingInvoiceRelation) UnmarshalBinary(b []byte) error {
	var res PurchaseOrderIncomingInvoiceRelation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
