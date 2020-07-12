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

package document_archive

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

// NewDocumentArchiveEmployeeGetEmployeeParams creates a new DocumentArchiveEmployeeGetEmployeeParams object
// with the default values initialized.
func NewDocumentArchiveEmployeeGetEmployeeParams() *DocumentArchiveEmployeeGetEmployeeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveEmployeeGetEmployeeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentArchiveEmployeeGetEmployeeParamsWithTimeout creates a new DocumentArchiveEmployeeGetEmployeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentArchiveEmployeeGetEmployeeParamsWithTimeout(timeout time.Duration) *DocumentArchiveEmployeeGetEmployeeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveEmployeeGetEmployeeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewDocumentArchiveEmployeeGetEmployeeParamsWithContext creates a new DocumentArchiveEmployeeGetEmployeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentArchiveEmployeeGetEmployeeParamsWithContext(ctx context.Context) *DocumentArchiveEmployeeGetEmployeeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveEmployeeGetEmployeeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewDocumentArchiveEmployeeGetEmployeeParamsWithHTTPClient creates a new DocumentArchiveEmployeeGetEmployeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentArchiveEmployeeGetEmployeeParamsWithHTTPClient(client *http.Client) *DocumentArchiveEmployeeGetEmployeeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveEmployeeGetEmployeeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*DocumentArchiveEmployeeGetEmployeeParams contains all the parameters to send to the API endpoint
for the document archive employee get employee operation typically these are written to a http.Request
*/
type DocumentArchiveEmployeeGetEmployeeParams struct {

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
	  Element ID

	*/
	ID int32
	/*PeriodDateFrom
	  From and including

	*/
	PeriodDateFrom *string
	/*PeriodDateTo
	  To and excluding

	*/
	PeriodDateTo *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithTimeout(timeout time.Duration) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithContext(ctx context.Context) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithHTTPClient(client *http.Client) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithCount(count *int64) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithFields(fields *string) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithFrom(from *int64) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithID(id int32) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetID(id int32) {
	o.ID = id
}

// WithPeriodDateFrom adds the periodDateFrom to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithPeriodDateFrom(periodDateFrom *string) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetPeriodDateFrom(periodDateFrom)
	return o
}

// SetPeriodDateFrom adds the periodDateFrom to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetPeriodDateFrom(periodDateFrom *string) {
	o.PeriodDateFrom = periodDateFrom
}

// WithPeriodDateTo adds the periodDateTo to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithPeriodDateTo(periodDateTo *string) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetPeriodDateTo(periodDateTo)
	return o
}

// SetPeriodDateTo adds the periodDateTo to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetPeriodDateTo(periodDateTo *string) {
	o.PeriodDateTo = periodDateTo
}

// WithSorting adds the sorting to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) WithSorting(sorting *string) *DocumentArchiveEmployeeGetEmployeeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the document archive employee get employee params
func (o *DocumentArchiveEmployeeGetEmployeeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentArchiveEmployeeGetEmployeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.PeriodDateFrom != nil {

		// query param periodDateFrom
		var qrPeriodDateFrom string
		if o.PeriodDateFrom != nil {
			qrPeriodDateFrom = *o.PeriodDateFrom
		}
		qPeriodDateFrom := qrPeriodDateFrom
		if qPeriodDateFrom != "" {
			if err := r.SetQueryParam("periodDateFrom", qPeriodDateFrom); err != nil {
				return err
			}
		}

	}

	if o.PeriodDateTo != nil {

		// query param periodDateTo
		var qrPeriodDateTo string
		if o.PeriodDateTo != nil {
			qrPeriodDateTo = *o.PeriodDateTo
		}
		qPeriodDateTo := qrPeriodDateTo
		if qPeriodDateTo != "" {
			if err := r.SetQueryParam("periodDateTo", qPeriodDateTo); err != nil {
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
