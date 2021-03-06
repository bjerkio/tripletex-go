// Code generated by go-swagger; DO NOT EDIT.

package salesmodules

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

// NewCompanySalesmodulesPostParams creates a new CompanySalesmodulesPostParams object
// with the default values initialized.
func NewCompanySalesmodulesPostParams() *CompanySalesmodulesPostParams {
	var ()
	return &CompanySalesmodulesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompanySalesmodulesPostParamsWithTimeout creates a new CompanySalesmodulesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompanySalesmodulesPostParamsWithTimeout(timeout time.Duration) *CompanySalesmodulesPostParams {
	var ()
	return &CompanySalesmodulesPostParams{

		timeout: timeout,
	}
}

// NewCompanySalesmodulesPostParamsWithContext creates a new CompanySalesmodulesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompanySalesmodulesPostParamsWithContext(ctx context.Context) *CompanySalesmodulesPostParams {
	var ()
	return &CompanySalesmodulesPostParams{

		Context: ctx,
	}
}

// NewCompanySalesmodulesPostParamsWithHTTPClient creates a new CompanySalesmodulesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompanySalesmodulesPostParamsWithHTTPClient(client *http.Client) *CompanySalesmodulesPostParams {
	var ()
	return &CompanySalesmodulesPostParams{
		HTTPClient: client,
	}
}

/*CompanySalesmodulesPostParams contains all the parameters to send to the API endpoint
for the company salesmodules post operation typically these are written to a http.Request
*/
type CompanySalesmodulesPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.SalesModuleDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) WithTimeout(timeout time.Duration) *CompanySalesmodulesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) WithContext(ctx context.Context) *CompanySalesmodulesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) WithHTTPClient(client *http.Client) *CompanySalesmodulesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) WithBody(body *models.SalesModuleDTO) *CompanySalesmodulesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the company salesmodules post params
func (o *CompanySalesmodulesPostParams) SetBody(body *models.SalesModuleDTO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CompanySalesmodulesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
