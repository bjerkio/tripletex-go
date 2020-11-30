// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeePostReader is a Reader for the EmployeePost structure.
type EmployeePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmployeePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeePostCreated creates a EmployeePostCreated with default headers values
func NewEmployeePostCreated() *EmployeePostCreated {
	return &EmployeePostCreated{}
}

/*EmployeePostCreated handles this case with default header values.

successfully created
*/
type EmployeePostCreated struct {
	Payload *models.ResponseWrapperEmployee
}

func (o *EmployeePostCreated) Error() string {
	return fmt.Sprintf("[POST /employee][%d] employeePostCreated  %+v", 201, o.Payload)
}

func (o *EmployeePostCreated) GetPayload() *models.ResponseWrapperEmployee {
	return o.Payload
}

func (o *EmployeePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEmployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
