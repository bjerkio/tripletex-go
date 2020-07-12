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

package reconciliation

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

// NewBankReconciliationGetParams creates a new BankReconciliationGetParams object
// with the default values initialized.
func NewBankReconciliationGetParams() *BankReconciliationGetParams {
	var ()
	return &BankReconciliationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBankReconciliationGetParamsWithTimeout creates a new BankReconciliationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankReconciliationGetParamsWithTimeout(timeout time.Duration) *BankReconciliationGetParams {
	var ()
	return &BankReconciliationGetParams{

		timeout: timeout,
	}
}

// NewBankReconciliationGetParamsWithContext creates a new BankReconciliationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankReconciliationGetParamsWithContext(ctx context.Context) *BankReconciliationGetParams {
	var ()
	return &BankReconciliationGetParams{

		Context: ctx,
	}
}

// NewBankReconciliationGetParamsWithHTTPClient creates a new BankReconciliationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankReconciliationGetParamsWithHTTPClient(client *http.Client) *BankReconciliationGetParams {
	var ()
	return &BankReconciliationGetParams{
		HTTPClient: client,
	}
}

/*BankReconciliationGetParams contains all the parameters to send to the API endpoint
for the bank reconciliation get operation typically these are written to a http.Request
*/
type BankReconciliationGetParams struct {

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

// WithTimeout adds the timeout to the bank reconciliation get params
func (o *BankReconciliationGetParams) WithTimeout(timeout time.Duration) *BankReconciliationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank reconciliation get params
func (o *BankReconciliationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank reconciliation get params
func (o *BankReconciliationGetParams) WithContext(ctx context.Context) *BankReconciliationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank reconciliation get params
func (o *BankReconciliationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank reconciliation get params
func (o *BankReconciliationGetParams) WithHTTPClient(client *http.Client) *BankReconciliationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank reconciliation get params
func (o *BankReconciliationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the bank reconciliation get params
func (o *BankReconciliationGetParams) WithFields(fields *string) *BankReconciliationGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the bank reconciliation get params
func (o *BankReconciliationGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the bank reconciliation get params
func (o *BankReconciliationGetParams) WithID(id int32) *BankReconciliationGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bank reconciliation get params
func (o *BankReconciliationGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BankReconciliationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
