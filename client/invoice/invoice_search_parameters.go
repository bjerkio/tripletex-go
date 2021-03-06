// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewInvoiceSearchParams creates a new InvoiceSearchParams object
// with the default values initialized.
func NewInvoiceSearchParams() *InvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInvoiceSearchParamsWithTimeout creates a new InvoiceSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvoiceSearchParamsWithTimeout(timeout time.Duration) *InvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewInvoiceSearchParamsWithContext creates a new InvoiceSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvoiceSearchParamsWithContext(ctx context.Context) *InvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewInvoiceSearchParamsWithHTTPClient creates a new InvoiceSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvoiceSearchParamsWithHTTPClient(client *http.Client) *InvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &InvoiceSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*InvoiceSearchParams contains all the parameters to send to the API endpoint
for the invoice search operation typically these are written to a http.Request
*/
type InvoiceSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  Equals

	*/
	CustomerID *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*ID
	  List of IDs

	*/
	ID *string
	/*InvoiceDateFrom
	  From and including

	*/
	InvoiceDateFrom string
	/*InvoiceDateTo
	  To and excluding

	*/
	InvoiceDateTo string
	/*InvoiceNumber
	  Equals

	*/
	InvoiceNumber *string
	/*Kid
	  Equals

	*/
	Kid *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*VoucherID
	  Equals

	*/
	VoucherID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invoice search params
func (o *InvoiceSearchParams) WithTimeout(timeout time.Duration) *InvoiceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoice search params
func (o *InvoiceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoice search params
func (o *InvoiceSearchParams) WithContext(ctx context.Context) *InvoiceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoice search params
func (o *InvoiceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invoice search params
func (o *InvoiceSearchParams) WithHTTPClient(client *http.Client) *InvoiceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invoice search params
func (o *InvoiceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the invoice search params
func (o *InvoiceSearchParams) WithCount(count *int64) *InvoiceSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the invoice search params
func (o *InvoiceSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the invoice search params
func (o *InvoiceSearchParams) WithCustomerID(customerID *string) *InvoiceSearchParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the invoice search params
func (o *InvoiceSearchParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithFields adds the fields to the invoice search params
func (o *InvoiceSearchParams) WithFields(fields *string) *InvoiceSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the invoice search params
func (o *InvoiceSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the invoice search params
func (o *InvoiceSearchParams) WithFrom(from *int64) *InvoiceSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the invoice search params
func (o *InvoiceSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the invoice search params
func (o *InvoiceSearchParams) WithID(id *string) *InvoiceSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the invoice search params
func (o *InvoiceSearchParams) SetID(id *string) {
	o.ID = id
}

// WithInvoiceDateFrom adds the invoiceDateFrom to the invoice search params
func (o *InvoiceSearchParams) WithInvoiceDateFrom(invoiceDateFrom string) *InvoiceSearchParams {
	o.SetInvoiceDateFrom(invoiceDateFrom)
	return o
}

// SetInvoiceDateFrom adds the invoiceDateFrom to the invoice search params
func (o *InvoiceSearchParams) SetInvoiceDateFrom(invoiceDateFrom string) {
	o.InvoiceDateFrom = invoiceDateFrom
}

// WithInvoiceDateTo adds the invoiceDateTo to the invoice search params
func (o *InvoiceSearchParams) WithInvoiceDateTo(invoiceDateTo string) *InvoiceSearchParams {
	o.SetInvoiceDateTo(invoiceDateTo)
	return o
}

// SetInvoiceDateTo adds the invoiceDateTo to the invoice search params
func (o *InvoiceSearchParams) SetInvoiceDateTo(invoiceDateTo string) {
	o.InvoiceDateTo = invoiceDateTo
}

// WithInvoiceNumber adds the invoiceNumber to the invoice search params
func (o *InvoiceSearchParams) WithInvoiceNumber(invoiceNumber *string) *InvoiceSearchParams {
	o.SetInvoiceNumber(invoiceNumber)
	return o
}

// SetInvoiceNumber adds the invoiceNumber to the invoice search params
func (o *InvoiceSearchParams) SetInvoiceNumber(invoiceNumber *string) {
	o.InvoiceNumber = invoiceNumber
}

// WithKid adds the kid to the invoice search params
func (o *InvoiceSearchParams) WithKid(kid *string) *InvoiceSearchParams {
	o.SetKid(kid)
	return o
}

// SetKid adds the kid to the invoice search params
func (o *InvoiceSearchParams) SetKid(kid *string) {
	o.Kid = kid
}

// WithSorting adds the sorting to the invoice search params
func (o *InvoiceSearchParams) WithSorting(sorting *string) *InvoiceSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the invoice search params
func (o *InvoiceSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithVoucherID adds the voucherID to the invoice search params
func (o *InvoiceSearchParams) WithVoucherID(voucherID *string) *InvoiceSearchParams {
	o.SetVoucherID(voucherID)
	return o
}

// SetVoucherID adds the voucherId to the invoice search params
func (o *InvoiceSearchParams) SetVoucherID(voucherID *string) {
	o.VoucherID = voucherID
}

// WriteToRequest writes these params to a swagger request
func (o *InvoiceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.CustomerID != nil {

		// query param customerId
		var qrCustomerID string
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := qrCustomerID
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

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

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	// query param invoiceDateFrom
	qrInvoiceDateFrom := o.InvoiceDateFrom
	qInvoiceDateFrom := qrInvoiceDateFrom
	if qInvoiceDateFrom != "" {
		if err := r.SetQueryParam("invoiceDateFrom", qInvoiceDateFrom); err != nil {
			return err
		}
	}

	// query param invoiceDateTo
	qrInvoiceDateTo := o.InvoiceDateTo
	qInvoiceDateTo := qrInvoiceDateTo
	if qInvoiceDateTo != "" {
		if err := r.SetQueryParam("invoiceDateTo", qInvoiceDateTo); err != nil {
			return err
		}
	}

	if o.InvoiceNumber != nil {

		// query param invoiceNumber
		var qrInvoiceNumber string
		if o.InvoiceNumber != nil {
			qrInvoiceNumber = *o.InvoiceNumber
		}
		qInvoiceNumber := qrInvoiceNumber
		if qInvoiceNumber != "" {
			if err := r.SetQueryParam("invoiceNumber", qInvoiceNumber); err != nil {
				return err
			}
		}

	}

	if o.Kid != nil {

		// query param kid
		var qrKid string
		if o.Kid != nil {
			qrKid = *o.Kid
		}
		qKid := qrKid
		if qKid != "" {
			if err := r.SetQueryParam("kid", qKid); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if o.VoucherID != nil {

		// query param voucherId
		var qrVoucherID string
		if o.VoucherID != nil {
			qrVoucherID = *o.VoucherID
		}
		qVoucherID := qrVoucherID
		if qVoucherID != "" {
			if err := r.SetQueryParam("voucherId", qVoucherID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
