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

	"github.com/bjerkio/tripletex-go/models"
)

// NewInventoryStocktakingProductlinePostParams creates a new InventoryStocktakingProductlinePostParams object
// with the default values initialized.
func NewInventoryStocktakingProductlinePostParams() *InventoryStocktakingProductlinePostParams {
	var ()
	return &InventoryStocktakingProductlinePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryStocktakingProductlinePostParamsWithTimeout creates a new InventoryStocktakingProductlinePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInventoryStocktakingProductlinePostParamsWithTimeout(timeout time.Duration) *InventoryStocktakingProductlinePostParams {
	var ()
	return &InventoryStocktakingProductlinePostParams{

		timeout: timeout,
	}
}

// NewInventoryStocktakingProductlinePostParamsWithContext creates a new InventoryStocktakingProductlinePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewInventoryStocktakingProductlinePostParamsWithContext(ctx context.Context) *InventoryStocktakingProductlinePostParams {
	var ()
	return &InventoryStocktakingProductlinePostParams{

		Context: ctx,
	}
}

// NewInventoryStocktakingProductlinePostParamsWithHTTPClient creates a new InventoryStocktakingProductlinePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInventoryStocktakingProductlinePostParamsWithHTTPClient(client *http.Client) *InventoryStocktakingProductlinePostParams {
	var ()
	return &InventoryStocktakingProductlinePostParams{
		HTTPClient: client,
	}
}

/*InventoryStocktakingProductlinePostParams contains all the parameters to send to the API endpoint
for the inventory stocktaking productline post operation typically these are written to a http.Request
*/
type InventoryStocktakingProductlinePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProductLine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) WithTimeout(timeout time.Duration) *InventoryStocktakingProductlinePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) WithContext(ctx context.Context) *InventoryStocktakingProductlinePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) WithHTTPClient(client *http.Client) *InventoryStocktakingProductlinePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) WithBody(body *models.ProductLine) *InventoryStocktakingProductlinePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the inventory stocktaking productline post params
func (o *InventoryStocktakingProductlinePostParams) SetBody(body *models.ProductLine) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryStocktakingProductlinePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
