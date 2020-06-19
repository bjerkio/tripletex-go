// Code generated by go-swagger; DO NOT EDIT.

package payslip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// SalaryPayslipGetReader is a Reader for the SalaryPayslipGet structure.
type SalaryPayslipGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalaryPayslipGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalaryPayslipGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSalaryPayslipGetOK creates a SalaryPayslipGetOK with default headers values
func NewSalaryPayslipGetOK() *SalaryPayslipGetOK {
	return &SalaryPayslipGetOK{}
}

/*SalaryPayslipGetOK handles this case with default header values.

successful operation
*/
type SalaryPayslipGetOK struct {
	Payload *models.ResponseWrapperPayslip
}

func (o *SalaryPayslipGetOK) Error() string {
	return fmt.Sprintf("[GET /salary/payslip/{id}][%d] salaryPayslipGetOK  %+v", 200, o.Payload)
}

func (o *SalaryPayslipGetOK) GetPayload() *models.ResponseWrapperPayslip {
	return o.Payload
}

func (o *SalaryPayslipGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPayslip)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
