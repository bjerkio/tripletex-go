// Code generated by go-swagger; DO NOT EDIT.

package order_group

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

// NewOrderOrderGroupPutParams creates a new OrderOrderGroupPutParams object
// with the default values initialized.
func NewOrderOrderGroupPutParams() *OrderOrderGroupPutParams {
	var (
		removeExistingOrderLinesDefault = bool(false)
	)
	return &OrderOrderGroupPutParams{
		RemoveExistingOrderLines: &removeExistingOrderLinesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderOrderGroupPutParamsWithTimeout creates a new OrderOrderGroupPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderOrderGroupPutParamsWithTimeout(timeout time.Duration) *OrderOrderGroupPutParams {
	var (
		removeExistingOrderLinesDefault = bool(false)
	)
	return &OrderOrderGroupPutParams{
		RemoveExistingOrderLines: &removeExistingOrderLinesDefault,

		timeout: timeout,
	}
}

// NewOrderOrderGroupPutParamsWithContext creates a new OrderOrderGroupPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderOrderGroupPutParamsWithContext(ctx context.Context) *OrderOrderGroupPutParams {
	var (
		removeExistingOrderLinesDefault = bool(false)
	)
	return &OrderOrderGroupPutParams{
		RemoveExistingOrderLines: &removeExistingOrderLinesDefault,

		Context: ctx,
	}
}

// NewOrderOrderGroupPutParamsWithHTTPClient creates a new OrderOrderGroupPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderOrderGroupPutParamsWithHTTPClient(client *http.Client) *OrderOrderGroupPutParams {
	var (
		removeExistingOrderLinesDefault = bool(false)
	)
	return &OrderOrderGroupPutParams{
		RemoveExistingOrderLines: &removeExistingOrderLinesDefault,
		HTTPClient:               client,
	}
}

/*OrderOrderGroupPutParams contains all the parameters to send to the API endpoint
for the order order group put operation typically these are written to a http.Request
*/
type OrderOrderGroupPutParams struct {

	/*OrderLineIds
	  List of IDs

	*/
	OrderLineIds *string
	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.OrderGroup
	/*RemoveExistingOrderLines
	  Should existing orderLines be removed from this orderGroup

	*/
	RemoveExistingOrderLines *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order order group put params
func (o *OrderOrderGroupPutParams) WithTimeout(timeout time.Duration) *OrderOrderGroupPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order order group put params
func (o *OrderOrderGroupPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order order group put params
func (o *OrderOrderGroupPutParams) WithContext(ctx context.Context) *OrderOrderGroupPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order order group put params
func (o *OrderOrderGroupPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order order group put params
func (o *OrderOrderGroupPutParams) WithHTTPClient(client *http.Client) *OrderOrderGroupPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order order group put params
func (o *OrderOrderGroupPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderLineIds adds the orderLineIds to the order order group put params
func (o *OrderOrderGroupPutParams) WithOrderLineIds(orderLineIds *string) *OrderOrderGroupPutParams {
	o.SetOrderLineIds(orderLineIds)
	return o
}

// SetOrderLineIds adds the orderLineIds to the order order group put params
func (o *OrderOrderGroupPutParams) SetOrderLineIds(orderLineIds *string) {
	o.OrderLineIds = orderLineIds
}

// WithBody adds the body to the order order group put params
func (o *OrderOrderGroupPutParams) WithBody(body *models.OrderGroup) *OrderOrderGroupPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the order order group put params
func (o *OrderOrderGroupPutParams) SetBody(body *models.OrderGroup) {
	o.Body = body
}

// WithRemoveExistingOrderLines adds the removeExistingOrderLines to the order order group put params
func (o *OrderOrderGroupPutParams) WithRemoveExistingOrderLines(removeExistingOrderLines *bool) *OrderOrderGroupPutParams {
	o.SetRemoveExistingOrderLines(removeExistingOrderLines)
	return o
}

// SetRemoveExistingOrderLines adds the removeExistingOrderLines to the order order group put params
func (o *OrderOrderGroupPutParams) SetRemoveExistingOrderLines(removeExistingOrderLines *bool) {
	o.RemoveExistingOrderLines = removeExistingOrderLines
}

// WriteToRequest writes these params to a swagger request
func (o *OrderOrderGroupPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrderLineIds != nil {

		// query param OrderLineIds
		var qrOrderLineIds string
		if o.OrderLineIds != nil {
			qrOrderLineIds = *o.OrderLineIds
		}
		qOrderLineIds := qrOrderLineIds
		if qOrderLineIds != "" {
			if err := r.SetQueryParam("OrderLineIds", qOrderLineIds); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.RemoveExistingOrderLines != nil {

		// query param removeExistingOrderLines
		var qrRemoveExistingOrderLines bool
		if o.RemoveExistingOrderLines != nil {
			qrRemoveExistingOrderLines = *o.RemoveExistingOrderLines
		}
		qRemoveExistingOrderLines := swag.FormatBool(qrRemoveExistingOrderLines)
		if qRemoveExistingOrderLines != "" {
			if err := r.SetQueryParam("removeExistingOrderLines", qRemoveExistingOrderLines); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
