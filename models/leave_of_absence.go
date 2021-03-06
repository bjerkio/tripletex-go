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

// LeaveOfAbsence leave of absence
//
// swagger:model LeaveOfAbsence
type LeaveOfAbsence struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// employment
	Employment *Employment `json:"employment,omitempty"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is wage deduction
	IsWageDeduction *bool `json:"isWageDeduction,omitempty"`

	// Existing leave of absence ID used by the current accounting system
	// Max Length: 255
	LeaveOfAbsenceID string `json:"leaveOfAbsenceId,omitempty"`

	// percentage
	// Required: true
	Percentage *float64 `json:"percentage"`

	// start date
	// Required: true
	StartDate *string `json:"startDate"`

	// Define the leave of absence type.
	// Enum: [LEAVE_OF_ABSENCE FURLOUGH PARENTAL_BENEFITS MILITARY_SERVICE EDUCATIONAL COMPASSIONATE]
	Type string `json:"type,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this leave of absence
func (m *LeaveOfAbsence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaveOfAbsenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
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

func (m *LeaveOfAbsence) validateChanges(formats strfmt.Registry) error {

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

func (m *LeaveOfAbsence) validateEmployment(formats strfmt.Registry) error {

	if swag.IsZero(m.Employment) { // not required
		return nil
	}

	if m.Employment != nil {
		if err := m.Employment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employment")
			}
			return err
		}
	}

	return nil
}

func (m *LeaveOfAbsence) validateLeaveOfAbsenceID(formats strfmt.Registry) error {

	if swag.IsZero(m.LeaveOfAbsenceID) { // not required
		return nil
	}

	if err := validate.MaxLength("leaveOfAbsenceId", "body", string(m.LeaveOfAbsenceID), 255); err != nil {
		return err
	}

	return nil
}

func (m *LeaveOfAbsence) validatePercentage(formats strfmt.Registry) error {

	if err := validate.Required("percentage", "body", m.Percentage); err != nil {
		return err
	}

	return nil
}

func (m *LeaveOfAbsence) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	return nil
}

var leaveOfAbsenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LEAVE_OF_ABSENCE","FURLOUGH","PARENTAL_BENEFITS","MILITARY_SERVICE","EDUCATIONAL","COMPASSIONATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		leaveOfAbsenceTypeTypePropEnum = append(leaveOfAbsenceTypeTypePropEnum, v)
	}
}

const (

	// LeaveOfAbsenceTypeLEAVEOFABSENCE captures enum value "LEAVE_OF_ABSENCE"
	LeaveOfAbsenceTypeLEAVEOFABSENCE string = "LEAVE_OF_ABSENCE"

	// LeaveOfAbsenceTypeFURLOUGH captures enum value "FURLOUGH"
	LeaveOfAbsenceTypeFURLOUGH string = "FURLOUGH"

	// LeaveOfAbsenceTypePARENTALBENEFITS captures enum value "PARENTAL_BENEFITS"
	LeaveOfAbsenceTypePARENTALBENEFITS string = "PARENTAL_BENEFITS"

	// LeaveOfAbsenceTypeMILITARYSERVICE captures enum value "MILITARY_SERVICE"
	LeaveOfAbsenceTypeMILITARYSERVICE string = "MILITARY_SERVICE"

	// LeaveOfAbsenceTypeEDUCATIONAL captures enum value "EDUCATIONAL"
	LeaveOfAbsenceTypeEDUCATIONAL string = "EDUCATIONAL"

	// LeaveOfAbsenceTypeCOMPASSIONATE captures enum value "COMPASSIONATE"
	LeaveOfAbsenceTypeCOMPASSIONATE string = "COMPASSIONATE"
)

// prop value enum
func (m *LeaveOfAbsence) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, leaveOfAbsenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LeaveOfAbsence) validateType(formats strfmt.Registry) error {

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
func (m *LeaveOfAbsence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaveOfAbsence) UnmarshalBinary(b []byte) error {
	var res LeaveOfAbsence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
