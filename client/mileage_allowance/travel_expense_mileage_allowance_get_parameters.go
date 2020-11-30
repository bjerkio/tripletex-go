// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

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

// NewTravelExpenseMileageAllowanceGetParams creates a new TravelExpenseMileageAllowanceGetParams object
// with the default values initialized.
func NewTravelExpenseMileageAllowanceGetParams() *TravelExpenseMileageAllowanceGetParams {
	var ()
	return &TravelExpenseMileageAllowanceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseMileageAllowanceGetParamsWithTimeout creates a new TravelExpenseMileageAllowanceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseMileageAllowanceGetParamsWithTimeout(timeout time.Duration) *TravelExpenseMileageAllowanceGetParams {
	var ()
	return &TravelExpenseMileageAllowanceGetParams{

		timeout: timeout,
	}
}

// NewTravelExpenseMileageAllowanceGetParamsWithContext creates a new TravelExpenseMileageAllowanceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseMileageAllowanceGetParamsWithContext(ctx context.Context) *TravelExpenseMileageAllowanceGetParams {
	var ()
	return &TravelExpenseMileageAllowanceGetParams{

		Context: ctx,
	}
}

// NewTravelExpenseMileageAllowanceGetParamsWithHTTPClient creates a new TravelExpenseMileageAllowanceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseMileageAllowanceGetParamsWithHTTPClient(client *http.Client) *TravelExpenseMileageAllowanceGetParams {
	var ()
	return &TravelExpenseMileageAllowanceGetParams{
		HTTPClient: client,
	}
}

/*TravelExpenseMileageAllowanceGetParams contains all the parameters to send to the API endpoint
for the travel expense mileage allowance get operation typically these are written to a http.Request
*/
type TravelExpenseMileageAllowanceGetParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) WithTimeout(timeout time.Duration) *TravelExpenseMileageAllowanceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) WithContext(ctx context.Context) *TravelExpenseMileageAllowanceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) WithHTTPClient(client *http.Client) *TravelExpenseMileageAllowanceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) WithFields(fields *string) *TravelExpenseMileageAllowanceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) WithID(id int32) *TravelExpenseMileageAllowanceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense mileage allowance get params
func (o *TravelExpenseMileageAllowanceGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseMileageAllowanceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
