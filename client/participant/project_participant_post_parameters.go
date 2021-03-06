// Code generated by go-swagger; DO NOT EDIT.

package participant

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

// NewProjectParticipantPostParams creates a new ProjectParticipantPostParams object
// with the default values initialized.
func NewProjectParticipantPostParams() *ProjectParticipantPostParams {
	var ()
	return &ProjectParticipantPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectParticipantPostParamsWithTimeout creates a new ProjectParticipantPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectParticipantPostParamsWithTimeout(timeout time.Duration) *ProjectParticipantPostParams {
	var ()
	return &ProjectParticipantPostParams{

		timeout: timeout,
	}
}

// NewProjectParticipantPostParamsWithContext creates a new ProjectParticipantPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectParticipantPostParamsWithContext(ctx context.Context) *ProjectParticipantPostParams {
	var ()
	return &ProjectParticipantPostParams{

		Context: ctx,
	}
}

// NewProjectParticipantPostParamsWithHTTPClient creates a new ProjectParticipantPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectParticipantPostParamsWithHTTPClient(client *http.Client) *ProjectParticipantPostParams {
	var ()
	return &ProjectParticipantPostParams{
		HTTPClient: client,
	}
}

/*ProjectParticipantPostParams contains all the parameters to send to the API endpoint
for the project participant post operation typically these are written to a http.Request
*/
type ProjectParticipantPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.ProjectParticipant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project participant post params
func (o *ProjectParticipantPostParams) WithTimeout(timeout time.Duration) *ProjectParticipantPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project participant post params
func (o *ProjectParticipantPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project participant post params
func (o *ProjectParticipantPostParams) WithContext(ctx context.Context) *ProjectParticipantPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project participant post params
func (o *ProjectParticipantPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project participant post params
func (o *ProjectParticipantPostParams) WithHTTPClient(client *http.Client) *ProjectParticipantPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project participant post params
func (o *ProjectParticipantPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project participant post params
func (o *ProjectParticipantPostParams) WithBody(body *models.ProjectParticipant) *ProjectParticipantPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project participant post params
func (o *ProjectParticipantPostParams) SetBody(body *models.ProjectParticipant) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectParticipantPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
