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
)

// NewLedgerAccountListDeleteByIdsParams creates a new LedgerAccountListDeleteByIdsParams object
// with the default values initialized.
func NewLedgerAccountListDeleteByIdsParams() *LedgerAccountListDeleteByIdsParams {
	var ()
	return &LedgerAccountListDeleteByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLedgerAccountListDeleteByIdsParamsWithTimeout creates a new LedgerAccountListDeleteByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLedgerAccountListDeleteByIdsParamsWithTimeout(timeout time.Duration) *LedgerAccountListDeleteByIdsParams {
	var ()
	return &LedgerAccountListDeleteByIdsParams{

		timeout: timeout,
	}
}

// NewLedgerAccountListDeleteByIdsParamsWithContext creates a new LedgerAccountListDeleteByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLedgerAccountListDeleteByIdsParamsWithContext(ctx context.Context) *LedgerAccountListDeleteByIdsParams {
	var ()
	return &LedgerAccountListDeleteByIdsParams{

		Context: ctx,
	}
}

// NewLedgerAccountListDeleteByIdsParamsWithHTTPClient creates a new LedgerAccountListDeleteByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLedgerAccountListDeleteByIdsParamsWithHTTPClient(client *http.Client) *LedgerAccountListDeleteByIdsParams {
	var ()
	return &LedgerAccountListDeleteByIdsParams{
		HTTPClient: client,
	}
}

/*LedgerAccountListDeleteByIdsParams contains all the parameters to send to the API endpoint
for the ledger account list delete by ids operation typically these are written to a http.Request
*/
type LedgerAccountListDeleteByIdsParams struct {

	/*Ids
	  ID of the elements

	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) WithTimeout(timeout time.Duration) *LedgerAccountListDeleteByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) WithContext(ctx context.Context) *LedgerAccountListDeleteByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) WithHTTPClient(client *http.Client) *LedgerAccountListDeleteByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) WithIds(ids string) *LedgerAccountListDeleteByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the ledger account list delete by ids params
func (o *LedgerAccountListDeleteByIdsParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *LedgerAccountListDeleteByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
