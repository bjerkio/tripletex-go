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

	"github.com/bjerkio/tripletex-go/models"
)

// NewProductListPostListParams creates a new ProductListPostListParams object
// with the default values initialized.
func NewProductListPostListParams() *ProductListPostListParams {
	var ()
	return &ProductListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductListPostListParamsWithTimeout creates a new ProductListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductListPostListParamsWithTimeout(timeout time.Duration) *ProductListPostListParams {
	var ()
	return &ProductListPostListParams{

		timeout: timeout,
	}
}

// NewProductListPostListParamsWithContext creates a new ProductListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductListPostListParamsWithContext(ctx context.Context) *ProductListPostListParams {
	var ()
	return &ProductListPostListParams{

		Context: ctx,
	}
}

// NewProductListPostListParamsWithHTTPClient creates a new ProductListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductListPostListParamsWithHTTPClient(client *http.Client) *ProductListPostListParams {
	var ()
	return &ProductListPostListParams{
		HTTPClient: client,
	}
}

/*ProductListPostListParams contains all the parameters to send to the API endpoint
for the product list post list operation typically these are written to a http.Request
*/
type ProductListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.Product

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product list post list params
func (o *ProductListPostListParams) WithTimeout(timeout time.Duration) *ProductListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product list post list params
func (o *ProductListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product list post list params
func (o *ProductListPostListParams) WithContext(ctx context.Context) *ProductListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product list post list params
func (o *ProductListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product list post list params
func (o *ProductListPostListParams) WithHTTPClient(client *http.Client) *ProductListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product list post list params
func (o *ProductListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product list post list params
func (o *ProductListPostListParams) WithBody(body []*models.Product) *ProductListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product list post list params
func (o *ProductListPostListParams) SetBody(body []*models.Product) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
