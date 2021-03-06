// Code generated by go-swagger; DO NOT EDIT.

package leave_of_absence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// EmployeeEmploymentLeaveOfAbsencePostReader is a Reader for the EmployeeEmploymentLeaveOfAbsencePost structure.
type EmployeeEmploymentLeaveOfAbsencePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentLeaveOfAbsencePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmployeeEmploymentLeaveOfAbsencePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeeEmploymentLeaveOfAbsencePostCreated creates a EmployeeEmploymentLeaveOfAbsencePostCreated with default headers values
func NewEmployeeEmploymentLeaveOfAbsencePostCreated() *EmployeeEmploymentLeaveOfAbsencePostCreated {
	return &EmployeeEmploymentLeaveOfAbsencePostCreated{}
}

/*EmployeeEmploymentLeaveOfAbsencePostCreated handles this case with default header values.

successfully created
*/
type EmployeeEmploymentLeaveOfAbsencePostCreated struct {
	Payload *models.ResponseWrapperLeaveOfAbsence
}

func (o *EmployeeEmploymentLeaveOfAbsencePostCreated) Error() string {
	return fmt.Sprintf("[POST /employee/employment/leaveOfAbsence][%d] employeeEmploymentLeaveOfAbsencePostCreated  %+v", 201, o.Payload)
}

func (o *EmployeeEmploymentLeaveOfAbsencePostCreated) GetPayload() *models.ResponseWrapperLeaveOfAbsence {
	return o.Payload
}

func (o *EmployeeEmploymentLeaveOfAbsencePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperLeaveOfAbsence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
