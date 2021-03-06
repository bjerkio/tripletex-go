// Code generated by go-swagger; DO NOT EDIT.

package bank

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

// NewBankSearchParams creates a new BankSearchParams object
// with the default values initialized.
func NewBankSearchParams() *BankSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewBankSearchParamsWithTimeout creates a new BankSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankSearchParamsWithTimeout(timeout time.Duration) *BankSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewBankSearchParamsWithContext creates a new BankSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankSearchParamsWithContext(ctx context.Context) *BankSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewBankSearchParamsWithHTTPClient creates a new BankSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankSearchParamsWithHTTPClient(client *http.Client) *BankSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*BankSearchParams contains all the parameters to send to the API endpoint
for the bank search operation typically these are written to a http.Request
*/
type BankSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*ID
	  List of IDs

	*/
	ID *string
	/*IsAutoPaySupported
	  Equals

	*/
	IsAutoPaySupported *bool
	/*IsBankReconciliationSupport
	  Equals

	*/
	IsBankReconciliationSupport *bool
	/*RegisterNumbers
	  Bank register number (four digits)

	*/
	RegisterNumbers *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bank search params
func (o *BankSearchParams) WithTimeout(timeout time.Duration) *BankSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank search params
func (o *BankSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank search params
func (o *BankSearchParams) WithContext(ctx context.Context) *BankSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank search params
func (o *BankSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank search params
func (o *BankSearchParams) WithHTTPClient(client *http.Client) *BankSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank search params
func (o *BankSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the bank search params
func (o *BankSearchParams) WithCount(count *int64) *BankSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the bank search params
func (o *BankSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the bank search params
func (o *BankSearchParams) WithFields(fields *string) *BankSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the bank search params
func (o *BankSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the bank search params
func (o *BankSearchParams) WithFrom(from *int64) *BankSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the bank search params
func (o *BankSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the bank search params
func (o *BankSearchParams) WithID(id *string) *BankSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bank search params
func (o *BankSearchParams) SetID(id *string) {
	o.ID = id
}

// WithIsAutoPaySupported adds the isAutoPaySupported to the bank search params
func (o *BankSearchParams) WithIsAutoPaySupported(isAutoPaySupported *bool) *BankSearchParams {
	o.SetIsAutoPaySupported(isAutoPaySupported)
	return o
}

// SetIsAutoPaySupported adds the isAutoPaySupported to the bank search params
func (o *BankSearchParams) SetIsAutoPaySupported(isAutoPaySupported *bool) {
	o.IsAutoPaySupported = isAutoPaySupported
}

// WithIsBankReconciliationSupport adds the isBankReconciliationSupport to the bank search params
func (o *BankSearchParams) WithIsBankReconciliationSupport(isBankReconciliationSupport *bool) *BankSearchParams {
	o.SetIsBankReconciliationSupport(isBankReconciliationSupport)
	return o
}

// SetIsBankReconciliationSupport adds the isBankReconciliationSupport to the bank search params
func (o *BankSearchParams) SetIsBankReconciliationSupport(isBankReconciliationSupport *bool) {
	o.IsBankReconciliationSupport = isBankReconciliationSupport
}

// WithRegisterNumbers adds the registerNumbers to the bank search params
func (o *BankSearchParams) WithRegisterNumbers(registerNumbers *string) *BankSearchParams {
	o.SetRegisterNumbers(registerNumbers)
	return o
}

// SetRegisterNumbers adds the registerNumbers to the bank search params
func (o *BankSearchParams) SetRegisterNumbers(registerNumbers *string) {
	o.RegisterNumbers = registerNumbers
}

// WithSorting adds the sorting to the bank search params
func (o *BankSearchParams) WithSorting(sorting *string) *BankSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the bank search params
func (o *BankSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *BankSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.IsAutoPaySupported != nil {

		// query param isAutoPaySupported
		var qrIsAutoPaySupported bool
		if o.IsAutoPaySupported != nil {
			qrIsAutoPaySupported = *o.IsAutoPaySupported
		}
		qIsAutoPaySupported := swag.FormatBool(qrIsAutoPaySupported)
		if qIsAutoPaySupported != "" {
			if err := r.SetQueryParam("isAutoPaySupported", qIsAutoPaySupported); err != nil {
				return err
			}
		}

	}

	if o.IsBankReconciliationSupport != nil {

		// query param isBankReconciliationSupport
		var qrIsBankReconciliationSupport bool
		if o.IsBankReconciliationSupport != nil {
			qrIsBankReconciliationSupport = *o.IsBankReconciliationSupport
		}
		qIsBankReconciliationSupport := swag.FormatBool(qrIsBankReconciliationSupport)
		if qIsBankReconciliationSupport != "" {
			if err := r.SetQueryParam("isBankReconciliationSupport", qIsBankReconciliationSupport); err != nil {
				return err
			}
		}

	}

	if o.RegisterNumbers != nil {

		// query param registerNumbers
		var qrRegisterNumbers string
		if o.RegisterNumbers != nil {
			qrRegisterNumbers = *o.RegisterNumbers
		}
		qRegisterNumbers := qrRegisterNumbers
		if qRegisterNumbers != "" {
			if err := r.SetQueryParam("registerNumbers", qRegisterNumbers); err != nil {
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
