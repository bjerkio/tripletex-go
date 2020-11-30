// Code generated by go-swagger; DO NOT EDIT.

package entry

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

// NewTimesheetEntryRecentProjectsGetRecentProjectsParams creates a new TimesheetEntryRecentProjectsGetRecentProjectsParams object
// with the default values initialized.
func NewTimesheetEntryRecentProjectsGetRecentProjectsParams() *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetEntryRecentProjectsGetRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithTimeout creates a new TimesheetEntryRecentProjectsGetRecentProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithTimeout(timeout time.Duration) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetEntryRecentProjectsGetRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithContext creates a new TimesheetEntryRecentProjectsGetRecentProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithContext(ctx context.Context) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetEntryRecentProjectsGetRecentProjectsParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithHTTPClient creates a new TimesheetEntryRecentProjectsGetRecentProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetEntryRecentProjectsGetRecentProjectsParamsWithHTTPClient(client *http.Client) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TimesheetEntryRecentProjectsGetRecentProjectsParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*TimesheetEntryRecentProjectsGetRecentProjectsParams contains all the parameters to send to the API endpoint
for the timesheet entry recent projects get recent projects operation typically these are written to a http.Request
*/
type TimesheetEntryRecentProjectsGetRecentProjectsParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*EmployeeID
	  ID of employee with recent project hours Defaults to ID of token owner.

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

// WithTimeout adds the timeout to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithTimeout(timeout time.Duration) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithContext(ctx context.Context) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithHTTPClient(client *http.Client) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithCount(count *int64) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetCount(count *int64) {
	o.Count = count
}

// WithEmployeeID adds the employeeID to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithEmployeeID(employeeID *int32) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithFields(fields *string) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithFrom(from *int64) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WithSorting(sorting *string) *TimesheetEntryRecentProjectsGetRecentProjectsParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the timesheet entry recent projects get recent projects params
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetEntryRecentProjectsGetRecentProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
