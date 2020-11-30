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
	"github.com/go-openapi/swag"
)

// NewPurchaseOrderSearchParams creates a new PurchaseOrderSearchParams object
// with the default values initialized.
func NewPurchaseOrderSearchParams() *PurchaseOrderSearchParams {
	var (
		countDefault             = int64(1000)
		fromDefault              = int64(0)
		withDeviationOnlyDefault = bool(false)
	)
	return &PurchaseOrderSearchParams{
		Count:             &countDefault,
		From:              &fromDefault,
		WithDeviationOnly: &withDeviationOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPurchaseOrderSearchParamsWithTimeout creates a new PurchaseOrderSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurchaseOrderSearchParamsWithTimeout(timeout time.Duration) *PurchaseOrderSearchParams {
	var (
		countDefault             = int64(1000)
		fromDefault              = int64(0)
		withDeviationOnlyDefault = bool(false)
	)
	return &PurchaseOrderSearchParams{
		Count:             &countDefault,
		From:              &fromDefault,
		WithDeviationOnly: &withDeviationOnlyDefault,

		timeout: timeout,
	}
}

// NewPurchaseOrderSearchParamsWithContext creates a new PurchaseOrderSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurchaseOrderSearchParamsWithContext(ctx context.Context) *PurchaseOrderSearchParams {
	var (
		countDefault             = int64(1000)
		fromDefault              = int64(0)
		withDeviationOnlyDefault = bool(false)
	)
	return &PurchaseOrderSearchParams{
		Count:             &countDefault,
		From:              &fromDefault,
		WithDeviationOnly: &withDeviationOnlyDefault,

		Context: ctx,
	}
}

// NewPurchaseOrderSearchParamsWithHTTPClient creates a new PurchaseOrderSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurchaseOrderSearchParamsWithHTTPClient(client *http.Client) *PurchaseOrderSearchParams {
	var (
		countDefault             = int64(1000)
		fromDefault              = int64(0)
		withDeviationOnlyDefault = bool(false)
	)
	return &PurchaseOrderSearchParams{
		Count:             &countDefault,
		From:              &fromDefault,
		WithDeviationOnly: &withDeviationOnlyDefault,
		HTTPClient:        client,
	}
}

/*PurchaseOrderSearchParams contains all the parameters to send to the API endpoint
for the purchase order search operation typically these are written to a http.Request
*/
type PurchaseOrderSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*CreationDateFrom
	  Format is yyyy-MM-dd (from and incl.).

	*/
	CreationDateFrom *string
	/*CreationDateTo
	  Format is yyyy-MM-dd (to and incl.).

	*/
	CreationDateTo *string
	/*DeliveryDateFrom
	  Format is yyyy-MM-dd (from and incl.).

	*/
	DeliveryDateFrom *string
	/*DeliveryDateTo
	  Format is yyyy-MM-dd (to and incl.).

	*/
	DeliveryDateTo *string
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
	/*Number
	  Equals

	*/
	Number *string
	/*ProjectID
	  List of IDs

	*/
	ProjectID *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*SupplierID
	  List of IDs

	*/
	SupplierID *string
	/*WithDeviationOnly
	  Equals

	*/
	WithDeviationOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purchase order search params
