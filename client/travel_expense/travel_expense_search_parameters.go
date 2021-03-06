// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

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

// NewTravelExpenseSearchParams creates a new TravelExpenseSearchParams object
// with the default values initialized.
func NewTravelExpenseSearchParams() *TravelExpenseSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
		stateDefault = string("ALL")
	)
	return &TravelExpenseSearchParams{
		Count: &countDefault,
		From:  &fromDefault,
		State: &stateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTravelExpenseSearchParamsWithTimeout creates a new TravelExpenseSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTravelExpenseSearchParamsWithTimeout(timeout time.Duration) *TravelExpenseSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
		stateDefault = string("ALL")
	)
	return &TravelExpenseSearchParams{
		Count: &countDefault,
		From:  &fromDefault,
		State: &stateDefault,

		timeout: timeout,
	}
}

// NewTravelExpenseSearchParamsWithContext creates a new TravelExpenseSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewTravelExpenseSearchParamsWithContext(ctx context.Context) *TravelExpenseSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
		stateDefault = string("ALL")
	)
	return &TravelExpenseSearchParams{
		Count: &countDefault,
		From:  &fromDefault,
		State: &stateDefault,

		Context: ctx,
	}
}

// NewTravelExpenseSearchParamsWithHTTPClient creates a new TravelExpenseSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTravelExpenseSearchParamsWithHTTPClient(client *http.Client) *TravelExpenseSearchParams {
	var (
		countDefault = int64(1000)
		fromDefault  = int64(0)
		stateDefault = string("ALL")
	)
	return &TravelExpenseSearchParams{
		Count:      &countDefault,
		From:       &fromDefault,
		State:      &stateDefault,
		HTTPClient: client,
	}
}

