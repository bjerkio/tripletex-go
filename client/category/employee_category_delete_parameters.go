// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewEmployeeCategoryDeleteParams creates a new EmployeeCategoryDeleteParams object
// with the default values initialized.
func NewEmployeeCategoryDeleteParams() *EmployeeCategoryDeleteParams {
	var ()
	return &EmployeeCategoryDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeCategoryDeleteParamsWithTimeout creates a new EmployeeCategoryDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeCategoryDeleteParamsWithTimeout(timeout time.Duration) *EmployeeCategoryDeleteParams {
	var ()
	return &EmployeeCategoryDeleteParams{

		timeout: timeout,
	}
}

// NewEmployeeCategoryDeleteParamsWithContext creates a new EmployeeCategoryDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeCategoryDeleteParamsWithContext(ctx context.Context) *EmployeeCategoryDeleteParams {
	var ()
	return &EmployeeCategoryDeleteParams{

		Context: ctx,
	}
}

// NewEmployeeCategoryDeleteParamsWithHTTPClient creates a new EmployeeCategoryDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeCategoryDeleteParamsWithHTTPClient(client *http.Client) *EmployeeCategoryDeleteParams {
	var ()
	return &EmployeeCategoryDeleteParams{
		HTTPClient: client,
	}
}

/*EmployeeCategoryDeleteParams contains all the parameters to send to the API endpoint
for the employee category delete operation typically these are written to a http.Request
*/
type EmployeeCategoryDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee category delete params
func (o *EmployeeCategoryDeleteParams) WithTimeout(timeout time.Duration) *EmployeeCategoryDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee category delete params
func (o *EmployeeCategoryDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee category delete params
func (o *EmployeeCategoryDeleteParams) WithContext(ctx context.Context) *EmployeeCategoryDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee category delete params
func (o *EmployeeCategoryDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee category delete params
func (o *EmployeeCategoryDeleteParams) WithHTTPClient(client *http.Client) *EmployeeCategoryDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee category delete params
func (o *EmployeeCategoryDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the employee category delete params
func (o *EmployeeCategoryDeleteParams) WithID(id int32) *EmployeeCategoryDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the employee category delete params
func (o *EmployeeCategoryDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeCategoryDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
