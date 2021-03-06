// Code generated by go-swagger; DO NOT EDIT.

package project_specific_rates

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
)

// NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams creates a new ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams object
// with the default values initialized.
func NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams() *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithTimeout creates a new ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithTimeout(timeout time.Duration) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams{

		timeout: timeout,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithContext creates a new ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithContext(ctx context.Context) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams{

		Context: ctx,
	}
}

// NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithHTTPClient creates a new ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParamsWithHTTPClient(client *http.Client) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	var ()
	return &ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams{
		HTTPClient: client,
	}
}

/*ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams contains all the parameters to send to the API endpoint
for the project hourly rates project specific rates list delete by ids operation typically these are written to a http.Request
*/
type ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams struct {

	/*Ids
	  ID of the elements

	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) WithTimeout(timeout time.Duration) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) WithContext(ctx context.Context) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) WithHTTPClient(client *http.Client) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) WithIds(ids string) *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the project hourly rates project specific rates list delete by ids params
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectHourlyRatesProjectSpecificRatesListDeleteByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
