// Code generated by go-swagger; DO NOT EDIT.

package period

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

// NewProjectPeriodOverallStatusOverallStatusParams creates a new ProjectPeriodOverallStatusOverallStatusParams object
// with the default values initialized.
func NewProjectPeriodOverallStatusOverallStatusParams() *ProjectPeriodOverallStatusOverallStatusParams {
	var ()
	return &ProjectPeriodOverallStatusOverallStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectPeriodOverallStatusOverallStatusParamsWithTimeout creates a new ProjectPeriodOverallStatusOverallStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectPeriodOverallStatusOverallStatusParamsWithTimeout(timeout time.Duration) *ProjectPeriodOverallStatusOverallStatusParams {
	var ()
	return &ProjectPeriodOverallStatusOverallStatusParams{

		timeout: timeout,
	}
}

// NewProjectPeriodOverallStatusOverallStatusParamsWithContext creates a new ProjectPeriodOverallStatusOverallStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectPeriodOverallStatusOverallStatusParamsWithContext(ctx context.Context) *ProjectPeriodOverallStatusOverallStatusParams {
	var ()
	return &ProjectPeriodOverallStatusOverallStatusParams{

		Context: ctx,
	}
}

// NewProjectPeriodOverallStatusOverallStatusParamsWithHTTPClient creates a new ProjectPeriodOverallStatusOverallStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectPeriodOverallStatusOverallStatusParamsWithHTTPClient(client *http.Client) *ProjectPeriodOverallStatusOverallStatusParams {
	var ()
	return &ProjectPeriodOverallStatusOverallStatusParams{
		HTTPClient: client,
	}
}

/*ProjectPeriodOverallStatusOverallStatusParams contains all the parameters to send to the API endpoint
for the project period overall status overall status operation typically these are written to a http.Request
*/
type ProjectPeriodOverallStatusOverallStatusParams struct {

	/*DateFrom
	  Format is yyyy-MM-dd (from and incl.).

	*/
	DateFrom string
	/*DateTo
	  Format is yyyy-MM-dd (to and excl.).

	*/
	DateTo string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithTimeout(timeout time.Duration) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithContext(ctx context.Context) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithHTTPClient(client *http.Client) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateFrom adds the dateFrom to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithDateFrom(dateFrom string) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetDateFrom(dateFrom string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithDateTo(dateTo string) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetDateTo(dateTo string) {
	o.DateTo = dateTo
}

// WithFields adds the fields to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithFields(fields *string) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) WithID(id int32) *ProjectPeriodOverallStatusOverallStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project period overall status overall status params
func (o *ProjectPeriodOverallStatusOverallStatusParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectPeriodOverallStatusOverallStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
