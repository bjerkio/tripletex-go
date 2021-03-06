// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// SalaryTransactionGetReader is a Reader for the SalaryTransactionGet structure.
type SalaryTransactionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalaryTransactionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalaryTransactionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSalaryTransactionGetOK creates a SalaryTransactionGetOK with default headers values
func NewSalaryTransactionGetOK() *SalaryTransactionGetOK {
	return &SalaryTransactionGetOK{}
}

/*SalaryTransactionGetOK handles this case with default header values.

successful operation
*/
type SalaryTransactionGetOK struct {
	Payload *models.ResponseWrapperSalaryTransaction
}

func (o *SalaryTransactionGetOK) Error() string {
	return fmt.Sprintf("[GET /salary/transaction/{id}][%d] salaryTransactionGetOK  %+v", 200, o.Payload)
}

func (o *SalaryTransactionGetOK) GetPayload() *models.ResponseWrapperSalaryTransaction {
	return o.Payload
}

func (o *SalaryTransactionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSalaryTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
