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

// VoucherMessage voucher message
//
// swagger:model VoucherMessage
type VoucherMessage struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// The message
	Content string `json:"content,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// The timestamp of the message
	SendTime string `json:"sendTime,omitempty"`

	// The employee that sent this message
	// Read Only: true
	Sender *Employee `json:"sender,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// The voucher to connect the message to, only set on create
	VoucherID int32 `json:"voucherId,omitempty"`
}

// Validate validates this voucher message
func (m *VoucherMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoucherMessage) validateChanges(formats strfmt.Registry) error {

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

func (m *VoucherMessage) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoucherMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoucherMessage) UnmarshalBinary(b []byte) error {
	var res VoucherMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}