// Code generated by go-swagger; DO NOT EDIT.

package division

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

	"github.com/bjerkio/tripletex-go/models"
)

// NewDivisionPutParams creates a new DivisionPutParams object
// with the default values initialized.
func NewDivisionPutParams() *DivisionPutParams {
	var ()
	return &DivisionPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDivisionPutParamsWithTimeout creates a new DivisionPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDivisionPutParamsWithTimeout(timeout time.Duration) *DivisionPutParams {
	var ()
	return &DivisionPutParams{

		timeout: timeout,
	}
}

// NewDivisionPutParamsWithContext creates a new DivisionPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewDivisionPutParamsWithContext(ctx context.Context) *DivisionPutParams {
	var ()
	return &DivisionPutParams{

		Context: ctx,
	}
}

// NewDivisionPutParamsWithHTTPClient creates a new DivisionPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDivisionPutParamsWithHTTPClient(client *http.Client) *DivisionPutParams {
	var ()
	return &DivisionPutParams{
		HTTPClient: client,
	}
}

/*DivisionPutParams contains all the parameters to send to the API endpoint
for the division put operation typically these are written to a http.Request
*/
type DivisionPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Division
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the division put params
func (o *DivisionPutParams) WithTimeout(timeout time.Duration) *DivisionPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the division put params
func (o *DivisionPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the division put params
func (o *DivisionPutParams) WithContext(ctx context.Context) *DivisionPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the division put params
func (o *DivisionPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the division put params
func (o *DivisionPutParams) WithHTTPClient(client *http.Client) *DivisionPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the division put params
func (o *DivisionPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the division put params
func (o *DivisionPutParams) WithBody(body *models.Division) *DivisionPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the division put params
func (o *DivisionPutParams) SetBody(body *models.Division) {
	o.Body = body
}

// WithID adds the id to the division put params
func (o *DivisionPutParams) WithID(id int32) *DivisionPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the division put params
func (o *DivisionPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DivisionPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
