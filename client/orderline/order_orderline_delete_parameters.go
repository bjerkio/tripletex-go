// Code generated by go-swagger; DO NOT EDIT.

package orderline

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

// NewOrderOrderlineDeleteParams creates a new OrderOrderlineDeleteParams object
// with the default values initialized.
func NewOrderOrderlineDeleteParams() *OrderOrderlineDeleteParams {
	var ()
	return &OrderOrderlineDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderOrderlineDeleteParamsWithTimeout creates a new OrderOrderlineDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderOrderlineDeleteParamsWithTimeout(timeout time.Duration) *OrderOrderlineDeleteParams {
	var ()
	return &OrderOrderlineDeleteParams{

		timeout: timeout,
	}
}

// NewOrderOrderlineDeleteParamsWithContext creates a new OrderOrderlineDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderOrderlineDeleteParamsWithContext(ctx context.Context) *OrderOrderlineDeleteParams {
	var ()
	return &OrderOrderlineDeleteParams{

		Context: ctx,
	}
}

// NewOrderOrderlineDeleteParamsWithHTTPClient creates a new OrderOrderlineDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderOrderlineDeleteParamsWithHTTPClient(client *http.Client) *OrderOrderlineDeleteParams {
	var ()
	return &OrderOrderlineDeleteParams{
		HTTPClient: client,
	}
}

/*OrderOrderlineDeleteParams contains all the parameters to send to the API endpoint
for the order orderline delete operation typically these are written to a http.Request
*/
type OrderOrderlineDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order orderline delete params
func (o *OrderOrderlineDeleteParams) WithTimeout(timeout time.Duration) *OrderOrderlineDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order orderline delete params
func (o *OrderOrderlineDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order orderline delete params
func (o *OrderOrderlineDeleteParams) WithContext(ctx context.Context) *OrderOrderlineDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order orderline delete params
func (o *OrderOrderlineDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order orderline delete params
func (o *OrderOrderlineDeleteParams) WithHTTPClient(client *http.Client) *OrderOrderlineDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order orderline delete params
func (o *OrderOrderlineDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the order orderline delete params
func (o *OrderOrderlineDeleteParams) WithID(id int32) *OrderOrderlineDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the order orderline delete params
func (o *OrderOrderlineDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrderOrderlineDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
