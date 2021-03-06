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
)

// NewPurchaseOrderDeviationDeliverDeliverParams creates a new PurchaseOrderDeviationDeliverDeliverParams object
// with the default values initialized.
func NewPurchaseOrderDeviationDeliverDeliverParams() *PurchaseOrderDeviationDeliverDeliverParams {
	var ()
	return &PurchaseOrderDeviationDeliverDeliverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderDeviationDeliverDeliverParamsWithTimeout creates a new PurchaseOrderDeviationDeliverDeliverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderDeviationDeliverDeliverParamsWithTimeout(timeout time.Duration) *PurchaseOrderDeviationDeliverDeliverParams {
	var ()
	return &PurchaseOrderDeviationDeliverDeliverParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderDeviationDeliverDeliverParamsWithContext creates a new PurchaseOrderDeviationDeliverDeliverParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderDeviationDeliverDeliverParamsWithContext(ctx context.Context) *PurchaseOrderDeviationDeliverDeliverParams {
	var ()
	return &PurchaseOrderDeviationDeliverDeliverParams{

		Context: ctx,
	}
}

// NewPurchaseOrderDeviationDeliverDeliverParamsWithHTTPClient creates a new PurchaseOrderDeviationDeliverDeliverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderDeviationDeliverDeliverParamsWithHTTPClient(client *http.Client) *PurchaseOrderDeviationDeliverDeliverParams {
	var ()
	return &PurchaseOrderDeviationDeliverDeliverParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderDeviationDeliverDeliverParams contains all the parameters to send to the API endpoint
for the purchase order deviation deliver deliver operation typically these are written to a http.Request
*/
type PurchaseOrderDeviationDeliverDeliverParams struct {

	/*ID
	  Purchase Order ID.

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) WithTimeout(timeout time.Duration) *PurchaseOrderDeviationDeliverDeliverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) WithContext(ctx context.Context) *PurchaseOrderDeviationDeliverDeliverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) WithHTTPClient(client *http.Client) *PurchaseOrderDeviationDeliverDeliverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) WithID(id int32) *PurchaseOrderDeviationDeliverDeliverParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the purchase order deviation deliver deliver params
func (o *PurchaseOrderDeviationDeliverDeliverParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderDeviationDeliverDeliverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
