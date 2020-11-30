// Code generated by go-swagger; DO NOT EDIT.

package salary_type_specification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetSalaryTypeSpecificationPutReader is a Reader for the TimesheetSalaryTypeSpecificationPut structure.
type TimesheetSalaryTypeSpecificationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetSalaryTypeSpecificationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetSalaryTypeSpecificationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetSalaryTypeSpecificationPutOK creates a TimesheetSalaryTypeSpecificationPutOK with default headers values
func NewTimesheetSalaryTypeSpecificationPutOK() *TimesheetSalaryTypeSpecificationPutOK {
	return &TimesheetSalaryTypeSpecificationPutOK{}
}

/*TimesheetSalaryTypeSpecificationPutOK handles this case with default header values.

successful operation
*/
type TimesheetSalaryTypeSpecificationPutOK struct {
	Payload *models.ResponseWrapperTimesheetSalaryTypeSpecification
}

func (o *TimesheetSalaryTypeSpecificationPutOK) Error() string {
	return fmt.Sprintf("[PUT /timesheet/salaryTypeSpecification/{id}][%d] timesheetSalaryTypeSpecificationPutOK  %+v", 200, o.Payload)
}

func (o *TimesheetSalaryTypeSpecificationPutOK) GetPayload() *models.ResponseWrapperTimesheetSalaryTypeSpecification {
	return o.Payload
}

func (o *TimesheetSalaryTypeSpecificationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTimesheetSalaryTypeSpecification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
