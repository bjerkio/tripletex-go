// Code generated by go-swagger; DO NOT EDIT.

package voucher_status

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

	"github.com/bjerkio/tripletex-go/v2/models"
)

// NewVoucherStatusPostParams creates a new VoucherStatusPostParams object
// with the default values initialized.
func NewVoucherStatusPostParams() *VoucherStatusPostParams {
	var ()
	return &VoucherStatusPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoucherStatusPostParamsWithTimeout creates a new VoucherStatusPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoucherStatusPostParamsWithTimeout(timeout time.Duration) *VoucherStatusPostParams {
	var ()
	return &VoucherStatusPostParams{

		timeout: timeout,
	}
}

// NewVoucherStatusPostParamsWithContext creates a new VoucherStatusPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoucherStatusPostParamsWithContext(ctx context.Context) *VoucherStatusPostParams {
	var ()
	return &VoucherStatusPostParams{

		Context: ctx,
	}
}

// NewVoucherStatusPostParamsWithHTTPClient creates a new VoucherStatusPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoucherStatusPostParamsWithHTTPClient(client *http.Client) *VoucherStatusPostParams {
	var ()
	return &VoucherStatusPostParams{
		HTTPClient: client,
	}
}

/*VoucherStatusPostParams contains all the parameters to send to the API endpoint
for the voucher status post operation typically these are written to a http.Request
*/
type VoucherStatusPostParams struct {

	/*Body
	  JSON representing the new object to be created. Should not have ID and version set.

	*/
	Body *models.VoucherStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voucher status post params
func (o *VoucherStatusPostParams) WithTimeout(timeout time.Duration) *VoucherStatusPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voucher status post params
func (o *VoucherStatusPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voucher status post params
func (o *VoucherStatusPostParams) WithContext(ctx context.Context) *VoucherStatusPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voucher status post params
func (o *VoucherStatusPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voucher status post params
func (o *VoucherStatusPostParams) WithHTTPClient(client *http.Client) *VoucherStatusPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voucher status post params
func (o *VoucherStatusPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the voucher status post params
func (o *VoucherStatusPostParams) WithBody(body *models.VoucherStatus) *VoucherStatusPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the voucher status post params
func (o *VoucherStatusPostParams) SetBody(body *models.VoucherStatus) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VoucherStatusPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}