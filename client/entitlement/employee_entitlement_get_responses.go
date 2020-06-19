// Code generated by go-swagger; DO NOT EDIT.

package entitlement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeeEntitlementGetReader is a Reader for the EmployeeEntitlementGet structure.
type EmployeeEntitlementGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEntitlementGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEntitlementGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeEntitlementGetOK creates a EmployeeEntitlementGetOK with default headers values
func NewEmployeeEntitlementGetOK() *EmployeeEntitlementGetOK {
	return &EmployeeEntitlementGetOK{}
}

/*EmployeeEntitlementGetOK handles this case with default header values.

successful operation
*/
type EmployeeEntitlementGetOK struct {
	Payload *models.ResponseWrapperEntitlement
}

func (o *EmployeeEntitlementGetOK) Error() string {
	return fmt.Sprintf("[GET /employee/entitlement/{id}][%d] employeeEntitlementGetOK  %+v", 200, o.Payload)
}

func (o *EmployeeEntitlementGetOK) GetPayload() *models.ResponseWrapperEntitlement {
	return o.Payload
}

func (o *EmployeeEntitlementGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperEntitlement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
