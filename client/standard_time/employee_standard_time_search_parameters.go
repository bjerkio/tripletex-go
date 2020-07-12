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

package standard_time

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

// NewEmployeeStandardTimeSearchParams creates a new EmployeeStandardTimeSearchParams object
// with the default values initialized.
func NewEmployeeStandardTimeSearchParams() *EmployeeStandardTimeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeStandardTimeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeStandardTimeSearchParamsWithTimeout creates a new EmployeeStandardTimeSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeStandardTimeSearchParamsWithTimeout(timeout time.Duration) *EmployeeStandardTimeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeStandardTimeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewEmployeeStandardTimeSearchParamsWithContext creates a new EmployeeStandardTimeSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeStandardTimeSearchParamsWithContext(ctx context.Context) *EmployeeStandardTimeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeStandardTimeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewEmployeeStandardTimeSearchParamsWithHTTPClient creates a new EmployeeStandardTimeSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeStandardTimeSearchParamsWithHTTPClient(client *http.Client) *EmployeeStandardTimeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeStandardTimeSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*EmployeeStandardTimeSearchParams contains all the parameters to send to the API endpoint
for the employee standard time search operation typically these are written to a http.Request
*/
type EmployeeStandardTimeSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmployeeID
	  Employee ID. Defaults to ID of token owner.

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
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithTimeout(timeout time.Duration) *EmployeeStandardTimeSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithContext(ctx context.Context) *EmployeeStandardTimeSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithHTTPClient(client *http.Client) *EmployeeStandardTimeSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithCount(count *int64) *EmployeeStandardTimeSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmployeeID adds the employeeID to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithEmployeeID(employeeID *int32) *EmployeeStandardTimeSearchParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithFields(fields *string) *EmployeeStandardTimeSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithFrom(from *int64) *EmployeeStandardTimeSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) WithSorting(sorting *string) *EmployeeStandardTimeSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the employee standard time search params
func (o *EmployeeStandardTimeSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeStandardTimeSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
