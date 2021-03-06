// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewProductGroupGetParams creates a new ProductGroupGetParams object
// with the default values initialized.
func NewProductGroupGetParams() *ProductGroupGetParams {
	var ()
	return &ProductGroupGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupGetParamsWithTimeout creates a new ProductGroupGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupGetParamsWithTimeout(timeout time.Duration) *ProductGroupGetParams {
	var ()
	return &ProductGroupGetParams{

		timeout: timeout,
	}
}

// NewProductGroupGetParamsWithContext creates a new ProductGroupGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupGetParamsWithContext(ctx context.Context) *ProductGroupGetParams {
	var ()
	return &ProductGroupGetParams{

		Context: ctx,
	}
}

// NewProductGroupGetParamsWithHTTPClient creates a new ProductGroupGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupGetParamsWithHTTPClient(client *http.Client) *ProductGroupGetParams {
	var ()
	return &ProductGroupGetParams{
		HTTPClient: client,
	}
}

/*ProductGroupGetParams contains all the parameters to send to the API endpoint
for the product group get operation typically these are written to a http.Request
*/
type ProductGroupGetParams struct {

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

// WithTimeout adds the timeout to the product group get params
func (o *ProductGroupGetParams) WithTimeout(timeout time.Duration) *ProductGroupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group get params
func (o *ProductGroupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group get params
func (o *ProductGroupGetParams) WithContext(ctx context.Context) *ProductGroupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group get params
func (o *ProductGroupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group get params
func (o *ProductGroupGetParams) WithHTTPClient(client *http.Client) *ProductGroupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group get params
func (o *ProductGroupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the product group get params
func (o *ProductGroupGetParams) WithFields(fields *string) *ProductGroupGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the product group get params
func (o *ProductGroupGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the product group get params
func (o *ProductGroupGetParams) WithID(id int32) *ProductGroupGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the product group get params
func (o *ProductGroupGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
