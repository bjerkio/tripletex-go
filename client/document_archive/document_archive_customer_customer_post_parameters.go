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

// NewDocumentArchiveCustomerCustomerPostParams creates a new DocumentArchiveCustomerCustomerPostParams object
// with the default values initialized.
func NewDocumentArchiveCustomerCustomerPostParams() *DocumentArchiveCustomerCustomerPostParams {
	var ()
	return &DocumentArchiveCustomerCustomerPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentArchiveCustomerCustomerPostParamsWithTimeout creates a new DocumentArchiveCustomerCustomerPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentArchiveCustomerCustomerPostParamsWithTimeout(timeout time.Duration) *DocumentArchiveCustomerCustomerPostParams {
	var ()
	return &DocumentArchiveCustomerCustomerPostParams{

		timeout: timeout,
	}
}

// NewDocumentArchiveCustomerCustomerPostParamsWithContext creates a new DocumentArchiveCustomerCustomerPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentArchiveCustomerCustomerPostParamsWithContext(ctx context.Context) *DocumentArchiveCustomerCustomerPostParams {
	var ()
	return &DocumentArchiveCustomerCustomerPostParams{

		Context: ctx,
	}
}

// NewDocumentArchiveCustomerCustomerPostParamsWithHTTPClient creates a new DocumentArchiveCustomerCustomerPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentArchiveCustomerCustomerPostParamsWithHTTPClient(client *http.Client) *DocumentArchiveCustomerCustomerPostParams {
	var ()
	return &DocumentArchiveCustomerCustomerPostParams{
		HTTPClient: client,
	}
}

/*DocumentArchiveCustomerCustomerPostParams contains all the parameters to send to the API endpoint
for the document archive customer customer post operation typically these are written to a http.Request
*/
type DocumentArchiveCustomerCustomerPostParams struct {

	/*File
	  The file

	*/
	File runtime.NamedReadCloser
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) WithTimeout(timeout time.Duration) *DocumentArchiveCustomerCustomerPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) WithContext(ctx context.Context) *DocumentArchiveCustomerCustomerPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) WithHTTPClient(client *http.Client) *DocumentArchiveCustomerCustomerPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) WithFile(file runtime.NamedReadCloser) *DocumentArchiveCustomerCustomerPostParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) WithID(id int32) *DocumentArchiveCustomerCustomerPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document archive customer customer post params
func (o *DocumentArchiveCustomerCustomerPostParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentArchiveCustomerCustomerPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
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
