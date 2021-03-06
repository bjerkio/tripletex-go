// Code generated by go-swagger; DO NOT EDIT.

package supplier_invoice

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

// NewSupplierInvoiceGetParams creates a new SupplierInvoiceGetParams object
// with the default values initialized.
func NewSupplierInvoiceGetParams() *SupplierInvoiceGetParams {
	var ()
	return &SupplierInvoiceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierInvoiceGetParamsWithTimeout creates a new SupplierInvoiceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierInvoiceGetParamsWithTimeout(timeout time.Duration) *SupplierInvoiceGetParams {
	var ()
	return &SupplierInvoiceGetParams{

		timeout: timeout,
	}
}

// NewSupplierInvoiceGetParamsWithContext creates a new SupplierInvoiceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierInvoiceGetParamsWithContext(ctx context.Context) *SupplierInvoiceGetParams {
	var ()
	return &SupplierInvoiceGetParams{

		Context: ctx,
	}
}

// NewSupplierInvoiceGetParamsWithHTTPClient creates a new SupplierInvoiceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierInvoiceGetParamsWithHTTPClient(client *http.Client) *SupplierInvoiceGetParams {
	var ()
	return &SupplierInvoiceGetParams{
		HTTPClient: client,
	}
}

/*SupplierInvoiceGetParams contains all the parameters to send to the API endpoint
for the supplier invoice get operation typically these are written to a http.Request
*/
type SupplierInvoiceGetParams struct {

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

// WithTimeout adds the timeout to the supplier invoice get params
func (o *SupplierInvoiceGetParams) WithTimeout(timeout time.Duration) *SupplierInvoiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier invoice get params
func (o *SupplierInvoiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier invoice get params
func (o *SupplierInvoiceGetParams) WithContext(ctx context.Context) *SupplierInvoiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier invoice get params
func (o *SupplierInvoiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier invoice get params
func (o *SupplierInvoiceGetParams) WithHTTPClient(client *http.Client) *SupplierInvoiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier invoice get params
func (o *SupplierInvoiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the supplier invoice get params
func (o *SupplierInvoiceGetParams) WithFields(fields *string) *SupplierInvoiceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the supplier invoice get params
func (o *SupplierInvoiceGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the supplier invoice get params
func (o *SupplierInvoiceGetParams) WithID(id int32) *SupplierInvoiceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the supplier invoice get params
func (o *SupplierInvoiceGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierInvoiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
