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

// NewOrderOrderlineGetParams creates a new OrderOrderlineGetParams object
// with the default values initialized.
func NewOrderOrderlineGetParams() *OrderOrderlineGetParams {
	var ()
	return &OrderOrderlineGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderOrderlineGetParamsWithTimeout creates a new OrderOrderlineGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderOrderlineGetParamsWithTimeout(timeout time.Duration) *OrderOrderlineGetParams {
	var ()
	return &OrderOrderlineGetParams{

		timeout: timeout,
	}
}

// NewOrderOrderlineGetParamsWithContext creates a new OrderOrderlineGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderOrderlineGetParamsWithContext(ctx context.Context) *OrderOrderlineGetParams {
	var ()
	return &OrderOrderlineGetParams{

		Context: ctx,
	}
}

// NewOrderOrderlineGetParamsWithHTTPClient creates a new OrderOrderlineGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderOrderlineGetParamsWithHTTPClient(client *http.Client) *OrderOrderlineGetParams {
	var ()
	return &OrderOrderlineGetParams{
		HTTPClient: client,
	}
}

/*OrderOrderlineGetParams contains all the parameters to send to the API endpoint
for the order orderline get operation typically these are written to a http.Request
*/
type OrderOrderlineGetParams struct {

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

// WithTimeout adds the timeout to the order orderline get params
func (o *OrderOrderlineGetParams) WithTimeout(timeout time.Duration) *OrderOrderlineGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order orderline get params
func (o *OrderOrderlineGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order orderline get params
func (o *OrderOrderlineGetParams) WithContext(ctx context.Context) *OrderOrderlineGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order orderline get params
func (o *OrderOrderlineGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order orderline get params
func (o *OrderOrderlineGetParams) WithHTTPClient(client *http.Client) *OrderOrderlineGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order orderline get params
func (o *OrderOrderlineGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the order orderline get params
func (o *OrderOrderlineGetParams) WithFields(fields *string) *OrderOrderlineGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the order orderline get params
func (o *OrderOrderlineGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the order orderline get params
func (o *OrderOrderlineGetParams) WithID(id int32) *OrderOrderlineGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the order orderline get params
func (o *OrderOrderlineGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrderOrderlineGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
