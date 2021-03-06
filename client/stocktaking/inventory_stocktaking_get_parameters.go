// Code generated by go-swagger; DO NOT EDIT.

package stocktaking

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

// NewInventoryStocktakingGetParams creates a new InventoryStocktakingGetParams object
// with the default values initialized.
func NewInventoryStocktakingGetParams() *InventoryStocktakingGetParams {
	var ()
	return &InventoryStocktakingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingGetParamsWithTimeout creates a new InventoryStocktakingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingGetParamsWithTimeout(timeout time.Duration) *InventoryStocktakingGetParams {
	var ()
	return &InventoryStocktakingGetParams{

		timeout: timeout,
	}
}

// NewInventoryStocktakingGetParamsWithContext creates a new InventoryStocktakingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingGetParamsWithContext(ctx context.Context) *InventoryStocktakingGetParams {
	var ()
	return &InventoryStocktakingGetParams{

		Context: ctx,
	}
}

// NewInventoryStocktakingGetParamsWithHTTPClient creates a new InventoryStocktakingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingGetParamsWithHTTPClient(client *http.Client) *InventoryStocktakingGetParams {
	var ()
	return &InventoryStocktakingGetParams{
		HTTPClient: client,
	}
}

/*InventoryStocktakingGetParams contains all the parameters to send to the API endpoint
for the inventory stocktaking get operation typically these are written to a http.Request
*/
type InventoryStocktakingGetParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) WithTimeout(timeout time.Duration) *InventoryStocktakingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) WithContext(ctx context.Context) *InventoryStocktakingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) WithHTTPClient(client *http.Client) *InventoryStocktakingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) WithFields(fields *string) *InventoryStocktakingGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) WithID(id int32) *InventoryStocktakingGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory stocktaking get params
func (o *InventoryStocktakingGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
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
