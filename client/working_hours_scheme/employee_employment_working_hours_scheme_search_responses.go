// Code generated by go-swagger; DO NOT EDIT.

package working_hours_scheme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// EmployeeEmploymentWorkingHoursSchemeSearchReader is a Reader for the EmployeeEmploymentWorkingHoursSchemeSearch structure.
type EmployeeEmploymentWorkingHoursSchemeSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentWorkingHoursSchemeSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEmploymentWorkingHoursSchemeSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeeEmploymentWorkingHoursSchemeSearchOK creates a EmployeeEmploymentWorkingHoursSchemeSearchOK with default headers values
func NewEmployeeEmploymentWorkingHoursSchemeSearchOK() *EmployeeEmploymentWorkingHoursSchemeSearchOK {
	return &EmployeeEmploymentWorkingHoursSchemeSearchOK{}
}

/*EmployeeEmploymentWorkingHoursSchemeSearchOK handles this case with default header values.

successful operation
*/
type EmployeeEmploymentWorkingHoursSchemeSearchOK struct {
	Payload *models.ListResponseWorkingHoursScheme
}

func (o *EmployeeEmploymentWorkingHoursSchemeSearchOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/workingHoursScheme][%d] employeeEmploymentWorkingHoursSchemeSearchOK  %+v", 200, o.Payload)
}

func (o *EmployeeEmploymentWorkingHoursSchemeSearchOK) GetPayload() *models.ListResponseWorkingHoursScheme {
	return o.Payload
}

func (o *EmployeeEmploymentWorkingHoursSchemeSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseWorkingHoursScheme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
