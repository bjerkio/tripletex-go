// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// NewSupplierInvoiceSearchParams creates a new SupplierInvoiceSearchParams object
// with the default values initialized.
func NewSupplierInvoiceSearchParams() *SupplierInvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &SupplierInvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierInvoiceSearchParamsWithTimeout creates a new SupplierInvoiceSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierInvoiceSearchParamsWithTimeout(timeout time.Duration) *SupplierInvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &SupplierInvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewSupplierInvoiceSearchParamsWithContext creates a new SupplierInvoiceSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierInvoiceSearchParamsWithContext(ctx context.Context) *SupplierInvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &SupplierInvoiceSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewSupplierInvoiceSearchParamsWithHTTPClient creates a new SupplierInvoiceSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierInvoiceSearchParamsWithHTTPClient(client *http.Client) *SupplierInvoiceSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &SupplierInvoiceSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*SupplierInvoiceSearchParams contains all the parameters to send to the API endpoint
for the supplier invoice search operation typically these are written to a http.Request
*/
type SupplierInvoiceSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
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
	/*SupplierID
	  Equals

	*/
	SupplierID *string
	/*VoucherID
	  Equals

	*/
	VoucherID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithTimeout(timeout time.Duration) *SupplierInvoiceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithContext(ctx context.Context) *SupplierInvoiceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithHTTPClient(client *http.Client) *SupplierInvoiceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithCount(count *int64) *SupplierInvoiceSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithFields(fields *string) *SupplierInvoiceSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithFrom(from *int64) *SupplierInvoiceSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithID(id *string) *SupplierInvoiceSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetID(id *string) {
	o.ID = id
}

// WithInvoiceDateFrom adds the invoiceDateFrom to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithInvoiceDateFrom(invoiceDateFrom string) *SupplierInvoiceSearchParams {
	o.SetInvoiceDateFrom(invoiceDateFrom)
	return o
}

// SetInvoiceDateFrom adds the invoiceDateFrom to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetInvoiceDateFrom(invoiceDateFrom string) {
	o.InvoiceDateFrom = invoiceDateFrom
}

// WithInvoiceDateTo adds the invoiceDateTo to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithInvoiceDateTo(invoiceDateTo string) *SupplierInvoiceSearchParams {
	o.SetInvoiceDateTo(invoiceDateTo)
	return o
}

// SetInvoiceDateTo adds the invoiceDateTo to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetInvoiceDateTo(invoiceDateTo string) {
	o.InvoiceDateTo = invoiceDateTo
}

// WithInvoiceNumber adds the invoiceNumber to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithInvoiceNumber(invoiceNumber *string) *SupplierInvoiceSearchParams {
	o.SetInvoiceNumber(invoiceNumber)
	return o
}

// SetInvoiceNumber adds the invoiceNumber to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetInvoiceNumber(invoiceNumber *string) {
	o.InvoiceNumber = invoiceNumber
}

// WithKid adds the kid to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithKid(kid *string) *SupplierInvoiceSearchParams {
	o.SetKid(kid)
	return o
}

// SetKid adds the kid to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetKid(kid *string) {
	o.Kid = kid
}

// WithSorting adds the sorting to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithSorting(sorting *string) *SupplierInvoiceSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierID adds the supplierID to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithSupplierID(supplierID *string) *SupplierInvoiceSearchParams {
	o.SetSupplierID(supplierID)
	return o
}

// SetSupplierID adds the supplierId to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetSupplierID(supplierID *string) {
	o.SupplierID = supplierID
}

// WithVoucherID adds the voucherID to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) WithVoucherID(voucherID *string) *SupplierInvoiceSearchParams {
	o.SetVoucherID(voucherID)
	return o
}

// SetVoucherID adds the voucherId to the supplier invoice search params
func (o *SupplierInvoiceSearchParams) SetVoucherID(voucherID *string) {
	o.VoucherID = voucherID
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierInvoiceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SupplierID != nil {

		// query param supplierId
		var qrSupplierID string
		if o.SupplierID != nil {
			qrSupplierID = *o.SupplierID
		}
		qSupplierID := qrSupplierID
		if qSupplierID != "" {
			if err := r.SetQueryParam("supplierId", qSupplierID); err != nil {
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
