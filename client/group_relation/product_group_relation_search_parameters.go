// Code generated by go-swagger; DO NOT EDIT.

package group_relation

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

// NewProductGroupRelationSearchParams creates a new ProductGroupRelationSearchParams object
// with the default values initialized.
func NewProductGroupRelationSearchParams() *ProductGroupRelationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductGroupRelationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewProductGroupRelationSearchParamsWithTimeout creates a new ProductGroupRelationSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductGroupRelationSearchParamsWithTimeout(timeout time.Duration) *ProductGroupRelationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductGroupRelationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		timeout: timeout,
	}
}

// NewProductGroupRelationSearchParamsWithContext creates a new ProductGroupRelationSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductGroupRelationSearchParamsWithContext(ctx context.Context) *ProductGroupRelationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductGroupRelationSearchParams{
		Count: &countDefault,
		From:  &fromDefault,

		Context: ctx,
	}
}

// NewProductGroupRelationSearchParamsWithHTTPClient creates a new ProductGroupRelationSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductGroupRelationSearchParamsWithHTTPClient(client *http.Client) *ProductGroupRelationSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
	)
	return &ProductGroupRelationSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		HTTPClient: client,
	}
}

/*ProductGroupRelationSearchParams contains all the parameters to send to the API endpoint
for the product group relation search operation typically these are written to a http.Request
*/
type ProductGroupRelationSearchParams struct {

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
	  List of IDs

	*/
	ID *string
	/*ProductGroupID
	  List of IDs

	*/
	ProductGroupID *string
	/*ProductID
	  List of IDs

	*/
	ProductID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithTimeout(timeout time.Duration) *ProductGroupRelationSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithContext(ctx context.Context) *ProductGroupRelationSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithHTTPClient(client *http.Client) *ProductGroupRelationSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithCount(count *int64) *ProductGroupRelationSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithFields(fields *string) *ProductGroupRelationSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithFrom(from *int64) *ProductGroupRelationSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithID(id *string) *ProductGroupRelationSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetID(id *string) {
	o.ID = id
}

// WithProductGroupID adds the productGroupID to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithProductGroupID(productGroupID *string) *ProductGroupRelationSearchParams {
	o.SetProductGroupID(productGroupID)
	return o
}

// SetProductGroupID adds the productGroupId to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetProductGroupID(productGroupID *string) {
	o.ProductGroupID = productGroupID
}

// WithProductID adds the productID to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithProductID(productID *string) *ProductGroupRelationSearchParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetProductID(productID *string) {
	o.ProductID = productID
}

// WithSorting adds the sorting to the product group relation search params
func (o *ProductGroupRelationSearchParams) WithSorting(sorting *string) *ProductGroupRelationSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the product group relation search params
func (o *ProductGroupRelationSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *ProductGroupRelationSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProductGroupID != nil {

		// query param productGroupId
		var qrProductGroupID string
		if o.ProductGroupID != nil {
			qrProductGroupID = *o.ProductGroupID
		}
		qProductGroupID := qrProductGroupID
		if qProductGroupID != "" {
			if err := r.SetQueryParam("productGroupId", qProductGroupID); err != nil {
				return err
			}
		}

	}

	if o.ProductID != nil {

		// query param productId
		var qrProductID string
		if o.ProductID != nil {
			qrProductID = *o.ProductID
		}
		qProductID := qrProductID
		if qProductID != "" {
			if err := r.SetQueryParam("productId", qProductID); err != nil {
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
