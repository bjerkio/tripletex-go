// Code generated by go-swagger; DO NOT EDIT.

package occupation_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeeEmploymentOccupationCodeSearchReader is a Reader for the EmployeeEmploymentOccupationCodeSearch structure.
type EmployeeEmploymentOccupationCodeSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentOccupationCodeSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEmploymentOccupationCodeSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeEmploymentOccupationCodeSearchOK creates a EmployeeEmploymentOccupationCodeSearchOK with default headers values
func NewEmployeeEmploymentOccupationCodeSearchOK() *EmployeeEmploymentOccupationCodeSearchOK {
	return &EmployeeEmploymentOccupationCodeSearchOK{}
}

/*EmployeeEmploymentOccupationCodeSearchOK handles this case with default header values.

successful operation
*/
type EmployeeEmploymentOccupationCodeSearchOK struct {
	Payload *models.ListResponseOccupationCode
}

func (o *EmployeeEmploymentOccupationCodeSearchOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/occupationCode][%d] employeeEmploymentOccupationCodeSearchOK  %+v", 200, o.Payload)
}

func (o *EmployeeEmploymentOccupationCodeSearchOK) GetPayload() *models.ListResponseOccupationCode {
	return o.Payload
}

func (o *EmployeeEmploymentOccupationCodeSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseOccupationCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
