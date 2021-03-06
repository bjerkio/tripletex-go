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

// VoucherStatus voucher status
//
// swagger:model VoucherStatus
type VoucherStatus struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// comment
	Comment string `json:"comment,omitempty"`

	// Link to external object
	// Max Length: 255
	ExternalObjectURL string `json:"externalObjectUrl,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// 1 or 0 predefined status message
	// Enum: [NONE ONGOING NEEDS_APPROVAL WITHDRAWN SETTLED]
	Message string `json:"message,omitempty"`

	// reference number to external object
	// Max Length: 255
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// Process status
	// Enum: [WAITING DONE SKIPPED ERROR NONE PROCESSING]
	Status string `json:"status,omitempty"`

	// Time of last update
	// Read Only: true
	Timestamp string `json:"timestamp,omitempty"`

	// The type of process
	// Read Only: true
	// Enum: [TRIPLETEX DEMO_TYPE DEBT_COLLECTION_3_4 DEBT_COLLECTION_4]
	Type string `json:"type,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// The voucher.
	Voucher *Voucher `json:"voucher,omitempty"`
}

// Validate validates this voucher status
func (m *VoucherStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalObjectURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *VoucherStatus) validateChanges(formats strfmt.Registry) error {

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

func (m *VoucherStatus) validateExternalObjectURL(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalObjectURL) { // not required
		return nil
	}

	if err := validate.MaxLength("externalObjectUrl", "body", string(m.ExternalObjectURL), 255); err != nil {
		return err
	}

	return nil
}

var voucherStatusTypeMessagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","ONGOING","NEEDS_APPROVAL","WITHDRAWN","SETTLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voucherStatusTypeMessagePropEnum = append(voucherStatusTypeMessagePropEnum, v)
	}
}

const (

	// VoucherStatusMessageNONE captures enum value "NONE"
	VoucherStatusMessageNONE string = "NONE"

	// VoucherStatusMessageONGOING captures enum value "ONGOING"
	VoucherStatusMessageONGOING string = "ONGOING"

	// VoucherStatusMessageNEEDSAPPROVAL captures enum value "NEEDS_APPROVAL"
	VoucherStatusMessageNEEDSAPPROVAL string = "NEEDS_APPROVAL"

	// VoucherStatusMessageWITHDRAWN captures enum value "WITHDRAWN"
	VoucherStatusMessageWITHDRAWN string = "WITHDRAWN"

	// VoucherStatusMessageSETTLED captures enum value "SETTLED"
	VoucherStatusMessageSETTLED string = "SETTLED"
)

// prop value enum
func (m *VoucherStatus) validateMessageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voucherStatusTypeMessagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoucherStatus) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageEnum("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *VoucherStatus) validateReferenceNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("referenceNumber", "body", string(m.ReferenceNumber), 255); err != nil {
		return err
	}

	return nil
}

var voucherStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WAITING","DONE","SKIPPED","ERROR","NONE","PROCESSING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voucherStatusTypeStatusPropEnum = append(voucherStatusTypeStatusPropEnum, v)
	}
}

const (

	// VoucherStatusStatusWAITING captures enum value "WAITING"
	VoucherStatusStatusWAITING string = "WAITING"

	// VoucherStatusStatusDONE captures enum value "DONE"
	VoucherStatusStatusDONE string = "DONE"

	// VoucherStatusStatusSKIPPED captures enum value "SKIPPED"
	VoucherStatusStatusSKIPPED string = "SKIPPED"

	// VoucherStatusStatusERROR captures enum value "ERROR"
	VoucherStatusStatusERROR string = "ERROR"

	// VoucherStatusStatusNONE captures enum value "NONE"
	VoucherStatusStatusNONE string = "NONE"

	// VoucherStatusStatusPROCESSING captures enum value "PROCESSING"
	VoucherStatusStatusPROCESSING string = "PROCESSING"
)

// prop value enum
func (m *VoucherStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voucherStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoucherStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var voucherStatusTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRIPLETEX","DEMO_TYPE","DEBT_COLLECTION_3_4","DEBT_COLLECTION_4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voucherStatusTypeTypePropEnum = append(voucherStatusTypeTypePropEnum, v)
	}
}

const (

	// VoucherStatusTypeTRIPLETEX captures enum value "TRIPLETEX"
	VoucherStatusTypeTRIPLETEX string = "TRIPLETEX"

	// VoucherStatusTypeDEMOTYPE captures enum value "DEMO_TYPE"
	VoucherStatusTypeDEMOTYPE string = "DEMO_TYPE"

	// VoucherStatusTypeDEBTCOLLECTION34 captures enum value "DEBT_COLLECTION_3_4"
	VoucherStatusTypeDEBTCOLLECTION34 string = "DEBT_COLLECTION_3_4"

	// VoucherStatusTypeDEBTCOLLECTION4 captures enum value "DEBT_COLLECTION_4"
	VoucherStatusTypeDEBTCOLLECTION4 string = "DEBT_COLLECTION_4"
)

// prop value enum
func (m *VoucherStatus) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voucherStatusTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoucherStatus) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *VoucherStatus) validateVoucher(formats strfmt.Registry) error {

	if swag.IsZero(m.Voucher) { // not required
		return nil
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
func (m *VoucherStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoucherStatus) UnmarshalBinary(b []byte) error {
	var res VoucherStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
