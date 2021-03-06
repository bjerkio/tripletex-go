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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewSalarySettingsHolidayListPutListParams creates a new SalarySettingsHolidayListPutListParams object
// with the default values initialized.
func NewSalarySettingsHolidayListPutListParams() *SalarySettingsHolidayListPutListParams {
	var ()
	return &SalarySettingsHolidayListPutListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalarySettingsHolidayListPutListParamsWithTimeout creates a new SalarySettingsHolidayListPutListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalarySettingsHolidayListPutListParamsWithTimeout(timeout time.Duration) *SalarySettingsHolidayListPutListParams {
	var ()
	return &SalarySettingsHolidayListPutListParams{

		timeout: timeout,
	}
}

// NewSalarySettingsHolidayListPutListParamsWithContext creates a new SalarySettingsHolidayListPutListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalarySettingsHolidayListPutListParamsWithContext(ctx context.Context) *SalarySettingsHolidayListPutListParams {
	var ()
	return &SalarySettingsHolidayListPutListParams{

		Context: ctx,
	}
}

// NewSalarySettingsHolidayListPutListParamsWithHTTPClient creates a new SalarySettingsHolidayListPutListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalarySettingsHolidayListPutListParamsWithHTTPClient(client *http.Client) *SalarySettingsHolidayListPutListParams {
	var ()
	return &SalarySettingsHolidayListPutListParams{
		HTTPClient: client,
	}
}

/*SalarySettingsHolidayListPutListParams contains all the parameters to send to the API endpoint
for the salary settings holiday list put list operation typically these are written to a http.Request
*/
type SalarySettingsHolidayListPutListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.CompanyHoliday

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) WithTimeout(timeout time.Duration) *SalarySettingsHolidayListPutListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) WithContext(ctx context.Context) *SalarySettingsHolidayListPutListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) WithHTTPClient(client *http.Client) *SalarySettingsHolidayListPutListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) WithBody(body []*models.CompanyHoliday) *SalarySettingsHolidayListPutListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the salary settings holiday list put list params
func (o *SalarySettingsHolidayListPutListParams) SetBody(body []*models.CompanyHoliday) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SalarySettingsHolidayListPutListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
