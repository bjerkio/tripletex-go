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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewTimesheetEntryListPutListParams creates a new TimesheetEntryListPutListParams object
// with the default values initialized.
func NewTimesheetEntryListPutListParams() *TimesheetEntryListPutListParams {
	var ()
	return &TimesheetEntryListPutListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimesheetEntryListPutListParamsWithTimeout creates a new TimesheetEntryListPutListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimesheetEntryListPutListParamsWithTimeout(timeout time.Duration) *TimesheetEntryListPutListParams {
	var ()
	return &TimesheetEntryListPutListParams{

		timeout: timeout,
	}
}

// NewTimesheetEntryListPutListParamsWithContext creates a new TimesheetEntryListPutListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimesheetEntryListPutListParamsWithContext(ctx context.Context) *TimesheetEntryListPutListParams {
	var ()
	return &TimesheetEntryListPutListParams{

		Context: ctx,
	}
}

// NewTimesheetEntryListPutListParamsWithHTTPClient creates a new TimesheetEntryListPutListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimesheetEntryListPutListParamsWithHTTPClient(client *http.Client) *TimesheetEntryListPutListParams {
	var ()
	return &TimesheetEntryListPutListParams{
		HTTPClient: client,
	}
}

/*TimesheetEntryListPutListParams contains all the parameters to send to the API endpoint
for the timesheet entry list put list operation typically these are written to a http.Request
*/
type TimesheetEntryListPutListParams struct {

	/*Body
	  List of timesheet entry objects to update

	*/
	Body []*models.TimesheetEntry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) WithTimeout(timeout time.Duration) *TimesheetEntryListPutListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) WithContext(ctx context.Context) *TimesheetEntryListPutListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) WithHTTPClient(client *http.Client) *TimesheetEntryListPutListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) WithBody(body []*models.TimesheetEntry) *TimesheetEntryListPutListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the timesheet entry list put list params
func (o *TimesheetEntryListPutListParams) SetBody(body []*models.TimesheetEntry) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TimesheetEntryListPutListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
