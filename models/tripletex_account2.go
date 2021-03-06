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

// TripletexAccount2 tripletex account2
//
// swagger:model TripletexAccount2
type TripletexAccount2 struct {

	// Is this a test account or a paying account?
	// Required: true
	// Enum: [TEST PAYING]
	AccountType *string `json:"accountType"`

	// accounting office
	AccountingOffice *bool `json:"accountingOffice,omitempty"`

	// Employee to create. Department on this object will also be created if supplied. If null a dummy user and department will be created instead
	Administrator *Employee `json:"administrator,omitempty"`

	// Password for the administrator user to create. Not a part of the administrator employee object since this is a value that never can be read (it is salted and hashed before storing)
	// Required: true
	AdministratorPassword *string `json:"administratorPassword"`

	// auditor
	Auditor *bool `json:"auditor,omitempty"`

	// If true, the users created will be allowed to log in without validating their email address. ONLY USE THIS IF YOU ALREADY HAVE VALIDATED THE USER EMAILS.
	AutoValidateUserLogin *bool `json:"autoValidateUserLogin,omitempty"`

	// Main bank account
	BankAccount string `json:"bankAccount,omitempty"`

	// The chart of accounts to use for the new company
	// Enum: [DEFAULT MAMUT_STD_PAYROLL MAMUT_NARF_PAYROLL AGRO_FORRETNING_PAYROLL AGRO_LANDBRUK_PAYROLL AGRO_FISKE_PAYROLL AGRO_FORSOKSRING_PAYROLL AGRO_IDRETTSLAG_PAYROLL AGRO_FORENING_PAYROLL]
	ChartOfAccountsType string `json:"chartOfAccountsType,omitempty"`

	// Information about the company to create. Supply as much info as you have, but at least name, type and address.
	// Required: true
	Company *Company `json:"company"`

	// Create an API token for the administrator user for the consumer token used during this call. The token will be returned in the response.
	CreateAdministratorAPIToken *bool `json:"createAdministratorApiToken,omitempty"`

	// Create an API token for the company to use to call their clients, only possible for accounting and auditor accounts. The token will be returned in the response.
	CreateCompanyOwnedAPIToken *bool `json:"createCompanyOwnedApiToken,omitempty"`

	// Should the company we are creating be able to create new Tripletex accounts?
	MayCreateTripletexAccounts *bool `json:"mayCreateTripletexAccounts,omitempty"`

	// Sales modules (functionality in the application) to activate for the newly created account. Some modules have extra costs.
	// Required: true
	Modules []*SalesModuleDTO `json:"modules"`

	// Used to calculate prices.
	// Required: true
	// Enum: [INTERVAL_0_100 INTERVAL_101_500 INTERVAL_0_500 INTERVAL_501_1000 INTERVAL_1001_2000 INTERVAL_2001_3500 INTERVAL_3501_5000 INTERVAL_5001_10000 INTERVAL_UNLIMITED]
	NumberOfVouchers *string `json:"numberOfVouchers"`

	// Swedish post account number (PlusGirot)
	PostAccount string `json:"postAccount,omitempty"`

	// reseller
	Reseller *bool `json:"reseller,omitempty"`

	// Should the regular creation emails be sent to the company created and its users? If false you probably want to set autoValidateUserLogin to true
	SendEmails *bool `json:"sendEmails,omitempty"`

	// VAT type
	// Enum: [VAT_REGISTERED VAT_NOT_REGISTERED VAT_APPLICANT]
	VatStatusType string `json:"vatStatusType,omitempty"`
}

// Validate validates this tripletex account2
func (m *TripletexAccount2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministrator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministratorPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChartOfAccountsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfVouchers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatStatusType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tripletexAccount2TypeAccountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TEST","PAYING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tripletexAccount2TypeAccountTypePropEnum = append(tripletexAccount2TypeAccountTypePropEnum, v)
	}
}

const (

	// TripletexAccount2AccountTypeTEST captures enum value "TEST"
	TripletexAccount2AccountTypeTEST string = "TEST"

	// TripletexAccount2AccountTypePAYING captures enum value "PAYING"
	TripletexAccount2AccountTypePAYING string = "PAYING"
)

