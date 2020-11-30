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

// BankReconciliationMatchGetReader is a Reader for the BankReconciliationMatchGet structure.
type BankReconciliationMatchGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationMatchGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationMatchGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankReconciliationMatchGetOK creates a BankReconciliationMatchGetOK with default headers values
func NewBankReconciliationMatchGetOK() *BankReconciliationMatchGetOK {
	return &BankReconciliationMatchGetOK{}
}

/*BankReconciliationMatchGetOK handles this case with default header values.

successful operation
*/
type BankReconciliationMatchGetOK struct {
	Payload *models.ResponseWrapperBankReconciliationMatch
}

func (o *BankReconciliationMatchGetOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/match/{id}][%d] bankReconciliationMatchGetOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationMatchGetOK) GetPayload() *models.ResponseWrapperBankReconciliationMatch {
	return o.Payload
}

func (o *BankReconciliationMatchGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliationMatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
