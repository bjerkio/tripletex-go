// Code generated by go-swagger; DO NOT EDIT.

package reminder

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

// NewReminderSearchParams creates a new ReminderSearchParams object
// with the default values initialized.
func NewReminderSearchParams() *ReminderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ReminderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReminderSearchParamsWithTimeout creates a new ReminderSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReminderSearchParamsWithTimeout(timeout time.Duration) *ReminderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ReminderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewReminderSearchParamsWithContext creates a new ReminderSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewReminderSearchParamsWithContext(ctx context.Context) *ReminderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ReminderSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewReminderSearchParamsWithHTTPClient creates a new ReminderSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReminderSearchParamsWithHTTPClient(client *http.Client) *ReminderSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ReminderSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*ReminderSearchParams contains all the parameters to send to the API endpoint
for the reminder search operation typically these are written to a http.Request
*/
type ReminderSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CustomerID
	  Equals

	*/
	CustomerID *int32
	/*DateFrom
	  From and including

	*/
	DateFrom string
	/*DateTo
	  To and excluding

	*/
	DateTo string
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
	/*InvoiceID
	  Equals

	*/
	InvoiceID *int32
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*TermOfPaymentFrom
	  From and including

	*/
	TermOfPaymentFrom *string
	/*TermOfPaymentTo
	  To and excluding

	*/
	TermOfPaymentTo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reminder search params
func (o *ReminderSearchParams) WithTimeout(timeout time.Duration) *ReminderSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reminder search params
func (o *ReminderSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reminder search params
func (o *ReminderSearchParams) WithContext(ctx context.Context) *ReminderSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reminder search params
func (o *ReminderSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reminder search params
func (o *ReminderSearchParams) WithHTTPClient(client *http.Client) *ReminderSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reminder search params
func (o *ReminderSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the reminder search params
func (o *ReminderSearchParams) WithCount(count *int64) *ReminderSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the reminder search params
func (o *ReminderSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithCustomerID adds the customerID to the reminder search params
func (o *ReminderSearchParams) WithCustomerID(customerID *int32) *ReminderSearchParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the reminder search params
func (o *ReminderSearchParams) SetCustomerID(customerID *int32) {
	o.CustomerID = customerID
}

// WithDateFrom adds the dateFrom to the reminder search params
func (o *ReminderSearchParams) WithDateFrom(dateFrom string) *ReminderSearchParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the reminder search params
func (o *ReminderSearchParams) SetDateFrom(dateFrom string) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the reminder search params
func (o *ReminderSearchParams) WithDateTo(dateTo string) *ReminderSearchParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the reminder search params
func (o *ReminderSearchParams) SetDateTo(dateTo string) {
	o.DateTo = dateTo
}

// WithFields adds the fields to the reminder search params
func (o *ReminderSearchParams) WithFields(fields *string) *ReminderSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the reminder search params
func (o *ReminderSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the reminder search params
func (o *ReminderSearchParams) WithFrom(from *int64) *ReminderSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the reminder search params
func (o *ReminderSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the reminder search params
func (o *ReminderSearchParams) WithID(id *string) *ReminderSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reminder search params
func (o *ReminderSearchParams) SetID(id *string) {
	o.ID = id
}

// WithInvoiceID adds the invoiceID to the reminder search params
func (o *ReminderSearchParams) WithInvoiceID(invoiceID *int32) *ReminderSearchParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the reminder search params
func (o *ReminderSearchParams) SetInvoiceID(invoiceID *int32) {
	o.InvoiceID = invoiceID
}

// WithSorting adds the sorting to the reminder search params
func (o *ReminderSearchParams) WithSorting(sorting *string) *ReminderSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the reminder search params
func (o *ReminderSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithTermOfPaymentFrom adds the termOfPaymentFrom to the reminder search params
func (o *ReminderSearchParams) WithTermOfPaymentFrom(termOfPaymentFrom *string) *ReminderSearchParams {
	o.SetTermOfPaymentFrom(termOfPaymentFrom)
	return o
}

// SetTermOfPaymentFrom adds the termOfPaymentFrom to the reminder search params
func (o *ReminderSearchParams) SetTermOfPaymentFrom(termOfPaymentFrom *string) {
	o.TermOfPaymentFrom = termOfPaymentFrom
}

// WithTermOfPaymentTo adds the termOfPaymentTo to the reminder search params
func (o *ReminderSearchParams) WithTermOfPaymentTo(termOfPaymentTo *string) *ReminderSearchParams {
	o.SetTermOfPaymentTo(termOfPaymentTo)
	return o
}

// SetTermOfPaymentTo adds the termOfPaymentTo to the reminder search params
func (o *ReminderSearchParams) SetTermOfPaymentTo(termOfPaymentTo *string) {
	o.TermOfPaymentTo = termOfPaymentTo
}

// WriteToRequest writes these params to a swagger request
func (o *ReminderSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrCustomerID int32
		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := swag.FormatInt32(qrCustomerID)
		if qCustomerID != "" {
			if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
				return err
			}
		}

	}

	// query param dateFrom
	qrDateFrom := o.DateFrom
	qDateFrom := qrDateFrom
	if qDateFrom != "" {
		if err := r.SetQueryParam("dateFrom", qDateFrom); err != nil {
			return err
		}
	}

	// query param dateTo
	qrDateTo := o.DateTo
	qDateTo := qrDateTo
	if qDateTo != "" {
		if err := r.SetQueryParam("dateTo", qDateTo); err != nil {
			return err
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

	if o.InvoiceID != nil {

		// query param invoiceId
		var qrInvoiceID int32
		if o.InvoiceID != nil {
			qrInvoiceID = *o.InvoiceID
		}
		qInvoiceID := swag.FormatInt32(qrInvoiceID)
		if qInvoiceID != "" {
			if err := r.SetQueryParam("invoiceId", qInvoiceID); err != nil {
				return err
			}
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

	if o.TermOfPaymentFrom != nil {

		// query param termOfPaymentFrom
		var qrTermOfPaymentFrom string
		if o.TermOfPaymentFrom != nil {
			qrTermOfPaymentFrom = *o.TermOfPaymentFrom
		}
		qTermOfPaymentFrom := qrTermOfPaymentFrom
		if qTermOfPaymentFrom != "" {
			if err := r.SetQueryParam("termOfPaymentFrom", qTermOfPaymentFrom); err != nil {
				return err
			}
		}

	}

	if o.TermOfPaymentTo != nil {

		// query param termOfPaymentTo
		var qrTermOfPaymentTo string
		if o.TermOfPaymentTo != nil {
			qrTermOfPaymentTo = *o.TermOfPaymentTo
		}
		qTermOfPaymentTo := qrTermOfPaymentTo
		if qTermOfPaymentTo != "" {
			if err := r.SetQueryParam("termOfPaymentTo", qTermOfPaymentTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
