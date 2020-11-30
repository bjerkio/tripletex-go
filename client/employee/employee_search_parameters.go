// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewEmployeeSearchParams creates a new EmployeeSearchParams object
// with the default values initialized.
func NewEmployeeSearchParams() *EmployeeSearchParams {
	var (
		countDefault           = int64(1000)
		fromDefault            = int64(0)
		includeContactsDefault = bool(false)
	)
	return &EmployeeSearchParams{
		Count:           &countDefault,
		From:            &fromDefault,
		IncludeContacts: &includeContactsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeSearchParamsWithTimeout creates a new EmployeeSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeSearchParamsWithTimeout(timeout time.Duration) *EmployeeSearchParams {
	var (
		countDefault           = int64(1000)
		fromDefault            = int64(0)
		includeContactsDefault = bool(false)
	)
	return &EmployeeSearchParams{
		Count:           &countDefault,
		From:            &fromDefault,
		IncludeContacts: &includeContactsDefault,

		timeout: timeout,
	}
}

// NewEmployeeSearchParamsWithContext creates a new EmployeeSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeSearchParamsWithContext(ctx context.Context) *EmployeeSearchParams {
	var (
		countDefault           = int64(1000)
		fromDefault            = int64(0)
		includeContactsDefault = bool(false)
	)
	return &EmployeeSearchParams{
		Count:           &countDefault,
		From:            &fromDefault,
		IncludeContacts: &includeContactsDefault,

		Context: ctx,
	}
}

// NewEmployeeSearchParamsWithHTTPClient creates a new EmployeeSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeSearchParamsWithHTTPClient(client *http.Client) *EmployeeSearchParams {
	var (
		countDefault           = int64(1000)
		fromDefault            = int64(0)
		includeContactsDefault = bool(false)
	)
	return &EmployeeSearchParams{
		Count:           &countDefault,
		From:            &fromDefault,
		IncludeContacts: &includeContactsDefault,
		HTTPClient:      client,
	}
}

