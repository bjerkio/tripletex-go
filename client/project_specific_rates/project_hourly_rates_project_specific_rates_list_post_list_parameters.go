// Code generated by go-swagger; DO NOT EDIT.

package project_specific_rates

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

// NewProjectHourlyRatesProjectSpecificRatesListPostListParams creates a new ProjectHourlyRatesProjectSpecificRatesListPostListParams object
// with the default values initialized.
func NewProjectHourlyRatesProjectSpecificRatesListPostListParams() *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithTimeout creates a new ProjectHourlyRatesProjectSpecificRatesListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithTimeout(timeout time.Duration) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListPostListParams{

		timeout: timeout,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithContext creates a new ProjectHourlyRatesProjectSpecificRatesListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithContext(ctx context.Context) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListPostListParams{

		Context: ctx,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithHTTPClient creates a new ProjectHourlyRatesProjectSpecificRatesListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectHourlyRatesProjectSpecificRatesListPostListParamsWithHTTPClient(client *http.Client) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListPostListParams{
		HTTPClient: client,
	}
}

/*ProjectHourlyRatesProjectSpecificRatesListPostListParams contains all the parameters to send to the API endpoint
for the project hourly rates project specific rates list post list operation typically these are written to a http.Request
*/
type ProjectHourlyRatesProjectSpecificRatesListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.ProjectSpecificRate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) WithTimeout(timeout time.Duration) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) WithContext(ctx context.Context) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) WithHTTPClient(client *http.Client) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) WithBody(body []*models.ProjectSpecificRate) *ProjectHourlyRatesProjectSpecificRatesListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project hourly rates project specific rates list post list params
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) SetBody(body []*models.ProjectSpecificRate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectHourlyRatesProjectSpecificRatesListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
