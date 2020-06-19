// Code generated by go-swagger; DO NOT EDIT.

package employment_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeReader is a Reader for the EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentType structure.
type EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK creates a EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK with default headers values
func NewEmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK() *EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK {
	return &EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK{}
}

/*EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK handles this case with default header values.

successful operation
*/
type EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK struct {
	Payload *models.ListResponseEmploymentType
}

func (o *EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/employmentType/maritimeEmploymentType][%d] employeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK  %+v", 200, o.Payload)
}

func (o *EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK) GetPayload() *models.ListResponseEmploymentType {
	return o.Payload
}

func (o *EmployeeEmploymentEmploymentTypeMaritimeEmploymentTypeGetMaritimeEmploymentTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseEmploymentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
