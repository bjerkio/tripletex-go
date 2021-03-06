// Code generated by go-swagger; DO NOT EDIT.

package account

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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewLedgerAccountPutParams creates a new LedgerAccountPutParams object
// with the default values initialized.
func NewLedgerAccountPutParams() *LedgerAccountPutParams {
	var ()
	return &LedgerAccountPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerAccountPutParamsWithTimeout creates a new LedgerAccountPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerAccountPutParamsWithTimeout(timeout time.Duration) *LedgerAccountPutParams {
	var ()
	return &LedgerAccountPutParams{

		timeout: timeout,
	}
}

// NewLedgerAccountPutParamsWithContext creates a new LedgerAccountPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerAccountPutParamsWithContext(ctx context.Context) *LedgerAccountPutParams {
	var ()
	return &LedgerAccountPutParams{

		Context: ctx,
	}
}

// NewLedgerAccountPutParamsWithHTTPClient creates a new LedgerAccountPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerAccountPutParamsWithHTTPClient(client *http.Client) *LedgerAccountPutParams {
	var ()
	return &LedgerAccountPutParams{
		HTTPClient: client,
	}
}

/*LedgerAccountPutParams contains all the parameters to send to the API endpoint
for the ledger account put operation typically these are written to a http.Request
*/
type LedgerAccountPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Account
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger account put params
func (o *LedgerAccountPutParams) WithTimeout(timeout time.Duration) *LedgerAccountPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger account put params
func (o *LedgerAccountPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger account put params
func (o *LedgerAccountPutParams) WithContext(ctx context.Context) *LedgerAccountPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger account put params
func (o *LedgerAccountPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger account put params
func (o *LedgerAccountPutParams) WithHTTPClient(client *http.Client) *LedgerAccountPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger account put params
func (o *LedgerAccountPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ledger account put params
func (o *LedgerAccountPutParams) WithBody(body *models.Account) *LedgerAccountPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ledger account put params
func (o *LedgerAccountPutParams) SetBody(body *models.Account) {
	o.Body = body
}

// WithID adds the id to the ledger account put params
func (o *LedgerAccountPutParams) WithID(id int32) *LedgerAccountPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger account put params
func (o *LedgerAccountPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerAccountPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
