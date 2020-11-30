// Code generated by go-swagger; DO NOT EDIT.

package hourly_rates

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

	"github.com/bjerkio/tripletex-go/models"
)

// NewProjectHourlyRatesPutParams creates a new ProjectHourlyRatesPutParams object
// with the default values initialized.
func NewProjectHourlyRatesPutParams() *ProjectHourlyRatesPutParams {
	var ()
	return &ProjectHourlyRatesPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectHourlyRatesPutParamsWithTimeout creates a new ProjectHourlyRatesPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectHourlyRatesPutParamsWithTimeout(timeout time.Duration) *ProjectHourlyRatesPutParams {
	var ()
	return &ProjectHourlyRatesPutParams{

		timeout: timeout,
	}
}

// NewProjectHourlyRatesPutParamsWithContext creates a new ProjectHourlyRatesPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectHourlyRatesPutParamsWithContext(ctx context.Context) *ProjectHourlyRatesPutParams {
	var ()
	return &ProjectHourlyRatesPutParams{

		Context: ctx,
	}
}

// NewProjectHourlyRatesPutParamsWithHTTPClient creates a new ProjectHourlyRatesPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectHourlyRatesPutParamsWithHTTPClient(client *http.Client) *ProjectHourlyRatesPutParams {
	var ()
	return &ProjectHourlyRatesPutParams{
		HTTPClient: client,
	}
}

/*ProjectHourlyRatesPutParams contains all the parameters to send to the API endpoint
for the project hourly rates put operation typically these are written to a http.Request
*/
type ProjectHourlyRatesPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.ProjectHourlyRate
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) WithTimeout(timeout time.Duration) *ProjectHourlyRatesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) WithContext(ctx context.Context) *ProjectHourlyRatesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) WithHTTPClient(client *http.Client) *ProjectHourlyRatesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) WithBody(body *models.ProjectHourlyRate) *ProjectHourlyRatesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) SetBody(body *models.ProjectHourlyRate) {
	o.Body = body
}

// WithID adds the id to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) WithID(id int32) *ProjectHourlyRatesPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project hourly rates put params
func (o *ProjectHourlyRatesPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectHourlyRatesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
