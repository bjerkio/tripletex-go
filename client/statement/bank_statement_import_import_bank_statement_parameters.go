// Code generated by go-swagger; DO NOT EDIT.

package statement

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

// NewBankStatementImportImportBankStatementParams creates a new BankStatementImportImportBankStatementParams object
// with the default values initialized.
func NewBankStatementImportImportBankStatementParams() *BankStatementImportImportBankStatementParams {
	var ()
	return &BankStatementImportImportBankStatementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBankStatementImportImportBankStatementParamsWithTimeout creates a new BankStatementImportImportBankStatementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBankStatementImportImportBankStatementParamsWithTimeout(timeout time.Duration) *BankStatementImportImportBankStatementParams {
	var ()
	return &BankStatementImportImportBankStatementParams{

		timeout: timeout,
	}
}

// NewBankStatementImportImportBankStatementParamsWithContext creates a new BankStatementImportImportBankStatementParams object
// with the default values initialized, and the ability to set a context for a request
func NewBankStatementImportImportBankStatementParamsWithContext(ctx context.Context) *BankStatementImportImportBankStatementParams {
	var ()
	return &BankStatementImportImportBankStatementParams{

		Context: ctx,
	}
}

// NewBankStatementImportImportBankStatementParamsWithHTTPClient creates a new BankStatementImportImportBankStatementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBankStatementImportImportBankStatementParamsWithHTTPClient(client *http.Client) *BankStatementImportImportBankStatementParams {
	var ()
	return &BankStatementImportImportBankStatementParams{
		HTTPClient: client,
	}
}

/*BankStatementImportImportBankStatementParams contains all the parameters to send to the API endpoint
for the bank statement import import bank statement operation typically these are written to a http.Request
*/
type BankStatementImportImportBankStatementParams struct {

	/*AccountID
	  Account ID

	*/
	AccountID int32
	/*BankID
	  Bank ID

	*/
	BankID int32
	/*ExternalID
	  External ID

	*/
	ExternalID *string
	/*File
	  The bank statement file

	*/
	File runtime.NamedReadCloser
	/*FileFormat
	  File format

	*/
	FileFormat string
	/*FromDate
	  Format is yyyy-MM-dd (from and incl.).

	*/
	FromDate string
	/*ToDate
	  Format is yyyy-MM-dd (to and excl.).

	*/
	ToDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithTimeout(timeout time.Duration) *BankStatementImportImportBankStatementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithContext(ctx context.Context) *BankStatementImportImportBankStatementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithHTTPClient(client *http.Client) *BankStatementImportImportBankStatementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithAccountID(accountID int32) *BankStatementImportImportBankStatementParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithBankID adds the bankID to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithBankID(bankID int32) *BankStatementImportImportBankStatementParams {
	o.SetBankID(bankID)
	return o
}

// SetBankID adds the bankId to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetBankID(bankID int32) {
	o.BankID = bankID
}

// WithExternalID adds the externalID to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithExternalID(externalID *string) *BankStatementImportImportBankStatementParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetExternalID(externalID *string) {
	o.ExternalID = externalID
}

// WithFile adds the file to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithFile(file runtime.NamedReadCloser) *BankStatementImportImportBankStatementParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithFileFormat adds the fileFormat to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithFileFormat(fileFormat string) *BankStatementImportImportBankStatementParams {
	o.SetFileFormat(fileFormat)
	return o
}

// SetFileFormat adds the fileFormat to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetFileFormat(fileFormat string) {
	o.FileFormat = fileFormat
}

// WithFromDate adds the fromDate to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithFromDate(fromDate string) *BankStatementImportImportBankStatementParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetFromDate(fromDate string) {
	o.FromDate = fromDate
}

// WithToDate adds the toDate to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) WithToDate(toDate string) *BankStatementImportImportBankStatementParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the bank statement import import bank statement params
func (o *BankStatementImportImportBankStatementParams) SetToDate(toDate string) {
	o.ToDate = toDate
}

// WriteToRequest writes these params to a swagger request
func (o *BankStatementImportImportBankStatementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := swag.FormatInt32(qrAccountID)
	if qAccountID != "" {
		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
	}

	// query param bankId
	qrBankID := o.BankID
	qBankID := swag.FormatInt32(qrBankID)
	if qBankID != "" {
		if err := r.SetQueryParam("bankId", qBankID); err != nil {
			return err
		}
	}

	if o.ExternalID != nil {

		// query param externalId
		var qrExternalID string
		if o.ExternalID != nil {
			qrExternalID = *o.ExternalID
		}
		qExternalID := qrExternalID
		if qExternalID != "" {
			if err := r.SetQueryParam("externalId", qExternalID); err != nil {
				return err
			}
		}

	}

	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// query param fileFormat
	qrFileFormat := o.FileFormat
	qFileFormat := qrFileFormat
	if qFileFormat != "" {
		if err := r.SetQueryParam("fileFormat", qFileFormat); err != nil {
			return err
		}
	}

	// query param fromDate
	qrFromDate := o.FromDate
	qFromDate := qrFromDate
	if qFromDate != "" {
		if err := r.SetQueryParam("fromDate", qFromDate); err != nil {
			return err
		}
	}

	// query param toDate
	qrToDate := o.ToDate
	qToDate := qrToDate
	if qToDate != "" {
		if err := r.SetQueryParam("toDate", qToDate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
