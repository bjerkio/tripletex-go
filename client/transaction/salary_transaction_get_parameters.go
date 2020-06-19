// Code generated by go-swagger; DO NOT EDIT.

package transaction

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

// NewSalaryTransactionGetParams creates a new SalaryTransactionGetParams object
// with the default values initialized.
func NewSalaryTransactionGetParams() *SalaryTransactionGetParams {
	var ()
	return &SalaryTransactionGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalaryTransactionGetParamsWithTimeout creates a new SalaryTransactionGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalaryTransactionGetParamsWithTimeout(timeout time.Duration) *SalaryTransactionGetParams {
	var ()
	return &SalaryTransactionGetParams{

		timeout: timeout,
	}
}

// NewSalaryTransactionGetParamsWithContext creates a new SalaryTransactionGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalaryTransactionGetParamsWithContext(ctx context.Context) *SalaryTransactionGetParams {
	var ()
	return &SalaryTransactionGetParams{

		Context: ctx,
	}
}

// NewSalaryTransactionGetParamsWithHTTPClient creates a new SalaryTransactionGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalaryTransactionGetParamsWithHTTPClient(client *http.Client) *SalaryTransactionGetParams {
	var ()
	return &SalaryTransactionGetParams{
		HTTPClient: client,
	}
}

/*SalaryTransactionGetParams contains all the parameters to send to the API endpoint
for the salary transaction get operation typically these are written to a http.Request
*/
type SalaryTransactionGetParams struct {

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

// WithTimeout adds the timeout to the salary transaction get params
func (o *SalaryTransactionGetParams) WithTimeout(timeout time.Duration) *SalaryTransactionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salary transaction get params
func (o *SalaryTransactionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salary transaction get params
func (o *SalaryTransactionGetParams) WithContext(ctx context.Context) *SalaryTransactionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salary transaction get params
func (o *SalaryTransactionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salary transaction get params
func (o *SalaryTransactionGetParams) WithHTTPClient(client *http.Client) *SalaryTransactionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salary transaction get params
func (o *SalaryTransactionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the salary transaction get params
func (o *SalaryTransactionGetParams) WithFields(fields *string) *SalaryTransactionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the salary transaction get params
func (o *SalaryTransactionGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the salary transaction get params
func (o *SalaryTransactionGetParams) WithID(id int32) *SalaryTransactionGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the salary transaction get params
func (o *SalaryTransactionGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalaryTransactionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
