// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

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

// NewTravelExpenseApproveApproveParams creates a new TravelExpenseApproveApproveParams object
// with the default values initialized.
func NewTravelExpenseApproveApproveParams() *TravelExpenseApproveApproveParams {
	var ()
	return &TravelExpenseApproveApproveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseApproveApproveParamsWithTimeout creates a new TravelExpenseApproveApproveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseApproveApproveParamsWithTimeout(timeout time.Duration) *TravelExpenseApproveApproveParams {
	var ()
	return &TravelExpenseApproveApproveParams{

		timeout: timeout,
	}
}

// NewTravelExpenseApproveApproveParamsWithContext creates a new TravelExpenseApproveApproveParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseApproveApproveParamsWithContext(ctx context.Context) *TravelExpenseApproveApproveParams {
	var ()
	return &TravelExpenseApproveApproveParams{

		Context: ctx,
	}
}

// NewTravelExpenseApproveApproveParamsWithHTTPClient creates a new TravelExpenseApproveApproveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseApproveApproveParamsWithHTTPClient(client *http.Client) *TravelExpenseApproveApproveParams {
	var ()
	return &TravelExpenseApproveApproveParams{
		HTTPClient: client,
	}
}

/*TravelExpenseApproveApproveParams contains all the parameters to send to the API endpoint
for the travel expense approve approve operation typically these are written to a http.Request
*/
type TravelExpenseApproveApproveParams struct {

	/*ID
	  ID of the elements

	*/
	ID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) WithTimeout(timeout time.Duration) *TravelExpenseApproveApproveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) WithContext(ctx context.Context) *TravelExpenseApproveApproveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) WithHTTPClient(client *http.Client) *TravelExpenseApproveApproveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) WithID(id *string) *TravelExpenseApproveApproveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the travel expense approve approve params
func (o *TravelExpenseApproveApproveParams) SetID(id *string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseApproveApproveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
