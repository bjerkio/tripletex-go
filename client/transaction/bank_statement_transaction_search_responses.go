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

// BankStatementTransactionSearchReader is a Reader for the BankStatementTransactionSearch structure.
type BankStatementTransactionSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankStatementTransactionSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankStatementTransactionSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankStatementTransactionSearchOK creates a BankStatementTransactionSearchOK with default headers values
func NewBankStatementTransactionSearchOK() *BankStatementTransactionSearchOK {
	return &BankStatementTransactionSearchOK{}
}

/*BankStatementTransactionSearchOK handles this case with default header values.

successful operation
*/
type BankStatementTransactionSearchOK struct {
	Payload *models.ListResponseBankTransaction
}

func (o *BankStatementTransactionSearchOK) Error() string {
	return fmt.Sprintf("[GET /bank/statement/transaction][%d] bankStatementTransactionSearchOK  %+v", 200, o.Payload)
}

func (o *BankStatementTransactionSearchOK) GetPayload() *models.ListResponseBankTransaction {
	return o.Payload
}

func (o *BankStatementTransactionSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
