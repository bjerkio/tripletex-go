// Code generated by go-swagger; DO NOT EDIT.

package holiday

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

// NewSalarySettingsHolidayPutParams creates a new SalarySettingsHolidayPutParams object
// with the default values initialized.
func NewSalarySettingsHolidayPutParams() *SalarySettingsHolidayPutParams {
	var ()
	return &SalarySettingsHolidayPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalarySettingsHolidayPutParamsWithTimeout creates a new SalarySettingsHolidayPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalarySettingsHolidayPutParamsWithTimeout(timeout time.Duration) *SalarySettingsHolidayPutParams {
	var ()
	return &SalarySettingsHolidayPutParams{

		timeout: timeout,
	}
}

// NewSalarySettingsHolidayPutParamsWithContext creates a new SalarySettingsHolidayPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalarySettingsHolidayPutParamsWithContext(ctx context.Context) *SalarySettingsHolidayPutParams {
	var ()
	return &SalarySettingsHolidayPutParams{

		Context: ctx,
	}
}

// NewSalarySettingsHolidayPutParamsWithHTTPClient creates a new SalarySettingsHolidayPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalarySettingsHolidayPutParamsWithHTTPClient(client *http.Client) *SalarySettingsHolidayPutParams {
	var ()
	return &SalarySettingsHolidayPutParams{
		HTTPClient: client,
	}
}

/*SalarySettingsHolidayPutParams contains all the parameters to send to the API endpoint
for the salary settings holiday put operation typically these are written to a http.Request
*/
type SalarySettingsHolidayPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.CompanyHoliday
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) WithTimeout(timeout time.Duration) *SalarySettingsHolidayPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) WithContext(ctx context.Context) *SalarySettingsHolidayPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) WithHTTPClient(client *http.Client) *SalarySettingsHolidayPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) WithBody(body *models.CompanyHoliday) *SalarySettingsHolidayPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) SetBody(body *models.CompanyHoliday) {
	o.Body = body
}

// WithID adds the id to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) WithID(id int32) *SalarySettingsHolidayPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the salary settings holiday put params
func (o *SalarySettingsHolidayPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalarySettingsHolidayPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
