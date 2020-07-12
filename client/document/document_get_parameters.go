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

package document

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

// NewDocumentGetParams creates a new DocumentGetParams object
// with the default values initialized.
func NewDocumentGetParams() *DocumentGetParams {
	var ()
	return &DocumentGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentGetParamsWithTimeout creates a new DocumentGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentGetParamsWithTimeout(timeout time.Duration) *DocumentGetParams {
	var ()
	return &DocumentGetParams{

		timeout: timeout,
	}
}

// NewDocumentGetParamsWithContext creates a new DocumentGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentGetParamsWithContext(ctx context.Context) *DocumentGetParams {
	var ()
	return &DocumentGetParams{

		Context: ctx,
	}
}

// NewDocumentGetParamsWithHTTPClient creates a new DocumentGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentGetParamsWithHTTPClient(client *http.Client) *DocumentGetParams {
	var ()
	return &DocumentGetParams{
		HTTPClient: client,
	}
}

/*DocumentGetParams contains all the parameters to send to the API endpoint
for the document get operation typically these are written to a http.Request
*/
type DocumentGetParams struct {

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

// WithTimeout adds the timeout to the document get params
func (o *DocumentGetParams) WithTimeout(timeout time.Duration) *DocumentGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document get params
func (o *DocumentGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document get params
func (o *DocumentGetParams) WithContext(ctx context.Context) *DocumentGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document get params
func (o *DocumentGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document get params
func (o *DocumentGetParams) WithHTTPClient(client *http.Client) *DocumentGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document get params
func (o *DocumentGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the document get params
func (o *DocumentGetParams) WithFields(fields *string) *DocumentGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the document get params
func (o *DocumentGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the document get params
func (o *DocumentGetParams) WithID(id int32) *DocumentGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document get params
func (o *DocumentGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
