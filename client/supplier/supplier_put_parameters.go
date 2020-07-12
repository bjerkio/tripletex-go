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

package supplier

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

// NewSupplierPutParams creates a new SupplierPutParams object
// with the default values initialized.
func NewSupplierPutParams() *SupplierPutParams {
	var ()
	return &SupplierPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierPutParamsWithTimeout creates a new SupplierPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierPutParamsWithTimeout(timeout time.Duration) *SupplierPutParams {
	var ()
	return &SupplierPutParams{

		timeout: timeout,
	}
}

// NewSupplierPutParamsWithContext creates a new SupplierPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierPutParamsWithContext(ctx context.Context) *SupplierPutParams {
	var ()
	return &SupplierPutParams{

		Context: ctx,
	}
}

// NewSupplierPutParamsWithHTTPClient creates a new SupplierPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierPutParamsWithHTTPClient(client *http.Client) *SupplierPutParams {
	var ()
	return &SupplierPutParams{
		HTTPClient: client,
	}
}

/*SupplierPutParams contains all the parameters to send to the API endpoint
for the supplier put operation typically these are written to a http.Request
*/
type SupplierPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Supplier
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier put params
func (o *SupplierPutParams) WithTimeout(timeout time.Duration) *SupplierPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier put params
func (o *SupplierPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier put params
func (o *SupplierPutParams) WithContext(ctx context.Context) *SupplierPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier put params
func (o *SupplierPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier put params
func (o *SupplierPutParams) WithHTTPClient(client *http.Client) *SupplierPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier put params
func (o *SupplierPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the supplier put params
func (o *SupplierPutParams) WithBody(body *models.Supplier) *SupplierPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the supplier put params
func (o *SupplierPutParams) SetBody(body *models.Supplier) {
	o.Body = body
}

// WithID adds the id to the supplier put params
func (o *SupplierPutParams) WithID(id int32) *SupplierPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the supplier put params
func (o *SupplierPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
