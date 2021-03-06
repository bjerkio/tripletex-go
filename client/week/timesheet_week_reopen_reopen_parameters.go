// Code generated by go-swagger; DO NOT EDIT.

package week

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

// NewTimesheetWeekReopenReopenParams creates a new TimesheetWeekReopenReopenParams object
// with the default values initialized.
func NewTimesheetWeekReopenReopenParams() *TimesheetWeekReopenReopenParams {
	var ()
	return &TimesheetWeekReopenReopenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetWeekReopenReopenParamsWithTimeout creates a new TimesheetWeekReopenReopenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetWeekReopenReopenParamsWithTimeout(timeout time.Duration) *TimesheetWeekReopenReopenParams {
	var ()
	return &TimesheetWeekReopenReopenParams{

		timeout: timeout,
	}
}

// NewTimesheetWeekReopenReopenParamsWithContext creates a new TimesheetWeekReopenReopenParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetWeekReopenReopenParamsWithContext(ctx context.Context) *TimesheetWeekReopenReopenParams {
	var ()
	return &TimesheetWeekReopenReopenParams{

		Context: ctx,
	}
}

// NewTimesheetWeekReopenReopenParamsWithHTTPClient creates a new TimesheetWeekReopenReopenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetWeekReopenReopenParamsWithHTTPClient(client *http.Client) *TimesheetWeekReopenReopenParams {
	var ()
	return &TimesheetWeekReopenReopenParams{
		HTTPClient: client,
	}
}

/*TimesheetWeekReopenReopenParams contains all the parameters to send to the API endpoint
for the timesheet week reopen reopen operation typically these are written to a http.Request
*/
type TimesheetWeekReopenReopenParams struct {

	/*EmployeeID
	  Equals

	*/
	EmployeeID *int32
	/*ID
	  Equals

	*/
	ID *int32
	/*WeekYear
	  ISO-8601 week-year

	*/
	WeekYear *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithTimeout(timeout time.Duration) *TimesheetWeekReopenReopenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithContext(ctx context.Context) *TimesheetWeekReopenReopenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithHTTPClient(client *http.Client) *TimesheetWeekReopenReopenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeID adds the employeeID to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithEmployeeID(employeeID *int32) *TimesheetWeekReopenReopenParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithID adds the id to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithID(id *int32) *TimesheetWeekReopenReopenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetID(id *int32) {
	o.ID = id
}

// WithWeekYear adds the weekYear to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) WithWeekYear(weekYear *string) *TimesheetWeekReopenReopenParams {
	o.SetWeekYear(weekYear)
	return o
}

// SetWeekYear adds the weekYear to the timesheet week reopen reopen params
func (o *TimesheetWeekReopenReopenParams) SetWeekYear(weekYear *string) {
	o.WeekYear = weekYear
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetWeekReopenReopenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.WeekYear != nil {

		// query param weekYear
		var qrWeekYear string
		if o.WeekYear != nil {
			qrWeekYear = *o.WeekYear
		}
		qWeekYear := qrWeekYear
		if qWeekYear != "" {
			if err := r.SetQueryParam("weekYear", qWeekYear); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
