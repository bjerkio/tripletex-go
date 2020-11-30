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
)

// NewSupplierInvoiceRejectRejectManyParams creates a new SupplierInvoiceRejectRejectManyParams object
// with the default values initialized.
func NewSupplierInvoiceRejectRejectManyParams() *SupplierInvoiceRejectRejectManyParams {
	var ()
	return &SupplierInvoiceRejectRejectManyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierInvoiceRejectRejectManyParamsWithTimeout creates a new SupplierInvoiceRejectRejectManyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierInvoiceRejectRejectManyParamsWithTimeout(timeout time.Duration) *SupplierInvoiceRejectRejectManyParams {
	var ()
	return &SupplierInvoiceRejectRejectManyParams{

		timeout: timeout,
	}
}

// NewSupplierInvoiceRejectRejectManyParamsWithContext creates a new SupplierInvoiceRejectRejectManyParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierInvoiceRejectRejectManyParamsWithContext(ctx context.Context) *SupplierInvoiceRejectRejectManyParams {
	var ()
	return &SupplierInvoiceRejectRejectManyParams{

		Context: ctx,
	}
}

// NewSupplierInvoiceRejectRejectManyParamsWithHTTPClient creates a new SupplierInvoiceRejectRejectManyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierInvoiceRejectRejectManyParamsWithHTTPClient(client *http.Client) *SupplierInvoiceRejectRejectManyParams {
	var ()
	return &SupplierInvoiceRejectRejectManyParams{
		HTTPClient: client,
	}
}

/*SupplierInvoiceRejectRejectManyParams contains all the parameters to send to the API endpoint
for the supplier invoice reject reject many operation typically these are written to a http.Request
*/
type SupplierInvoiceRejectRejectManyParams struct {

	/*Comment*/
	Comment string
	/*InvoiceIds
	  ID of the elements

	*/
	InvoiceIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) WithTimeout(timeout time.Duration) *SupplierInvoiceRejectRejectManyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) WithContext(ctx context.Context) *SupplierInvoiceRejectRejectManyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) WithHTTPClient(client *http.Client) *SupplierInvoiceRejectRejectManyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) WithComment(comment string) *SupplierInvoiceRejectRejectManyParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) SetComment(comment string) {
	o.Comment = comment
}

// WithInvoiceIds adds the invoiceIds to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) WithInvoiceIds(invoiceIds *string) *SupplierInvoiceRejectRejectManyParams {
	o.SetInvoiceIds(invoiceIds)
	return o
}

// SetInvoiceIds adds the invoiceIds to the supplier invoice reject reject many params
func (o *SupplierInvoiceRejectRejectManyParams) SetInvoiceIds(invoiceIds *string) {
	o.InvoiceIds = invoiceIds
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierInvoiceRejectRejectManyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param comment
	qrComment := o.Comment
	qComment := qrComment
	if qComment != "" {
		if err := r.SetQueryParam("comment", qComment); err != nil {
			return err
		}
	}

	if o.InvoiceIds != nil {

		// query param invoiceIds
		var qrInvoiceIds string
		if o.InvoiceIds != nil {
			qrInvoiceIds = *o.InvoiceIds
		}
		qInvoiceIds := qrInvoiceIds
		if qInvoiceIds != "" {
			if err := r.SetQueryParam("invoiceIds", qInvoiceIds); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}