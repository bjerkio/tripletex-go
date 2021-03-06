// Code generated by go-swagger; DO NOT EDIT.

package compilation

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
)

// NewSalaryCompilationPdfDownloadPdfParams creates a new SalaryCompilationPdfDownloadPdfParams object
// with the default values initialized.
func NewSalaryCompilationPdfDownloadPdfParams() *SalaryCompilationPdfDownloadPdfParams {
	var ()
	return &SalaryCompilationPdfDownloadPdfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalaryCompilationPdfDownloadPdfParamsWithTimeout creates a new SalaryCompilationPdfDownloadPdfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalaryCompilationPdfDownloadPdfParamsWithTimeout(timeout time.Duration) *SalaryCompilationPdfDownloadPdfParams {
	var ()
	return &SalaryCompilationPdfDownloadPdfParams{

		timeout: timeout,
	}
}

// NewSalaryCompilationPdfDownloadPdfParamsWithContext creates a new SalaryCompilationPdfDownloadPdfParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalaryCompilationPdfDownloadPdfParamsWithContext(ctx context.Context) *SalaryCompilationPdfDownloadPdfParams {
	var ()
	return &SalaryCompilationPdfDownloadPdfParams{

		Context: ctx,
	}
}

// NewSalaryCompilationPdfDownloadPdfParamsWithHTTPClient creates a new SalaryCompilationPdfDownloadPdfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalaryCompilationPdfDownloadPdfParamsWithHTTPClient(client *http.Client) *SalaryCompilationPdfDownloadPdfParams {
	var ()
	return &SalaryCompilationPdfDownloadPdfParams{
		HTTPClient: client,
	}
}

/*SalaryCompilationPdfDownloadPdfParams contains all the parameters to send to the API endpoint
for the salary compilation pdf download pdf operation typically these are written to a http.Request
*/
type SalaryCompilationPdfDownloadPdfParams struct {

	/*EmployeeID
	  Element ID

	*/
	EmployeeID int32
	/*Year
	  Must be between 1900-2100. Defaults to previous year.

	*/
	Year *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) WithTimeout(timeout time.Duration) *SalaryCompilationPdfDownloadPdfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) WithContext(ctx context.Context) *SalaryCompilationPdfDownloadPdfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) WithHTTPClient(client *http.Client) *SalaryCompilationPdfDownloadPdfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeID adds the employeeID to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) WithEmployeeID(employeeID int32) *SalaryCompilationPdfDownloadPdfParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) SetEmployeeID(employeeID int32) {
	o.EmployeeID = employeeID
}

// WithYear adds the year to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) WithYear(year *int32) *SalaryCompilationPdfDownloadPdfParams {
	o.SetYear(year)
	return o
}

// SetYear adds the year to the salary compilation pdf download pdf params
func (o *SalaryCompilationPdfDownloadPdfParams) SetYear(year *int32) {
	o.Year = year
}

// WriteToRequest writes these params to a swagger request
func (o *SalaryCompilationPdfDownloadPdfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param employeeId
	qrEmployeeID := o.EmployeeID
	qEmployeeID := swag.FormatInt32(qrEmployeeID)
	if qEmployeeID != "" {
		if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
			return err
		}
	}

	if o.Year != nil {

		// query param year
		var qrYear int32
		if o.Year != nil {
			qrYear = *o.Year
		}
		qYear := swag.FormatInt32(qrYear)
		if qYear != "" {
			if err := r.SetQueryParam("year", qYear); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
