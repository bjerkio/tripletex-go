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

// NewTimesheetMonthUnapproveUnapproveParams creates a new TimesheetMonthUnapproveUnapproveParams object
// with the default values initialized.
func NewTimesheetMonthUnapproveUnapproveParams() *TimesheetMonthUnapproveUnapproveParams {
	var ()
	return &TimesheetMonthUnapproveUnapproveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetMonthUnapproveUnapproveParamsWithTimeout creates a new TimesheetMonthUnapproveUnapproveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetMonthUnapproveUnapproveParamsWithTimeout(timeout time.Duration) *TimesheetMonthUnapproveUnapproveParams {
	var ()
	return &TimesheetMonthUnapproveUnapproveParams{

		timeout: timeout,
	}
}

// NewTimesheetMonthUnapproveUnapproveParamsWithContext creates a new TimesheetMonthUnapproveUnapproveParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetMonthUnapproveUnapproveParamsWithContext(ctx context.Context) *TimesheetMonthUnapproveUnapproveParams {
	var ()
	return &TimesheetMonthUnapproveUnapproveParams{

		Context: ctx,
	}
}

// NewTimesheetMonthUnapproveUnapproveParamsWithHTTPClient creates a new TimesheetMonthUnapproveUnapproveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetMonthUnapproveUnapproveParamsWithHTTPClient(client *http.Client) *TimesheetMonthUnapproveUnapproveParams {
	var ()
	return &TimesheetMonthUnapproveUnapproveParams{
		HTTPClient: client,
	}
}

/*TimesheetMonthUnapproveUnapproveParams contains all the parameters to send to the API endpoint
for the timesheet month unapprove unapprove operation typically these are written to a http.Request
*/
type TimesheetMonthUnapproveUnapproveParams struct {

	/*EmployeeIds
	  List of IDs. Defaults to ID of token owner.

	*/
	EmployeeIds *string
	/*ID
	  Element ID

	*/
	ID *int32
	/*MonthYear
	  2020-01

	*/
	MonthYear *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithTimeout(timeout time.Duration) *TimesheetMonthUnapproveUnapproveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithContext(ctx context.Context) *TimesheetMonthUnapproveUnapproveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithHTTPClient(client *http.Client) *TimesheetMonthUnapproveUnapproveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeIds adds the employeeIds to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithEmployeeIds(employeeIds *string) *TimesheetMonthUnapproveUnapproveParams {
	o.SetEmployeeIds(employeeIds)
	return o
}

// SetEmployeeIds adds the employeeIds to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetEmployeeIds(employeeIds *string) {
	o.EmployeeIds = employeeIds
}

// WithID adds the id to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithID(id *int32) *TimesheetMonthUnapproveUnapproveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetID(id *int32) {
	o.ID = id
}

// WithMonthYear adds the monthYear to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) WithMonthYear(monthYear *string) *TimesheetMonthUnapproveUnapproveParams {
	o.SetMonthYear(monthYear)
	return o
}

// SetMonthYear adds the monthYear to the timesheet month unapprove unapprove params
func (o *TimesheetMonthUnapproveUnapproveParams) SetMonthYear(monthYear *string) {
	o.MonthYear = monthYear
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetMonthUnapproveUnapproveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmployeeIds != nil {

		// query param employeeIds
		var qrEmployeeIds string
		if o.EmployeeIds != nil {
			qrEmployeeIds = *o.EmployeeIds
		}
		qEmployeeIds := qrEmployeeIds
		if qEmployeeIds != "" {
			if err := r.SetQueryParam("employeeIds", qEmployeeIds); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int32
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt32(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.MonthYear != nil {

		// query param monthYear
		var qrMonthYear string
		if o.MonthYear != nil {
			qrMonthYear = *o.MonthYear
		}
		qMonthYear := qrMonthYear
		if qMonthYear != "" {
			if err := r.SetQueryParam("monthYear", qMonthYear); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
