// Code generated by go-swagger; DO NOT EDIT.

package purchase_order

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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewPurchaseOrderPostParams creates a new PurchaseOrderPostParams object
// with the default values initialized.
func NewPurchaseOrderPostParams() *PurchaseOrderPostParams {
	var ()
	return &PurchaseOrderPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderPostParamsWithTimeout creates a new PurchaseOrderPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderPostParamsWithTimeout(timeout time.Duration) *PurchaseOrderPostParams {
	var ()
	return &PurchaseOrderPostParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderPostParamsWithContext creates a new PurchaseOrderPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderPostParamsWithContext(ctx context.Context) *PurchaseOrderPostParams {
	var ()
	return &PurchaseOrderPostParams{

		Context: ctx,
	}
}

// NewPurchaseOrderPostParamsWithHTTPClient creates a new PurchaseOrderPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderPostParamsWithHTTPClient(client *http.Client) *PurchaseOrderPostParams {
	var ()
	return &PurchaseOrderPostParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderPostParams contains all the parameters to send to the API endpoint
for the purchase order post operation typically these are written to a http.Request
*/
type PurchaseOrderPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.PurchaseOrder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order post params
func (o *PurchaseOrderPostParams) WithTimeout(timeout time.Duration) *PurchaseOrderPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order post params
func (o *PurchaseOrderPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order post params
func (o *PurchaseOrderPostParams) WithContext(ctx context.Context) *PurchaseOrderPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order post params
func (o *PurchaseOrderPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order post params
func (o *PurchaseOrderPostParams) WithHTTPClient(client *http.Client) *PurchaseOrderPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order post params
func (o *PurchaseOrderPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the purchase order post params
func (o *PurchaseOrderPostParams) WithBody(body *models.PurchaseOrder) *PurchaseOrderPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the purchase order post params
func (o *PurchaseOrderPostParams) SetBody(body *models.PurchaseOrder) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
