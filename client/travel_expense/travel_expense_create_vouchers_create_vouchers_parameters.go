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

package travel_expense

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

// NewTravelExpenseCreateVouchersCreateVouchersParams creates a new TravelExpenseCreateVouchersCreateVouchersParams object
// with the default values initialized.
func NewTravelExpenseCreateVouchersCreateVouchersParams() *TravelExpenseCreateVouchersCreateVouchersParams {
	var ()
	return &TravelExpenseCreateVouchersCreateVouchersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseCreateVouchersCreateVouchersParamsWithTimeout creates a new TravelExpenseCreateVouchersCreateVouchersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseCreateVouchersCreateVouchersParamsWithTimeout(timeout time.Duration) *TravelExpenseCreateVouchersCreateVouchersParams {
	var ()
	return &TravelExpenseCreateVouchersCreateVouchersParams{

		timeout: timeout,
	}
}

// NewTravelExpenseCreateVouchersCreateVouchersParamsWithContext creates a new TravelExpenseCreateVouchersCreateVouchersParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseCreateVouchersCreateVouchersParamsWithContext(ctx context.Context) *TravelExpenseCreateVouchersCreateVouchersParams {
	var ()
	return &TravelExpenseCreateVouchersCreateVouchersParams{

		Context: ctx,
	}
}

// NewTravelExpenseCreateVouchersCreateVouchersParamsWithHTTPClient creates a new TravelExpenseCreateVouchersCreateVouchersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseCreateVouchersCreateVouchersParamsWithHTTPClient(client *http.Client) *TravelExpenseCreateVouchersCreateVouchersParams {
	var ()
	return &TravelExpenseCreateVouchersCreateVouchersParams{
		HTTPClient: client,
	}
}

/*TravelExpenseCreateVouchersCreateVouchersParams contains all the parameters to send to the API endpoint
for the travel expense create vouchers create vouchers operation typically these are written to a http.Request
*/
type TravelExpenseCreateVouchersCreateVouchersParams struct {

	/*Date
	  yyyy-MM-dd. Defaults to today.

	*/
	Date string
	/*ID
	  ID of the elements

	*/
	ID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WithTimeout(timeout time.Duration) *TravelExpenseCreateVouchersCreateVouchersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WithContext(ctx context.Context) *TravelExpenseCreateVouchersCreateVouchersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WithHTTPClient(client *http.Client) *TravelExpenseCreateVouchersCreateVouchersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WithDate(date string) *TravelExpenseCreateVouchersCreateVouchersParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) SetDate(date string) {
	o.Date = date
}

// WithID adds the id to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WithID(id *string) *TravelExpenseCreateVouchersCreateVouchersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense create vouchers create vouchers params
func (o *TravelExpenseCreateVouchersCreateVouchersParams) SetID(id *string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseCreateVouchersCreateVouchersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param date
	qrDate := o.Date
	qDate := qrDate
	if qDate != "" {
		if err := r.SetQueryParam("date", qDate); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
