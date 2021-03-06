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

// NewSupplierInvoicePdfDownloadPdfParams creates a new SupplierInvoicePdfDownloadPdfParams object
// with the default values initialized.
func NewSupplierInvoicePdfDownloadPdfParams() *SupplierInvoicePdfDownloadPdfParams {
	var ()
	return &SupplierInvoicePdfDownloadPdfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierInvoicePdfDownloadPdfParamsWithTimeout creates a new SupplierInvoicePdfDownloadPdfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierInvoicePdfDownloadPdfParamsWithTimeout(timeout time.Duration) *SupplierInvoicePdfDownloadPdfParams {
	var ()
	return &SupplierInvoicePdfDownloadPdfParams{

		timeout: timeout,
	}
}

// NewSupplierInvoicePdfDownloadPdfParamsWithContext creates a new SupplierInvoicePdfDownloadPdfParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierInvoicePdfDownloadPdfParamsWithContext(ctx context.Context) *SupplierInvoicePdfDownloadPdfParams {
	var ()
	return &SupplierInvoicePdfDownloadPdfParams{

		Context: ctx,
	}
}

// NewSupplierInvoicePdfDownloadPdfParamsWithHTTPClient creates a new SupplierInvoicePdfDownloadPdfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierInvoicePdfDownloadPdfParamsWithHTTPClient(client *http.Client) *SupplierInvoicePdfDownloadPdfParams {
	var ()
	return &SupplierInvoicePdfDownloadPdfParams{
		HTTPClient: client,
	}
}

/*SupplierInvoicePdfDownloadPdfParams contains all the parameters to send to the API endpoint
for the supplier invoice pdf download pdf operation typically these are written to a http.Request
*/
type SupplierInvoicePdfDownloadPdfParams struct {

	/*InvoiceID
	  Invoice ID from which document is downloaded.

	*/
	InvoiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) WithTimeout(timeout time.Duration) *SupplierInvoicePdfDownloadPdfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) WithContext(ctx context.Context) *SupplierInvoicePdfDownloadPdfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) WithHTTPClient(client *http.Client) *SupplierInvoicePdfDownloadPdfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceID adds the invoiceID to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) WithInvoiceID(invoiceID int32) *SupplierInvoicePdfDownloadPdfParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the supplier invoice pdf download pdf params
func (o *SupplierInvoicePdfDownloadPdfParams) SetInvoiceID(invoiceID int32) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierInvoicePdfDownloadPdfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", swag.FormatInt32(o.InvoiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
