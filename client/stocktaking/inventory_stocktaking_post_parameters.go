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

	"github.com/bjerkio/tripletex-go/models"
)

// NewInventoryStocktakingPostParams creates a new InventoryStocktakingPostParams object
// with the default values initialized.
func NewInventoryStocktakingPostParams() *InventoryStocktakingPostParams {
	var ()
	return &InventoryStocktakingPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingPostParamsWithTimeout creates a new InventoryStocktakingPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingPostParamsWithTimeout(timeout time.Duration) *InventoryStocktakingPostParams {
	var ()
	return &InventoryStocktakingPostParams{

		timeout: timeout,
	}
}

// NewInventoryStocktakingPostParamsWithContext creates a new InventoryStocktakingPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingPostParamsWithContext(ctx context.Context) *InventoryStocktakingPostParams {
	var ()
	return &InventoryStocktakingPostParams{

		Context: ctx,
	}
}

// NewInventoryStocktakingPostParamsWithHTTPClient creates a new InventoryStocktakingPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingPostParamsWithHTTPClient(client *http.Client) *InventoryStocktakingPostParams {
	var ()
	return &InventoryStocktakingPostParams{
		HTTPClient: client,
	}
}

/*InventoryStocktakingPostParams contains all the parameters to send to the API endpoint
for the inventory stocktaking post operation typically these are written to a http.Request
*/
type InventoryStocktakingPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Stocktaking

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) WithTimeout(timeout time.Duration) *InventoryStocktakingPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) WithContext(ctx context.Context) *InventoryStocktakingPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) WithHTTPClient(client *http.Client) *InventoryStocktakingPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) WithBody(body *models.Stocktaking) *InventoryStocktakingPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the inventory stocktaking post params
func (o *InventoryStocktakingPostParams) SetBody(body *models.Stocktaking) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
