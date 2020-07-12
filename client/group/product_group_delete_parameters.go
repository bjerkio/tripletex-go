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

package group

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

// NewProductGroupDeleteParams creates a new ProductGroupDeleteParams object
// with the default values initialized.
func NewProductGroupDeleteParams() *ProductGroupDeleteParams {
	var ()
	return &ProductGroupDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupDeleteParamsWithTimeout creates a new ProductGroupDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupDeleteParamsWithTimeout(timeout time.Duration) *ProductGroupDeleteParams {
	var ()
	return &ProductGroupDeleteParams{

		timeout: timeout,
	}
}

// NewProductGroupDeleteParamsWithContext creates a new ProductGroupDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupDeleteParamsWithContext(ctx context.Context) *ProductGroupDeleteParams {
	var ()
	return &ProductGroupDeleteParams{

		Context: ctx,
	}
}

// NewProductGroupDeleteParamsWithHTTPClient creates a new ProductGroupDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupDeleteParamsWithHTTPClient(client *http.Client) *ProductGroupDeleteParams {
	var ()
	return &ProductGroupDeleteParams{
		HTTPClient: client,
	}
}

/*ProductGroupDeleteParams contains all the parameters to send to the API endpoint
for the product group delete operation typically these are written to a http.Request
*/
type ProductGroupDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group delete params
func (o *ProductGroupDeleteParams) WithTimeout(timeout time.Duration) *ProductGroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group delete params
func (o *ProductGroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group delete params
func (o *ProductGroupDeleteParams) WithContext(ctx context.Context) *ProductGroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group delete params
func (o *ProductGroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group delete params
func (o *ProductGroupDeleteParams) WithHTTPClient(client *http.Client) *ProductGroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group delete params
func (o *ProductGroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the product group delete params
func (o *ProductGroupDeleteParams) WithID(id int32) *ProductGroupDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the product group delete params
func (o *ProductGroupDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
