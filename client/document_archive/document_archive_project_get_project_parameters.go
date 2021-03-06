// Code generated by go-swagger; DO NOT EDIT.

package document_archive

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

// NewDocumentArchiveProjectGetProjectParams creates a new DocumentArchiveProjectGetProjectParams object
// with the default values initialized.
func NewDocumentArchiveProjectGetProjectParams() *DocumentArchiveProjectGetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveProjectGetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentArchiveProjectGetProjectParamsWithTimeout creates a new DocumentArchiveProjectGetProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentArchiveProjectGetProjectParamsWithTimeout(timeout time.Duration) *DocumentArchiveProjectGetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveProjectGetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewDocumentArchiveProjectGetProjectParamsWithContext creates a new DocumentArchiveProjectGetProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentArchiveProjectGetProjectParamsWithContext(ctx context.Context) *DocumentArchiveProjectGetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveProjectGetProjectParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewDocumentArchiveProjectGetProjectParamsWithHTTPClient creates a new DocumentArchiveProjectGetProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentArchiveProjectGetProjectParamsWithHTTPClient(client *http.Client) *DocumentArchiveProjectGetProjectParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &DocumentArchiveProjectGetProjectParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*DocumentArchiveProjectGetProjectParams contains all the parameters to send to the API endpoint
for the document archive project get project operation typically these are written to a http.Request
*/
type DocumentArchiveProjectGetProjectParams struct {

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
	/*ID
	  Element ID

	*/
	ID int32
	/*PeriodDateFrom
	  From and including

	*/
	PeriodDateFrom *string
	/*PeriodDateTo
	  To and excluding

	*/
	PeriodDateTo *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithTimeout(timeout time.Duration) *DocumentArchiveProjectGetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithContext(ctx context.Context) *DocumentArchiveProjectGetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithHTTPClient(client *http.Client) *DocumentArchiveProjectGetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithCount(count *int64) *DocumentArchiveProjectGetProjectParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithFields(fields *string) *DocumentArchiveProjectGetProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithFrom(from *int64) *DocumentArchiveProjectGetProjectParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithID(id int32) *DocumentArchiveProjectGetProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetID(id int32) {
	o.ID = id
}

// WithPeriodDateFrom adds the periodDateFrom to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithPeriodDateFrom(periodDateFrom *string) *DocumentArchiveProjectGetProjectParams {
	o.SetPeriodDateFrom(periodDateFrom)
	return o
}

// SetPeriodDateFrom adds the periodDateFrom to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetPeriodDateFrom(periodDateFrom *string) {
	o.PeriodDateFrom = periodDateFrom
}

// WithPeriodDateTo adds the periodDateTo to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithPeriodDateTo(periodDateTo *string) *DocumentArchiveProjectGetProjectParams {
	o.SetPeriodDateTo(periodDateTo)
	return o
}

// SetPeriodDateTo adds the periodDateTo to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetPeriodDateTo(periodDateTo *string) {
	o.PeriodDateTo = periodDateTo
}

// WithSorting adds the sorting to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) WithSorting(sorting *string) *DocumentArchiveProjectGetProjectParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the document archive project get project params
func (o *DocumentArchiveProjectGetProjectParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentArchiveProjectGetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.PeriodDateFrom != nil {

		// query param periodDateFrom
		var qrPeriodDateFrom string
		if o.PeriodDateFrom != nil {
			qrPeriodDateFrom = *o.PeriodDateFrom
		}
		qPeriodDateFrom := qrPeriodDateFrom
		if qPeriodDateFrom != "" {
			if err := r.SetQueryParam("periodDateFrom", qPeriodDateFrom); err != nil {
				return err
			}
		}

	}

	if o.PeriodDateTo != nil {

		// query param periodDateTo
		var qrPeriodDateTo string
		if o.PeriodDateTo != nil {
			qrPeriodDateTo = *o.PeriodDateTo
		}
		qPeriodDateTo := qrPeriodDateTo
		if qPeriodDateTo != "" {
			if err := r.SetQueryParam("periodDateTo", qPeriodDateTo); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
