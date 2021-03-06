// Code generated by go-swagger; DO NOT EDIT.

package voucher

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

// NewLedgerVoucherSendToInboxSendToInboxParams creates a new LedgerVoucherSendToInboxSendToInboxParams object
// with the default values initialized.
func NewLedgerVoucherSendToInboxSendToInboxParams() *LedgerVoucherSendToInboxSendToInboxParams {
	var ()
	return &LedgerVoucherSendToInboxSendToInboxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerVoucherSendToInboxSendToInboxParamsWithTimeout creates a new LedgerVoucherSendToInboxSendToInboxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerVoucherSendToInboxSendToInboxParamsWithTimeout(timeout time.Duration) *LedgerVoucherSendToInboxSendToInboxParams {
	var ()
	return &LedgerVoucherSendToInboxSendToInboxParams{

		timeout: timeout,
	}
}

// NewLedgerVoucherSendToInboxSendToInboxParamsWithContext creates a new LedgerVoucherSendToInboxSendToInboxParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerVoucherSendToInboxSendToInboxParamsWithContext(ctx context.Context) *LedgerVoucherSendToInboxSendToInboxParams {
	var ()
	return &LedgerVoucherSendToInboxSendToInboxParams{

		Context: ctx,
	}
}

// NewLedgerVoucherSendToInboxSendToInboxParamsWithHTTPClient creates a new LedgerVoucherSendToInboxSendToInboxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerVoucherSendToInboxSendToInboxParamsWithHTTPClient(client *http.Client) *LedgerVoucherSendToInboxSendToInboxParams {
	var ()
	return &LedgerVoucherSendToInboxSendToInboxParams{
		HTTPClient: client,
	}
}

/*LedgerVoucherSendToInboxSendToInboxParams contains all the parameters to send to the API endpoint
for the ledger voucher send to inbox send to inbox operation typically these are written to a http.Request
*/
type LedgerVoucherSendToInboxSendToInboxParams struct {

	/*Comment
	  Description of why the voucher was rejected. This parameter is only used if the approval feature is enabled.

	*/
	Comment *string
	/*ID
	  ID of voucher that should be sent to inbox.

	*/
	ID int32
	/*Version
	  Version of voucher that should be sent to inbox.

	*/
	Version *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithTimeout(timeout time.Duration) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithContext(ctx context.Context) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithHTTPClient(client *http.Client) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithComment(comment *string) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithID adds the id to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithID(id int32) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetID(id int32) {
	o.ID = id
}

// WithVersion adds the version to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) WithVersion(version *int32) *LedgerVoucherSendToInboxSendToInboxParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the ledger voucher send to inbox send to inbox params
func (o *LedgerVoucherSendToInboxSendToInboxParams) SetVersion(version *int32) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerVoucherSendToInboxSendToInboxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int32
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt32(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
