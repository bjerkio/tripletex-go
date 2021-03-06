// Code generated by go-swagger; DO NOT EDIT.

package posting

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

// NewLedgerPostingGetParams creates a new LedgerPostingGetParams object
// with the default values initialized.
func NewLedgerPostingGetParams() *LedgerPostingGetParams {
	var ()
	return &LedgerPostingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerPostingGetParamsWithTimeout creates a new LedgerPostingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerPostingGetParamsWithTimeout(timeout time.Duration) *LedgerPostingGetParams {
	var ()
	return &LedgerPostingGetParams{

		timeout: timeout,
	}
}

// NewLedgerPostingGetParamsWithContext creates a new LedgerPostingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerPostingGetParamsWithContext(ctx context.Context) *LedgerPostingGetParams {
	var ()
	return &LedgerPostingGetParams{

		Context: ctx,
	}
}

// NewLedgerPostingGetParamsWithHTTPClient creates a new LedgerPostingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerPostingGetParamsWithHTTPClient(client *http.Client) *LedgerPostingGetParams {
	var ()
	return &LedgerPostingGetParams{
		HTTPClient: client,
	}
}

/*LedgerPostingGetParams contains all the parameters to send to the API endpoint
for the ledger posting get operation typically these are written to a http.Request
*/
type LedgerPostingGetParams struct {

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

// WithTimeout adds the timeout to the ledger posting get params
func (o *LedgerPostingGetParams) WithTimeout(timeout time.Duration) *LedgerPostingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger posting get params
func (o *LedgerPostingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger posting get params
func (o *LedgerPostingGetParams) WithContext(ctx context.Context) *LedgerPostingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger posting get params
func (o *LedgerPostingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger posting get params
func (o *LedgerPostingGetParams) WithHTTPClient(client *http.Client) *LedgerPostingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger posting get params
func (o *LedgerPostingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ledger posting get params
func (o *LedgerPostingGetParams) WithFields(fields *string) *LedgerPostingGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ledger posting get params
func (o *LedgerPostingGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the ledger posting get params
func (o *LedgerPostingGetParams) WithID(id int32) *LedgerPostingGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger posting get params
func (o *LedgerPostingGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerPostingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
