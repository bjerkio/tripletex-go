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

	"github.com/bjerkio/tripletex-go/models"
)

// NewSalarySettingsHolidayListPostListParams creates a new SalarySettingsHolidayListPostListParams object
// with the default values initialized.
func NewSalarySettingsHolidayListPostListParams() *SalarySettingsHolidayListPostListParams {
	var ()
	return &SalarySettingsHolidayListPostListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalarySettingsHolidayListPostListParamsWithTimeout creates a new SalarySettingsHolidayListPostListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalarySettingsHolidayListPostListParamsWithTimeout(timeout time.Duration) *SalarySettingsHolidayListPostListParams {
	var ()
	return &SalarySettingsHolidayListPostListParams{

		timeout: timeout,
	}
}

// NewSalarySettingsHolidayListPostListParamsWithContext creates a new SalarySettingsHolidayListPostListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalarySettingsHolidayListPostListParamsWithContext(ctx context.Context) *SalarySettingsHolidayListPostListParams {
	var ()
	return &SalarySettingsHolidayListPostListParams{

		Context: ctx,
	}
}

// NewSalarySettingsHolidayListPostListParamsWithHTTPClient creates a new SalarySettingsHolidayListPostListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalarySettingsHolidayListPostListParamsWithHTTPClient(client *http.Client) *SalarySettingsHolidayListPostListParams {
	var ()
	return &SalarySettingsHolidayListPostListParams{
		HTTPClient: client,
	}
}

/*SalarySettingsHolidayListPostListParams contains all the parameters to send to the API endpoint
for the salary settings holiday list post list operation typically these are written to a http.Request
*/
type SalarySettingsHolidayListPostListParams struct {

	/*Body
	  JSON representing a list of new object to be created. Should not have ID and version set.

	*/
	Body []*models.CompanyHoliday

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) WithTimeout(timeout time.Duration) *SalarySettingsHolidayListPostListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) WithContext(ctx context.Context) *SalarySettingsHolidayListPostListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) WithHTTPClient(client *http.Client) *SalarySettingsHolidayListPostListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) WithBody(body []*models.CompanyHoliday) *SalarySettingsHolidayListPostListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the salary settings holiday list post list params
func (o *SalarySettingsHolidayListPostListParams) SetBody(body []*models.CompanyHoliday) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SalarySettingsHolidayListPostListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
