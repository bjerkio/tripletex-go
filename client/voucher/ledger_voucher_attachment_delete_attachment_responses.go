// Code generated by go-swagger; DO NOT EDIT.

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LedgerVoucherAttachmentDeleteAttachmentReader is a Reader for the LedgerVoucherAttachmentDeleteAttachment structure.
type LedgerVoucherAttachmentDeleteAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherAttachmentDeleteAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLedgerVoucherAttachmentDeleteAttachmentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLedgerVoucherAttachmentDeleteAttachmentDefault creates a LedgerVoucherAttachmentDeleteAttachmentDefault with default headers values
func NewLedgerVoucherAttachmentDeleteAttachmentDefault(code int) *LedgerVoucherAttachmentDeleteAttachmentDefault {
	return &LedgerVoucherAttachmentDeleteAttachmentDefault{
		_statusCode: code,
	}
}

/*LedgerVoucherAttachmentDeleteAttachmentDefault handles this case with default header values.

successful operation
*/
type LedgerVoucherAttachmentDeleteAttachmentDefault struct {
	_statusCode int
}

// Code gets the status code for the ledger voucher attachment delete attachment default response
func (o *LedgerVoucherAttachmentDeleteAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *LedgerVoucherAttachmentDeleteAttachmentDefault) Error() string {
	return fmt.Sprintf("[DELETE /ledger/voucher/{id}/attachment][%d] LedgerVoucherAttachmentDeleteAttachment default ", o._statusCode)
}

func (o *LedgerVoucherAttachmentDeleteAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
