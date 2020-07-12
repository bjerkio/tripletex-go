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

package orderline

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

	"github.com/bjerkio/tripletex-go/models"
)

// NewProjectOrderlinePostParams creates a new ProjectOrderlinePostParams object
// with the default values initialized.
func NewProjectOrderlinePostParams() *ProjectOrderlinePostParams {
	var ()
	return &ProjectOrderlinePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectOrderlinePostParamsWithTimeout creates a new ProjectOrderlinePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectOrderlinePostParamsWithTimeout(timeout time.Duration) *ProjectOrderlinePostParams {
	var ()
	return &ProjectOrderlinePostParams{

		timeout: timeout,
	}
}

// NewProjectOrderlinePostParamsWithContext creates a new ProjectOrderlinePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectOrderlinePostParamsWithContext(ctx context.Context) *ProjectOrderlinePostParams {
	var ()
	return &ProjectOrderlinePostParams{

		Context: ctx,
	}
}

// NewProjectOrderlinePostParamsWithHTTPClient creates a new ProjectOrderlinePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectOrderlinePostParamsWithHTTPClient(client *http.Client) *ProjectOrderlinePostParams {
	var ()
	return &ProjectOrderlinePostParams{
		HTTPClient: client,
	}
}

/*ProjectOrderlinePostParams contains all the parameters to send to the API endpoint
for the project orderline post operation typically these are written to a http.Request
*/
type ProjectOrderlinePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProjectOrderLine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project orderline post params
func (o *ProjectOrderlinePostParams) WithTimeout(timeout time.Duration) *ProjectOrderlinePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project orderline post params
func (o *ProjectOrderlinePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project orderline post params
func (o *ProjectOrderlinePostParams) WithContext(ctx context.Context) *ProjectOrderlinePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project orderline post params
func (o *ProjectOrderlinePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project orderline post params
func (o *ProjectOrderlinePostParams) WithHTTPClient(client *http.Client) *ProjectOrderlinePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project orderline post params
func (o *ProjectOrderlinePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project orderline post params
func (o *ProjectOrderlinePostParams) WithBody(body *models.ProjectOrderLine) *ProjectOrderlinePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project orderline post params
func (o *ProjectOrderlinePostParams) SetBody(body *models.ProjectOrderLine) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectOrderlinePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
