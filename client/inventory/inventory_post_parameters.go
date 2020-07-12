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

package inventory

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

// NewInventoryPostParams creates a new InventoryPostParams object
// with the default values initialized.
func NewInventoryPostParams() *InventoryPostParams {
	var ()
	return &InventoryPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryPostParamsWithTimeout creates a new InventoryPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryPostParamsWithTimeout(timeout time.Duration) *InventoryPostParams {
	var ()
	return &InventoryPostParams{

		timeout: timeout,
	}
}

// NewInventoryPostParamsWithContext creates a new InventoryPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryPostParamsWithContext(ctx context.Context) *InventoryPostParams {
	var ()
	return &InventoryPostParams{

		Context: ctx,
	}
}

// NewInventoryPostParamsWithHTTPClient creates a new InventoryPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryPostParamsWithHTTPClient(client *http.Client) *InventoryPostParams {
	var ()
	return &InventoryPostParams{
		HTTPClient: client,
	}
}

/*InventoryPostParams contains all the parameters to send to the API endpoint
for the inventory post operation typically these are written to a http.Request
*/
type InventoryPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Inventory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory post params
func (o *InventoryPostParams) WithTimeout(timeout time.Duration) *InventoryPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory post params
func (o *InventoryPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory post params
func (o *InventoryPostParams) WithContext(ctx context.Context) *InventoryPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory post params
func (o *InventoryPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory post params
func (o *InventoryPostParams) WithHTTPClient(client *http.Client) *InventoryPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory post params
func (o *InventoryPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the inventory post params
func (o *InventoryPostParams) WithBody(body *models.Inventory) *InventoryPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the inventory post params
func (o *InventoryPostParams) SetBody(body *models.Inventory) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
