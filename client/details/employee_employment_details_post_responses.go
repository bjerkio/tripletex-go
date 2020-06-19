// Code generated by go-swagger; DO NOT EDIT.

package details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeeEmploymentDetailsPostReader is a Reader for the EmployeeEmploymentDetailsPost structure.
type EmployeeEmploymentDetailsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentDetailsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmployeeEmploymentDetailsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeEmploymentDetailsPostCreated creates a EmployeeEmploymentDetailsPostCreated with default headers values
func NewEmployeeEmploymentDetailsPostCreated() *EmployeeEmploymentDetailsPostCreated {
	return &EmployeeEmploymentDetailsPostCreated{}
}

/*EmployeeEmploymentDetailsPostCreated handles this case with default header values.

successfully created
*/
type EmployeeEmploymentDetailsPostCreated struct {
	Payload *models.ResponseWrapperEmploymentDetails
}

func (o *EmployeeEmploymentDetailsPostCreated) Error() string {
	return fmt.Sprintf("[POST /employee/employment/details][%d] employeeEmploymentDetailsPostCreated  %+v", 201, o.Payload)
}

func (o *EmployeeEmploymentDetailsPostCreated) GetPayload() *models.ResponseWrapperEmploymentDetails {
	return o.Payload
}

func (o *EmployeeEmploymentDetailsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEmploymentDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
