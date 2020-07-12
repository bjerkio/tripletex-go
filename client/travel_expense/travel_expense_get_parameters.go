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
	"github.com/go-openapi/swag"
)

// NewTravelExpenseGetParams creates a new TravelExpenseGetParams object
// with the default values initialized.
func NewTravelExpenseGetParams() *TravelExpenseGetParams {
	var ()
	return &TravelExpenseGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseGetParamsWithTimeout creates a new TravelExpenseGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseGetParamsWithTimeout(timeout time.Duration) *TravelExpenseGetParams {
	var ()
	return &TravelExpenseGetParams{

		timeout: timeout,
	}
}

// NewTravelExpenseGetParamsWithContext creates a new TravelExpenseGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseGetParamsWithContext(ctx context.Context) *TravelExpenseGetParams {
	var ()
	return &TravelExpenseGetParams{

		Context: ctx,
	}
}

// NewTravelExpenseGetParamsWithHTTPClient creates a new TravelExpenseGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseGetParamsWithHTTPClient(client *http.Client) *TravelExpenseGetParams {
	var ()
	return &TravelExpenseGetParams{
		HTTPClient: client,
	}
}

/*TravelExpenseGetParams contains all the parameters to send to the API endpoint
for the travel expense get operation typically these are written to a http.Request
*/
type TravelExpenseGetParams struct {

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

// WithTimeout adds the timeout to the travel expense get params
func (o *TravelExpenseGetParams) WithTimeout(timeout time.Duration) *TravelExpenseGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense get params
func (o *TravelExpenseGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense get params
func (o *TravelExpenseGetParams) WithContext(ctx context.Context) *TravelExpenseGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense get params
func (o *TravelExpenseGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense get params
func (o *TravelExpenseGetParams) WithHTTPClient(client *http.Client) *TravelExpenseGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense get params
func (o *TravelExpenseGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the travel expense get params
func (o *TravelExpenseGetParams) WithFields(fields *string) *TravelExpenseGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the travel expense get params
func (o *TravelExpenseGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the travel expense get params
func (o *TravelExpenseGetParams) WithID(id int32) *TravelExpenseGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense get params
func (o *TravelExpenseGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
