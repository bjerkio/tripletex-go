// Code generated by go-swagger; DO NOT EDIT.

package session

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
)

// NewTokenSessionWhoAmIWhoAmIParams creates a new TokenSessionWhoAmIWhoAmIParams object
// with the default values initialized.
func NewTokenSessionWhoAmIWhoAmIParams() *TokenSessionWhoAmIWhoAmIParams {
	var ()
	return &TokenSessionWhoAmIWhoAmIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenSessionWhoAmIWhoAmIParamsWithTimeout creates a new TokenSessionWhoAmIWhoAmIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenSessionWhoAmIWhoAmIParamsWithTimeout(timeout time.Duration) *TokenSessionWhoAmIWhoAmIParams {
	var ()
	return &TokenSessionWhoAmIWhoAmIParams{

		timeout: timeout,
	}
}

// NewTokenSessionWhoAmIWhoAmIParamsWithContext creates a new TokenSessionWhoAmIWhoAmIParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokenSessionWhoAmIWhoAmIParamsWithContext(ctx context.Context) *TokenSessionWhoAmIWhoAmIParams {
	var ()
	return &TokenSessionWhoAmIWhoAmIParams{

		Context: ctx,
	}
}

// NewTokenSessionWhoAmIWhoAmIParamsWithHTTPClient creates a new TokenSessionWhoAmIWhoAmIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenSessionWhoAmIWhoAmIParamsWithHTTPClient(client *http.Client) *TokenSessionWhoAmIWhoAmIParams {
	var ()
	return &TokenSessionWhoAmIWhoAmIParams{
		HTTPClient: client,
	}
}

/*TokenSessionWhoAmIWhoAmIParams contains all the parameters to send to the API endpoint
for the token session who am i who am i operation typically these are written to a http.Request
*/
type TokenSessionWhoAmIWhoAmIParams struct {

	/*Fields
	  Fields filter pattern

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) WithTimeout(timeout time.Duration) *TokenSessionWhoAmIWhoAmIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) WithContext(ctx context.Context) *TokenSessionWhoAmIWhoAmIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) WithHTTPClient(client *http.Client) *TokenSessionWhoAmIWhoAmIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) WithFields(fields *string) *TokenSessionWhoAmIWhoAmIParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the token session who am i who am i params
func (o *TokenSessionWhoAmIWhoAmIParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *TokenSessionWhoAmIWhoAmIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
