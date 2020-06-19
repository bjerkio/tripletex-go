// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// BankReconciliationMatchSearchReader is a Reader for the BankReconciliationMatchSearch structure.
type BankReconciliationMatchSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationMatchSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationMatchSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBankReconciliationMatchSearchOK creates a BankReconciliationMatchSearchOK with default headers values
func NewBankReconciliationMatchSearchOK() *BankReconciliationMatchSearchOK {
	return &BankReconciliationMatchSearchOK{}
}

/*BankReconciliationMatchSearchOK handles this case with default header values.

successful operation
*/
type BankReconciliationMatchSearchOK struct {
	Payload *models.ListResponseBankReconciliationMatch
}

func (o *BankReconciliationMatchSearchOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/match][%d] bankReconciliationMatchSearchOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationMatchSearchOK) GetPayload() *models.ListResponseBankReconciliationMatch {
	return o.Payload
}

func (o *BankReconciliationMatchSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankReconciliationMatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
