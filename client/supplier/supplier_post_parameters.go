// Code generated by go-swagger; DO NOT EDIT.

package supplier

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

// NewSupplierPostParams creates a new SupplierPostParams object
// with the default values initialized.
func NewSupplierPostParams() *SupplierPostParams {
	var ()
	return &SupplierPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSupplierPostParamsWithTimeout creates a new SupplierPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSupplierPostParamsWithTimeout(timeout time.Duration) *SupplierPostParams {
	var ()
	return &SupplierPostParams{

		timeout: timeout,
	}
}

// NewSupplierPostParamsWithContext creates a new SupplierPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSupplierPostParamsWithContext(ctx context.Context) *SupplierPostParams {
	var ()
	return &SupplierPostParams{

		Context: ctx,
	}
}

// NewSupplierPostParamsWithHTTPClient creates a new SupplierPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSupplierPostParamsWithHTTPClient(client *http.Client) *SupplierPostParams {
	var ()
	return &SupplierPostParams{
		HTTPClient: client,
	}
}

/*SupplierPostParams contains all the parameters to send to the API endpoint
for the supplier post operation typically these are written to a http.Request
*/
type SupplierPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.Supplier

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the supplier post params
func (o *SupplierPostParams) WithTimeout(timeout time.Duration) *SupplierPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supplier post params
func (o *SupplierPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supplier post params
func (o *SupplierPostParams) WithContext(ctx context.Context) *SupplierPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supplier post params
func (o *SupplierPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supplier post params
func (o *SupplierPostParams) WithHTTPClient(client *http.Client) *SupplierPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supplier post params
func (o *SupplierPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the supplier post params
func (o *SupplierPostParams) WithBody(body *models.Supplier) *SupplierPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the supplier post params
func (o *SupplierPostParams) SetBody(body *models.Supplier) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SupplierPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
