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

package productline

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

// NewInventoryStocktakingProductlineDeleteParams creates a new InventoryStocktakingProductlineDeleteParams object
// with the default values initialized.
func NewInventoryStocktakingProductlineDeleteParams() *InventoryStocktakingProductlineDeleteParams {
	var ()
	return &InventoryStocktakingProductlineDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingProductlineDeleteParamsWithTimeout creates a new InventoryStocktakingProductlineDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingProductlineDeleteParamsWithTimeout(timeout time.Duration) *InventoryStocktakingProductlineDeleteParams {
	var ()
	return &InventoryStocktakingProductlineDeleteParams{

		timeout: timeout,
	}
}

// NewInventoryStocktakingProductlineDeleteParamsWithContext creates a new InventoryStocktakingProductlineDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingProductlineDeleteParamsWithContext(ctx context.Context) *InventoryStocktakingProductlineDeleteParams {
	var ()
	return &InventoryStocktakingProductlineDeleteParams{

		Context: ctx,
	}
}

// NewInventoryStocktakingProductlineDeleteParamsWithHTTPClient creates a new InventoryStocktakingProductlineDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingProductlineDeleteParamsWithHTTPClient(client *http.Client) *InventoryStocktakingProductlineDeleteParams {
	var ()
	return &InventoryStocktakingProductlineDeleteParams{
		HTTPClient: client,
	}
}

/*InventoryStocktakingProductlineDeleteParams contains all the parameters to send to the API endpoint
for the inventory stocktaking productline delete operation typically these are written to a http.Request
*/
type InventoryStocktakingProductlineDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) WithTimeout(timeout time.Duration) *InventoryStocktakingProductlineDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) WithContext(ctx context.Context) *InventoryStocktakingProductlineDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) WithHTTPClient(client *http.Client) *InventoryStocktakingProductlineDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) WithID(id int32) *InventoryStocktakingProductlineDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory stocktaking productline delete params
func (o *InventoryStocktakingProductlineDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingProductlineDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
