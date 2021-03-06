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

// NewProductGroupListPostListParams creates a new ProductGroupListPostListParams object
// with the default values initialized.
func NewProductGroupListPostListParams() *ProductGroupListPostListParams {
	var ()
	return &ProductGroupListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupListPostListParamsWithTimeout creates a new ProductGroupListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupListPostListParamsWithTimeout(timeout time.Duration) *ProductGroupListPostListParams {
	var ()
	return &ProductGroupListPostListParams{

		timeout: timeout,
	}
}

// NewProductGroupListPostListParamsWithContext creates a new ProductGroupListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupListPostListParamsWithContext(ctx context.Context) *ProductGroupListPostListParams {
	var ()
	return &ProductGroupListPostListParams{

		Context: ctx,
	}
}

// NewProductGroupListPostListParamsWithHTTPClient creates a new ProductGroupListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupListPostListParamsWithHTTPClient(client *http.Client) *ProductGroupListPostListParams {
	var ()
	return &ProductGroupListPostListParams{
		HTTPClient: client,
	}
}

/*ProductGroupListPostListParams contains all the parameters to send to the API endpoint
for the product group list post list operation typically these are written to a http.Request
*/
type ProductGroupListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.ProductGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group list post list params
func (o *ProductGroupListPostListParams) WithTimeout(timeout time.Duration) *ProductGroupListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group list post list params
func (o *ProductGroupListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group list post list params
func (o *ProductGroupListPostListParams) WithContext(ctx context.Context) *ProductGroupListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group list post list params
func (o *ProductGroupListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group list post list params
func (o *ProductGroupListPostListParams) WithHTTPClient(client *http.Client) *ProductGroupListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group list post list params
func (o *ProductGroupListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the product group list post list params
func (o *ProductGroupListPostListParams) WithBody(body []*models.ProductGroup) *ProductGroupListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the product group list post list params
func (o *ProductGroupListPostListParams) SetBody(body []*models.ProductGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
