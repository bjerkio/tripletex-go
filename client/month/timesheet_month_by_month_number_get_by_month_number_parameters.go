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

package month

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

// NewTimesheetMonthByMonthNumberGetByMonthNumberParams creates a new TimesheetMonthByMonthNumberGetByMonthNumberParams object
// with the default values initialized.
func NewTimesheetMonthByMonthNumberGetByMonthNumberParams() *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetMonthByMonthNumberGetByMonthNumberParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithTimeout creates a new TimesheetMonthByMonthNumberGetByMonthNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithTimeout(timeout time.Duration) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetMonthByMonthNumberGetByMonthNumberParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithContext creates a new TimesheetMonthByMonthNumberGetByMonthNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithContext(ctx context.Context) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetMonthByMonthNumberGetByMonthNumberParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithHTTPClient creates a new TimesheetMonthByMonthNumberGetByMonthNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetMonthByMonthNumberGetByMonthNumberParamsWithHTTPClient(client *http.Client) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetMonthByMonthNumberGetByMonthNumberParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*TimesheetMonthByMonthNumberGetByMonthNumberParams contains all the parameters to send to the API endpoint
for the timesheet month by month number get by month number operation typically these are written to a http.Request
*/
type TimesheetMonthByMonthNumberGetByMonthNumberParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmployeeIds
	  List of IDs. Defaults to ID of token owner.

	*/
	EmployeeIds string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*MonthYear
	  2020-01

	*/
	MonthYear string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithTimeout(timeout time.Duration) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithContext(ctx context.Context) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithHTTPClient(client *http.Client) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithCount(count *int64) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmployeeIds adds the employeeIds to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithEmployeeIds(employeeIds string) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetEmployeeIds(employeeIds)
	return o
}

// SetEmployeeIds adds the employeeIds to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetEmployeeIds(employeeIds string) {
	o.EmployeeIds = employeeIds
}

// WithFields adds the fields to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithFields(fields *string) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithFrom(from *int64) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetFrom(from *int64) {
	o.From = from
}

// WithMonthYear adds the monthYear to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithMonthYear(monthYear string) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetMonthYear(monthYear)
	return o
}

// SetMonthYear adds the monthYear to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetMonthYear(monthYear string) {
	o.MonthYear = monthYear
}

// WithSorting adds the sorting to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WithSorting(sorting *string) *TimesheetMonthByMonthNumberGetByMonthNumberParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the timesheet month by month number get by month number params
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetMonthByMonthNumberGetByMonthNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param employeeIds
	qrEmployeeIds := o.EmployeeIds
	qEmployeeIds := qrEmployeeIds
	if qEmployeeIds != "" {
		if err := r.SetQueryParam("employeeIds", qEmployeeIds); err != nil {
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

	// query param monthYear
	qrMonthYear := o.MonthYear
	qMonthYear := qrMonthYear
	if qMonthYear != "" {
		if err := r.SetQueryParam("monthYear", qMonthYear); err != nil {
			return err
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
