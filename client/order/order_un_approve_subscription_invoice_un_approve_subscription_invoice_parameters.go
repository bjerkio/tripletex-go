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

package order

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

// NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams creates a new OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams object
// with the default values initialized.
func NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams() *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	var ()
	return &OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithTimeout creates a new OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithTimeout(timeout time.Duration) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	var ()
	return &OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams{

		timeout: timeout,
	}
}

// NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithContext creates a new OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithContext(ctx context.Context) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	var ()
	return &OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams{

		Context: ctx,
	}
}

// NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithHTTPClient creates a new OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParamsWithHTTPClient(client *http.Client) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	var ()
	return &OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams{
		HTTPClient: client,
	}
}

/*OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams contains all the parameters to send to the API endpoint
for the order un approve subscription invoice un approve subscription invoice operation typically these are written to a http.Request
*/
type OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams struct {

	/*ID
	  ID of order to unapprove for subscription invoicing.

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) WithTimeout(timeout time.Duration) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) WithContext(ctx context.Context) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) WithHTTPClient(client *http.Client) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) WithID(id int32) *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the order un approve subscription invoice un approve subscription invoice params
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrderUnApproveSubscriptionInvoiceUnApproveSubscriptionInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
