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

// EmploymentDetails employment details
//
// swagger:model EmploymentDetails
type EmploymentDetails struct {

	// annual salary
	AnnualSalary float64 `json:"annualSalary,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// date
	Date string `json:"date,omitempty"`

	// employment
	Employment *Employment `json:"employment,omitempty"`

	// Define the employment type.
	// Enum: [ORDINARY MARITIME FREELANCE]
	EmploymentType string `json:"employmentType,omitempty"`

	// hourly wage
	HourlyWage float64 `json:"hourlyWage,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// maritime employment
	MaritimeEmployment *MaritimeEmployment `json:"maritimeEmployment,omitempty"`

	// To find the right value to enter in this field, you could go to GET /employee/employment/occupationCode to get a list of valid ID's.
	OccupationCode *OccupationCode `json:"occupationCode,omitempty"`

	// payroll tax municipality Id
	PayrollTaxMunicipalityID *Municipality `json:"payrollTaxMunicipalityId,omitempty"`

	// percentage of full time equivalent
	// Required: true
	PercentageOfFullTimeEquivalent *float64 `json:"percentageOfFullTimeEquivalent"`

	// Define the remuneration type.
	// Enum: [MONTHLY_WAGE HOURLY_WAGE COMMISION_PERCENTAGE FEE PIECEWORK_WAGE]
	RemunerationType string `json:"remunerationType,omitempty"`

	// shift duration hours
	ShiftDurationHours float64 `json:"shiftDurationHours,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// Define the working hours scheme type. If you enter a value for SHIFT WORK, you must also enter value for shiftDurationHours
	// Enum: [NOT_SHIFT ROUND_THE_CLOCK SHIFT_365 OFFSHORE_336 CONTINUOUS OTHER_SHIFT]
	WorkingHoursScheme string `json:"workingHoursScheme,omitempty"`
}

// Validate validates this employment details
func (m *EmploymentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaritimeEmployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOccupationCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayrollTaxMunicipalityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentageOfFullTimeEquivalent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemunerationType(formats); err != nil {
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

func (m *EmploymentDetails) validateChanges(formats strfmt.Registry) error {

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

func (m *EmploymentDetails) validateEmployment(formats strfmt.Registry) error {

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

var employmentDetailsTypeEmploymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ORDINARY","MARITIME","FREELANCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		employmentDetailsTypeEmploymentTypePropEnum = append(employmentDetailsTypeEmploymentTypePropEnum, v)
	}
}

const (

	// EmploymentDetailsEmploymentTypeORDINARY captures enum value "ORDINARY"
	EmploymentDetailsEmploymentTypeORDINARY string = "ORDINARY"

	// EmploymentDetailsEmploymentTypeMARITIME captures enum value "MARITIME"
	EmploymentDetailsEmploymentTypeMARITIME string = "MARITIME"

	// EmploymentDetailsEmploymentTypeFREELANCE captures enum value "FREELANCE"
	EmploymentDetailsEmploymentTypeFREELANCE string = "FREELANCE"
)

// prop value enum
func (m *EmploymentDetails) validateEmploymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, employmentDetailsTypeEmploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmploymentDetails) validateEmploymentType(formats strfmt.Registry) error {

	if swag.IsZero(m.EmploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEmploymentTypeEnum("employmentType", "body", m.EmploymentType); err != nil {
		return err
	}

	return nil
}

func (m *EmploymentDetails) validateMaritimeEmployment(formats strfmt.Registry) error {

	if swag.IsZero(m.MaritimeEmployment) { // not required
		return nil
	}

	if m.MaritimeEmployment != nil {
		if err := m.MaritimeEmployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maritimeEmployment")
			}
			return err
		}
	}

	return nil
}

func (m *EmploymentDetails) validateOccupationCode(formats strfmt.Registry) error {

	if swag.IsZero(m.OccupationCode) { // not required
		return nil
	}

	if m.OccupationCode != nil {
		if err := m.OccupationCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("occupationCode")
			}
			return err
		}
	}

	return nil
}

func (m *EmploymentDetails) validatePayrollTaxMunicipalityID(formats strfmt.Registry) error {

	if swag.IsZero(m.PayrollTaxMunicipalityID) { // not required
		return nil
	}

	if m.PayrollTaxMunicipalityID != nil {
		if err := m.PayrollTaxMunicipalityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payrollTaxMunicipalityId")
			}
			return err
		}
	}

	return nil
}

func (m *EmploymentDetails) validatePercentageOfFullTimeEquivalent(formats strfmt.Registry) error {

	if err := validate.Required("percentageOfFullTimeEquivalent", "body", m.PercentageOfFullTimeEquivalent); err != nil {
		return err
	}

	return nil
}

var employmentDetailsTypeRemunerationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MONTHLY_WAGE","HOURLY_WAGE","COMMISION_PERCENTAGE","FEE","PIECEWORK_WAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		employmentDetailsTypeRemunerationTypePropEnum = append(employmentDetailsTypeRemunerationTypePropEnum, v)
	}
}

const (

	// EmploymentDetailsRemunerationTypeMONTHLYWAGE captures enum value "MONTHLY_WAGE"
	EmploymentDetailsRemunerationTypeMONTHLYWAGE string = "MONTHLY_WAGE"

	// EmploymentDetailsRemunerationTypeHOURLYWAGE captures enum value "HOURLY_WAGE"
	EmploymentDetailsRemunerationTypeHOURLYWAGE string = "HOURLY_WAGE"

	// EmploymentDetailsRemunerationTypeCOMMISIONPERCENTAGE captures enum value "COMMISION_PERCENTAGE"
	EmploymentDetailsRemunerationTypeCOMMISIONPERCENTAGE string = "COMMISION_PERCENTAGE"

	// EmploymentDetailsRemunerationTypeFEE captures enum value "FEE"
	EmploymentDetailsRemunerationTypeFEE string = "FEE"

	// EmploymentDetailsRemunerationTypePIECEWORKWAGE captures enum value "PIECEWORK_WAGE"
	EmploymentDetailsRemunerationTypePIECEWORKWAGE string = "PIECEWORK_WAGE"
)

// prop value enum
func (m *EmploymentDetails) validateRemunerationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, employmentDetailsTypeRemunerationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmploymentDetails) validateRemunerationType(formats strfmt.Registry) error {

	if swag.IsZero(m.RemunerationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemunerationTypeEnum("remunerationType", "body", m.RemunerationType); err != nil {
		return err
	}

	return nil
}

var employmentDetailsTypeWorkingHoursSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_SHIFT","ROUND_THE_CLOCK","SHIFT_365","OFFSHORE_336","CONTINUOUS","OTHER_SHIFT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		employmentDetailsTypeWorkingHoursSchemePropEnum = append(employmentDetailsTypeWorkingHoursSchemePropEnum, v)
	}
}

const (

	// EmploymentDetailsWorkingHoursSchemeNOTSHIFT captures enum value "NOT_SHIFT"
	EmploymentDetailsWorkingHoursSchemeNOTSHIFT string = "NOT_SHIFT"

	// EmploymentDetailsWorkingHoursSchemeROUNDTHECLOCK captures enum value "ROUND_THE_CLOCK"
	EmploymentDetailsWorkingHoursSchemeROUNDTHECLOCK string = "ROUND_THE_CLOCK"

	// EmploymentDetailsWorkingHoursSchemeSHIFT365 captures enum value "SHIFT_365"
	EmploymentDetailsWorkingHoursSchemeSHIFT365 string = "SHIFT_365"

	// EmploymentDetailsWorkingHoursSchemeOFFSHORE336 captures enum value "OFFSHORE_336"
	EmploymentDetailsWorkingHoursSchemeOFFSHORE336 string = "OFFSHORE_336"

	// EmploymentDetailsWorkingHoursSchemeCONTINUOUS captures enum value "CONTINUOUS"
	EmploymentDetailsWorkingHoursSchemeCONTINUOUS string = "CONTINUOUS"

	// EmploymentDetailsWorkingHoursSchemeOTHERSHIFT captures enum value "OTHER_SHIFT"
	EmploymentDetailsWorkingHoursSchemeOTHERSHIFT string = "OTHER_SHIFT"
)

// prop value enum
func (m *EmploymentDetails) validateWorkingHoursSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, employmentDetailsTypeWorkingHoursSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmploymentDetails) validateWorkingHoursScheme(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkingHoursScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateWorkingHoursSchemeEnum("workingHoursScheme", "body", m.WorkingHoursScheme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmploymentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmploymentDetails) UnmarshalBinary(b []byte) error {
	var res EmploymentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
