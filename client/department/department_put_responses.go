// Code generated by go-swagger; DO NOT EDIT.

package department

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// DepartmentPutReader is a Reader for the DepartmentPut structure.
type DepartmentPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DepartmentPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDepartmentPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDepartmentPutOK creates a DepartmentPutOK with default headers values
func NewDepartmentPutOK() *DepartmentPutOK {
	return &DepartmentPutOK{}
}

/*DepartmentPutOK handles this case with default header values.

successful operation
*/
type DepartmentPutOK struct {
	Payload *models.ResponseWrapperDepartment
}

func (o *DepartmentPutOK) Error() string {
	return fmt.Sprintf("[PUT /department/{id}][%d] departmentPutOK  %+v", 200, o.Payload)
}

func (o *DepartmentPutOK) GetPayload() *models.ResponseWrapperDepartment {
	return o.Payload
}

func (o *DepartmentPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