func (o *PurchaseOrderSearchParams) WithTimeout(timeout time.Duration) *PurchaseOrderSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purchase order search params
func (o *PurchaseOrderSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purchase order search params
func (o *PurchaseOrderSearchParams) WithContext(ctx context.Context) *PurchaseOrderSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purchase order search params
func (o *PurchaseOrderSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purchase order search params
func (o *PurchaseOrderSearchParams) WithHTTPClient(client *http.Client) *PurchaseOrderSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purchase order search params
func (o *PurchaseOrderSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the purchase order search params
func (o *PurchaseOrderSearchParams) WithCount(count *int64) *PurchaseOrderSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the purchase order search params
func (o *PurchaseOrderSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithCreationDateFrom adds the creationDateFrom to the purchase order search params
func (o *PurchaseOrderSearchParams) WithCreationDateFrom(creationDateFrom *string) *PurchaseOrderSearchParams {
	o.SetCreationDateFrom(creationDateFrom)
	return o
}

// SetCreationDateFrom adds the creationDateFrom to the purchase order search params
func (o *PurchaseOrderSearchParams) SetCreationDateFrom(creationDateFrom *string) {
	o.CreationDateFrom = creationDateFrom
}

// WithCreationDateTo adds the creationDateTo to the purchase order search params
func (o *PurchaseOrderSearchParams) WithCreationDateTo(creationDateTo *string) *PurchaseOrderSearchParams {
	o.SetCreationDateTo(creationDateTo)
	return o
}

// SetCreationDateTo adds the creationDateTo to the purchase order search params
func (o *PurchaseOrderSearchParams) SetCreationDateTo(creationDateTo *string) {
	o.CreationDateTo = creationDateTo
}

// WithDeliveryDateFrom adds the deliveryDateFrom to the purchase order search params
func (o *PurchaseOrderSearchParams) WithDeliveryDateFrom(deliveryDateFrom *string) *PurchaseOrderSearchParams {
	o.SetDeliveryDateFrom(deliveryDateFrom)
	return o
}

// SetDeliveryDateFrom adds the deliveryDateFrom to the purchase order search params
func (o *PurchaseOrderSearchParams) SetDeliveryDateFrom(deliveryDateFrom *string) {
	o.DeliveryDateFrom = deliveryDateFrom
}

// WithDeliveryDateTo adds the deliveryDateTo to the purchase order search params
func (o *PurchaseOrderSearchParams) WithDeliveryDateTo(deliveryDateTo *string) *PurchaseOrderSearchParams {
	o.SetDeliveryDateTo(deliveryDateTo)
	return o
}

// SetDeliveryDateTo adds the deliveryDateTo to the purchase order search params
func (o *PurchaseOrderSearchParams) SetDeliveryDateTo(deliveryDateTo *string) {
	o.DeliveryDateTo = deliveryDateTo
}

// WithFields adds the fields to the purchase order search params
func (o *PurchaseOrderSearchParams) WithFields(fields *string) *PurchaseOrderSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the purchase order search params
func (o *PurchaseOrderSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the purchase order search params
func (o *PurchaseOrderSearchParams) WithFrom(from *int64) *PurchaseOrderSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the purchase order search params
func (o *PurchaseOrderSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithID adds the id to the purchase order search params
func (o *PurchaseOrderSearchParams) WithID(id *string) *PurchaseOrderSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the purchase order search params
func (o *PurchaseOrderSearchParams) SetID(id *string) {
	o.ID = id
}

// WithIsClosed adds the isClosed to the purchase order search params
func (o *PurchaseOrderSearchParams) WithIsClosed(isClosed *bool) *PurchaseOrderSearchParams {
	o.SetIsClosed(isClosed)
	return o
}

// SetIsClosed adds the isClosed to the purchase order search params
func (o *PurchaseOrderSearchParams) SetIsClosed(isClosed *bool) {
	o.IsClosed = isClosed
}

// WithNumber adds the number to the purchase order search params
func (o *PurchaseOrderSearchParams) WithNumber(number *string) *PurchaseOrderSearchParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the purchase order search params
func (o *PurchaseOrderSearchParams) SetNumber(number *string) {
	o.Number = number
}

// WithProjectID adds the projectID to the purchase order search params
func (o *PurchaseOrderSearchParams) WithProjectID(projectID *string) *PurchaseOrderSearchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the purchase order search params
func (o *PurchaseOrderSearchParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSorting adds the sorting to the purchase order search params
func (o *PurchaseOrderSearchParams) WithSorting(sorting *string) *PurchaseOrderSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the purchase order search params
func (o *PurchaseOrderSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithSupplierID adds the supplierID to the purchase order search params
func (o *PurchaseOrderSearchParams) WithSupplierID(supplierID *string) *PurchaseOrderSearchParams {
	o.SetSupplierID(supplierID)
	return o
}

// SetSupplierID adds the supplierId to the purchase order search params
func (o *PurchaseOrderSearchParams) SetSupplierID(supplierID *string) {
	o.SupplierID = supplierID
}

// WithWithDeviationOnly adds the withDeviationOnly to the purchase order search params
func (o *PurchaseOrderSearchParams) WithWithDeviationOnly(withDeviationOnly *bool) *PurchaseOrderSearchParams {
	o.SetWithDeviationOnly(withDeviationOnly)
	return o
}

// SetWithDeviationOnly adds the withDeviationOnly to the purchase order search params
func (o *PurchaseOrderSearchParams) SetWithDeviationOnly(withDeviationOnly *bool) {
	o.WithDeviationOnly = withDeviationOnly
}

// WriteToRequest writes these params to a swagger request
func (o *PurchaseOrderSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CreationDateFrom != nil {

		// query param creationDateFrom
		var qrCreationDateFrom string
		if o.CreationDateFrom != nil {
			qrCreationDateFrom = *o.CreationDateFrom
		}
		qCreationDateFrom := qrCreationDateFrom
		if qCreationDateFrom != "" {
			if err := r.SetQueryParam("creationDateFrom", qCreationDateFrom); err != nil {
				return err
			}
		}

	}

	if o.CreationDateTo != nil {

		// query param creationDateTo
		var qrCreationDateTo string
		if o.CreationDateTo != nil {
			qrCreationDateTo = *o.CreationDateTo
		}
		qCreationDateTo := qrCreationDateTo
		if qCreationDateTo != "" {
			if err := r.SetQueryParam("creationDateTo", qCreationDateTo); err != nil {
				return err
			}
		}

	}

	if o.DeliveryDateFrom != nil {

		// query param deliveryDateFrom
		var qrDeliveryDateFrom string
		if o.DeliveryDateFrom != nil {
			qrDeliveryDateFrom = *o.DeliveryDateFrom
		}
		qDeliveryDateFrom := qrDeliveryDateFrom
		if qDeliveryDateFrom != "" {
			if err := r.SetQueryParam("deliveryDateFrom", qDeliveryDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DeliveryDateTo != nil {

		// query param deliveryDateTo
		var qrDeliveryDateTo string
		if o.DeliveryDateTo != nil {
			qrDeliveryDateTo = *o.DeliveryDateTo
		}
		qDeliveryDateTo := qrDeliveryDateTo
		if qDeliveryDateTo != "" {
			if err := r.SetQueryParam("deliveryDateTo", qDeliveryDateTo); err != nil {
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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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

	if o.SupplierID != nil {

		// query param supplierId
		var qrSupplierID string
		if o.SupplierID != nil {
			qrSupplierID = *o.SupplierID
		}
		qSupplierID := qrSupplierID
		if qSupplierID != "" {
			if err := r.SetQueryParam("supplierId", qSupplierID); err != nil {
				return err
			}
		}

	}

	if o.WithDeviationOnly != nil {

		// query param withDeviationOnly
		var qrWithDeviationOnly bool
		if o.WithDeviationOnly != nil {
			qrWithDeviationOnly = *o.WithDeviationOnly
		}
		qWithDeviationOnly := swag.FormatBool(qrWithDeviationOnly)
		if qWithDeviationOnly != "" {
			if err := r.SetQueryParam("withDeviationOnly", qWithDeviationOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
