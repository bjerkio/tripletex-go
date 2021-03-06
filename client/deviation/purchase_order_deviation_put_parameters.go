// Code generated by go-swagger; DO NOT EDIT.

package deviation

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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewPurchaseOrderDeviationPutParams creates a new PurchaseOrderDeviationPutParams object
// with the default values initialized.
func NewPurchaseOrderDeviationPutParams() *PurchaseOrderDeviationPutParams {
	var ()
	return &PurchaseOrderDeviationPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderDeviationPutParamsWithTimeout creates a new PurchaseOrderDeviationPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderDeviationPutParamsWithTimeout(timeout time.Duration) *PurchaseOrderDeviationPutParams {
	var ()
	return &PurchaseOrderDeviationPutParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderDeviationPutParamsWithContext creates a new PurchaseOrderDeviationPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderDeviationPutParamsWithContext(ctx context.Context) *PurchaseOrderDeviationPutParams {
	var ()
	return &PurchaseOrderDeviationPutParams{

		Context: ctx,
	}
}

// NewPurchaseOrderDeviationPutParamsWithHTTPClient creates a new PurchaseOrderDeviationPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderDeviationPutParamsWithHTTPClient(client *http.Client) *PurchaseOrderDeviationPutParams {
	var ()
	return &PurchaseOrderDeviationPutParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderDeviationPutParams contains all the parameters to send to the API endpoint
for the purchase order deviation put operation typically these are written to a http.Request
*/
type PurchaseOrderDeviationPutParams struct {

	/*Body
	  Partial object describing what should be updated

	*/
	Body *models.Deviation
	/*ID
	  Element ID

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) WithTimeout(timeout time.Duration) *PurchaseOrderDeviationPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) WithContext(ctx context.Context) *PurchaseOrderDeviationPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) WithHTTPClient(client *http.Client) *PurchaseOrderDeviationPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) WithBody(body *models.Deviation) *PurchaseOrderDeviationPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) SetBody(body *models.Deviation) {
	o.Body = body
}

// WithID adds the id to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) WithID(id int32) *PurchaseOrderDeviationPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the purchase order deviation put params
func (o *PurchaseOrderDeviationPutParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderDeviationPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