/*TravelExpenseSearchParams contains all the parameters to send to the API endpoint
for the travel expense search operation typically these are written to a http.Request
*/
type TravelExpenseSearchParams struct {

	/*Count
	  Number of elements to return

	*/
	Count *int64
	/*DepartmentID
	  Equals

	*/
	DepartmentID *string
	/*DepartureDateFrom
	  From and including

	*/
	DepartureDateFrom *string
	/*EmployeeID
	  Equals

	*/
	EmployeeID *string
	/*Fields
	  Fields filter pattern

	*/
	Fields *string
	/*From
	  From index

	*/
	From *int64
	/*ProjectID
	  Equals

	*/
	ProjectID *string
	/*ProjectManagerID
	  Equals

	*/
	ProjectManagerID *string
	/*ReturnDateTo
	  To and excluding

	*/
	ReturnDateTo *string
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string
	/*State
	  category

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the travel expense search params
func (o *TravelExpenseSearchParams) WithTimeout(timeout time.Duration) *TravelExpenseSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the travel expense search params
func (o *TravelExpenseSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the travel expense search params
func (o *TravelExpenseSearchParams) WithContext(ctx context.Context) *TravelExpenseSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the travel expense search params
func (o *TravelExpenseSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the travel expense search params
func (o *TravelExpenseSearchParams) WithHTTPClient(client *http.Client) *TravelExpenseSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the travel expense search params
func (o *TravelExpenseSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the travel expense search params
func (o *TravelExpenseSearchParams) WithCount(count *int64) *TravelExpenseSearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the travel expense search params
func (o *TravelExpenseSearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithDepartmentID adds the departmentID to the travel expense search params
func (o *TravelExpenseSearchParams) WithDepartmentID(departmentID *string) *TravelExpenseSearchParams {
	o.SetDepartmentID(departmentID)
	return o
}

// SetDepartmentID adds the departmentId to the travel expense search params
func (o *TravelExpenseSearchParams) SetDepartmentID(departmentID *string) {
	o.DepartmentID = departmentID
}

// WithDepartureDateFrom adds the departureDateFrom to the travel expense search params
func (o *TravelExpenseSearchParams) WithDepartureDateFrom(departureDateFrom *string) *TravelExpenseSearchParams {
	o.SetDepartureDateFrom(departureDateFrom)
	return o
}

// SetDepartureDateFrom adds the departureDateFrom to the travel expense search params
func (o *TravelExpenseSearchParams) SetDepartureDateFrom(departureDateFrom *string) {
	o.DepartureDateFrom = departureDateFrom
}

// WithEmployeeID adds the employeeID to the travel expense search params
func (o *TravelExpenseSearchParams) WithEmployeeID(employeeID *string) *TravelExpenseSearchParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the travel expense search params
func (o *TravelExpenseSearchParams) SetEmployeeID(employeeID *string) {
	o.EmployeeID = employeeID
}

// WithFields adds the fields to the travel expense search params
func (o *TravelExpenseSearchParams) WithFields(fields *string) *TravelExpenseSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the travel expense search params
func (o *TravelExpenseSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the travel expense search params
func (o *TravelExpenseSearchParams) WithFrom(from *int64) *TravelExpenseSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the travel expense search params
func (o *TravelExpenseSearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithProjectID adds the projectID to the travel expense search params
func (o *TravelExpenseSearchParams) WithProjectID(projectID *string) *TravelExpenseSearchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the travel expense search params
func (o *TravelExpenseSearchParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithProjectManagerID adds the projectManagerID to the travel expense search params
func (o *TravelExpenseSearchParams) WithProjectManagerID(projectManagerID *string) *TravelExpenseSearchParams {
	o.SetProjectManagerID(projectManagerID)
	return o
}

// SetProjectManagerID adds the projectManagerId to the travel expense search params
func (o *TravelExpenseSearchParams) SetProjectManagerID(projectManagerID *string) {
	o.ProjectManagerID = projectManagerID
}

// WithReturnDateTo adds the returnDateTo to the travel expense search params
func (o *TravelExpenseSearchParams) WithReturnDateTo(returnDateTo *string) *TravelExpenseSearchParams {
	o.SetReturnDateTo(returnDateTo)
	return o
}

// SetReturnDateTo adds the returnDateTo to the travel expense search params
func (o *TravelExpenseSearchParams) SetReturnDateTo(returnDateTo *string) {
	o.ReturnDateTo = returnDateTo
}

// WithSorting adds the sorting to the travel expense search params
func (o *TravelExpenseSearchParams) WithSorting(sorting *string) *TravelExpenseSearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the travel expense search params
func (o *TravelExpenseSearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithState adds the state to the travel expense search params
func (o *TravelExpenseSearchParams) WithState(state *string) *TravelExpenseSearchParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the travel expense search params
func (o *TravelExpenseSearchParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *TravelExpenseSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DepartmentID != nil {

		// query param departmentId
		var qrDepartmentID string
		if o.DepartmentID != nil {
			qrDepartmentID = *o.DepartmentID
		}
		qDepartmentID := qrDepartmentID
		if qDepartmentID != "" {
			if err := r.SetQueryParam("departmentId", qDepartmentID); err != nil {
				return err
			}
		}

	}

	if o.DepartureDateFrom != nil {

		// query param departureDateFrom
		var qrDepartureDateFrom string
		if o.DepartureDateFrom != nil {
			qrDepartureDateFrom = *o.DepartureDateFrom
		}
		qDepartureDateFrom := qrDepartureDateFrom
		if qDepartureDateFrom != "" {
			if err := r.SetQueryParam("departureDateFrom", qDepartureDateFrom); err != nil {
				return err
			}
		}

	}

	if o.EmployeeID != nil {

		// query param employeeId
		var qrEmployeeID string
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := qrEmployeeID
		if qEmployeeID != "" {
			if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
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

	if o.ProjectManagerID != nil {

		// query param projectManagerId
		var qrProjectManagerID string
		if o.ProjectManagerID != nil {
			qrProjectManagerID = *o.ProjectManagerID
		}
		qProjectManagerID := qrProjectManagerID
		if qProjectManagerID != "" {
			if err := r.SetQueryParam("projectManagerId", qProjectManagerID); err != nil {
				return err
			}
		}

	}

	if o.ReturnDateTo != nil {

		// query param returnDateTo
		var qrReturnDateTo string
		if o.ReturnDateTo != nil {
			qrReturnDateTo = *o.ReturnDateTo
		}
		qReturnDateTo := qrReturnDateTo
		if qReturnDateTo != "" {
			if err := r.SetQueryParam("returnDateTo", qReturnDateTo); err != nil {
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

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
