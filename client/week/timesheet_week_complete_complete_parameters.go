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

// NewTimesheetWeekCompleteCompleteParams creates a new TimesheetWeekCompleteCompleteParams object
// with the default values initialized.
func NewTimesheetWeekCompleteCompleteParams() *TimesheetWeekCompleteCompleteParams {
	var ()
	return &TimesheetWeekCompleteCompleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetWeekCompleteCompleteParamsWithTimeout creates a new TimesheetWeekCompleteCompleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetWeekCompleteCompleteParamsWithTimeout(timeout time.Duration) *TimesheetWeekCompleteCompleteParams {
	var ()
	return &TimesheetWeekCompleteCompleteParams{

		timeout: timeout,
	}
}

// NewTimesheetWeekCompleteCompleteParamsWithContext creates a new TimesheetWeekCompleteCompleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetWeekCompleteCompleteParamsWithContext(ctx context.Context) *TimesheetWeekCompleteCompleteParams {
	var ()
	return &TimesheetWeekCompleteCompleteParams{

		Context: ctx,
	}
}

// NewTimesheetWeekCompleteCompleteParamsWithHTTPClient creates a new TimesheetWeekCompleteCompleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetWeekCompleteCompleteParamsWithHTTPClient(client *http.Client) *TimesheetWeekCompleteCompleteParams {
	var ()
	return &TimesheetWeekCompleteCompleteParams{
		HTTPClient: client,
	}
}

/*TimesheetWeekCompleteCompleteParams contains all the parameters to send to the API endpoint
for the timesheet week complete complete operation typically these are written to a http.Request
*/
type TimesheetWeekCompleteCompleteParams struct {

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

// WithTimeout adds the timeout to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithTimeout(timeout time.Duration) *TimesheetWeekCompleteCompleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithContext(ctx context.Context) *TimesheetWeekCompleteCompleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithHTTPClient(client *http.Client) *TimesheetWeekCompleteCompleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeID adds the employeeID to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithEmployeeID(employeeID *int32) *TimesheetWeekCompleteCompleteParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetEmployeeID(employeeID *int32) {
	o.EmployeeID = employeeID
}

// WithID adds the id to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithID(id *int32) *TimesheetWeekCompleteCompleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetID(id *int32) {
	o.ID = id
}

// WithWeekYear adds the weekYear to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) WithWeekYear(weekYear *string) *TimesheetWeekCompleteCompleteParams {
	o.SetWeekYear(weekYear)
	return o
}

// SetWeekYear adds the weekYear to the timesheet week complete complete params
func (o *TimesheetWeekCompleteCompleteParams) SetWeekYear(weekYear *string) {
	o.WeekYear = weekYear
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetWeekCompleteCompleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
