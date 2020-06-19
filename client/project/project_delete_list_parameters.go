// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewProjectDeleteListParams creates a new ProjectDeleteListParams object
// with the default values initialized.
func NewProjectDeleteListParams() *ProjectDeleteListParams {
	var ()
	return &ProjectDeleteListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectDeleteListParamsWithTimeout creates a new ProjectDeleteListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectDeleteListParamsWithTimeout(timeout time.Duration) *ProjectDeleteListParams {
	var ()
	return &ProjectDeleteListParams{

		timeout: timeout,
	}
}

// NewProjectDeleteListParamsWithContext creates a new ProjectDeleteListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectDeleteListParamsWithContext(ctx context.Context) *ProjectDeleteListParams {
	var ()
	return &ProjectDeleteListParams{

		Context: ctx,
	}
}

// NewProjectDeleteListParamsWithHTTPClient creates a new ProjectDeleteListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectDeleteListParamsWithHTTPClient(client *http.Client) *ProjectDeleteListParams {
	var ()
	return &ProjectDeleteListParams{
		HTTPClient: client,
	}
}

/*ProjectDeleteListParams contains all the parameters to send to the API endpoint
for the project delete list operation typically these are written to a http.Request
*/
type ProjectDeleteListParams struct {

	/*Body
	  JSON representing objects to delete. Should have ID and version set.

	*/
	Body []*models.Project

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project delete list params
func (o *ProjectDeleteListParams) WithTimeout(timeout time.Duration) *ProjectDeleteListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project delete list params
func (o *ProjectDeleteListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project delete list params
func (o *ProjectDeleteListParams) WithContext(ctx context.Context) *ProjectDeleteListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project delete list params
func (o *ProjectDeleteListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project delete list params
func (o *ProjectDeleteListParams) WithHTTPClient(client *http.Client) *ProjectDeleteListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project delete list params
func (o *ProjectDeleteListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project delete list params
func (o *ProjectDeleteListParams) WithBody(body []*models.Project) *ProjectDeleteListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project delete list params
func (o *ProjectDeleteListParams) SetBody(body []*models.Project) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectDeleteListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
