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

// NewSupplierInvoiceAddRecipientAddRecipientParams creates a new SupplierInvoiceAddRecipientAddRecipientParams object
// with the default values initialized.
func NewSupplierInvoiceAddRecipientAddRecipientParams() *SupplierInvoiceAddRecipientAddRecipientParams {
	var ()
	return &SupplierInvoiceAddRecipientAddRecipientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierInvoiceAddRecipientAddRecipientParamsWithTimeout creates a new SupplierInvoiceAddRecipientAddRecipientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierInvoiceAddRecipientAddRecipientParamsWithTimeout(timeout time.Duration) *SupplierInvoiceAddRecipientAddRecipientParams {
	var ()
	return &SupplierInvoiceAddRecipientAddRecipientParams{

		timeout: timeout,
	}
}

// NewSupplierInvoiceAddRecipientAddRecipientParamsWithContext creates a new SupplierInvoiceAddRecipientAddRecipientParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierInvoiceAddRecipientAddRecipientParamsWithContext(ctx context.Context) *SupplierInvoiceAddRecipientAddRecipientParams {
	var ()
	return &SupplierInvoiceAddRecipientAddRecipientParams{

		Context: ctx,
	}
}

// NewSupplierInvoiceAddRecipientAddRecipientParamsWithHTTPClient creates a new SupplierInvoiceAddRecipientAddRecipientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierInvoiceAddRecipientAddRecipientParamsWithHTTPClient(client *http.Client) *SupplierInvoiceAddRecipientAddRecipientParams {
	var ()
	return &SupplierInvoiceAddRecipientAddRecipientParams{
		HTTPClient: client,
	}
}

/*SupplierInvoiceAddRecipientAddRecipientParams contains all the parameters to send to the API endpoint
for the supplier invoice add recipient add recipient operation typically these are written to a http.Request
*/
type SupplierInvoiceAddRecipientAddRecipientParams struct {

	/*Comment
	  comment

	*/
	Comment *string
	/*EmployeeID
	  ID of the elements

	*/
	EmployeeID int32
	/*InvoiceID
	  Invoice ID.

	*/
	InvoiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithTimeout(timeout time.Duration) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithContext(ctx context.Context) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithHTTPClient(client *http.Client) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithComment(comment *string) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithEmployeeID adds the employeeID to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithEmployeeID(employeeID int32) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetEmployeeID(employeeID int32) {
	o.EmployeeID = employeeID
}

// WithInvoiceID adds the invoiceID to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WithInvoiceID(invoiceID int32) *SupplierInvoiceAddRecipientAddRecipientParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the supplier invoice add recipient add recipient params
func (o *SupplierInvoiceAddRecipientAddRecipientParams) SetInvoiceID(invoiceID int32) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierInvoiceAddRecipientAddRecipientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string
		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {
			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}

	}

	// query param employeeId
	qrEmployeeID := o.EmployeeID
	qEmployeeID := swag.FormatInt32(qrEmployeeID)
	if qEmployeeID != "" {
		if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", swag.FormatInt32(o.InvoiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
