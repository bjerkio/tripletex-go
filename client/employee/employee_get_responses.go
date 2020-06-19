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

// EmployeeGetReader is a Reader for the EmployeeGet structure.
type EmployeeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeGetOK creates a EmployeeGetOK with default headers values
func NewEmployeeGetOK() *EmployeeGetOK {
	return &EmployeeGetOK{}
}

/*EmployeeGetOK handles this case with default header values.

successful operation
*/
type EmployeeGetOK struct {
	Payload *models.ResponseWrapperEmployee
}

func (o *EmployeeGetOK) Error() string {
	return fmt.Sprintf("[GET /employee/{id}][%d] employeeGetOK  %+v", 200, o.Payload)
}

func (o *EmployeeGetOK) GetPayload() *models.ResponseWrapperEmployee {
	return o.Payload
}

func (o *EmployeeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEmployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
