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

package voucher

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

// NewLedgerVoucherSearchParams creates a new LedgerVoucherSearchParams object
// with the default values initialized.
func NewLedgerVoucherSearchParams() *LedgerVoucherSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerVoucherSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerVoucherSearchParamsWithTimeout creates a new LedgerVoucherSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerVoucherSearchParamsWithTimeout(timeout time.Duration) *LedgerVoucherSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerVoucherSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewLedgerVoucherSearchParamsWithContext creates a new LedgerVoucherSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerVoucherSearchParamsWithContext(ctx context.Context) *LedgerVoucherSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerVoucherSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewLedgerVoucherSearchParamsWithHTTPClient creates a new LedgerVoucherSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerVoucherSearchParamsWithHTTPClient(client *http.Client) *LedgerVoucherSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &LedgerVoucherSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*LedgerVoucherSearchParams contains all the parameters to send to the API endpoint
for the ledger voucher search operation typically these are written to a http.Request
*/
type LedgerVoucherSearchParams struct {

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
	/*Number
	  List of IDs

	*/
	Number *string
	/*NumberFrom
	  From and including

	*/
	NumberFrom *int32
	/*NumberTo
	  To and excluding

	*/
	NumberTo *int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*TypeID
	  List of IDs

	*/
	TypeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithTimeout(timeout time.Duration) *LedgerVoucherSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithContext(ctx context.Context) *LedgerVoucherSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithHTTPClient(client *http.Client) *LedgerVoucherSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithCount(count *int64) *LedgerVoucherSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithDateFrom adds the dateFrom to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithDateFrom(dateFrom string) *LedgerVoucherSearchParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetDateFrom(dateFrom string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithDateTo(dateTo string) *LedgerVoucherSearchParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetDateTo(dateTo string) {
	o.DateTo = dateTo
}

// WithFields adds the fields to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithFields(fields *string) *LedgerVoucherSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithFrom(from *int64) *LedgerVoucherSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithID(id *string) *LedgerVoucherSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetID(id *string) {
	o.ID = id
}

// WithNumber adds the number to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithNumber(number *string) *LedgerVoucherSearchParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetNumber(number *string) {
	o.Number = number
}

// WithNumberFrom adds the numberFrom to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithNumberFrom(numberFrom *int32) *LedgerVoucherSearchParams {
	o.SetNumberFrom(numberFrom)
	return o
}

// SetNumberFrom adds the numberFrom to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetNumberFrom(numberFrom *int32) {
	o.NumberFrom = numberFrom
}

// WithNumberTo adds the numberTo to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithNumberTo(numberTo *int32) *LedgerVoucherSearchParams {
	o.SetNumberTo(numberTo)
	return o
}

// SetNumberTo adds the numberTo to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetNumberTo(numberTo *int32) {
	o.NumberTo = numberTo
}

// WithSorting adds the sorting to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithSorting(sorting *string) *LedgerVoucherSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithTypeID adds the typeID to the ledger voucher search params
func (o *LedgerVoucherSearchParams) WithTypeID(typeID *string) *LedgerVoucherSearchParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the ledger voucher search params
func (o *LedgerVoucherSearchParams) SetTypeID(typeID *string) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerVoucherSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
				return err
			}
		}

	}

	if o.NumberFrom != nil {

		// query param numberFrom
		var qrNumberFrom int32
		if o.NumberFrom != nil {
			qrNumberFrom = *o.NumberFrom
		}
		qNumberFrom := swag.FormatInt32(qrNumberFrom)
		if qNumberFrom != "" {
			if err := r.SetQueryParam("numberFrom", qNumberFrom); err != nil {
				return err
			}
		}

	}

	if o.NumberTo != nil {

		// query param numberTo
		var qrNumberTo int32
		if o.NumberTo != nil {
			qrNumberTo = *o.NumberTo
		}
		qNumberTo := swag.FormatInt32(qrNumberTo)
		if qNumberTo != "" {
			if err := r.SetQueryParam("numberTo", qNumberTo); err != nil {
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

	if o.TypeID != nil {

		// query param typeId
		var qrTypeID string
		if o.TypeID != nil {
			qrTypeID = *o.TypeID
		}
		qTypeID := qrTypeID
		if qTypeID != "" {
			if err := r.SetQueryParam("typeId", qTypeID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
