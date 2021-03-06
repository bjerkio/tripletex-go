// Code generated by go-swagger; DO NOT EDIT.

package accommodation_allowance

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

// NewTravelExpenseAccommodationAllowanceDeleteParams creates a new TravelExpenseAccommodationAllowanceDeleteParams object
// with the default values initialized.
func NewTravelExpenseAccommodationAllowanceDeleteParams() *TravelExpenseAccommodationAllowanceDeleteParams {
	var ()
	return &TravelExpenseAccommodationAllowanceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseAccommodationAllowanceDeleteParamsWithTimeout creates a new TravelExpenseAccommodationAllowanceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseAccommodationAllowanceDeleteParamsWithTimeout(timeout time.Duration) *TravelExpenseAccommodationAllowanceDeleteParams {
	var ()
	return &TravelExpenseAccommodationAllowanceDeleteParams{

		timeout: timeout,
	}
}

// NewTravelExpenseAccommodationAllowanceDeleteParamsWithContext creates a new TravelExpenseAccommodationAllowanceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseAccommodationAllowanceDeleteParamsWithContext(ctx context.Context) *TravelExpenseAccommodationAllowanceDeleteParams {
	var ()
	return &TravelExpenseAccommodationAllowanceDeleteParams{

		Context: ctx,
	}
}

// NewTravelExpenseAccommodationAllowanceDeleteParamsWithHTTPClient creates a new TravelExpenseAccommodationAllowanceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseAccommodationAllowanceDeleteParamsWithHTTPClient(client *http.Client) *TravelExpenseAccommodationAllowanceDeleteParams {
	var ()
	return &TravelExpenseAccommodationAllowanceDeleteParams{
		HTTPClient: client,
	}
}

/*TravelExpenseAccommodationAllowanceDeleteParams contains all the parameters to send to the API endpoint
for the travel expense accommodation allowance delete operation typically these are written to a http.Request
*/
type TravelExpenseAccommodationAllowanceDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) WithTimeout(timeout time.Duration) *TravelExpenseAccommodationAllowanceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) WithContext(ctx context.Context) *TravelExpenseAccommodationAllowanceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) WithHTTPClient(client *http.Client) *TravelExpenseAccommodationAllowanceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) WithID(id int32) *TravelExpenseAccommodationAllowanceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense accommodation allowance delete params
func (o *TravelExpenseAccommodationAllowanceDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseAccommodationAllowanceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
