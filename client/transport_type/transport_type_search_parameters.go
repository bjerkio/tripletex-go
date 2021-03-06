// Code generated by go-swagger; DO NOT EDIT.

package transport_type

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

// NewTransportTypeSearchParams creates a new TransportTypeSearchParams object
// with the default values initialized.
func NewTransportTypeSearchParams() *TransportTypeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TransportTypeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTransportTypeSearchParamsWithTimeout creates a new TransportTypeSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransportTypeSearchParamsWithTimeout(timeout time.Duration) *TransportTypeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TransportTypeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewTransportTypeSearchParamsWithContext creates a new TransportTypeSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransportTypeSearchParamsWithContext(ctx context.Context) *TransportTypeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TransportTypeSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewTransportTypeSearchParamsWithHTTPClient creates a new TransportTypeSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransportTypeSearchParamsWithHTTPClient(client *http.Client) *TransportTypeSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &TransportTypeSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*TransportTypeSearchParams contains all the parameters to send to the API endpoint
for the transport type search operation typically these are written to a http.Request
*/
type TransportTypeSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*Name
	  Containing

	*/
	Name *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*SupplierID
	  List of valid suppliers ids.

	*/
	SupplierID []int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the transport type search params
func (o *TransportTypeSearchParams) WithTimeout(timeout time.Duration) *TransportTypeSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transport type search params
func (o *TransportTypeSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transport type search params
func (o *TransportTypeSearchParams) WithContext(ctx context.Context) *TransportTypeSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transport type search params
func (o *TransportTypeSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transport type search params
func (o *TransportTypeSearchParams) WithHTTPClient(client *http.Client) *TransportTypeSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transport type search params
func (o *TransportTypeSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the transport type search params
func (o *TransportTypeSearchParams) WithCount(count *int64) *TransportTypeSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the transport type search params
func (o *TransportTypeSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the transport type search params
func (o *TransportTypeSearchParams) WithFields(fields *string) *TransportTypeSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the transport type search params
func (o *TransportTypeSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the transport type search params
func (o *TransportTypeSearchParams) WithFrom(from *int64) *TransportTypeSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the transport type search params
func (o *TransportTypeSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithName adds the name to the transport type search params
func (o *TransportTypeSearchParams) WithName(name *string) *TransportTypeSearchParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the transport type search params
func (o *TransportTypeSearchParams) SetName(name *string) {
	o.Name = name
}

// WithSorting adds the sorting to the transport type search params
func (o *TransportTypeSearchParams) WithSorting(sorting *string) *TransportTypeSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the transport type search params
func (o *TransportTypeSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierID adds the supplierID to the transport type search params
func (o *TransportTypeSearchParams) WithSupplierID(supplierID []int32) *TransportTypeSearchParams {
	o.SetSupplierID(supplierID)
	return o
}

// SetSupplierID adds the supplierId to the transport type search params
func (o *TransportTypeSearchParams) SetSupplierID(supplierID []int32) {
	o.SupplierID = supplierID
}

// WriteToRequest writes these params to a swagger request
func (o *TransportTypeSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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

	var valuesSupplierID []string
	for _, v := range o.SupplierID {
		valuesSupplierID = append(valuesSupplierID, swag.FormatInt32(v))
	}

	joinedSupplierID := swag.JoinByFormat(valuesSupplierID, "multi")
	// query array param supplierId
	if err := r.SetQueryParam("supplierId", joinedSupplierID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
