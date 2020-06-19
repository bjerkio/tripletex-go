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

// Employment employment
//
// swagger:model Employment
type Employment struct {

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// division
	Division *Division `json:"division,omitempty"`

	// employee
	Employee *Employee `json:"employee,omitempty"`

	// Employment types tied to the employment
	EmploymentDetails []*EmploymentDetails `json:"employmentDetails"`

	// Existing employment ID used by the current accounting system
	// Max Length: 255
	EmploymentID string `json:"employmentId,omitempty"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// Determines if company is main employer for the employee. Default value is true.<br />Some values will be default set if not sent upon creation of employment: <br/> If isMainEmployer is NOT sent and tax deduction code loennFraHovedarbeidsgiver is sent, isMainEmployer will be set to true. <br /> If isMainEmployer is NOT sent and tax deduction code loennFraBiarbeidsgiver is sent, isMainEmployer will be set to false. <br /> If true and deduction code is NOT sent, value of tax deduction code will be set to loennFraHovedarbeidsgiver. <br /> If false and deduction code is NOT sent, value of tax deduction code will be set to loennFraBiarbeidsgiver. <br /> For other types of Tax Deduction Codes, isMainEmployer does not influence anything.
	IsMainEmployer *bool `json:"isMainEmployer,omitempty"`

	// last salary change date
	LastSalaryChangeDate string `json:"lastSalaryChangeDate,omitempty"`

	// Activate pensions and other benefits with no employment relationship.
	NoEmploymentRelationship *bool `json:"noEmploymentRelationship,omitempty"`

	// start date
	// Required: true
	StartDate *string `json:"startDate"`

	// EMPTY - represents that a tax deduction code is not set on the employment. It is illegal to set the field to this value.  <br /> Default value of this field is loennFraHovedarbeidsgiver or loennFraBiarbeidsgiver depending on boolean isMainEmployer
	// Enum: [loennFraHovedarbeidsgiver loennFraBiarbeidsgiver pensjon loennTilUtenrikstjenestemann loennKunTrygdeavgiftTilUtenlandskBorger loennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger introduksjonsstoenad ufoereytelserFraAndre EMPTY]
	TaxDeductionCode string `json:"taxDeductionCode,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this employment
func (m *Employment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmploymentDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxDeductionCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Employment) validateChanges(formats strfmt.Registry) error {

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

func (m *Employment) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Employment) validateEmployee(formats strfmt.Registry) error {

	if swag.IsZero(m.Employee) { // not required
		return nil
	}

	if m.Employee != nil {
		if err := m.Employee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee")
			}
			return err
		}
	}

	return nil
}

func (m *Employment) validateEmploymentDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.EmploymentDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.EmploymentDetails); i++ {
		if swag.IsZero(m.EmploymentDetails[i]) { // not required
			continue
		}

		if m.EmploymentDetails[i] != nil {
			if err := m.EmploymentDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("employmentDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Employment) validateEmploymentID(formats strfmt.Registry) error {

	if swag.IsZero(m.EmploymentID) { // not required
		return nil
	}

	if err := validate.MaxLength("employmentId", "body", string(m.EmploymentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Employment) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	return nil
}

var employmentTypeTaxDeductionCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loennFraHovedarbeidsgiver","loennFraBiarbeidsgiver","pensjon","loennTilUtenrikstjenestemann","loennKunTrygdeavgiftTilUtenlandskBorger","loennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger","introduksjonsstoenad","ufoereytelserFraAndre","EMPTY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		employmentTypeTaxDeductionCodePropEnum = append(employmentTypeTaxDeductionCodePropEnum, v)
	}
}

const (

	// EmploymentTaxDeductionCodeLoennFraHovedarbeidsgiver captures enum value "loennFraHovedarbeidsgiver"
	EmploymentTaxDeductionCodeLoennFraHovedarbeidsgiver string = "loennFraHovedarbeidsgiver"

	// EmploymentTaxDeductionCodeLoennFraBiarbeidsgiver captures enum value "loennFraBiarbeidsgiver"
	EmploymentTaxDeductionCodeLoennFraBiarbeidsgiver string = "loennFraBiarbeidsgiver"

	// EmploymentTaxDeductionCodePensjon captures enum value "pensjon"
	EmploymentTaxDeductionCodePensjon string = "pensjon"

	// EmploymentTaxDeductionCodeLoennTilUtenrikstjenestemann captures enum value "loennTilUtenrikstjenestemann"
	EmploymentTaxDeductionCodeLoennTilUtenrikstjenestemann string = "loennTilUtenrikstjenestemann"

	// EmploymentTaxDeductionCodeLoennKunTrygdeavgiftTilUtenlandskBorger captures enum value "loennKunTrygdeavgiftTilUtenlandskBorger"
	EmploymentTaxDeductionCodeLoennKunTrygdeavgiftTilUtenlandskBorger string = "loennKunTrygdeavgiftTilUtenlandskBorger"

	// EmploymentTaxDeductionCodeLoennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger captures enum value "loennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger"
	EmploymentTaxDeductionCodeLoennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger string = "loennKunTrygdeavgiftTilUtenlandskBorgerSomGrensegjenger"

	// EmploymentTaxDeductionCodeIntroduksjonsstoenad captures enum value "introduksjonsstoenad"
	EmploymentTaxDeductionCodeIntroduksjonsstoenad string = "introduksjonsstoenad"

	// EmploymentTaxDeductionCodeUfoereytelserFraAndre captures enum value "ufoereytelserFraAndre"
	EmploymentTaxDeductionCodeUfoereytelserFraAndre string = "ufoereytelserFraAndre"

	// EmploymentTaxDeductionCodeEMPTY captures enum value "EMPTY"
	EmploymentTaxDeductionCodeEMPTY string = "EMPTY"
)

// prop value enum
func (m *Employment) validateTaxDeductionCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, employmentTypeTaxDeductionCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Employment) validateTaxDeductionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxDeductionCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaxDeductionCodeEnum("taxDeductionCode", "body", m.TaxDeductionCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Employment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Employment) UnmarshalBinary(b []byte) error {
	var res Employment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
