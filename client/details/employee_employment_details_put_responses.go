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

// EmployeeEmploymentDetailsPutReader is a Reader for the EmployeeEmploymentDetailsPut structure.
type EmployeeEmploymentDetailsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentDetailsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEmploymentDetailsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeeEmploymentDetailsPutOK creates a EmployeeEmploymentDetailsPutOK with default headers values
func NewEmployeeEmploymentDetailsPutOK() *EmployeeEmploymentDetailsPutOK {
	return &EmployeeEmploymentDetailsPutOK{}
}

/*EmployeeEmploymentDetailsPutOK handles this case with default header values.

successful operation
*/
type EmployeeEmploymentDetailsPutOK struct {
	Payload *models.ResponseWrapperEmploymentDetails
}

func (o *EmployeeEmploymentDetailsPutOK) Error() string {
	return fmt.Sprintf("[PUT /employee/employment/details/{id}][%d] employeeEmploymentDetailsPutOK  %+v", 200, o.Payload)
}

func (o *EmployeeEmploymentDetailsPutOK) GetPayload() *models.ResponseWrapperEmploymentDetails {
	return o.Payload
}

func (o *EmployeeEmploymentDetailsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEmploymentDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
