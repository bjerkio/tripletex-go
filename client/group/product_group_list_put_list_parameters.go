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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewProductGroupListPutListParams creates a new ProductGroupListPutListParams object
// with the default values initialized.
func NewProductGroupListPutListParams() *ProductGroupListPutListParams {
	var ()
	return &ProductGroupListPutListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupListPutListParamsWithTimeout creates a new ProductGroupListPutListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupListPutListParamsWithTimeout(timeout time.Duration) *ProductGroupListPutListParams {
	var ()
	return &ProductGroupListPutListParams{

		timeout: timeout,
	}
}

// NewProductGroupListPutListParamsWithContext creates a new ProductGroupListPutListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupListPutListParamsWithContext(ctx context.Context) *ProductGroupListPutListParams {
	var ()
	return &ProductGroupListPutListParams{

		Context: ctx,
	}
}

// NewProductGroupListPutListParamsWithHTTPClient creates a new ProductGroupListPutListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupListPutListParamsWithHTTPClient(client *http.Client) *ProductGroupListPutListParams {
	var ()
	return &ProductGroupListPutListParams{
		HTTPClient: client,
	}
}

/*ProductGroupListPutListParams contains all the parameters to send to the API endpoint
for the product group list put list operation typically these are written to a http.Request
*/
type ProductGroupListPutListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.ProductGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group list put list params
func (o *ProductGroupListPutListParams) WithTimeout(timeout time.Duration) *ProductGroupListPutListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group list put list params
func (o *ProductGroupListPutListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group list put list params
func (o *ProductGroupListPutListParams) WithContext(ctx context.Context) *ProductGroupListPutListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group list put list params
func (o *ProductGroupListPutListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group list put list params
func (o *ProductGroupListPutListParams) WithHTTPClient(client *http.Client) *ProductGroupListPutListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group list put list params
func (o *ProductGroupListPutListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product group list put list params
func (o *ProductGroupListPutListParams) WithBody(body []*models.ProductGroup) *ProductGroupListPutListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product group list put list params
func (o *ProductGroupListPutListParams) SetBody(body []*models.ProductGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupListPutListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
