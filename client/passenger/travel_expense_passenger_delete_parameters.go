// Code generated by go-swagger; DO NOT EDIT.

package passenger

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

// NewTravelExpensePassengerDeleteParams creates a new TravelExpensePassengerDeleteParams object
// with the default values initialized.
func NewTravelExpensePassengerDeleteParams() *TravelExpensePassengerDeleteParams {
	var ()
	return &TravelExpensePassengerDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpensePassengerDeleteParamsWithTimeout creates a new TravelExpensePassengerDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpensePassengerDeleteParamsWithTimeout(timeout time.Duration) *TravelExpensePassengerDeleteParams {
	var ()
	return &TravelExpensePassengerDeleteParams{

		timeout: timeout,
	}
}

// NewTravelExpensePassengerDeleteParamsWithContext creates a new TravelExpensePassengerDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpensePassengerDeleteParamsWithContext(ctx context.Context) *TravelExpensePassengerDeleteParams {
	var ()
	return &TravelExpensePassengerDeleteParams{

		Context: ctx,
	}
}

// NewTravelExpensePassengerDeleteParamsWithHTTPClient creates a new TravelExpensePassengerDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpensePassengerDeleteParamsWithHTTPClient(client *http.Client) *TravelExpensePassengerDeleteParams {
	var ()
	return &TravelExpensePassengerDeleteParams{
		HTTPClient: client,
	}
}

/*TravelExpensePassengerDeleteParams contains all the parameters to send to the API endpoint
for the travel expense passenger delete operation typically these are written to a http.Request
*/
type TravelExpensePassengerDeleteParams struct {

	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) WithTimeout(timeout time.Duration) *TravelExpensePassengerDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) WithContext(ctx context.Context) *TravelExpensePassengerDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) WithHTTPClient(client *http.Client) *TravelExpensePassengerDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) WithID(id int32) *TravelExpensePassengerDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense passenger delete params
func (o *TravelExpensePassengerDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpensePassengerDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