/*EmployeeSearchParams contains all the parameters to send to the API endpoint
for the employee search operation typically these are written to a http.Request
*/
type EmployeeSearchParams struct {

	/*AllowInformationRegistration
	  Equals

	*/
	AllowInformationRegistration *bool
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*DepartmentID
	  List of IDs

	*/
	DepartmentID *string
	/*EmployeeNumber
	  Containing

	*/
	EmployeeNumber *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*FirstName
	  Containing

	*/
	FirstName *string
	/*From
	  From index

	*/
	From *int64
	/*HasSystemAccess
	  Equals

	*/
	HasSystemAccess *bool
	/*ID
	  List of IDs

	*/
	ID *string
	/*IncludeContacts
	  Equals

	*/
	IncludeContacts *bool
	/*LastName
	  Containing

	*/
	LastName *string
	/*OnlyProjectManagers
	  Equals

	*/
	OnlyProjectManagers *bool
	/*PeriodEnd
	  Equals

	*/
	PeriodEnd *string
	/*PeriodStart
	  Equals

	*/
	PeriodStart *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee search params
func (o *EmployeeSearchParams) WithTimeout(timeout time.Duration) *EmployeeSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee search params
func (o *EmployeeSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee search params
func (o *EmployeeSearchParams) WithContext(ctx context.Context) *EmployeeSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee search params
func (o *EmployeeSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee search params
func (o *EmployeeSearchParams) WithHTTPClient(client *http.Client) *EmployeeSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee search params
func (o *EmployeeSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowInformationRegistration adds the allowInformationRegistration to the employee search params
func (o *EmployeeSearchParams) WithAllowInformationRegistration(allowInformationRegistration *bool) *EmployeeSearchParams {
	o.SetAllowInformationRegistration(allowInformationRegistration)
	return o
}

// SetAllowInformationRegistration adds the allowInformationRegistration to the employee search params
func (o *EmployeeSearchParams) SetAllowInformationRegistration(allowInformationRegistration *bool) {
	o.AllowInformationRegistration = allowInformationRegistration
}

// WithCount adds the count to the employee search params
func (o *EmployeeSearchParams) WithCount(count *int64) *EmployeeSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the employee search params
func (o *EmployeeSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithDepartmentID adds the departmentID to the employee search params
func (o *EmployeeSearchParams) WithDepartmentID(departmentID *string) *EmployeeSearchParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the employee search params
func (o *EmployeeSearchParams) SetDepartmentID(departmentID *string) {
	o.DepartmentID = departmentID
}

// WithEmployeeNumber adds the employeeNumber to the employee search params
func (o *EmployeeSearchParams) WithEmployeeNumber(employeeNumber *string) *EmployeeSearchParams {
	o.SetEmployeeNumber(employeeNumber)
	return o
}

// SetEmployeeNumber adds the employeeNumber to the employee search params
func (o *EmployeeSearchParams) SetEmployeeNumber(employeeNumber *string) {
	o.EmployeeNumber = employeeNumber
}

// WithFields adds the fields to the employee search params
func (o *EmployeeSearchParams) WithFields(fields *string) *EmployeeSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the employee search params
func (o *EmployeeSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFirstName adds the firstName to the employee search params
func (o *EmployeeSearchParams) WithFirstName(firstName *string) *EmployeeSearchParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the employee search params
func (o *EmployeeSearchParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithFrom adds the from to the employee search params
func (o *EmployeeSearchParams) WithFrom(from *int64) *EmployeeSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the employee search params
func (o *EmployeeSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithHasSystemAccess adds the hasSystemAccess to the employee search params
func (o *EmployeeSearchParams) WithHasSystemAccess(hasSystemAccess *bool) *EmployeeSearchParams {
	o.SetHasSystemAccess(hasSystemAccess)
	return o
}

// SetHasSystemAccess adds the hasSystemAccess to the employee search params
func (o *EmployeeSearchParams) SetHasSystemAccess(hasSystemAccess *bool) {
	o.HasSystemAccess = hasSystemAccess
}

// WithID adds the id to the employee search params
func (o *EmployeeSearchParams) WithID(id *string) *EmployeeSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employee search params
func (o *EmployeeSearchParams) SetID(id *string) {
	o.ID = id
}

// WithIncludeContacts adds the includeContacts to the employee search params
func (o *EmployeeSearchParams) WithIncludeContacts(includeContacts *bool) *EmployeeSearchParams {
	o.SetIncludeContacts(includeContacts)
	return o
}

// SetIncludeContacts adds the includeContacts to the employee search params
func (o *EmployeeSearchParams) SetIncludeContacts(includeContacts *bool) {
	o.IncludeContacts = includeContacts
}

// WithLastName adds the lastName to the employee search params
func (o *EmployeeSearchParams) WithLastName(lastName *string) *EmployeeSearchParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the employee search params
func (o *EmployeeSearchParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithOnlyProjectManagers adds the onlyProjectManagers to the employee search params
func (o *EmployeeSearchParams) WithOnlyProjectManagers(onlyProjectManagers *bool) *EmployeeSearchParams {
	o.SetOnlyProjectManagers(onlyProjectManagers)
	return o
}

// SetOnlyProjectManagers adds the onlyProjectManagers to the employee search params
func (o *EmployeeSearchParams) SetOnlyProjectManagers(onlyProjectManagers *bool) {
	o.OnlyProjectManagers = onlyProjectManagers
}

// WithPeriodEnd adds the periodEnd to the employee search params
func (o *EmployeeSearchParams) WithPeriodEnd(periodEnd *string) *EmployeeSearchParams {
	o.SetPeriodEnd(periodEnd)
	return o
}

// SetPeriodEnd adds the periodEnd to the employee search params
func (o *EmployeeSearchParams) SetPeriodEnd(periodEnd *string) {
	o.PeriodEnd = periodEnd
}

// WithPeriodStart adds the periodStart to the employee search params
func (o *EmployeeSearchParams) WithPeriodStart(periodStart *string) *EmployeeSearchParams {
	o.SetPeriodStart(periodStart)
	return o
}

// SetPeriodStart adds the periodStart to the employee search params
func (o *EmployeeSearchParams) SetPeriodStart(periodStart *string) {
	o.PeriodStart = periodStart
}

// WithSorting adds the sorting to the employee search params
func (o *EmployeeSearchParams) WithSorting(sorting *string) *EmployeeSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the employee search params
func (o *EmployeeSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowInformationRegistration != nil {

		// query param allowInformationRegistration
		var qrAllowInformationRegistration bool
		if o.AllowInformationRegistration != nil {
			qrAllowInformationRegistration = *o.AllowInformationRegistration
		}
		qAllowInformationRegistration := swag.FormatBool(qrAllowInformationRegistration)
		if qAllowInformationRegistration != "" {
			if err := r.SetQueryParam("allowInformationRegistration", qAllowInformationRegistration); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.DepartmentID != nil {

		// query param departmentId
		var qrDepartmentID string
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := qrDepartmentID
		if qDepartmentID != "" {
			if err := r.SetQueryParam("departmentId", qDepartmentID); err != nil {
				return err
			}
		}

	}

	if o.EmployeeNumber != nil {

		// query param employeeNumber
		var qrEmployeeNumber string
		if o.EmployeeNumber != nil {
			qrEmployeeNumber = *o.EmployeeNumber
		}
		qEmployeeNumber := qrEmployeeNumber
		if qEmployeeNumber != "" {
			if err := r.SetQueryParam("employeeNumber", qEmployeeNumber); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.FirstName != nil {

		// query param firstName
		var qrFirstName string
		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {
			if err := r.SetQueryParam("firstName", qFirstName); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.HasSystemAccess != nil {

		// query param hasSystemAccess
		var qrHasSystemAccess bool
		if o.HasSystemAccess != nil {
			qrHasSystemAccess = *o.HasSystemAccess
		}
		qHasSystemAccess := swag.FormatBool(qrHasSystemAccess)
		if qHasSystemAccess != "" {
			if err := r.SetQueryParam("hasSystemAccess", qHasSystemAccess); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IncludeContacts != nil {

		// query param includeContacts
		var qrIncludeContacts bool
		if o.IncludeContacts != nil {
			qrIncludeContacts = *o.IncludeContacts
		}
		qIncludeContacts := swag.FormatBool(qrIncludeContacts)
		if qIncludeContacts != "" {
			if err := r.SetQueryParam("includeContacts", qIncludeContacts); err != nil {
				return err
			}
		}

	}

	if o.LastName != nil {

		// query param lastName
		var qrLastName string
		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {
			if err := r.SetQueryParam("lastName", qLastName); err != nil {
				return err
			}
		}

	}

	if o.OnlyProjectManagers != nil {

		// query param onlyProjectManagers
		var qrOnlyProjectManagers bool
		if o.OnlyProjectManagers != nil {
			qrOnlyProjectManagers = *o.OnlyProjectManagers
		}
		qOnlyProjectManagers := swag.FormatBool(qrOnlyProjectManagers)
		if qOnlyProjectManagers != "" {
			if err := r.SetQueryParam("onlyProjectManagers", qOnlyProjectManagers); err != nil {
				return err
			}
		}

	}

	if o.PeriodEnd != nil {

		// query param periodEnd
		var qrPeriodEnd string
		if o.PeriodEnd != nil {
			qrPeriodEnd = *o.PeriodEnd
		}
		qPeriodEnd := qrPeriodEnd
		if qPeriodEnd != "" {
			if err := r.SetQueryParam("periodEnd", qPeriodEnd); err != nil {
				return err
			}
		}

	}

	if o.PeriodStart != nil {

		// query param periodStart
		var qrPeriodStart string
		if o.PeriodStart != nil {
			qrPeriodStart = *o.PeriodStart
		}
		qPeriodStart := qrPeriodStart
		if qPeriodStart != "" {
			if err := r.SetQueryParam("periodStart", qPeriodStart); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
