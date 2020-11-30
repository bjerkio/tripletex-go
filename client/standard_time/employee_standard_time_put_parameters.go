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

	"github.com/bjerkio/tripletex-go/models"
)

// NewEmployeeStandardTimePutParams creates a new EmployeeStandardTimePutParams object
// with the default values initialized.
func NewEmployeeStandardTimePutParams() *EmployeeStandardTimePutParams {
	var ()
	return &EmployeeStandardTimePutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeStandardTimePutParamsWithTimeout creates a new EmployeeStandardTimePutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeStandardTimePutParamsWithTimeout(timeout time.Duration) *EmployeeStandardTimePutParams {
	var ()
	return &EmployeeStandardTimePutParams{

		timeout: timeout,
	}
}

// NewEmployeeStandardTimePutParamsWithContext creates a new EmployeeStandardTimePutParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeStandardTimePutParamsWithContext(ctx context.Context) *EmployeeStandardTimePutParams {
	var ()
	return &EmployeeStandardTimePutParams{

		Context: ctx,
	}
}

// NewEmployeeStandardTimePutParamsWithHTTPClient creates a new EmployeeStandardTimePutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeStandardTimePutParamsWithHTTPClient(client *http.Client) *EmployeeStandardTimePutParams {
	var ()
	return &EmployeeStandardTimePutParams{
		HTTPClient: client,
	}
}

/*EmployeeStandardTimePutParams contains all the parameters to send to the API endpoint
for the employee standard time put operation typically these are written to a http.Request
*/
type EmployeeStandardTimePutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.StandardTime
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee standard time put params
func (o *EmployeeStandardTimePutParams) WithTimeout(timeout time.Duration) *EmployeeStandardTimePutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee standard time put params
func (o *EmployeeStandardTimePutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee standard time put params
func (o *EmployeeStandardTimePutParams) WithContext(ctx context.Context) *EmployeeStandardTimePutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee standard time put params
func (o *EmployeeStandardTimePutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee standard time put params
func (o *EmployeeStandardTimePutParams) WithHTTPClient(client *http.Client) *EmployeeStandardTimePutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee standard time put params
func (o *EmployeeStandardTimePutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee standard time put params
func (o *EmployeeStandardTimePutParams) WithBody(body *models.StandardTime) *EmployeeStandardTimePutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee standard time put params
func (o *EmployeeStandardTimePutParams) SetBody(body *models.StandardTime) {
	o.Body = body
}

// WithID adds the id to the employee standard time put params
func (o *EmployeeStandardTimePutParams) WithID(id int32) *EmployeeStandardTimePutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employee standard time put params
func (o *EmployeeStandardTimePutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeStandardTimePutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
