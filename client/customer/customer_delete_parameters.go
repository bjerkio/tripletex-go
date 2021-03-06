// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// NewCustomerDeleteParams creates a new CustomerDeleteParams object
// with the default values initialized.
func NewCustomerDeleteParams() *CustomerDeleteParams {
	var ()
	return &CustomerDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerDeleteParamsWithTimeout creates a new CustomerDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerDeleteParamsWithTimeout(timeout time.Duration) *CustomerDeleteParams {
	var ()
	return &CustomerDeleteParams{

		timeout: timeout,
	}
}

// NewCustomerDeleteParamsWithContext creates a new CustomerDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerDeleteParamsWithContext(ctx context.Context) *CustomerDeleteParams {
	var ()
	return &CustomerDeleteParams{

		Context: ctx,
	}
}

// NewCustomerDeleteParamsWithHTTPClient creates a new CustomerDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerDeleteParamsWithHTTPClient(client *http.Client) *CustomerDeleteParams {
	var ()
	return &CustomerDeleteParams{
		HTTPClient: client,
	}
}

/*CustomerDeleteParams contains all the parameters to send to the API endpoint
for the customer delete operation typically these are written to a http.Request
*/
type CustomerDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer delete params
func (o *CustomerDeleteParams) WithTimeout(timeout time.Duration) *CustomerDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer delete params
func (o *CustomerDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer delete params
func (o *CustomerDeleteParams) WithContext(ctx context.Context) *CustomerDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer delete params
func (o *CustomerDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer delete params
func (o *CustomerDeleteParams) WithHTTPClient(client *http.Client) *CustomerDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer delete params
func (o *CustomerDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer delete params
func (o *CustomerDeleteParams) WithID(id int32) *CustomerDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer delete params
func (o *CustomerDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
