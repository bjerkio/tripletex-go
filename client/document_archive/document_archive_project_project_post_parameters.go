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

// NewDocumentArchiveProjectProjectPostParams creates a new DocumentArchiveProjectProjectPostParams object
// with the default values initialized.
func NewDocumentArchiveProjectProjectPostParams() *DocumentArchiveProjectProjectPostParams {
	var ()
	return &DocumentArchiveProjectProjectPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentArchiveProjectProjectPostParamsWithTimeout creates a new DocumentArchiveProjectProjectPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentArchiveProjectProjectPostParamsWithTimeout(timeout time.Duration) *DocumentArchiveProjectProjectPostParams {
	var ()
	return &DocumentArchiveProjectProjectPostParams{

		timeout: timeout,
	}
}

// NewDocumentArchiveProjectProjectPostParamsWithContext creates a new DocumentArchiveProjectProjectPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentArchiveProjectProjectPostParamsWithContext(ctx context.Context) *DocumentArchiveProjectProjectPostParams {
	var ()
	return &DocumentArchiveProjectProjectPostParams{

		Context: ctx,
	}
}

// NewDocumentArchiveProjectProjectPostParamsWithHTTPClient creates a new DocumentArchiveProjectProjectPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentArchiveProjectProjectPostParamsWithHTTPClient(client *http.Client) *DocumentArchiveProjectProjectPostParams {
	var ()
	return &DocumentArchiveProjectProjectPostParams{
		HTTPClient: client,
	}
}

/*DocumentArchiveProjectProjectPostParams contains all the parameters to send to the API endpoint
for the document archive project project post operation typically these are written to a http.Request
*/
type DocumentArchiveProjectProjectPostParams struct {

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

// WithTimeout adds the timeout to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) WithTimeout(timeout time.Duration) *DocumentArchiveProjectProjectPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) WithContext(ctx context.Context) *DocumentArchiveProjectProjectPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) WithHTTPClient(client *http.Client) *DocumentArchiveProjectProjectPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) WithFile(file runtime.NamedReadCloser) *DocumentArchiveProjectProjectPostParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) WithID(id int32) *DocumentArchiveProjectProjectPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document archive project project post params
func (o *DocumentArchiveProjectProjectPostParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentArchiveProjectProjectPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