// prop value enum
func (m *TripletexAccount2) validateAccountTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tripletexAccount2TypeAccountTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TripletexAccount2) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("accountType", "body", m.AccountType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountTypeEnum("accountType", "body", *m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *TripletexAccount2) validateAdministrator(formats strfmt.Registry) error {

	if swag.IsZero(m.Administrator) { // not required
		return nil
	}

	if m.Administrator != nil {
		if err := m.Administrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("administrator")
			}
			return err
		}
	}

	return nil
}

func (m *TripletexAccount2) validateAdministratorPassword(formats strfmt.Registry) error {

	if err := validate.Required("administratorPassword", "body", m.AdministratorPassword); err != nil {
		return err
	}

	return nil
}

var tripletexAccount2TypeChartOfAccountsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","MAMUT_STD_PAYROLL","MAMUT_NARF_PAYROLL","AGRO_FORRETNING_PAYROLL","AGRO_LANDBRUK_PAYROLL","AGRO_FISKE_PAYROLL","AGRO_FORSOKSRING_PAYROLL","AGRO_IDRETTSLAG_PAYROLL","AGRO_FORENING_PAYROLL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tripletexAccount2TypeChartOfAccountsTypePropEnum = append(tripletexAccount2TypeChartOfAccountsTypePropEnum, v)
	}
}

const (

	// TripletexAccount2ChartOfAccountsTypeDEFAULT captures enum value "DEFAULT"
	TripletexAccount2ChartOfAccountsTypeDEFAULT string = "DEFAULT"

	// TripletexAccount2ChartOfAccountsTypeMAMUTSTDPAYROLL captures enum value "MAMUT_STD_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeMAMUTSTDPAYROLL string = "MAMUT_STD_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeMAMUTNARFPAYROLL captures enum value "MAMUT_NARF_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeMAMUTNARFPAYROLL string = "MAMUT_NARF_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROFORRETNINGPAYROLL captures enum value "AGRO_FORRETNING_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROFORRETNINGPAYROLL string = "AGRO_FORRETNING_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROLANDBRUKPAYROLL captures enum value "AGRO_LANDBRUK_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROLANDBRUKPAYROLL string = "AGRO_LANDBRUK_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROFISKEPAYROLL captures enum value "AGRO_FISKE_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROFISKEPAYROLL string = "AGRO_FISKE_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROFORSOKSRINGPAYROLL captures enum value "AGRO_FORSOKSRING_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROFORSOKSRINGPAYROLL string = "AGRO_FORSOKSRING_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROIDRETTSLAGPAYROLL captures enum value "AGRO_IDRETTSLAG_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROIDRETTSLAGPAYROLL string = "AGRO_IDRETTSLAG_PAYROLL"

	// TripletexAccount2ChartOfAccountsTypeAGROFORENINGPAYROLL captures enum value "AGRO_FORENING_PAYROLL"
	TripletexAccount2ChartOfAccountsTypeAGROFORENINGPAYROLL string = "AGRO_FORENING_PAYROLL"
)

// prop value enum
func (m *TripletexAccount2) validateChartOfAccountsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tripletexAccount2TypeChartOfAccountsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TripletexAccount2) validateChartOfAccountsType(formats strfmt.Registry) error {

	if swag.IsZero(m.ChartOfAccountsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateChartOfAccountsTypeEnum("chartOfAccountsType", "body", m.ChartOfAccountsType); err != nil {
		return err
	}

	return nil
}

func (m *TripletexAccount2) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("company", "body", m.Company); err != nil {
		return err
	}

	if m.Company != nil {
		if err := m.Company.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("company")
			}
			return err
		}
	}

	return nil
}

