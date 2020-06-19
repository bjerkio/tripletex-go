// Code generated by go-swagger; DO NOT EDIT.

package payment_type_out

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

// NewLedgerPaymentTypeOutSearchParams creates a new LedgerPaymentTypeOutSearchParams object
// with the default values initialized.
func NewLedgerPaymentTypeOutSearchParams() *LedgerPaymentTypeOutSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPaymentTypeOutSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerPaymentTypeOutSearchParamsWithTimeout creates a new LedgerPaymentTypeOutSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerPaymentTypeOutSearchParamsWithTimeout(timeout time.Duration) *LedgerPaymentTypeOutSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPaymentTypeOutSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewLedgerPaymentTypeOutSearchParamsWithContext creates a new LedgerPaymentTypeOutSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerPaymentTypeOutSearchParamsWithContext(ctx context.Context) *LedgerPaymentTypeOutSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPaymentTypeOutSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewLedgerPaymentTypeOutSearchParamsWithHTTPClient creates a new LedgerPaymentTypeOutSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerPaymentTypeOutSearchParamsWithHTTPClient(client *http.Client) *LedgerPaymentTypeOutSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerPaymentTypeOutSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*LedgerPaymentTypeOutSearchParams contains all the parameters to send to the API endpoint
for the ledger payment type out search operation typically these are written to a http.Request
*/
type LedgerPaymentTypeOutSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Description
	  Containing

	*/
	Description *string
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
	/*IsInactive
	  Equals

	*/
	IsInactive *bool
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithTimeout(timeout time.Duration) *LedgerPaymentTypeOutSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithContext(ctx context.Context) *LedgerPaymentTypeOutSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithHTTPClient(client *http.Client) *LedgerPaymentTypeOutSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithCount(count *int64) *LedgerPaymentTypeOutSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithDescription adds the description to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithDescription(description *string) *LedgerPaymentTypeOutSearchParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithFields(fields *string) *LedgerPaymentTypeOutSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithFrom(from *int64) *LedgerPaymentTypeOutSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithID(id *string) *LedgerPaymentTypeOutSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetID(id *string) {
	o.ID = id
}

// WithIsInactive adds the isInactive to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithIsInactive(isInactive *bool) *LedgerPaymentTypeOutSearchParams {
	o.SetIsInactive(isInactive)
	return o
}

// SetIsInactive adds the isInactive to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetIsInactive(isInactive *bool) {
	o.IsInactive = isInactive
}

// WithSorting adds the sorting to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) WithSorting(sorting *string) *LedgerPaymentTypeOutSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the ledger payment type out search params
func (o *LedgerPaymentTypeOutSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerPaymentTypeOutSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
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

	if o.IsInactive != nil {

		// query param isInactive
		var qrIsInactive bool
		if o.IsInactive != nil {
			qrIsInactive = *o.IsInactive
		}
		qIsInactive := swag.FormatBool(qrIsInactive)
		if qIsInactive != "" {
			if err := r.SetQueryParam("isInactive", qIsInactive); err != nil {
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
