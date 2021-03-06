// Code generated by go-swagger; DO NOT EDIT.

package annual_account

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

// NewLedgerAnnualAccountSearchParams creates a new LedgerAnnualAccountSearchParams object
// with the default values initialized.
func NewLedgerAnnualAccountSearchParams() *LedgerAnnualAccountSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerAnnualAccountSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerAnnualAccountSearchParamsWithTimeout creates a new LedgerAnnualAccountSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerAnnualAccountSearchParamsWithTimeout(timeout time.Duration) *LedgerAnnualAccountSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerAnnualAccountSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewLedgerAnnualAccountSearchParamsWithContext creates a new LedgerAnnualAccountSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerAnnualAccountSearchParamsWithContext(ctx context.Context) *LedgerAnnualAccountSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerAnnualAccountSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewLedgerAnnualAccountSearchParamsWithHTTPClient creates a new LedgerAnnualAccountSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerAnnualAccountSearchParamsWithHTTPClient(client *http.Client) *LedgerAnnualAccountSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerAnnualAccountSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*LedgerAnnualAccountSearchParams contains all the parameters to send to the API endpoint
for the ledger annual account search operation typically these are written to a http.Request
*/
type LedgerAnnualAccountSearchParams struct {

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
	/*YearFrom
	  From and including

	*/
	YearFrom *int32
	/*YearTo
	  To and excluding

	*/
	YearTo *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithTimeout(timeout time.Duration) *LedgerAnnualAccountSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithContext(ctx context.Context) *LedgerAnnualAccountSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithHTTPClient(client *http.Client) *LedgerAnnualAccountSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithCount(count *int64) *LedgerAnnualAccountSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithFields(fields *string) *LedgerAnnualAccountSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithFrom(from *int64) *LedgerAnnualAccountSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithID(id *string) *LedgerAnnualAccountSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithSorting(sorting *string) *LedgerAnnualAccountSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithYearFrom adds the yearFrom to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithYearFrom(yearFrom *int32) *LedgerAnnualAccountSearchParams {
	o.SetYearFrom(yearFrom)
	return o
}

// SetYearFrom adds the yearFrom to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetYearFrom(yearFrom *int32) {
	o.YearFrom = yearFrom
}

// WithYearTo adds the yearTo to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) WithYearTo(yearTo *int32) *LedgerAnnualAccountSearchParams {
	o.SetYearTo(yearTo)
	return o
}

// SetYearTo adds the yearTo to the ledger annual account search params
func (o *LedgerAnnualAccountSearchParams) SetYearTo(yearTo *int32) {
	o.YearTo = yearTo
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerAnnualAccountSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.YearFrom != nil {

		// query param yearFrom
		var qrYearFrom int32
		if o.YearFrom != nil {
			qrYearFrom = *o.YearFrom
		}
		qYearFrom := swag.FormatInt32(qrYearFrom)
		if qYearFrom != "" {
			if err := r.SetQueryParam("yearFrom", qYearFrom); err != nil {
				return err
			}
		}

	}

	if o.YearTo != nil {

		// query param yearTo
		var qrYearTo int32
		if o.YearTo != nil {
			qrYearTo = *o.YearTo
		}
		qYearTo := swag.FormatInt32(qrYearTo)
		if qYearTo != "" {
			if err := r.SetQueryParam("yearTo", qYearTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
