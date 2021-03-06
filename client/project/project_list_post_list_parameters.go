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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewProjectListPostListParams creates a new ProjectListPostListParams object
// with the default values initialized.
func NewProjectListPostListParams() *ProjectListPostListParams {
	var ()
	return &ProjectListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectListPostListParamsWithTimeout creates a new ProjectListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectListPostListParamsWithTimeout(timeout time.Duration) *ProjectListPostListParams {
	var ()
	return &ProjectListPostListParams{

		timeout: timeout,
	}
}

// NewProjectListPostListParamsWithContext creates a new ProjectListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectListPostListParamsWithContext(ctx context.Context) *ProjectListPostListParams {
	var ()
	return &ProjectListPostListParams{

		Context: ctx,
	}
}

// NewProjectListPostListParamsWithHTTPClient creates a new ProjectListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectListPostListParamsWithHTTPClient(client *http.Client) *ProjectListPostListParams {
	var ()
	return &ProjectListPostListParams{
		HTTPClient: client,
	}
}

/*ProjectListPostListParams contains all the parameters to send to the API endpoint
for the project list post list operation typically these are written to a http.Request
*/
type ProjectListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.Project

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project list post list params
func (o *ProjectListPostListParams) WithTimeout(timeout time.Duration) *ProjectListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project list post list params
func (o *ProjectListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project list post list params
func (o *ProjectListPostListParams) WithContext(ctx context.Context) *ProjectListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project list post list params
func (o *ProjectListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project list post list params
func (o *ProjectListPostListParams) WithHTTPClient(client *http.Client) *ProjectListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project list post list params
func (o *ProjectListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project list post list params
func (o *ProjectListPostListParams) WithBody(body []*models.Project) *ProjectListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project list post list params
func (o *ProjectListPostListParams) SetBody(body []*models.Project) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
