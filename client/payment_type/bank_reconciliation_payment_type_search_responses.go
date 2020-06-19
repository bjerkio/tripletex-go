// Code generated by go-swagger; DO NOT EDIT.

package payment_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// BankReconciliationPaymentTypeSearchReader is a Reader for the BankReconciliationPaymentTypeSearch structure.
type BankReconciliationPaymentTypeSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BankReconciliationPaymentTypeSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBankReconciliationPaymentTypeSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBankReconciliationPaymentTypeSearchOK creates a BankReconciliationPaymentTypeSearchOK with default headers values
func NewBankReconciliationPaymentTypeSearchOK() *BankReconciliationPaymentTypeSearchOK {
	return &BankReconciliationPaymentTypeSearchOK{}
}

/*BankReconciliationPaymentTypeSearchOK handles this case with default header values.

successful operation
*/
type BankReconciliationPaymentTypeSearchOK struct {
	Payload *models.ListResponseBankReconciliationPaymentType
}

func (o *BankReconciliationPaymentTypeSearchOK) Error() string {
	return fmt.Sprintf("[GET /bank/reconciliation/paymentType][%d] bankReconciliationPaymentTypeSearchOK  %+v", 200, o.Payload)
}

func (o *BankReconciliationPaymentTypeSearchOK) GetPayload() *models.ListResponseBankReconciliationPaymentType {
	return o.Payload
}

func (o *BankReconciliationPaymentTypeSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseBankReconciliationPaymentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
