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

package posting

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

// NewLedgerPostingSearchParams creates a new LedgerPostingSearchParams object
// with the default values initialized.
func NewLedgerPostingSearchParams() *LedgerPostingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPostingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerPostingSearchParamsWithTimeout creates a new LedgerPostingSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerPostingSearchParamsWithTimeout(timeout time.Duration) *LedgerPostingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPostingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewLedgerPostingSearchParamsWithContext creates a new LedgerPostingSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerPostingSearchParamsWithContext(ctx context.Context) *LedgerPostingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPostingSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewLedgerPostingSearchParamsWithHTTPClient creates a new LedgerPostingSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerPostingSearchParamsWithHTTPClient(client *http.Client) *LedgerPostingSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPostingSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*LedgerPostingSearchParams contains all the parameters to send to the API endpoint
for the ledger posting search operation typically these are written to a http.Request
*/
type LedgerPostingSearchParams struct {

	/*AccountID
	  Element ID

	*/
	AccountID *int32
	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  Element ID

	*/
	CustomerID *int32
	/*DateFrom
	  Format is yyyy-MM-dd (from and incl.).

	*/
	DateFrom string
	/*DateTo
	  Format is yyyy-MM-dd (to and excl.).

	*/
	DateTo string
	/*DepartmentID
	  Element ID

	*/
	DepartmentID *int32
	/*EmployeeID
	  Element ID

	*/
	EmployeeID *int32
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*OpenPostings
	  Deprecated

	*/
	OpenPostings *string
	/*ProductID
	  Element ID

	*/
	ProductID *int32
	/*ProjectID
	  Element ID

	*/
	ProjectID *int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*SupplierID
	  Element ID

	*/
	SupplierID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger posting search params
func (o *LedgerPostingSearchParams) WithTimeout(timeout time.Duration) *LedgerPostingSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger posting search params
func (o *LedgerPostingSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger posting search params
func (o *LedgerPostingSearchParams) WithContext(ctx context.Context) *LedgerPostingSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger posting search params
func (o *LedgerPostingSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger posting search params
func (o *LedgerPostingSearchParams) WithHTTPClient(client *http.Client) *LedgerPostingSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger posting search params
func (o *LedgerPostingSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithAccountID(accountID *int32) *LedgerPostingSearchParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetAccountID(accountID *int32) {
	o.AccountID = accountID
}

// WithCount adds the count to the ledger posting search params
func (o *LedgerPostingSearchParams) WithCount(count *int64) *LedgerPostingSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the ledger posting search params
func (o *LedgerPostingSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithCustomerID(customerID *int32) *LedgerPostingSearchParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetCustomerID(customerID *int32) {
	o.CustomerID = customerID
}

// WithDateFrom adds the dateFrom to the ledger posting search params
func (o *LedgerPostingSearchParams) WithDateFrom(dateFrom string) *LedgerPostingSearchParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the ledger posting search params
func (o *LedgerPostingSearchParams) SetDateFrom(dateFrom string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the ledger posting search params
func (o *LedgerPostingSearchParams) WithDateTo(dateTo string) *LedgerPostingSearchParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the ledger posting search params
func (o *LedgerPostingSearchParams) SetDateTo(dateTo string) {
	o.DateTo = dateTo
}

// WithDepartmentID adds the departmentID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithDepartmentID(departmentID *int32) *LedgerPostingSearchParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetDepartmentID(departmentID *int32) {
	o.DepartmentID = departmentID
}

// WithEmployeeID adds the employeeID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithEmployeeID(employeeID *int32) *LedgerPostingSearchParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the ledger posting search params
func (o *LedgerPostingSearchParams) WithFields(fields *string) *LedgerPostingSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger posting search params
func (o *LedgerPostingSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the ledger posting search params
func (o *LedgerPostingSearchParams) WithFrom(from *int64) *LedgerPostingSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the ledger posting search params
func (o *LedgerPostingSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithOpenPostings adds the openPostings to the ledger posting search params
func (o *LedgerPostingSearchParams) WithOpenPostings(openPostings *string) *LedgerPostingSearchParams {
	o.SetOpenPostings(openPostings)
	return o
}

// SetOpenPostings adds the openPostings to the ledger posting search params
func (o *LedgerPostingSearchParams) SetOpenPostings(openPostings *string) {
	o.OpenPostings = openPostings
}

// WithProductID adds the productID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithProductID(productID *int32) *LedgerPostingSearchParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetProductID(productID *int32) {
	o.ProductID = productID
}

// WithProjectID adds the projectID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithProjectID(projectID *int32) *LedgerPostingSearchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WithSorting adds the sorting to the ledger posting search params
func (o *LedgerPostingSearchParams) WithSorting(sorting *string) *LedgerPostingSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the ledger posting search params
func (o *LedgerPostingSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierID adds the supplierID to the ledger posting search params
func (o *LedgerPostingSearchParams) WithSupplierID(supplierID *int32) *LedgerPostingSearchParams {
	o.SetSupplierID(supplierID)
	return o
}

// SetSupplierID adds the supplierId to the ledger posting search params
func (o *LedgerPostingSearchParams) SetSupplierID(supplierID *int32) {
	o.SupplierID = supplierID
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerPostingSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID int32
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := swag.FormatInt32(qrAccountID)
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
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

	if o.CustomerID != nil {

		// query param customerId
		var qrCustomerID int32
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := swag.FormatInt32(qrCustomerID)
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

	// query param dateFrom
	qrDateFrom := o.DateFrom
	qDateFrom := qrDateFrom
	if qDateFrom != "" {
		if err := r.SetQueryParam("dateFrom", qDateFrom); err != nil {
			return err
		}
	}

	// query param dateTo
	qrDateTo := o.DateTo
	qDateTo := qrDateTo
	if qDateTo != "" {
		if err := r.SetQueryParam("dateTo", qDateTo); err != nil {
			return err
		}
	}

	if o.DepartmentID != nil {

		// query param departmentId
		var qrDepartmentID int32
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := swag.FormatInt32(qrDepartmentID)
		if qDepartmentID != "" {
			if err := r.SetQueryParam("departmentId", qDepartmentID); err != nil {
				return err
			}
		}

	}

	if o.EmployeeID != nil {

		// query param employeeId
		var qrEmployeeID int32
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := swag.FormatInt32(qrEmployeeID)
		if qEmployeeID != "" {
			if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
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

	if o.OpenPostings != nil {

		// query param openPostings
		var qrOpenPostings string
		if o.OpenPostings != nil {
			qrOpenPostings = *o.OpenPostings
		}
		qOpenPostings := qrOpenPostings
		if qOpenPostings != "" {
			if err := r.SetQueryParam("openPostings", qOpenPostings); err != nil {
				return err
			}
		}

	}

	if o.ProductID != nil {

		// query param productId
		var qrProductID int32
		if o.ProductID != nil {
			qrProductID = *o.ProductID
		}
		qProductID := swag.FormatInt32(qrProductID)
		if qProductID != "" {
			if err := r.SetQueryParam("productId", qProductID); err != nil {
				return err
			}
		}

	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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

	if o.SupplierID != nil {

		// query param supplierId
		var qrSupplierID int32
		if o.SupplierID != nil {
			qrSupplierID = *o.SupplierID
		}
		qSupplierID := swag.FormatInt32(qrSupplierID)
		if qSupplierID != "" {
			if err := r.SetQueryParam("supplierId", qSupplierID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
