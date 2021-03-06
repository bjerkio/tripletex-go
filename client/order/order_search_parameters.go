// Code generated by go-swagger; DO NOT EDIT.

package order

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

// NewOrderSearchParams creates a new OrderSearchParams object
// with the default values initialized.
func NewOrderSearchParams() *OrderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &OrderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderSearchParamsWithTimeout creates a new OrderSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderSearchParamsWithTimeout(timeout time.Duration) *OrderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &OrderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewOrderSearchParamsWithContext creates a new OrderSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderSearchParamsWithContext(ctx context.Context) *OrderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &OrderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewOrderSearchParamsWithHTTPClient creates a new OrderSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderSearchParamsWithHTTPClient(client *http.Client) *OrderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &OrderSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*OrderSearchParams contains all the parameters to send to the API endpoint
for the order search operation typically these are written to a http.Request
*/
type OrderSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  List of IDs

	*/
	CustomerID *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*ID
	  List of IDs

	*/
	ID *string
	/*IsClosed
	  Equals

	*/
	IsClosed *bool
	/*IsSubscription
	  Equals

	*/
	IsSubscription *bool
	/*Number
	  Equals

	*/
	Number *string
	/*OrderDateFrom
	  From and including

	*/
	OrderDateFrom string
	/*OrderDateTo
	  To and excluding

	*/
	OrderDateTo string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order search params
func (o *OrderSearchParams) WithTimeout(timeout time.Duration) *OrderSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order search params
func (o *OrderSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order search params
func (o *OrderSearchParams) WithContext(ctx context.Context) *OrderSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order search params
func (o *OrderSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order search params
func (o *OrderSearchParams) WithHTTPClient(client *http.Client) *OrderSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order search params
func (o *OrderSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the order search params
func (o *OrderSearchParams) WithCount(count *int64) *OrderSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the order search params
func (o *OrderSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the order search params
func (o *OrderSearchParams) WithCustomerID(customerID *string) *OrderSearchParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the order search params
func (o *OrderSearchParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithFields adds the fields to the order search params
func (o *OrderSearchParams) WithFields(fields *string) *OrderSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the order search params
func (o *OrderSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the order search params
func (o *OrderSearchParams) WithFrom(from *int64) *OrderSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the order search params
func (o *OrderSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the order search params
func (o *OrderSearchParams) WithID(id *string) *OrderSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the order search params
func (o *OrderSearchParams) SetID(id *string) {
	o.ID = id
}

// WithIsClosed adds the isClosed to the order search params
func (o *OrderSearchParams) WithIsClosed(isClosed *bool) *OrderSearchParams {
	o.SetIsClosed(isClosed)
	return o
}

// SetIsClosed adds the isClosed to the order search params
func (o *OrderSearchParams) SetIsClosed(isClosed *bool) {
	o.IsClosed = isClosed
}

// WithIsSubscription adds the isSubscription to the order search params
func (o *OrderSearchParams) WithIsSubscription(isSubscription *bool) *OrderSearchParams {
	o.SetIsSubscription(isSubscription)
	return o
}

// SetIsSubscription adds the isSubscription to the order search params
func (o *OrderSearchParams) SetIsSubscription(isSubscription *bool) {
	o.IsSubscription = isSubscription
}

// WithNumber adds the number to the order search params
func (o *OrderSearchParams) WithNumber(number *string) *OrderSearchParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the order search params
func (o *OrderSearchParams) SetNumber(number *string) {
	o.Number = number
}

// WithOrderDateFrom adds the orderDateFrom to the order search params
func (o *OrderSearchParams) WithOrderDateFrom(orderDateFrom string) *OrderSearchParams {
	o.SetOrderDateFrom(orderDateFrom)
	return o
}

// SetOrderDateFrom adds the orderDateFrom to the order search params
func (o *OrderSearchParams) SetOrderDateFrom(orderDateFrom string) {
	o.OrderDateFrom = orderDateFrom
}

// WithOrderDateTo adds the orderDateTo to the order search params
func (o *OrderSearchParams) WithOrderDateTo(orderDateTo string) *OrderSearchParams {
	o.SetOrderDateTo(orderDateTo)
	return o
}

// SetOrderDateTo adds the orderDateTo to the order search params
func (o *OrderSearchParams) SetOrderDateTo(orderDateTo string) {
	o.OrderDateTo = orderDateTo
}

// WithSorting adds the sorting to the order search params
func (o *OrderSearchParams) WithSorting(sorting *string) *OrderSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the order search params
func (o *OrderSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *OrderSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.CustomerID != nil {

		// query param customerId
		var qrCustomerID string
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := qrCustomerID
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

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

	if o.From != nil {

		// query param from
		var qrFrom int64
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

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

	if o.IsClosed != nil {

		// query param isClosed
		var qrIsClosed bool
		if o.IsClosed != nil {
			qrIsClosed = *o.IsClosed
		}
		qIsClosed := swag.FormatBool(qrIsClosed)
		if qIsClosed != "" {
			if err := r.SetQueryParam("isClosed", qIsClosed); err != nil {
				return err
			}
		}

	}

	if o.IsSubscription != nil {

		// query param isSubscription
		var qrIsSubscription bool
		if o.IsSubscription != nil {
			qrIsSubscription = *o.IsSubscription
		}
		qIsSubscription := swag.FormatBool(qrIsSubscription)
		if qIsSubscription != "" {
			if err := r.SetQueryParam("isSubscription", qIsSubscription); err != nil {
				return err
			}
		}

	}

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
				return err
			}
		}

	}

	// query param orderDateFrom
	qrOrderDateFrom := o.OrderDateFrom
	qOrderDateFrom := qrOrderDateFrom
	if qOrderDateFrom != "" {
		if err := r.SetQueryParam("orderDateFrom", qOrderDateFrom); err != nil {
			return err
		}
	}

	// query param orderDateTo
	qrOrderDateTo := o.OrderDateTo
	qOrderDateTo := qrOrderDateTo
	if qOrderDateTo != "" {
		if err := r.SetQueryParam("orderDateTo", qOrderDateTo); err != nil {
			return err
		}
	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
