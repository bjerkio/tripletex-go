// Code generated by go-swagger; DO NOT EDIT.

package standard_time

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

// NewEmployeeStandardTimePostParams creates a new EmployeeStandardTimePostParams object
// with the default values initialized.
func NewEmployeeStandardTimePostParams() *EmployeeStandardTimePostParams {
	var ()
	return &EmployeeStandardTimePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeStandardTimePostParamsWithTimeout creates a new EmployeeStandardTimePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeStandardTimePostParamsWithTimeout(timeout time.Duration) *EmployeeStandardTimePostParams {
	var ()
	return &EmployeeStandardTimePostParams{

		timeout: timeout,
	}
}

// NewEmployeeStandardTimePostParamsWithContext creates a new EmployeeStandardTimePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeStandardTimePostParamsWithContext(ctx context.Context) *EmployeeStandardTimePostParams {
	var ()
	return &EmployeeStandardTimePostParams{

		Context: ctx,
	}
}

// NewEmployeeStandardTimePostParamsWithHTTPClient creates a new EmployeeStandardTimePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeStandardTimePostParamsWithHTTPClient(client *http.Client) *EmployeeStandardTimePostParams {
	var ()
	return &EmployeeStandardTimePostParams{
		HTTPClient: client,
	}
}

/*EmployeeStandardTimePostParams contains all the parameters to send to the API endpoint
for the employee standard time post operation typically these are written to a http.Request
*/
type EmployeeStandardTimePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.StandardTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee standard time post params
func (o *EmployeeStandardTimePostParams) WithTimeout(timeout time.Duration) *EmployeeStandardTimePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee standard time post params
func (o *EmployeeStandardTimePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee standard time post params
func (o *EmployeeStandardTimePostParams) WithContext(ctx context.Context) *EmployeeStandardTimePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee standard time post params
func (o *EmployeeStandardTimePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee standard time post params
func (o *EmployeeStandardTimePostParams) WithHTTPClient(client *http.Client) *EmployeeStandardTimePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee standard time post params
func (o *EmployeeStandardTimePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee standard time post params
func (o *EmployeeStandardTimePostParams) WithBody(body *models.StandardTime) *EmployeeStandardTimePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee standard time post params
func (o *EmployeeStandardTimePostParams) SetBody(body *models.StandardTime) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeStandardTimePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
