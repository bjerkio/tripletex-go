// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams creates a new EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams object
// with the default values initialized.
func NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams() *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	var (
		addToExistingDefault = bool(false)
	)
	return &EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams{
		AddToExisting: &addToExistingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithTimeout creates a new EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithTimeout(timeout time.Duration) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	var (
		addToExistingDefault = bool(false)
	)
	return &EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams{
		AddToExisting: &addToExistingDefault,

		timeout: timeout,
	}
}

// NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithContext creates a new EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithContext(ctx context.Context) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	var (
		addToExistingDefault = bool(false)
	)
	return &EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams{
		AddToExisting: &addToExistingDefault,

		Context: ctx,
	}
}

// NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithHTTPClient creates a new EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParamsWithHTTPClient(client *http.Client) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	var (
		addToExistingDefault = bool(false)
	)
	return &EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams{
		AddToExisting: &addToExistingDefault,
		HTTPClient:    client,
	}
}

/*EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams contains all the parameters to send to the API endpoint
for the employee entitlement grant client entitlements by template grant client entitlements by template operation typically these are written to a http.Request
*/
type EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams struct {

	/*AddToExisting
	  Add template to existing entitlements

	*/
	AddToExisting *bool
	/*CustomerID
	  Client ID

	*/
	CustomerID int32
	/*EmployeeID
	  Employee ID

	*/
	EmployeeID int32
	/*Template
	  Template

	*/
	Template string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithTimeout(timeout time.Duration) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithContext(ctx context.Context) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithHTTPClient(client *http.Client) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddToExisting adds the addToExisting to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithAddToExisting(addToExisting *bool) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetAddToExisting(addToExisting)
	return o
}

// SetAddToExisting adds the addToExisting to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetAddToExisting(addToExisting *bool) {
	o.AddToExisting = addToExisting
}

// WithCustomerID adds the customerID to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithCustomerID(customerID int32) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetCustomerID(customerID int32) {
	o.CustomerID = customerID
}

// WithEmployeeID adds the employeeID to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithEmployeeID(employeeID int32) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetEmployeeID(employeeID int32) {
	o.EmployeeID = employeeID
}

// WithTemplate adds the template to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WithTemplate(template string) *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the employee entitlement grant client entitlements by template grant client entitlements by template params
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) SetTemplate(template string) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *EmployeeEntitlementGrantClientEntitlementsByTemplateGrantClientEntitlementsByTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddToExisting != nil {

		// query param addToExisting
		var qrAddToExisting bool
		if o.AddToExisting != nil {
			qrAddToExisting = *o.AddToExisting
		}
		qAddToExisting := swag.FormatBool(qrAddToExisting)
		if qAddToExisting != "" {
			if err := r.SetQueryParam("addToExisting", qAddToExisting); err != nil {
				return err
			}
		}

	}

	// query param customerId
	qrCustomerID := o.CustomerID
	qCustomerID := swag.FormatInt32(qrCustomerID)
	if qCustomerID != "" {
		if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
			return err
		}
	}

	// query param employeeId
	qrEmployeeID := o.EmployeeID
	qEmployeeID := swag.FormatInt32(qrEmployeeID)
	if qEmployeeID != "" {
		if err := r.SetQueryParam("employeeId", qEmployeeID); err != nil {
			return err
		}
	}

	// query param template
	qrTemplate := o.Template
	qTemplate := qrTemplate
	if qTemplate != "" {
		if err := r.SetQueryParam("template", qTemplate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
