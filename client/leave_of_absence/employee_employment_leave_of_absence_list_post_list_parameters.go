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

// NewEmployeeEmploymentLeaveOfAbsenceListPostListParams creates a new EmployeeEmploymentLeaveOfAbsenceListPostListParams object
// with the default values initialized.
func NewEmployeeEmploymentLeaveOfAbsenceListPostListParams() *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsenceListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithTimeout creates a new EmployeeEmploymentLeaveOfAbsenceListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithTimeout(timeout time.Duration) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsenceListPostListParams{

		timeout: timeout,
	}
}

// NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithContext creates a new EmployeeEmploymentLeaveOfAbsenceListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithContext(ctx context.Context) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsenceListPostListParams{

		Context: ctx,
	}
}

// NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithHTTPClient creates a new EmployeeEmploymentLeaveOfAbsenceListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEmploymentLeaveOfAbsenceListPostListParamsWithHTTPClient(client *http.Client) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	var ()
	return &EmployeeEmploymentLeaveOfAbsenceListPostListParams{
		HTTPClient: client,
	}
}

/*EmployeeEmploymentLeaveOfAbsenceListPostListParams contains all the parameters to send to the API endpoint
for the employee employment leave of absence list post list operation typically these are written to a http.Request
*/
type EmployeeEmploymentLeaveOfAbsenceListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.LeaveOfAbsence

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) WithTimeout(timeout time.Duration) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) WithContext(ctx context.Context) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) WithHTTPClient(client *http.Client) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) WithBody(body []*models.LeaveOfAbsence) *EmployeeEmploymentLeaveOfAbsenceListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the employee employment leave of absence list post list params
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) SetBody(body []*models.LeaveOfAbsence) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEmploymentLeaveOfAbsenceListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
