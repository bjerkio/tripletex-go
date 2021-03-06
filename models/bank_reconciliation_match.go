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

// BankReconciliationMatch bank reconciliation match
//
// swagger:model BankReconciliationMatch
type BankReconciliationMatch struct {

	// bank reconciliation
	// Required: true
	BankReconciliation *BankReconciliation `json:"bankReconciliation"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// id
	ID int32 `json:"id,omitempty"`

	// Match postings
	Postings []*Posting `json:"postings"`

	// Match transactions
	Transactions []*BankTransaction `json:"transactions"`

	// Type of match, MANUAL and APPROVED_SUGGESTION are considered part of reconciliation.
	// Enum: [MANUAL PENDING_SUGGESTION REJECTED_SUGGESTION APPROVED_SUGGESTION ADJUSTMENT AUTO_MATCHED REJECTED_AUTO_MATCH]
	Type string `json:"type,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this bank reconciliation match
func (m *BankReconciliationMatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankReconciliation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
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

func (m *BankReconciliationMatch) validateBankReconciliation(formats strfmt.Registry) error {

	if err := validate.Required("bankReconciliation", "body", m.BankReconciliation); err != nil {
		return err
	}

	if m.BankReconciliation != nil {
		if err := m.BankReconciliation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bankReconciliation")
			}
			return err
		}
	}

	return nil
}

func (m *BankReconciliationMatch) validateChanges(formats strfmt.Registry) error {

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

func (m *BankReconciliationMatch) validatePostings(formats strfmt.Registry) error {

	if swag.IsZero(m.Postings) { // not required
		return nil
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

func (m *BankReconciliationMatch) validateTransactions(formats strfmt.Registry) error {

	if swag.IsZero(m.Transactions) { // not required
		return nil
	}

	for i := 0; i < len(m.Transactions); i++ {
		if swag.IsZero(m.Transactions[i]) { // not required
			continue
		}

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var bankReconciliationMatchTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","PENDING_SUGGESTION","REJECTED_SUGGESTION","APPROVED_SUGGESTION","ADJUSTMENT","AUTO_MATCHED","REJECTED_AUTO_MATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bankReconciliationMatchTypeTypePropEnum = append(bankReconciliationMatchTypeTypePropEnum, v)
	}
}

const (

	// BankReconciliationMatchTypeMANUAL captures enum value "MANUAL"
	BankReconciliationMatchTypeMANUAL string = "MANUAL"

	// BankReconciliationMatchTypePENDINGSUGGESTION captures enum value "PENDING_SUGGESTION"
	BankReconciliationMatchTypePENDINGSUGGESTION string = "PENDING_SUGGESTION"

	// BankReconciliationMatchTypeREJECTEDSUGGESTION captures enum value "REJECTED_SUGGESTION"
	BankReconciliationMatchTypeREJECTEDSUGGESTION string = "REJECTED_SUGGESTION"

	// BankReconciliationMatchTypeAPPROVEDSUGGESTION captures enum value "APPROVED_SUGGESTION"
	BankReconciliationMatchTypeAPPROVEDSUGGESTION string = "APPROVED_SUGGESTION"

	// BankReconciliationMatchTypeADJUSTMENT captures enum value "ADJUSTMENT"
	BankReconciliationMatchTypeADJUSTMENT string = "ADJUSTMENT"

	// BankReconciliationMatchTypeAUTOMATCHED captures enum value "AUTO_MATCHED"
	BankReconciliationMatchTypeAUTOMATCHED string = "AUTO_MATCHED"

	// BankReconciliationMatchTypeREJECTEDAUTOMATCH captures enum value "REJECTED_AUTO_MATCH"
	BankReconciliationMatchTypeREJECTEDAUTOMATCH string = "REJECTED_AUTO_MATCH"
)

// prop value enum
func (m *BankReconciliationMatch) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bankReconciliationMatchTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BankReconciliationMatch) validateType(formats strfmt.Registry) error {

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
func (m *BankReconciliationMatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankReconciliationMatch) UnmarshalBinary(b []byte) error {
	var res BankReconciliationMatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
