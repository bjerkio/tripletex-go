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
	"github.com/go-openapi/swag"
)

// NewPurchaseOrderGoodsReceiptLineGetParams creates a new PurchaseOrderGoodsReceiptLineGetParams object
// with the default values initialized.
func NewPurchaseOrderGoodsReceiptLineGetParams() *PurchaseOrderGoodsReceiptLineGetParams {
	var ()
	return &PurchaseOrderGoodsReceiptLineGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderGoodsReceiptLineGetParamsWithTimeout creates a new PurchaseOrderGoodsReceiptLineGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderGoodsReceiptLineGetParamsWithTimeout(timeout time.Duration) *PurchaseOrderGoodsReceiptLineGetParams {
	var ()
	return &PurchaseOrderGoodsReceiptLineGetParams{

		timeout: timeout,
	}
}

// NewPurchaseOrderGoodsReceiptLineGetParamsWithContext creates a new PurchaseOrderGoodsReceiptLineGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderGoodsReceiptLineGetParamsWithContext(ctx context.Context) *PurchaseOrderGoodsReceiptLineGetParams {
	var ()
	return &PurchaseOrderGoodsReceiptLineGetParams{

		Context: ctx,
	}
}

// NewPurchaseOrderGoodsReceiptLineGetParamsWithHTTPClient creates a new PurchaseOrderGoodsReceiptLineGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderGoodsReceiptLineGetParamsWithHTTPClient(client *http.Client) *PurchaseOrderGoodsReceiptLineGetParams {
	var ()
	return &PurchaseOrderGoodsReceiptLineGetParams{
		HTTPClient: client,
	}
}

/*PurchaseOrderGoodsReceiptLineGetParams contains all the parameters to send to the API endpoint
for the purchase order goods receipt line get operation typically these are written to a http.Request
*/
type PurchaseOrderGoodsReceiptLineGetParams struct {

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

// WithTimeout adds the timeout to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) WithTimeout(timeout time.Duration) *PurchaseOrderGoodsReceiptLineGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) WithContext(ctx context.Context) *PurchaseOrderGoodsReceiptLineGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) WithHTTPClient(client *http.Client) *PurchaseOrderGoodsReceiptLineGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) WithFields(fields *string) *PurchaseOrderGoodsReceiptLineGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) WithID(id int32) *PurchaseOrderGoodsReceiptLineGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the purchase order goods receipt line get params
func (o *PurchaseOrderGoodsReceiptLineGetParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderGoodsReceiptLineGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
