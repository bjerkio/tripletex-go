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

package product

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

// NewProductPutParams creates a new ProductPutParams object
// with the default values initialized.
func NewProductPutParams() *ProductPutParams {
	var ()
	return &ProductPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductPutParamsWithTimeout creates a new ProductPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductPutParamsWithTimeout(timeout time.Duration) *ProductPutParams {
	var ()
	return &ProductPutParams{

		timeout: timeout,
	}
}

// NewProductPutParamsWithContext creates a new ProductPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductPutParamsWithContext(ctx context.Context) *ProductPutParams {
	var ()
	return &ProductPutParams{

		Context: ctx,
	}
}

// NewProductPutParamsWithHTTPClient creates a new ProductPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductPutParamsWithHTTPClient(client *http.Client) *ProductPutParams {
	var ()
	return &ProductPutParams{
		HTTPClient: client,
	}
}

/*ProductPutParams contains all the parameters to send to the API endpoint
for the product put operation typically these are written to a http.Request
*/
type ProductPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Product
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product put params
func (o *ProductPutParams) WithTimeout(timeout time.Duration) *ProductPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product put params
func (o *ProductPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product put params
func (o *ProductPutParams) WithContext(ctx context.Context) *ProductPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product put params
func (o *ProductPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product put params
func (o *ProductPutParams) WithHTTPClient(client *http.Client) *ProductPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product put params
func (o *ProductPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product put params
func (o *ProductPutParams) WithBody(body *models.Product) *ProductPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product put params
func (o *ProductPutParams) SetBody(body *models.Product) {
	o.Body = body
}

// WithID adds the id to the product put params
func (o *ProductPutParams) WithID(id int32) *ProductPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the product put params
func (o *ProductPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProductPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
