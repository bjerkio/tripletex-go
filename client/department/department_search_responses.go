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

// DepartmentSearchReader is a Reader for the DepartmentSearch structure.
type DepartmentSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DepartmentSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDepartmentSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDepartmentSearchOK creates a DepartmentSearchOK with default headers values
func NewDepartmentSearchOK() *DepartmentSearchOK {
	return &DepartmentSearchOK{}
}

/*DepartmentSearchOK handles this case with default header values.

successful operation
*/
type DepartmentSearchOK struct {
	Payload *models.ListResponseDepartment
}

func (o *DepartmentSearchOK) Error() string {
	return fmt.Sprintf("[GET /department][%d] departmentSearchOK  %+v", 200, o.Payload)
}

func (o *DepartmentSearchOK) GetPayload() *models.ListResponseDepartment {
	return o.Payload
}

func (o *DepartmentSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDepartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
