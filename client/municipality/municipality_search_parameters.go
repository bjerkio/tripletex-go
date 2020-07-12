// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package municipality

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

// NewMunicipalitySearchParams creates a new MunicipalitySearchParams object
// with the default values initialized.
func NewMunicipalitySearchParams() *MunicipalitySearchParams {
	var (
		countDefault                  = int64(1000)
		fromDefault                   = int64(0)
		includePayrollTaxZonesDefault = bool(true)
	)
	return &MunicipalitySearchParams{
		Count:                  &countDefault,
		From:                   &fromDefault,
		IncludePayrollTaxZones: &includePayrollTaxZonesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMunicipalitySearchParamsWithTimeout creates a new MunicipalitySearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMunicipalitySearchParamsWithTimeout(timeout time.Duration) *MunicipalitySearchParams {
	var (
		countDefault                  = int64(1000)
		fromDefault                   = int64(0)
		includePayrollTaxZonesDefault = bool(true)
	)
	return &MunicipalitySearchParams{
		Count:                  &countDefault,
		From:                   &fromDefault,
		IncludePayrollTaxZones: &includePayrollTaxZonesDefault,

		timeout: timeout,
	}
}

// NewMunicipalitySearchParamsWithContext creates a new MunicipalitySearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewMunicipalitySearchParamsWithContext(ctx context.Context) *MunicipalitySearchParams {
	var (
		countDefault                  = int64(1000)
		fromDefault                   = int64(0)
		includePayrollTaxZonesDefault = bool(true)
	)
	return &MunicipalitySearchParams{
		Count:                  &countDefault,
		From:                   &fromDefault,
		IncludePayrollTaxZones: &includePayrollTaxZonesDefault,

		Context: ctx,
	}
}

// NewMunicipalitySearchParamsWithHTTPClient creates a new MunicipalitySearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMunicipalitySearchParamsWithHTTPClient(client *http.Client) *MunicipalitySearchParams {
	var (
		countDefault                  = int64(1000)
		fromDefault                   = int64(0)
		includePayrollTaxZonesDefault = bool(true)
	)
	return &MunicipalitySearchParams{
		Count:                  &countDefault,
		From:                   &fromDefault,
		IncludePayrollTaxZones: &includePayrollTaxZonesDefault,
		HTTPClient:             client,
	}
}

/*MunicipalitySearchParams contains all the parameters to send to the API endpoint
for the municipality search operation typically these are written to a http.Request
*/
type MunicipalitySearchParams struct {

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
	/*IncludePayrollTaxZones
	  Equals

	*/
	IncludePayrollTaxZones *bool
	/*Sorting
	  Sorting pattern

	*/
	Sorting *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the municipality search params
func (o *MunicipalitySearchParams) WithTimeout(timeout time.Duration) *MunicipalitySearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the municipality search params
func (o *MunicipalitySearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the municipality search params
func (o *MunicipalitySearchParams) WithContext(ctx context.Context) *MunicipalitySearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the municipality search params
func (o *MunicipalitySearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the municipality search params
func (o *MunicipalitySearchParams) WithHTTPClient(client *http.Client) *MunicipalitySearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the municipality search params
func (o *MunicipalitySearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the municipality search params
func (o *MunicipalitySearchParams) WithCount(count *int64) *MunicipalitySearchParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the municipality search params
func (o *MunicipalitySearchParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the municipality search params
func (o *MunicipalitySearchParams) WithFields(fields *string) *MunicipalitySearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the municipality search params
func (o *MunicipalitySearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the municipality search params
func (o *MunicipalitySearchParams) WithFrom(from *int64) *MunicipalitySearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the municipality search params
func (o *MunicipalitySearchParams) SetFrom(from *int64) {
	o.From = from
}

// WithIncludePayrollTaxZones adds the includePayrollTaxZones to the municipality search params
func (o *MunicipalitySearchParams) WithIncludePayrollTaxZones(includePayrollTaxZones *bool) *MunicipalitySearchParams {
	o.SetIncludePayrollTaxZones(includePayrollTaxZones)
	return o
}

// SetIncludePayrollTaxZones adds the includePayrollTaxZones to the municipality search params
func (o *MunicipalitySearchParams) SetIncludePayrollTaxZones(includePayrollTaxZones *bool) {
	o.IncludePayrollTaxZones = includePayrollTaxZones
}

// WithSorting adds the sorting to the municipality search params
func (o *MunicipalitySearchParams) WithSorting(sorting *string) *MunicipalitySearchParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the municipality search params
func (o *MunicipalitySearchParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WriteToRequest writes these params to a swagger request
func (o *MunicipalitySearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludePayrollTaxZones != nil {

		// query param includePayrollTaxZones
		var qrIncludePayrollTaxZones bool
		if o.IncludePayrollTaxZones != nil {
			qrIncludePayrollTaxZones = *o.IncludePayrollTaxZones
		}
		qIncludePayrollTaxZones := swag.FormatBool(qrIncludePayrollTaxZones)
		if qIncludePayrollTaxZones != "" {
			if err := r.SetQueryParam("includePayrollTaxZones", qIncludePayrollTaxZones); err != nil {
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
