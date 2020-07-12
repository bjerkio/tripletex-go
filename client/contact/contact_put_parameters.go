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

package contact

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

	"github.com/bjerkio/tripletex-go/models"
)

// NewContactPutParams creates a new ContactPutParams object
// with the default values initialized.
func NewContactPutParams() *ContactPutParams {
	var ()
	return &ContactPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContactPutParamsWithTimeout creates a new ContactPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContactPutParamsWithTimeout(timeout time.Duration) *ContactPutParams {
	var ()
	return &ContactPutParams{

		timeout: timeout,
	}
}

// NewContactPutParamsWithContext creates a new ContactPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewContactPutParamsWithContext(ctx context.Context) *ContactPutParams {
	var ()
	return &ContactPutParams{

		Context: ctx,
	}
}

// NewContactPutParamsWithHTTPClient creates a new ContactPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContactPutParamsWithHTTPClient(client *http.Client) *ContactPutParams {
	var ()
	return &ContactPutParams{
		HTTPClient: client,
	}
}

/*ContactPutParams contains all the parameters to send to the API endpoint
for the contact put operation typically these are written to a http.Request
*/
type ContactPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Contact
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contact put params
func (o *ContactPutParams) WithTimeout(timeout time.Duration) *ContactPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact put params
func (o *ContactPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact put params
func (o *ContactPutParams) WithContext(ctx context.Context) *ContactPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact put params
func (o *ContactPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact put params
func (o *ContactPutParams) WithHTTPClient(client *http.Client) *ContactPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact put params
func (o *ContactPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the contact put params
func (o *ContactPutParams) WithBody(body *models.Contact) *ContactPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the contact put params
func (o *ContactPutParams) SetBody(body *models.Contact) {
	o.Body = body
}

// WithID adds the id to the contact put params
func (o *ContactPutParams) WithID(id int32) *ContactPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact put params
func (o *ContactPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
