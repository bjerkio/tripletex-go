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

	"github.com/bjerkio/tripletex-go/models"
)

// NewPurchaseOrderDeviationListPutListParams creates a new PurchaseOrderDeviationListPutListParams object
// with the default values initialized.
func NewPurchaseOrderDeviationListPutListParams() *PurchaseOrderDeviationListPutListParams {
	var ()
	return &PurchaseOrderDeviationListPutListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderDeviationListPutListParamsWithTimeout creates a new PurchaseOrderDeviationListPutListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderDeviationListPutListParamsWithTimeout(timeout time.Duration) *PurchaseOrderDeviationListPutListParams {
	var ()
	return &PurchaseOrderDeviationListPutListParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderDeviationListPutListParamsWithContext creates a new PurchaseOrderDeviationListPutListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderDeviationListPutListParamsWithContext(ctx context.Context) *PurchaseOrderDeviationListPutListParams {
	var ()
	return &PurchaseOrderDeviationListPutListParams{

		Context: ctx,
	}
}

// NewPurchaseOrderDeviationListPutListParamsWithHTTPClient creates a new PurchaseOrderDeviationListPutListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderDeviationListPutListParamsWithHTTPClient(client *http.Client) *PurchaseOrderDeviationListPutListParams {
	var ()
	return &PurchaseOrderDeviationListPutListParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderDeviationListPutListParams contains all the parameters to send to the API endpoint
for the purchase order deviation list put list operation typically these are written to a http.Request
*/
type PurchaseOrderDeviationListPutListParams struct {

	/*Body
	  JSON representing updates to object. Should have ID and version set.

	*/
	Body []*models.Deviation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) WithTimeout(timeout time.Duration) *PurchaseOrderDeviationListPutListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) WithContext(ctx context.Context) *PurchaseOrderDeviationListPutListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) WithHTTPClient(client *http.Client) *PurchaseOrderDeviationListPutListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) WithBody(body []*models.Deviation) *PurchaseOrderDeviationListPutListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the purchase order deviation list put list params
func (o *PurchaseOrderDeviationListPutListParams) SetBody(body []*models.Deviation) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderDeviationListPutListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
