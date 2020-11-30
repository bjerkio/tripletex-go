// Code generated by go-swagger; DO NOT EDIT.

package payment_type_out

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

// NewLedgerPaymentTypeOutGetParams creates a new LedgerPaymentTypeOutGetParams object
// with the default values initialized.
func NewLedgerPaymentTypeOutGetParams() *LedgerPaymentTypeOutGetParams {
	var ()
	return &LedgerPaymentTypeOutGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerPaymentTypeOutGetParamsWithTimeout creates a new LedgerPaymentTypeOutGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerPaymentTypeOutGetParamsWithTimeout(timeout time.Duration) *LedgerPaymentTypeOutGetParams {
	var ()
	return &LedgerPaymentTypeOutGetParams{

		timeout: timeout,
	}
}

// NewLedgerPaymentTypeOutGetParamsWithContext creates a new LedgerPaymentTypeOutGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerPaymentTypeOutGetParamsWithContext(ctx context.Context) *LedgerPaymentTypeOutGetParams {
	var ()
	return &LedgerPaymentTypeOutGetParams{

		Context: ctx,
	}
}

// NewLedgerPaymentTypeOutGetParamsWithHTTPClient creates a new LedgerPaymentTypeOutGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerPaymentTypeOutGetParamsWithHTTPClient(client *http.Client) *LedgerPaymentTypeOutGetParams {
	var ()
	return &LedgerPaymentTypeOutGetParams{
		HTTPClient: client,
	}
}

/*LedgerPaymentTypeOutGetParams contains all the parameters to send to the API endpoint
for the ledger payment type out get operation typically these are written to a http.Request
*/
type LedgerPaymentTypeOutGetParams struct {

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

// WithTimeout adds the timeout to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) WithTimeout(timeout time.Duration) *LedgerPaymentTypeOutGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) WithContext(ctx context.Context) *LedgerPaymentTypeOutGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) WithHTTPClient(client *http.Client) *LedgerPaymentTypeOutGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) WithFields(fields *string) *LedgerPaymentTypeOutGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) WithID(id int32) *LedgerPaymentTypeOutGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger payment type out get params
func (o *LedgerPaymentTypeOutGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerPaymentTypeOutGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
