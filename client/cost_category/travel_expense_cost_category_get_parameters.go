// Code generated by go-swagger; DO NOT EDIT.

package cost_category

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

// NewTravelExpenseCostCategoryGetParams creates a new TravelExpenseCostCategoryGetParams object
// with the default values initialized.
func NewTravelExpenseCostCategoryGetParams() *TravelExpenseCostCategoryGetParams {
	var ()
	return &TravelExpenseCostCategoryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseCostCategoryGetParamsWithTimeout creates a new TravelExpenseCostCategoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseCostCategoryGetParamsWithTimeout(timeout time.Duration) *TravelExpenseCostCategoryGetParams {
	var ()
	return &TravelExpenseCostCategoryGetParams{

		timeout: timeout,
	}
}

// NewTravelExpenseCostCategoryGetParamsWithContext creates a new TravelExpenseCostCategoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseCostCategoryGetParamsWithContext(ctx context.Context) *TravelExpenseCostCategoryGetParams {
	var ()
	return &TravelExpenseCostCategoryGetParams{

		Context: ctx,
	}
}

// NewTravelExpenseCostCategoryGetParamsWithHTTPClient creates a new TravelExpenseCostCategoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseCostCategoryGetParamsWithHTTPClient(client *http.Client) *TravelExpenseCostCategoryGetParams {
	var ()
	return &TravelExpenseCostCategoryGetParams{
		HTTPClient: client,
	}
}

/*TravelExpenseCostCategoryGetParams contains all the parameters to send to the API endpoint
for the travel expense cost category get operation typically these are written to a http.Request
*/
type TravelExpenseCostCategoryGetParams struct {

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

// WithTimeout adds the timeout to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) WithTimeout(timeout time.Duration) *TravelExpenseCostCategoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) WithContext(ctx context.Context) *TravelExpenseCostCategoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) WithHTTPClient(client *http.Client) *TravelExpenseCostCategoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) WithFields(fields *string) *TravelExpenseCostCategoryGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) WithID(id int32) *TravelExpenseCostCategoryGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense cost category get params
func (o *TravelExpenseCostCategoryGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseCostCategoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
