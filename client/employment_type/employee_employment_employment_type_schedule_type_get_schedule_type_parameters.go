// Code generated by go-swagger; DO NOT EDIT.

package employment_type

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

// NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams creates a new EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams object
// with the default values initialized.
func NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams() *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithTimeout creates a new EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithTimeout(timeout time.Duration) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithContext creates a new EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithContext(ctx context.Context) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithHTTPClient creates a new EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParamsWithHTTPClient(client *http.Client) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams contains all the parameters to send to the API endpoint
for the employee employment employment type schedule type get schedule type operation typically these are written to a http.Request
*/
type EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithTimeout(timeout time.Duration) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithContext(ctx context.Context) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithHTTPClient(client *http.Client) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithCount(count *int64) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithFields(fields *string) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithFrom(from *int64) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WithSorting(sorting *string) *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the employee employment employment type schedule type get schedule type params
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEmploymentEmploymentTypeScheduleTypeGetScheduleTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
