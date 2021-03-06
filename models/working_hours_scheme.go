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

// WorkingHoursScheme working hours scheme
//
// swagger:model WorkingHoursScheme
type WorkingHoursScheme struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// code
	// Max Length: 100
	Code string `json:"code,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name n o
	// Max Length: 100
	NameNO string `json:"nameNO,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// Defines the working hours scheme option.
	// Required: true
	// Enum: [NOT_SHIFT ROUND_THE_CLOCK SHIFT_365 OFFSHORE_336 CONTINUOUS OTHER_SHIFT]
	WorkingHoursScheme *string `json:"workingHoursScheme"`
}

// Validate validates this working hours scheme
func (m *WorkingHoursScheme) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkingHoursScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkingHoursScheme) validateChanges(formats strfmt.Registry) error {

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

func (m *WorkingHoursScheme) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("code", "body", string(m.Code), 100); err != nil {
		return err
	}

	return nil
}

func (m *WorkingHoursScheme) validateNameNO(formats strfmt.Registry) error {

	if swag.IsZero(m.NameNO) { // not required
		return nil
	}

	if err := validate.MaxLength("nameNO", "body", string(m.NameNO), 100); err != nil {
		return err
	}

	return nil
}

var workingHoursSchemeTypeWorkingHoursSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_SHIFT","ROUND_THE_CLOCK","SHIFT_365","OFFSHORE_336","CONTINUOUS","OTHER_SHIFT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workingHoursSchemeTypeWorkingHoursSchemePropEnum = append(workingHoursSchemeTypeWorkingHoursSchemePropEnum, v)
	}
}

const (

	// WorkingHoursSchemeWorkingHoursSchemeNOTSHIFT captures enum value "NOT_SHIFT"
	WorkingHoursSchemeWorkingHoursSchemeNOTSHIFT string = "NOT_SHIFT"

	// WorkingHoursSchemeWorkingHoursSchemeROUNDTHECLOCK captures enum value "ROUND_THE_CLOCK"
	WorkingHoursSchemeWorkingHoursSchemeROUNDTHECLOCK string = "ROUND_THE_CLOCK"

	// WorkingHoursSchemeWorkingHoursSchemeSHIFT365 captures enum value "SHIFT_365"
	WorkingHoursSchemeWorkingHoursSchemeSHIFT365 string = "SHIFT_365"

	// WorkingHoursSchemeWorkingHoursSchemeOFFSHORE336 captures enum value "OFFSHORE_336"
	WorkingHoursSchemeWorkingHoursSchemeOFFSHORE336 string = "OFFSHORE_336"

	// WorkingHoursSchemeWorkingHoursSchemeCONTINUOUS captures enum value "CONTINUOUS"
	WorkingHoursSchemeWorkingHoursSchemeCONTINUOUS string = "CONTINUOUS"

	// WorkingHoursSchemeWorkingHoursSchemeOTHERSHIFT captures enum value "OTHER_SHIFT"
	WorkingHoursSchemeWorkingHoursSchemeOTHERSHIFT string = "OTHER_SHIFT"
)

// prop value enum
func (m *WorkingHoursScheme) validateWorkingHoursSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workingHoursSchemeTypeWorkingHoursSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkingHoursScheme) validateWorkingHoursScheme(formats strfmt.Registry) error {

	if err := validate.Required("workingHoursScheme", "body", m.WorkingHoursScheme); err != nil {
		return err
	}

	// value enum
	if err := m.validateWorkingHoursSchemeEnum("workingHoursScheme", "body", *m.WorkingHoursScheme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkingHoursScheme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkingHoursScheme) UnmarshalBinary(b []byte) error {
	var res WorkingHoursScheme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
