// Code generated by go-swagger; DO NOT EDIT.

package details

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

// NewEmployeeEmploymentDetailsPutParams creates a new EmployeeEmploymentDetailsPutParams object
// with the default values initialized.
func NewEmployeeEmploymentDetailsPutParams() *EmployeeEmploymentDetailsPutParams {
	var ()
	return &EmployeeEmploymentDetailsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEmploymentDetailsPutParamsWithTimeout creates a new EmployeeEmploymentDetailsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEmploymentDetailsPutParamsWithTimeout(timeout time.Duration) *EmployeeEmploymentDetailsPutParams {
	var ()
	return &EmployeeEmploymentDetailsPutParams{

		timeout: timeout,
	}
}

// NewEmployeeEmploymentDetailsPutParamsWithContext creates a new EmployeeEmploymentDetailsPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEmploymentDetailsPutParamsWithContext(ctx context.Context) *EmployeeEmploymentDetailsPutParams {
	var ()
	return &EmployeeEmploymentDetailsPutParams{

		Context: ctx,
	}
}

// NewEmployeeEmploymentDetailsPutParamsWithHTTPClient creates a new EmployeeEmploymentDetailsPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEmploymentDetailsPutParamsWithHTTPClient(client *http.Client) *EmployeeEmploymentDetailsPutParams {
	var ()
	return &EmployeeEmploymentDetailsPutParams{
		HTTPClient: client,
	}
}

/*EmployeeEmploymentDetailsPutParams contains all the parameters to send to the API endpoint
for the employee employment details put operation typically these are written to a http.Request
*/
type EmployeeEmploymentDetailsPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.EmploymentDetails
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) WithTimeout(timeout time.Duration) *EmployeeEmploymentDetailsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) WithContext(ctx context.Context) *EmployeeEmploymentDetailsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) WithHTTPClient(client *http.Client) *EmployeeEmploymentDetailsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) WithBody(body *models.EmploymentDetails) *EmployeeEmploymentDetailsPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) SetBody(body *models.EmploymentDetails) {
	o.Body = body
}

// WithID adds the id to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) WithID(id int32) *EmployeeEmploymentDetailsPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employee employment details put params
func (o *EmployeeEmploymentDetailsPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEmploymentDetailsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
