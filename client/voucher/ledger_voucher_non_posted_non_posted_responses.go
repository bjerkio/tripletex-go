// Code generated by go-swagger; DO NOT EDIT.

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// LedgerVoucherNonPostedNonPostedReader is a Reader for the LedgerVoucherNonPostedNonPosted structure.
type LedgerVoucherNonPostedNonPostedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherNonPostedNonPostedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerVoucherNonPostedNonPostedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLedgerVoucherNonPostedNonPostedOK creates a LedgerVoucherNonPostedNonPostedOK with default headers values
func NewLedgerVoucherNonPostedNonPostedOK() *LedgerVoucherNonPostedNonPostedOK {
	return &LedgerVoucherNonPostedNonPostedOK{}
}

/*LedgerVoucherNonPostedNonPostedOK handles this case with default header values.

successful operation
*/
type LedgerVoucherNonPostedNonPostedOK struct {
	Payload *models.ListResponseVoucher
}

func (o *LedgerVoucherNonPostedNonPostedOK) Error() string {
	return fmt.Sprintf("[GET /ledger/voucher/>nonPosted][%d] ledgerVoucherNonPostedNonPostedOK  %+v", 200, o.Payload)
}

func (o *LedgerVoucherNonPostedNonPostedOK) GetPayload() *models.ListResponseVoucher {
	return o.Payload
}

func (o *LedgerVoucherNonPostedNonPostedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseVoucher)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
