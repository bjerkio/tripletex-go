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

package statement

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

// NewBankStatementSearchParams creates a new BankStatementSearchParams object
// with the default values initialized.
func NewBankStatementSearchParams() *BankStatementSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankStatementSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewBankStatementSearchParamsWithTimeout creates a new BankStatementSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankStatementSearchParamsWithTimeout(timeout time.Duration) *BankStatementSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankStatementSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewBankStatementSearchParamsWithContext creates a new BankStatementSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankStatementSearchParamsWithContext(ctx context.Context) *BankStatementSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankStatementSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewBankStatementSearchParamsWithHTTPClient creates a new BankStatementSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankStatementSearchParamsWithHTTPClient(client *http.Client) *BankStatementSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &BankStatementSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*BankStatementSearchParams contains all the parameters to send to the API endpoint
for the bank statement search operation typically these are written to a http.Request
*/
type BankStatementSearchParams struct {

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

// WithTimeout adds the timeout to the bank statement search params
func (o *BankStatementSearchParams) WithTimeout(timeout time.Duration) *BankStatementSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank statement search params
func (o *BankStatementSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank statement search params
func (o *BankStatementSearchParams) WithContext(ctx context.Context) *BankStatementSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank statement search params
func (o *BankStatementSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank statement search params
func (o *BankStatementSearchParams) WithHTTPClient(client *http.Client) *BankStatementSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank statement search params
func (o *BankStatementSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the bank statement search params
func (o *BankStatementSearchParams) WithCount(count *int64) *BankStatementSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the bank statement search params
func (o *BankStatementSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the bank statement search params
func (o *BankStatementSearchParams) WithFields(fields *string) *BankStatementSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the bank statement search params
func (o *BankStatementSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the bank statement search params
func (o *BankStatementSearchParams) WithFrom(from *int64) *BankStatementSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the bank statement search params
func (o *BankStatementSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the bank statement search params
func (o *BankStatementSearchParams) WithID(id *string) *BankStatementSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bank statement search params
func (o *BankStatementSearchParams) SetID(id *string) {
	o.ID = id
}

// WithSorting adds the sorting to the bank statement search params
func (o *BankStatementSearchParams) WithSorting(sorting *string) *BankStatementSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the bank statement search params
func (o *BankStatementSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *BankStatementSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
