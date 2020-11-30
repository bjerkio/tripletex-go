// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TripletexCompanyModules tripletex company modules
//
// swagger:model TripletexCompanyModules
type TripletexCompanyModules struct {

	// approvehourlists
	// Read Only: true
	Approvehourlists *bool `json:"approvehourlists,omitempty"`

	// approveinvoices
	// Read Only: true
	Approveinvoices *bool `json:"approveinvoices,omitempty"`

	// approvemonthlyhourlists
	// Read Only: true
	Approvemonthlyhourlists *bool `json:"approvemonthlyhourlists,omitempty"`

	// approvetravelreports
	// Read Only: true
	Approvetravelreports *bool `json:"approvetravelreports,omitempty"`

	// auto invoicing
	// Read Only: true
	AutoInvoicing *bool `json:"autoInvoicing,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// company Id
	// Read Only: true
	CompanyID int32 `json:"companyId,omitempty"`

	// completemonthlyhourlists
	// Read Only: true
	Completemonthlyhourlists *bool `json:"completemonthlyhourlists,omitempty"`

	// completeweeklyhourlists
	// Read Only: true
	Completeweeklyhourlists *bool `json:"completeweeklyhourlists,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// module accountant connect client
	// Read Only: true
	ModuleAccountantConnectClient *bool `json:"moduleAccountantConnectClient,omitempty"`

	// module accounting external
	// Read Only: true
	ModuleAccountingExternal *bool `json:"moduleAccountingExternal,omitempty"`

	// module accounting reports
	// Read Only: true
	ModuleAccountingReports *bool `json:"moduleAccountingReports,omitempty"`

	// module activity hourly wage wage code
	// Read Only: true
	ModuleActivityHourlyWageWageCode *bool `json:"moduleActivityHourlyWageWageCode,omitempty"`

	// module addition project markup
	// Read Only: true
	ModuleAdditionProjectMarkup *bool `json:"moduleAdditionProjectMarkup,omitempty"`

	// module agro
	// Read Only: true
	ModuleAgro *bool `json:"moduleAgro,omitempty"`

	// module ahlsell partner
	// Read Only: true
	ModuleAhlsellPartner *bool `json:"moduleAhlsellPartner,omitempty"`

	// module amortization
	// Read Only: true
	ModuleAmortization *bool `json:"moduleAmortization,omitempty"`

	// module api20
	// Read Only: true
	ModuleApi20 *bool `json:"moduleApi20,omitempty"`

	// module approve department voucher
	// Read Only: true
	ModuleApproveDepartmentVoucher *bool `json:"moduleApproveDepartmentVoucher,omitempty"`

	// module approve project voucher
	// Read Only: true
	ModuleApproveProjectVoucher *bool `json:"moduleApproveProjectVoucher,omitempty"`

	// module approve voucher
	// Read Only: true
	ModuleApproveVoucher *bool `json:"moduleApproveVoucher,omitempty"`

	// module archive
	// Read Only: true
	ModuleArchive *bool `json:"moduleArchive,omitempty"`

	// module auto bank reconciliation
	// Read Only: true
	ModuleAutoBankReconciliation *bool `json:"moduleAutoBankReconciliation,omitempty"`

	// module auto customer number
	// Read Only: true
	ModuleAutoCustomerNumber *bool `json:"moduleAutoCustomerNumber,omitempty"`

	// module auto project number
	// Read Only: true
	ModuleAutoProjectNumber *bool `json:"moduleAutoProjectNumber,omitempty"`

	// module auto vendor number
	// Read Only: true
	ModuleAutoVendorNumber *bool `json:"moduleAutoVendorNumber,omitempty"`

	// module boligmappa
	// Read Only: true
	ModuleBoligmappa *bool `json:"moduleBoligmappa,omitempty"`

	// module c r m
	// Read Only: true
	ModuleCRM *bool `json:"moduleCRM,omitempty"`

	// module cash credit aprila
	// Read Only: true
	ModuleCashCreditAprila *bool `json:"moduleCashCreditAprila,omitempty"`

	// module change debt collector
	// Read Only: true
	ModuleChangeDebtCollector *bool `json:"moduleChangeDebtCollector,omitempty"`

	// module contact
	// Read Only: true
	ModuleContact *bool `json:"moduleContact,omitempty"`

	// module control schema required hour tracking
	// Read Only: true
	ModuleControlSchemaRequiredHourTracking *bool `json:"moduleControlSchemaRequiredHourTracking,omitempty"`

	// module control schema required invoicing
	// Read Only: true
	ModuleControlSchemaRequiredInvoicing *bool `json:"moduleControlSchemaRequiredInvoicing,omitempty"`

	// module currency
	// Read Only: true
	ModuleCurrency *bool `json:"moduleCurrency,omitempty"`

	// module customer categories
	// Read Only: true
	ModuleCustomerCategories *bool `json:"moduleCustomerCategories,omitempty"`

	// module customer category1
	// Read Only: true
	ModuleCustomerCategory1 *bool `json:"moduleCustomerCategory1,omitempty"`

	// module customer category2
	// Read Only: true
	ModuleCustomerCategory2 *bool `json:"moduleCustomerCategory2,omitempty"`

	// module customer category3
	// Read Only: true
	ModuleCustomerCategory3 *bool `json:"moduleCustomerCategory3,omitempty"`

	// module department accounting
	// Read Only: true
	ModuleDepartmentAccounting *bool `json:"moduleDepartmentAccounting,omitempty"`

	// module divisions
	// Read Only: true
	ModuleDivisions *bool `json:"moduleDivisions,omitempty"`

	// module ehf
	// Read Only: true
	ModuleEhf *bool `json:"moduleEhf,omitempty"`

	// module electro
	// Read Only: true
	ModuleElectro *bool `json:"moduleElectro,omitempty"`

	// module elektro union
	// Read Only: true
	ModuleElektroUnion *bool `json:"moduleElektroUnion,omitempty"`

	// module elproffen
	// Read Only: true
	ModuleElproffen *bool `json:"moduleElproffen,omitempty"`

	// module email
	// Read Only: true
	ModuleEmail *bool `json:"moduleEmail,omitempty"`

	// module employee accounting
	// Read Only: true
	ModuleEmployeeAccounting *bool `json:"moduleEmployeeAccounting,omitempty"`

	// module employee category
	// Read Only: true
	ModuleEmployeeCategory *bool `json:"moduleEmployeeCategory,omitempty"`

	// module encrypted pay slip
	// Read Only: true
	ModuleEncryptedPaySlip *bool `json:"moduleEncryptedPaySlip,omitempty"`

	// module factoring aprila
	// Read Only: true
	ModuleFactoringAprila *bool `json:"moduleFactoringAprila,omitempty"`

	// module finance tax
	// Read Only: true
	ModuleFinanceTax *bool `json:"moduleFinanceTax,omitempty"`

	// module gtin
	// Read Only: true
	ModuleGtin *bool `json:"moduleGtin,omitempty"`

	// module holyday plan
	// Read Only: true
	ModuleHolydayPlan *bool `json:"moduleHolydayPlan,omitempty"`

	// module invoice
	// Read Only: true
	ModuleInvoice *bool `json:"moduleInvoice,omitempty"`

	// module invoice fee comment
	// Read Only: true
	ModuleInvoiceFeeComment *bool `json:"moduleInvoiceFeeComment,omitempty"`

	// module invoice import
	// Read Only: true
	ModuleInvoiceImport *bool `json:"moduleInvoiceImport,omitempty"`

	// module invoice option autoinvoice ehf
	// Read Only: true
	ModuleInvoiceOptionAutoinvoiceEhf *bool `json:"moduleInvoiceOptionAutoinvoiceEhf,omitempty"`

	// module invoice option avtale giro
	// Read Only: true
	ModuleInvoiceOptionAvtaleGiro *bool `json:"moduleInvoiceOptionAvtaleGiro,omitempty"`

	// module invoice option efaktura
	// Read Only: true
	ModuleInvoiceOptionEfaktura *bool `json:"moduleInvoiceOptionEfaktura,omitempty"`

	// module invoice option paper
	// Read Only: true
	ModuleInvoiceOptionPaper *bool `json:"moduleInvoiceOptionPaper,omitempty"`

	// module invoice option vipps
	// Read Only: true
	ModuleInvoiceOptionVipps *bool `json:"moduleInvoiceOptionVipps,omitempty"`

	// module invoice scanning
	// Read Only: true
	ModuleInvoiceScanning *bool `json:"moduleInvoiceScanning,omitempty"`

	// module mamut
	// Read Only: true
	ModuleMamut *bool `json:"moduleMamut,omitempty"`

	// module mesan
	// Read Only: true
	ModuleMesan *bool `json:"moduleMesan,omitempty"`

	// module nets eboks
	// Read Only: true
	ModuleNetsEboks *bool `json:"moduleNetsEboks,omitempty"`

	// module nets print invoice
	// Read Only: true
	ModuleNetsPrintInvoice *bool `json:"moduleNetsPrintInvoice,omitempty"`

	// module nets print salary
	// Read Only: true
	ModuleNetsPrintSalary *bool `json:"moduleNetsPrintSalary,omitempty"`

	// module nrf
	// Read Only: true
	ModuleNrf *bool `json:"moduleNrf,omitempty"`

	// module ocr
	// Read Only: true
	ModuleOcr *bool `json:"moduleOcr,omitempty"`

	// module ocr auto pay
	// Read Only: true
	ModuleOcrAutoPay *bool `json:"moduleOcrAutoPay,omitempty"`

	// module offer
	// Read Only: true
	ModuleOffer *bool `json:"moduleOffer,omitempty"`

	// module onninen123
	// Read Only: true
	ModuleOnninen123 *bool `json:"moduleOnninen123,omitempty"`

	// module order discount
	// Read Only: true
	ModuleOrderDiscount *bool `json:"moduleOrderDiscount,omitempty"`

	// module order ext
	// Read Only: true
	ModuleOrderExt *bool `json:"moduleOrderExt,omitempty"`

	// module order line cost
	// Read Only: true
	ModuleOrderLineCost *bool `json:"moduleOrderLineCost,omitempty"`

	// module order markup
	// Read Only: true
	ModuleOrderMarkup *bool `json:"moduleOrderMarkup,omitempty"`

	// module order number
	// Read Only: true
	ModuleOrderNumber *bool `json:"moduleOrderNumber,omitempty"`

	// module order out
	// Read Only: true
	ModuleOrderOut *bool `json:"moduleOrderOut,omitempty"`

	// module payroll accounting
	// Read Only: true
	ModulePayrollAccounting *bool `json:"modulePayrollAccounting,omitempty"`

	// module pensionreport
	// Read Only: true
	ModulePensionreport *bool `json:"modulePensionreport,omitempty"`

	// module product accounting
	// Read Only: true
	ModuleProductAccounting *bool `json:"moduleProductAccounting,omitempty"`

	// module product invoice
	// Read Only: true
	ModuleProductInvoice *bool `json:"moduleProductInvoice,omitempty"`

	// module project accounting
	// Read Only: true
	ModuleProjectAccounting *bool `json:"moduleProjectAccounting,omitempty"`

	// module project budget
	// Read Only: true
	ModuleProjectBudget *bool `json:"moduleProjectBudget,omitempty"`

	// module project budget reference fee
	// Read Only: true
	ModuleProjectBudgetReferenceFee *bool `json:"moduleProjectBudgetReferenceFee,omitempty"`

	// module project participants
	// Read Only: true
	ModuleProjectParticipants *bool `json:"moduleProjectParticipants,omitempty"`

	// module provision salary
	// Read Only: true
	ModuleProvisionSalary *bool `json:"moduleProvisionSalary,omitempty"`

	// module resource groups
	// Read Only: true
	ModuleResourceGroups *bool `json:"moduleResourceGroups,omitempty"`

	// module result budget
	// Read Only: true
	ModuleResultBudget *bool `json:"moduleResultBudget,omitempty"`

	// module rorkjop
	// Read Only: true
	ModuleRorkjop *bool `json:"moduleRorkjop,omitempty"`

	// module smart scan
	// Read Only: true
	ModuleSmartScan *bool `json:"moduleSmartScan,omitempty"`

	// module stop watch
	// Read Only: true
	ModuleStopWatch *bool `json:"moduleStopWatch,omitempty"`

	// module subscription address list
	// Read Only: true
	ModuleSubscriptionAddressList *bool `json:"moduleSubscriptionAddressList,omitempty"`

	// module subscriptions periodisation
	// Read Only: true
	ModuleSubscriptionsPeriodisation *bool `json:"moduleSubscriptionsPeriodisation,omitempty"`

	// module swedish
	// Read Only: true
	ModuleSwedish *bool `json:"moduleSwedish,omitempty"`

	// module time balance
	// Read Only: true
	ModuleTimeBalance *bool `json:"moduleTimeBalance,omitempty"`

	// module travel expense
	// Read Only: true
	ModuleTravelExpense *bool `json:"moduleTravelExpense,omitempty"`

	// module travel expense rates
	// Read Only: true
	ModuleTravelExpenseRates *bool `json:"moduleTravelExpenseRates,omitempty"`

	// module vacation balance
	// Read Only: true
	ModuleVacationBalance *bool `json:"moduleVacationBalance,omitempty"`

	// module voucher automation
	// Read Only: true
	ModuleVoucherAutomation *bool `json:"moduleVoucherAutomation,omitempty"`

	// module voucher scanning
	// Read Only: true
	ModuleVoucherScanning *bool `json:"moduleVoucherScanning,omitempty"`

	// module voucher types
	// Read Only: true
	ModuleVoucherTypes *bool `json:"moduleVoucherTypes,omitempty"`

	// module wage amortization
	// Read Only: true
	ModuleWageAmortization *bool `json:"moduleWageAmortization,omitempty"`

	// module wage export
	// Read Only: true
	ModuleWageExport *bool `json:"moduleWageExport,omitempty"`

	// module wage project accounting
	// Read Only: true
	ModuleWageProjectAccounting *bool `json:"moduleWageProjectAccounting,omitempty"`

	// module warehouse
	// Read Only: true
	ModuleWarehouse *bool `json:"moduleWarehouse,omitempty"`

	// module working hours
	// Read Only: true
	ModuleWorkingHours *bool `json:"moduleWorkingHours,omitempty"`

	// moduleaccountinginternal
	// Read Only: true
	Moduleaccountinginternal *bool `json:"moduleaccountinginternal,omitempty"`

	// modulebudget
	// Read Only: true
	Modulebudget *bool `json:"modulebudget,omitempty"`

	// modulebunches
	// Read Only: true
	Modulebunches *bool `json:"modulebunches,omitempty"`

	// modulecustomer
	// Read Only: true
	Modulecustomer *bool `json:"modulecustomer,omitempty"`

	// moduledepartment
	// Read Only: true
	Moduledepartment *bool `json:"moduledepartment,omitempty"`

	// moduleemployee
	// Read Only: true
	Moduleemployee *bool `json:"moduleemployee,omitempty"`

	// modulehistorical
	// Read Only: true
	Modulehistorical *bool `json:"modulehistorical,omitempty"`

	// modulehourlist
	// Read Only: true
	Modulehourlist *bool `json:"modulehourlist,omitempty"`

	// modulenote
	// Read Only: true
	Modulenote *bool `json:"modulenote,omitempty"`

	// moduleproduct
	// Read Only: true
	Moduleproduct *bool `json:"moduleproduct,omitempty"`

	// moduleproject
	// Read Only: true
	Moduleproject *bool `json:"moduleproject,omitempty"`

	// moduleprojectcategory
	// Read Only: true
	Moduleprojectcategory *bool `json:"moduleprojectcategory,omitempty"`

	// moduleprojecteconomy
	// Read Only: true
	Moduleprojecteconomy *bool `json:"moduleprojecteconomy,omitempty"`

	// moduleprojectlocation
	// Read Only: true
	Moduleprojectlocation *bool `json:"moduleprojectlocation,omitempty"`

	// moduleprojectprognosis
	// Read Only: true
	Moduleprojectprognosis *bool `json:"moduleprojectprognosis,omitempty"`

	// moduleprojectsubcontract
	// Read Only: true
	Moduleprojectsubcontract *bool `json:"moduleprojectsubcontract,omitempty"`

	// modulereferencefee
	// Read Only: true
	Modulereferencefee *bool `json:"modulereferencefee,omitempty"`

	// moduleresourceallocation
	// Read Only: true
	Moduleresourceallocation *bool `json:"moduleresourceallocation,omitempty"`

	// modulesubscription
	// Read Only: true
	Modulesubscription *bool `json:"modulesubscription,omitempty"`

	// moduletask
	// Read Only: true
	Moduletask *bool `json:"moduletask,omitempty"`

	// monthly hourlist minus time warning
	// Read Only: true
	MonthlyHourlistMinusTimeWarning *bool `json:"monthlyHourlistMinusTimeWarning,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this tripletex company modules
func (m *TripletexCompanyModules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TripletexCompanyModules) validateChanges(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *TripletexCompanyModules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripletexCompanyModules) UnmarshalBinary(b []byte) error {
	var res TripletexCompanyModules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}