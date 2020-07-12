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

package statement

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

// NewBankStatementDeleteParams creates a new BankStatementDeleteParams object
// with the default values initialized.
func NewBankStatementDeleteParams() *BankStatementDeleteParams {
	var ()
	return &BankStatementDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBankStatementDeleteParamsWithTimeout creates a new BankStatementDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankStatementDeleteParamsWithTimeout(timeout time.Duration) *BankStatementDeleteParams {
	var ()
	return &BankStatementDeleteParams{

		timeout: timeout,
	}
}

// NewBankStatementDeleteParamsWithContext creates a new BankStatementDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankStatementDeleteParamsWithContext(ctx context.Context) *BankStatementDeleteParams {
	var ()
	return &BankStatementDeleteParams{

		Context: ctx,
	}
}

// NewBankStatementDeleteParamsWithHTTPClient creates a new BankStatementDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankStatementDeleteParamsWithHTTPClient(client *http.Client) *BankStatementDeleteParams {
	var ()
	return &BankStatementDeleteParams{
		HTTPClient: client,
	}
}

/*BankStatementDeleteParams contains all the parameters to send to the API endpoint
for the bank statement delete operation typically these are written to a http.Request
*/
type BankStatementDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bank statement delete params
func (o *BankStatementDeleteParams) WithTimeout(timeout time.Duration) *BankStatementDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank statement delete params
func (o *BankStatementDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank statement delete params
func (o *BankStatementDeleteParams) WithContext(ctx context.Context) *BankStatementDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank statement delete params
func (o *BankStatementDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank statement delete params
func (o *BankStatementDeleteParams) WithHTTPClient(client *http.Client) *BankStatementDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank statement delete params
func (o *BankStatementDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the bank statement delete params
func (o *BankStatementDeleteParams) WithID(id int32) *BankStatementDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bank statement delete params
func (o *BankStatementDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BankStatementDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
