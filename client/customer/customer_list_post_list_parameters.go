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

package customer

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

// NewCustomerListPostListParams creates a new CustomerListPostListParams object
// with the default values initialized.
func NewCustomerListPostListParams() *CustomerListPostListParams {
	var ()
	return &CustomerListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerListPostListParamsWithTimeout creates a new CustomerListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerListPostListParamsWithTimeout(timeout time.Duration) *CustomerListPostListParams {
	var ()
	return &CustomerListPostListParams{

		timeout: timeout,
	}
}

// NewCustomerListPostListParamsWithContext creates a new CustomerListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerListPostListParamsWithContext(ctx context.Context) *CustomerListPostListParams {
	var ()
	return &CustomerListPostListParams{

		Context: ctx,
	}
}

// NewCustomerListPostListParamsWithHTTPClient creates a new CustomerListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerListPostListParamsWithHTTPClient(client *http.Client) *CustomerListPostListParams {
	var ()
	return &CustomerListPostListParams{
		HTTPClient: client,
	}
}

/*CustomerListPostListParams contains all the parameters to send to the API endpoint
for the customer list post list operation typically these are written to a http.Request
*/
type CustomerListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.Customer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer list post list params
func (o *CustomerListPostListParams) WithTimeout(timeout time.Duration) *CustomerListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer list post list params
func (o *CustomerListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer list post list params
func (o *CustomerListPostListParams) WithContext(ctx context.Context) *CustomerListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer list post list params
func (o *CustomerListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer list post list params
func (o *CustomerListPostListParams) WithHTTPClient(client *http.Client) *CustomerListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer list post list params
func (o *CustomerListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the customer list post list params
func (o *CustomerListPostListParams) WithBody(body []*models.Customer) *CustomerListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the customer list post list params
func (o *CustomerListPostListParams) SetBody(body []*models.Customer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