func (m *TripletexAccount2) validateModules(formats strfmt.Registry) error {

	if err := validate.Required("modules", "body", m.Modules); err != nil {
		return err
	}

	for i := 0; i < len(m.Modules); i++ {
		if swag.IsZero(m.Modules[i]) { // not required
			continue
		}

		if m.Modules[i] != nil {
			if err := m.Modules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var tripletexAccount2TypeNumberOfVouchersPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INTERVAL_0_100","INTERVAL_101_500","INTERVAL_0_500","INTERVAL_501_1000","INTERVAL_1001_2000","INTERVAL_2001_3500","INTERVAL_3501_5000","INTERVAL_5001_10000","INTERVAL_UNLIMITED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tripletexAccount2TypeNumberOfVouchersPropEnum = append(tripletexAccount2TypeNumberOfVouchersPropEnum, v)
	}
}

const (

	// TripletexAccount2NumberOfVouchersINTERVAL0100 captures enum value "INTERVAL_0_100"
	TripletexAccount2NumberOfVouchersINTERVAL0100 string = "INTERVAL_0_100"

	// TripletexAccount2NumberOfVouchersINTERVAL101500 captures enum value "INTERVAL_101_500"
	TripletexAccount2NumberOfVouchersINTERVAL101500 string = "INTERVAL_101_500"

	// TripletexAccount2NumberOfVouchersINTERVAL0500 captures enum value "INTERVAL_0_500"
	TripletexAccount2NumberOfVouchersINTERVAL0500 string = "INTERVAL_0_500"

	// TripletexAccount2NumberOfVouchersINTERVAL5011000 captures enum value "INTERVAL_501_1000"
	TripletexAccount2NumberOfVouchersINTERVAL5011000 string = "INTERVAL_501_1000"

	// TripletexAccount2NumberOfVouchersINTERVAL10012000 captures enum value "INTERVAL_1001_2000"
	TripletexAccount2NumberOfVouchersINTERVAL10012000 string = "INTERVAL_1001_2000"

	// TripletexAccount2NumberOfVouchersINTERVAL20013500 captures enum value "INTERVAL_2001_3500"
	TripletexAccount2NumberOfVouchersINTERVAL20013500 string = "INTERVAL_2001_3500"

	// TripletexAccount2NumberOfVouchersINTERVAL35015000 captures enum value "INTERVAL_3501_5000"
	TripletexAccount2NumberOfVouchersINTERVAL35015000 string = "INTERVAL_3501_5000"

	// TripletexAccount2NumberOfVouchersINTERVAL500110000 captures enum value "INTERVAL_5001_10000"
	TripletexAccount2NumberOfVouchersINTERVAL500110000 string = "INTERVAL_5001_10000"

	// TripletexAccount2NumberOfVouchersINTERVALUNLIMITED captures enum value "INTERVAL_UNLIMITED"
	TripletexAccount2NumberOfVouchersINTERVALUNLIMITED string = "INTERVAL_UNLIMITED"
)

// prop value enum
func (m *TripletexAccount2) validateNumberOfVouchersEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tripletexAccount2TypeNumberOfVouchersPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TripletexAccount2) validateNumberOfVouchers(formats strfmt.Registry) error {

	if err := validate.Required("numberOfVouchers", "body", m.NumberOfVouchers); err != nil {
		return err
	}

	// value enum
	if err := m.validateNumberOfVouchersEnum("numberOfVouchers", "body", *m.NumberOfVouchers); err != nil {
		return err
	}

	return nil
}

var tripletexAccount2TypeVatStatusTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VAT_REGISTERED","VAT_NOT_REGISTERED","VAT_APPLICANT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tripletexAccount2TypeVatStatusTypePropEnum = append(tripletexAccount2TypeVatStatusTypePropEnum, v)
	}
}

const (

	// TripletexAccount2VatStatusTypeVATREGISTERED captures enum value "VAT_REGISTERED"
	TripletexAccount2VatStatusTypeVATREGISTERED string = "VAT_REGISTERED"

	// TripletexAccount2VatStatusTypeVATNOTREGISTERED captures enum value "VAT_NOT_REGISTERED"
	TripletexAccount2VatStatusTypeVATNOTREGISTERED string = "VAT_NOT_REGISTERED"

	// TripletexAccount2VatStatusTypeVATAPPLICANT captures enum value "VAT_APPLICANT"
	TripletexAccount2VatStatusTypeVATAPPLICANT string = "VAT_APPLICANT"
)

// prop value enum
func (m *TripletexAccount2) validateVatStatusTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tripletexAccount2TypeVatStatusTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TripletexAccount2) validateVatStatusType(formats strfmt.Registry) error {

	if swag.IsZero(m.VatStatusType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVatStatusTypeEnum("vatStatusType", "body", m.VatStatusType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TripletexAccount2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripletexAccount2) UnmarshalBinary(b []byte) error {
	var res TripletexAccount2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
