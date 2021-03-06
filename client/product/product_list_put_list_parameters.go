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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewProductListPutListParams creates a new ProductListPutListParams object
// with the default values initialized.
func NewProductListPutListParams() *ProductListPutListParams {
	var ()
	return &ProductListPutListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductListPutListParamsWithTimeout creates a new ProductListPutListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductListPutListParamsWithTimeout(timeout time.Duration) *ProductListPutListParams {
	var ()
	return &ProductListPutListParams{

		timeout: timeout,
	}
}

// NewProductListPutListParamsWithContext creates a new ProductListPutListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductListPutListParamsWithContext(ctx context.Context) *ProductListPutListParams {
	var ()
	return &ProductListPutListParams{

		Context: ctx,
	}
}

// NewProductListPutListParamsWithHTTPClient creates a new ProductListPutListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductListPutListParamsWithHTTPClient(client *http.Client) *ProductListPutListParams {
	var ()
	return &ProductListPutListParams{
		HTTPClient: client,
	}
}

/*ProductListPutListParams contains all the parameters to send to the API endpoint
for the product list put list operation typically these are written to a http.Request
*/
type ProductListPutListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.Product

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product list put list params
func (o *ProductListPutListParams) WithTimeout(timeout time.Duration) *ProductListPutListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product list put list params
func (o *ProductListPutListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product list put list params
func (o *ProductListPutListParams) WithContext(ctx context.Context) *ProductListPutListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product list put list params
func (o *ProductListPutListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product list put list params
func (o *ProductListPutListParams) WithHTTPClient(client *http.Client) *ProductListPutListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product list put list params
func (o *ProductListPutListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product list put list params
func (o *ProductListPutListParams) WithBody(body []*models.Product) *ProductListPutListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product list put list params
func (o *ProductListPutListParams) SetBody(body []*models.Product) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductListPutListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
