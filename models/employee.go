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

// Employee employee
//
// swagger:model Employee
type Employee struct {

	// Address tied to the employee
	Address *Address `json:"address,omitempty"`

	// Determines if salary information can be registered on the user including hours, travel expenses and employee expenses. The user may also be selected as a project member on projects.
	// Read Only: true
	AllowInformationRegistration *bool `json:"allowInformationRegistration,omitempty"`

	// bank account number
	// Max Length: 100
	BankAccountNumber string `json:"bankAccountNumber,omitempty"`

	// Bic (swift) field -- pilot program
	Bic string `json:"bic,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// comments
	Comments string `json:"comments,omitempty"`

	// Country of creditor bank field -- pilot program
	CreditorBankCountryID int32 `json:"creditorBankCountryId,omitempty"`

	// date of birth
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// department
	Department *Department `json:"department,omitempty"`

	// dnumber
	// Max Length: 11
	Dnumber string `json:"dnumber,omitempty"`

	// email
	// Max Length: 100
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// employee category
	EmployeeCategory *EmployeeCategory `json:"employeeCategory,omitempty"`

	// employee number
	// Max Length: 100
	// Min Length: 0
	EmployeeNumber *string `json:"employeeNumber,omitempty"`

	// Employments tied to the employee
	Employments []*Employment `json:"employments"`

	// first name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	FirstName *string `json:"firstName"`

	// holiday allowance earned
	HolidayAllowanceEarned *HolidayAllowanceEarned `json:"holidayAllowanceEarned,omitempty"`

	// IBAN field -- pilot program
	Iban string `json:"iban,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// international Id
	InternationalID *InternationalID `json:"internationalId,omitempty"`

	// is contact
	// Read Only: true
	IsContact *bool `json:"isContact,omitempty"`

	// last name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	LastName *string `json:"lastName"`

	// national identity number
	// Max Length: 100
	NationalIdentityNumber string `json:"nationalIdentityNumber,omitempty"`

	// phone number home
	// Max Length: 100
	PhoneNumberHome string `json:"phoneNumberHome,omitempty"`

	// phone number mobile
	// Max Length: 100
	PhoneNumberMobile string `json:"phoneNumberMobile,omitempty"`

	// The country of the mobile phone number. If not set, the country is derived as best as possible from phoneNumberMobile. NB! 8 digit numbers are assumed to be Norwegian.
	PhoneNumberMobileCountry *Country `json:"phoneNumberMobileCountry,omitempty"`

	// phone number work
	// Max Length: 100
	PhoneNumberWork string `json:"phoneNumberWork,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// Define the employee's user type.<br>STANDARD: Reduced access. Users with limited system entitlements.<br>EXTENDED: Users can be given all system entitlements.<br>NO_ACCESS: User with no log on access.<br>Users with access to Tripletex must confirm the email address.
	// Enum: [STANDARD EXTENDED NO_ACCESS]
	UserType string `json:"userType,omitempty"`

	// UsesAbroadPayment field -- pilot program. Determines if we should use domestic or abroad remittance. To be able to use abroad remittance, one has to: 1: have Autopay 2: have valid combination of the fields Iban, Bic (swift) and Country of creditor bank.
	UsesAbroadPayment *bool `json:"usesAbroadPayment,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this employee
func (m *Employee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDnumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployeeCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployeeNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHolidayAllowanceEarned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternationalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalIdentityNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberHome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberMobileCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberWork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Employee) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validateBankAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccountNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("bankAccountNumber", "body", string(m.BankAccountNumber), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateChanges(formats strfmt.Registry) error {

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

func (m *Employee) validateDepartment(formats strfmt.Registry) error {

	if swag.IsZero(m.Department) { // not required
		return nil
	}

	if m.Department != nil {
		if err := m.Department.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("department")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validateDnumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Dnumber) { // not required
		return nil
	}

	if err := validate.MaxLength("dnumber", "body", string(m.Dnumber), 11); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MaxLength("email", "body", string(m.Email), 100); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateEmployeeCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeCategory) { // not required
		return nil
	}

	if m.EmployeeCategory != nil {
		if err := m.EmployeeCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employeeCategory")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validateEmployeeNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeNumber) { // not required
		return nil
	}

	if err := validate.MinLength("employeeNumber", "body", string(*m.EmployeeNumber), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("employeeNumber", "body", string(*m.EmployeeNumber), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateEmployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Employments) { // not required
		return nil
	}

	for i := 0; i < len(m.Employments); i++ {
		if swag.IsZero(m.Employments[i]) { // not required
			continue
		}

		if m.Employments[i] != nil {
			if err := m.Employments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("employments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Employee) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("firstName", "body", string(*m.FirstName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("firstName", "body", string(*m.FirstName), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateHolidayAllowanceEarned(formats strfmt.Registry) error {

	if swag.IsZero(m.HolidayAllowanceEarned) { // not required
		return nil
	}

	if m.HolidayAllowanceEarned != nil {
		if err := m.HolidayAllowanceEarned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("holidayAllowanceEarned")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validateInternationalID(formats strfmt.Registry) error {

	if swag.IsZero(m.InternationalID) { // not required
		return nil
	}

	if m.InternationalID != nil {
		if err := m.InternationalID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internationalId")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MinLength("lastName", "body", string(*m.LastName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("lastName", "body", string(*m.LastName), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validateNationalIdentityNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.NationalIdentityNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("nationalIdentityNumber", "body", string(m.NationalIdentityNumber), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validatePhoneNumberHome(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberHome) { // not required
		return nil
	}

	if err := validate.MaxLength("phoneNumberHome", "body", string(m.PhoneNumberHome), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validatePhoneNumberMobile(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberMobile) { // not required
		return nil
	}

	if err := validate.MaxLength("phoneNumberMobile", "body", string(m.PhoneNumberMobile), 100); err != nil {
		return err
	}

	return nil
}

func (m *Employee) validatePhoneNumberMobileCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberMobileCountry) { // not required
		return nil
	}

	if m.PhoneNumberMobileCountry != nil {
		if err := m.PhoneNumberMobileCountry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumberMobileCountry")
			}
			return err
		}
	}

	return nil
}

func (m *Employee) validatePhoneNumberWork(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberWork) { // not required
		return nil
	}

	if err := validate.MaxLength("phoneNumberWork", "body", string(m.PhoneNumberWork), 100); err != nil {
		return err
	}

	return nil
}

var employeeTypeUserTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STANDARD","EXTENDED","NO_ACCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		employeeTypeUserTypePropEnum = append(employeeTypeUserTypePropEnum, v)
	}
}

const (

	// EmployeeUserTypeSTANDARD captures enum value "STANDARD"
	EmployeeUserTypeSTANDARD string = "STANDARD"

	// EmployeeUserTypeEXTENDED captures enum value "EXTENDED"
	EmployeeUserTypeEXTENDED string = "EXTENDED"

	// EmployeeUserTypeNOACCESS captures enum value "NO_ACCESS"
	EmployeeUserTypeNOACCESS string = "NO_ACCESS"
)

// prop value enum
func (m *Employee) validateUserTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, employeeTypeUserTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Employee) validateUserType(formats strfmt.Registry) error {

	if swag.IsZero(m.UserType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserTypeEnum("userType", "body", m.UserType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Employee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Employee) UnmarshalBinary(b []byte) error {
	var res Employee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
