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

// LeaveOfAbsenceType leave of absence type
//
// swagger:model LeaveOfAbsenceType
type LeaveOfAbsenceType struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// code
	// Max Length: 100
	Code string `json:"code,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// Defines the leave of absence type option.
	// Required: true
	// Enum: [LEAVE_OF_ABSENCE FURLOUGH PARENTAL_BENEFITS MILITARY_SERVICE EDUCATIONAL COMPASSIONATE]
	LeaveOfAbsenceType *string `json:"leaveOfAbsenceType"`

	// name n o
	// Max Length: 100
	NameNO string `json:"nameNO,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this leave of absence type
func (m *LeaveOfAbsenceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaveOfAbsenceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNO(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeaveOfAbsenceType) validateChanges(formats strfmt.Registry) error {

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

func (m *LeaveOfAbsenceType) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("code", "body", string(m.Code), 100); err != nil {
		return err
	}

	return nil
}

var leaveOfAbsenceTypeTypeLeaveOfAbsenceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LEAVE_OF_ABSENCE","FURLOUGH","PARENTAL_BENEFITS","MILITARY_SERVICE","EDUCATIONAL","COMPASSIONATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		leaveOfAbsenceTypeTypeLeaveOfAbsenceTypePropEnum = append(leaveOfAbsenceTypeTypeLeaveOfAbsenceTypePropEnum, v)
	}
}

const (

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypeLEAVEOFABSENCE captures enum value "LEAVE_OF_ABSENCE"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypeLEAVEOFABSENCE string = "LEAVE_OF_ABSENCE"

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypeFURLOUGH captures enum value "FURLOUGH"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypeFURLOUGH string = "FURLOUGH"

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypePARENTALBENEFITS captures enum value "PARENTAL_BENEFITS"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypePARENTALBENEFITS string = "PARENTAL_BENEFITS"

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypeMILITARYSERVICE captures enum value "MILITARY_SERVICE"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypeMILITARYSERVICE string = "MILITARY_SERVICE"

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypeEDUCATIONAL captures enum value "EDUCATIONAL"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypeEDUCATIONAL string = "EDUCATIONAL"

	// LeaveOfAbsenceTypeLeaveOfAbsenceTypeCOMPASSIONATE captures enum value "COMPASSIONATE"
	LeaveOfAbsenceTypeLeaveOfAbsenceTypeCOMPASSIONATE string = "COMPASSIONATE"
)

// prop value enum
func (m *LeaveOfAbsenceType) validateLeaveOfAbsenceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, leaveOfAbsenceTypeTypeLeaveOfAbsenceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LeaveOfAbsenceType) validateLeaveOfAbsenceType(formats strfmt.Registry) error {

	if err := validate.Required("leaveOfAbsenceType", "body", m.LeaveOfAbsenceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateLeaveOfAbsenceTypeEnum("leaveOfAbsenceType", "body", *m.LeaveOfAbsenceType); err != nil {
		return err
	}

	return nil
}

func (m *LeaveOfAbsenceType) validateNameNO(formats strfmt.Registry) error {

	if swag.IsZero(m.NameNO) { // not required
		return nil
	}

	if err := validate.MaxLength("nameNO", "body", string(m.NameNO), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeaveOfAbsenceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaveOfAbsenceType) UnmarshalBinary(b []byte) error {
	var res LeaveOfAbsenceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
