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

package document_archive

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

// NewDocumentArchiveDeleteParams creates a new DocumentArchiveDeleteParams object
// with the default values initialized.
func NewDocumentArchiveDeleteParams() *DocumentArchiveDeleteParams {
	var ()
	return &DocumentArchiveDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentArchiveDeleteParamsWithTimeout creates a new DocumentArchiveDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentArchiveDeleteParamsWithTimeout(timeout time.Duration) *DocumentArchiveDeleteParams {
	var ()
	return &DocumentArchiveDeleteParams{

		timeout: timeout,
	}
}

// NewDocumentArchiveDeleteParamsWithContext creates a new DocumentArchiveDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentArchiveDeleteParamsWithContext(ctx context.Context) *DocumentArchiveDeleteParams {
	var ()
	return &DocumentArchiveDeleteParams{

		Context: ctx,
	}
}

// NewDocumentArchiveDeleteParamsWithHTTPClient creates a new DocumentArchiveDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentArchiveDeleteParamsWithHTTPClient(client *http.Client) *DocumentArchiveDeleteParams {
	var ()
	return &DocumentArchiveDeleteParams{
		HTTPClient: client,
	}
}

/*DocumentArchiveDeleteParams contains all the parameters to send to the API endpoint
for the document archive delete operation typically these are written to a http.Request
*/
type DocumentArchiveDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the document archive delete params
func (o *DocumentArchiveDeleteParams) WithTimeout(timeout time.Duration) *DocumentArchiveDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document archive delete params
func (o *DocumentArchiveDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document archive delete params
func (o *DocumentArchiveDeleteParams) WithContext(ctx context.Context) *DocumentArchiveDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document archive delete params
func (o *DocumentArchiveDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document archive delete params
func (o *DocumentArchiveDeleteParams) WithHTTPClient(client *http.Client) *DocumentArchiveDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document archive delete params
func (o *DocumentArchiveDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the document archive delete params
func (o *DocumentArchiveDeleteParams) WithID(id int32) *DocumentArchiveDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document archive delete params
func (o *DocumentArchiveDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentArchiveDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
