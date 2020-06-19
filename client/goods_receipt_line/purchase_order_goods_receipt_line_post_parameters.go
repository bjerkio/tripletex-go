// Code generated by go-swagger; DO NOT EDIT.

package goods_receipt_line

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

// NewPurchaseOrderGoodsReceiptLinePostParams creates a new PurchaseOrderGoodsReceiptLinePostParams object
// with the default values initialized.
func NewPurchaseOrderGoodsReceiptLinePostParams() *PurchaseOrderGoodsReceiptLinePostParams {
	var ()
	return &PurchaseOrderGoodsReceiptLinePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderGoodsReceiptLinePostParamsWithTimeout creates a new PurchaseOrderGoodsReceiptLinePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderGoodsReceiptLinePostParamsWithTimeout(timeout time.Duration) *PurchaseOrderGoodsReceiptLinePostParams {
	var ()
	return &PurchaseOrderGoodsReceiptLinePostParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderGoodsReceiptLinePostParamsWithContext creates a new PurchaseOrderGoodsReceiptLinePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderGoodsReceiptLinePostParamsWithContext(ctx context.Context) *PurchaseOrderGoodsReceiptLinePostParams {
	var ()
	return &PurchaseOrderGoodsReceiptLinePostParams{

		Context: ctx,
	}
}

// NewPurchaseOrderGoodsReceiptLinePostParamsWithHTTPClient creates a new PurchaseOrderGoodsReceiptLinePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderGoodsReceiptLinePostParamsWithHTTPClient(client *http.Client) *PurchaseOrderGoodsReceiptLinePostParams {
	var ()
	return &PurchaseOrderGoodsReceiptLinePostParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderGoodsReceiptLinePostParams contains all the parameters to send to the API endpoint
for the purchase order goods receipt line post operation typically these are written to a http.Request
*/
type PurchaseOrderGoodsReceiptLinePostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.GoodsReceiptLine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) WithTimeout(timeout time.Duration) *PurchaseOrderGoodsReceiptLinePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) WithContext(ctx context.Context) *PurchaseOrderGoodsReceiptLinePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) WithHTTPClient(client *http.Client) *PurchaseOrderGoodsReceiptLinePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) WithBody(body *models.GoodsReceiptLine) *PurchaseOrderGoodsReceiptLinePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the purchase order goods receipt line post params
func (o *PurchaseOrderGoodsReceiptLinePostParams) SetBody(body *models.GoodsReceiptLine) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderGoodsReceiptLinePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
