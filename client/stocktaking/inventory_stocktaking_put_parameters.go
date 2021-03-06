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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewInventoryStocktakingPutParams creates a new InventoryStocktakingPutParams object
// with the default values initialized.
func NewInventoryStocktakingPutParams() *InventoryStocktakingPutParams {
	var ()
	return &InventoryStocktakingPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingPutParamsWithTimeout creates a new InventoryStocktakingPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingPutParamsWithTimeout(timeout time.Duration) *InventoryStocktakingPutParams {
	var ()
	return &InventoryStocktakingPutParams{

		timeout: timeout,
	}
}

// NewInventoryStocktakingPutParamsWithContext creates a new InventoryStocktakingPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingPutParamsWithContext(ctx context.Context) *InventoryStocktakingPutParams {
	var ()
	return &InventoryStocktakingPutParams{

		Context: ctx,
	}
}

// NewInventoryStocktakingPutParamsWithHTTPClient creates a new InventoryStocktakingPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingPutParamsWithHTTPClient(client *http.Client) *InventoryStocktakingPutParams {
	var ()
	return &InventoryStocktakingPutParams{
		HTTPClient: client,
	}
}

/*InventoryStocktakingPutParams contains all the parameters to send to the API endpoint
for the inventory stocktaking put operation typically these are written to a http.Request
*/
type InventoryStocktakingPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Stocktaking
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) WithTimeout(timeout time.Duration) *InventoryStocktakingPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) WithContext(ctx context.Context) *InventoryStocktakingPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) WithHTTPClient(client *http.Client) *InventoryStocktakingPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) WithBody(body *models.Stocktaking) *InventoryStocktakingPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) SetBody(body *models.Stocktaking) {
	o.Body = body
}

// WithID adds the id to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) WithID(id int32) *InventoryStocktakingPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory stocktaking put params
func (o *InventoryStocktakingPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
