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

package close_group

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

// NewLedgerCloseGroupSearchParams creates a new LedgerCloseGroupSearchParams object
// with the default values initialized.
func NewLedgerCloseGroupSearchParams() *LedgerCloseGroupSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerCloseGroupSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerCloseGroupSearchParamsWithTimeout creates a new LedgerCloseGroupSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerCloseGroupSearchParamsWithTimeout(timeout time.Duration) *LedgerCloseGroupSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerCloseGroupSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewLedgerCloseGroupSearchParamsWithContext creates a new LedgerCloseGroupSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerCloseGroupSearchParamsWithContext(ctx context.Context) *LedgerCloseGroupSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerCloseGroupSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewLedgerCloseGroupSearchParamsWithHTTPClient creates a new LedgerCloseGroupSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerCloseGroupSearchParamsWithHTTPClient(client *http.Client) *LedgerCloseGroupSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerCloseGroupSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*LedgerCloseGroupSearchParams contains all the parameters to send to the API endpoint
for the ledger close group search operation typically these are written to a http.Request
*/
type LedgerCloseGroupSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*DateFrom
	  From and including

	*/
	DateFrom string
	/*DateTo
	  To and excluding

	*/
	DateTo string
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

// WithTimeout adds the timeout to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithTimeout(timeout time.Duration) *LedgerCloseGroupSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithContext(ctx context.Context) *LedgerCloseGroupSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithHTTPClient(client *http.Client) *LedgerCloseGroupSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithCount(count *int64) *LedgerCloseGroupSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithDateFrom adds the dateFrom to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithDateFrom(dateFrom string) *LedgerCloseGroupSearchParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetDateFrom(dateFrom string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithDateTo(dateTo string) *LedgerCloseGroupSearchParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetDateTo(dateTo string) {
	o.DateTo = dateTo
}

// WithFields adds the fields to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithFields(fields *string) *LedgerCloseGroupSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithFrom(from *int64) *LedgerCloseGroupSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithID(id *string) *LedgerCloseGroupSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) WithSorting(sorting *string) *LedgerCloseGroupSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the ledger close group search params
func (o *LedgerCloseGroupSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerCloseGroupSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
