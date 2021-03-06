// Code generated by go-swagger; DO NOT EDIT.

package company

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

// NewCompanyPutParams creates a new CompanyPutParams object
// with the default values initialized.
func NewCompanyPutParams() *CompanyPutParams {
	var ()
	return &CompanyPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompanyPutParamsWithTimeout creates a new CompanyPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompanyPutParamsWithTimeout(timeout time.Duration) *CompanyPutParams {
	var ()
	return &CompanyPutParams{

		timeout: timeout,
	}
}

// NewCompanyPutParamsWithContext creates a new CompanyPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompanyPutParamsWithContext(ctx context.Context) *CompanyPutParams {
	var ()
	return &CompanyPutParams{

		Context: ctx,
	}
}

// NewCompanyPutParamsWithHTTPClient creates a new CompanyPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompanyPutParamsWithHTTPClient(client *http.Client) *CompanyPutParams {
	var ()
	return &CompanyPutParams{
		HTTPClient: client,
	}
}

/*CompanyPutParams contains all the parameters to send to the API endpoint
for the company put operation typically these are written to a http.Request
*/
type CompanyPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Company

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the company put params
func (o *CompanyPutParams) WithTimeout(timeout time.Duration) *CompanyPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the company put params
func (o *CompanyPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the company put params
func (o *CompanyPutParams) WithContext(ctx context.Context) *CompanyPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the company put params
func (o *CompanyPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the company put params
func (o *CompanyPutParams) WithHTTPClient(client *http.Client) *CompanyPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the company put params
func (o *CompanyPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the company put params
func (o *CompanyPutParams) WithBody(body *models.Company) *CompanyPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the company put params
func (o *CompanyPutParams) SetBody(body *models.Company) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CompanyPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
