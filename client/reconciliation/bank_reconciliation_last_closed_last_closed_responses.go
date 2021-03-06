// Code generated by go-swagger; DO NOT EDIT.

package reconciliation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// BankReconciliationLastClosedLastClosedReader is a Reader for the BankReconciliationLastClosedLastClosed structure.
type BankReconciliationLastClosedLastClosedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationLastClosedLastClosedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationLastClosedLastClosedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBankReconciliationLastClosedLastClosedOK creates a BankReconciliationLastClosedLastClosedOK with default headers values
func NewBankReconciliationLastClosedLastClosedOK() *BankReconciliationLastClosedLastClosedOK {
	return &BankReconciliationLastClosedLastClosedOK{}
}

/*BankReconciliationLastClosedLastClosedOK handles this case with default header values.

successful operation
*/
type BankReconciliationLastClosedLastClosedOK struct {
	Payload *models.ResponseWrapperBankReconciliation
}

func (o *BankReconciliationLastClosedLastClosedOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/>lastClosed][%d] bankReconciliationLastClosedLastClosedOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationLastClosedLastClosedOK) GetPayload() *models.ResponseWrapperBankReconciliation {
	return o.Payload
}

func (o *BankReconciliationLastClosedLastClosedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperBankReconciliation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
