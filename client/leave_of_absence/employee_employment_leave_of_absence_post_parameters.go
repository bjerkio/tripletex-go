// Code generated by go-swagger; DO NOT EDIT.

package leave_of_absence

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

// NewEmployeeEmploymentLeaveOfAbsencePostParams creates a new EmployeeEmploymentLeaveOfAbsencePostParams object
// with the default values initialized.
func NewEmployeeEmploymentLeaveOfAbsencePostParams() *EmployeeEmploymentLeaveOfAbsencePostParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsencePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEmploymentLeaveOfAbsencePostParamsWithTimeout creates a new EmployeeEmploymentLeaveOfAbsencePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEmploymentLeaveOfAbsencePostParamsWithTimeout(timeout time.Duration) *EmployeeEmploymentLeaveOfAbsencePostParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsencePostParams{

		timeout: timeout,
	}
}

// NewEmployeeEmploymentLeaveOfAbsencePostParamsWithContext creates a new EmployeeEmploymentLeaveOfAbsencePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEmploymentLeaveOfAbsencePostParamsWithContext(ctx context.Context) *EmployeeEmploymentLeaveOfAbsencePostParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsencePostParams{

		Context: ctx,
	}
}

// NewEmployeeEmploymentLeaveOfAbsencePostParamsWithHTTPClient creates a new EmployeeEmploymentLeaveOfAbsencePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEmploymentLeaveOfAbsencePostParamsWithHTTPClient(client *http.Client) *EmployeeEmploymentLeaveOfAbsencePostParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsencePostParams{
		HTTPClient: client,
	}
}

/*EmployeeEmploymentLeaveOfAbsencePostParams contains all the parameters to send to the API endpoint
for the employee employment leave of absence post operation typically these are written to a http.Request
*/
type EmployeeEmploymentLeaveOfAbsencePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.LeaveOfAbsence

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) WithTimeout(timeout time.Duration) *EmployeeEmploymentLeaveOfAbsencePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) WithContext(ctx context.Context) *EmployeeEmploymentLeaveOfAbsencePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) WithHTTPClient(client *http.Client) *EmployeeEmploymentLeaveOfAbsencePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) WithBody(body *models.LeaveOfAbsence) *EmployeeEmploymentLeaveOfAbsencePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee employment leave of absence post params
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) SetBody(body *models.LeaveOfAbsence) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEmploymentLeaveOfAbsencePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
