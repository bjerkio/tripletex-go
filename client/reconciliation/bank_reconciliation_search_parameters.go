// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

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

// NewBankReconciliationSearchParams creates a new BankReconciliationSearchParams object
// with the default values initialized.
func NewBankReconciliationSearchParams() *BankReconciliationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankReconciliationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewBankReconciliationSearchParamsWithTimeout creates a new BankReconciliationSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankReconciliationSearchParamsWithTimeout(timeout time.Duration) *BankReconciliationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankReconciliationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewBankReconciliationSearchParamsWithContext creates a new BankReconciliationSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankReconciliationSearchParamsWithContext(ctx context.Context) *BankReconciliationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankReconciliationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewBankReconciliationSearchParamsWithHTTPClient creates a new BankReconciliationSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankReconciliationSearchParamsWithHTTPClient(client *http.Client) *BankReconciliationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankReconciliationSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*BankReconciliationSearchParams contains all the parameters to send to the API endpoint
for the bank reconciliation search operation typically these are written to a http.Request
*/
type BankReconciliationSearchParams struct {

	/*AccountID
	  List of IDs

	*/
	AccountID *string
	/*AccountingPeriodID
	  List of IDs

	*/
	AccountingPeriodID *string
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
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithTimeout(timeout time.Duration) *BankReconciliationSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithContext(ctx context.Context) *BankReconciliationSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithHTTPClient(client *http.Client) *BankReconciliationSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithAccountID(accountID *string) *BankReconciliationSearchParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithAccountingPeriodID adds the accountingPeriodID to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithAccountingPeriodID(accountingPeriodID *string) *BankReconciliationSearchParams {
	o.SetAccountingPeriodID(accountingPeriodID)
	return o
}

// SetAccountingPeriodID adds the accountingPeriodId to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetAccountingPeriodID(accountingPeriodID *string) {
	o.AccountingPeriodID = accountingPeriodID
}

// WithCount adds the count to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithCount(count *int64) *BankReconciliationSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithFields(fields *string) *BankReconciliationSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithFrom(from *int64) *BankReconciliationSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithID(id *string) *BankReconciliationSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the bank reconciliation search params
func (o *BankReconciliationSearchParams) WithSorting(sorting *string) *BankReconciliationSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the bank reconciliation search params
func (o *BankReconciliationSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *BankReconciliationSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if o.AccountingPeriodID != nil {

		// query param accountingPeriodId
		var qrAccountingPeriodID string
		if o.AccountingPeriodID != nil {
			qrAccountingPeriodID = *o.AccountingPeriodID
		}
		qAccountingPeriodID := qrAccountingPeriodID
		if qAccountingPeriodID != "" {
			if err := r.SetQueryParam("accountingPeriodId", qAccountingPeriodID); err != nil {
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
