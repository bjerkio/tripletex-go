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

// NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams creates a new EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams object
// with the default values initialized.
func NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams() *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithTimeout creates a new EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithTimeout(timeout time.Duration) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithContext creates a new EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithContext(ctx context.Context) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithHTTPClient creates a new EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParamsWithHTTPClient(client *http.Client) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams contains all the parameters to send to the API endpoint
for the employee employment employment type employment end reason type get employment end reason type operation typically these are written to a http.Request
*/
type EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams struct {

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

// WithTimeout adds the timeout to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithTimeout(timeout time.Duration) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithContext(ctx context.Context) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithHTTPClient(client *http.Client) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithCount(count *int64) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithFields(fields *string) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithFrom(from *int64) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetFrom(from *int64) {
	o.From = from
}

// WithSorting adds the sorting to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WithSorting(sorting *string) *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the employee employment employment type employment end reason type get employment end reason type params
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEmploymentEmploymentTypeEmploymentEndReasonTypeGetEmploymentEndReasonTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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