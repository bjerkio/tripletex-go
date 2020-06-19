// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// BankStatementTransactionDetailsGetDetailsReader is a Reader for the BankStatementTransactionDetailsGetDetails structure.
type BankStatementTransactionDetailsGetDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankStatementTransactionDetailsGetDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankStatementTransactionDetailsGetDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBankStatementTransactionDetailsGetDetailsOK creates a BankStatementTransactionDetailsGetDetailsOK with default headers values
func NewBankStatementTransactionDetailsGetDetailsOK() *BankStatementTransactionDetailsGetDetailsOK {
	return &BankStatementTransactionDetailsGetDetailsOK{}
}

/*BankStatementTransactionDetailsGetDetailsOK handles this case with default header values.

successful operation
*/
type BankStatementTransactionDetailsGetDetailsOK struct {
	Payload *models.ResponseWrapperObject
}

func (o *BankStatementTransactionDetailsGetDetailsOK) Error() string {
	return fmt.Sprintf("[GET /bank/statement/transaction/{id}/details][%d] bankStatementTransactionDetailsGetDetailsOK  %+v", 200, o.Payload)
}

func (o *BankStatementTransactionDetailsGetDetailsOK) GetPayload() *models.ResponseWrapperObject {
	return o.Payload
}

func (o *BankStatementTransactionDetailsGetDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
