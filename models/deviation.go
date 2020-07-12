// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// Deviation deviation
//
// swagger:model Deviation
type Deviation struct {

	// action
	// Enum: [ACTION_IGNORE ACTION_GENERATE_RESTORDER ACTION_RETURN ACTION_RETURN_GENERATE_RESTORDER]
	Action string `json:"action,omitempty"`

	// cause
	// Enum: [CAUSE_DEFECT CAUSE_TOO_FEW CAUSE_TOO_MANY CAUSE_REPLACEMENT]
	Cause string `json:"cause,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// comment
	Comment string `json:"comment,omitempty"`

	// date
	// Required: true
	Date *string `json:"date"`

	// deviation
	// Read Only: true
	Deviation float64 `json:"deviation,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// purchase order line
	// Required: true
	PurchaseOrderLine *OrderLine `json:"purchaseOrderLine"`

	// quantity ordered
	// Read Only: true
	QuantityOrdered float64 `json:"quantityOrdered,omitempty"`

	// quantity received
	// Read Only: true
	QuantityReceived float64 `json:"quantityReceived,omitempty"`

	// received by
	// Read Only: true
	ReceivedBy string `json:"receivedBy,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this deviation
func (m *Deviation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderLine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deviationTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTION_IGNORE","ACTION_GENERATE_RESTORDER","ACTION_RETURN","ACTION_RETURN_GENERATE_RESTORDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviationTypeActionPropEnum = append(deviationTypeActionPropEnum, v)
	}
}

const (

	// DeviationActionACTIONIGNORE captures enum value "ACTION_IGNORE"
	DeviationActionACTIONIGNORE string = "ACTION_IGNORE"

	// DeviationActionACTIONGENERATERESTORDER captures enum value "ACTION_GENERATE_RESTORDER"
	DeviationActionACTIONGENERATERESTORDER string = "ACTION_GENERATE_RESTORDER"

	// DeviationActionACTIONRETURN captures enum value "ACTION_RETURN"
	DeviationActionACTIONRETURN string = "ACTION_RETURN"

	// DeviationActionACTIONRETURNGENERATERESTORDER captures enum value "ACTION_RETURN_GENERATE_RESTORDER"
	DeviationActionACTIONRETURNGENERATERESTORDER string = "ACTION_RETURN_GENERATE_RESTORDER"
)

// prop value enum
func (m *Deviation) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviationTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Deviation) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var deviationTypeCausePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CAUSE_DEFECT","CAUSE_TOO_FEW","CAUSE_TOO_MANY","CAUSE_REPLACEMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviationTypeCausePropEnum = append(deviationTypeCausePropEnum, v)
	}
}

const (

	// DeviationCauseCAUSEDEFECT captures enum value "CAUSE_DEFECT"
	DeviationCauseCAUSEDEFECT string = "CAUSE_DEFECT"

	// DeviationCauseCAUSETOOFEW captures enum value "CAUSE_TOO_FEW"
	DeviationCauseCAUSETOOFEW string = "CAUSE_TOO_FEW"

	// DeviationCauseCAUSETOOMANY captures enum value "CAUSE_TOO_MANY"
	DeviationCauseCAUSETOOMANY string = "CAUSE_TOO_MANY"

	// DeviationCauseCAUSEREPLACEMENT captures enum value "CAUSE_REPLACEMENT"
	DeviationCauseCAUSEREPLACEMENT string = "CAUSE_REPLACEMENT"
)

// prop value enum
func (m *Deviation) validateCauseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviationTypeCausePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Deviation) validateCause(formats strfmt.Registry) error {

	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	// value enum
	if err := m.validateCauseEnum("cause", "body", m.Cause); err != nil {
		return err
	}

	return nil
}

func (m *Deviation) validateChanges(formats strfmt.Registry) error {

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

func (m *Deviation) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *Deviation) validatePurchaseOrderLine(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderLine", "body", m.PurchaseOrderLine); err != nil {
		return err
	}

	if m.PurchaseOrderLine != nil {
		if err := m.PurchaseOrderLine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purchaseOrderLine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deviation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deviation) UnmarshalBinary(b []byte) error {
	var res Deviation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
