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

// EmployeeEntitlementSearchReader is a Reader for the EmployeeEntitlementSearch structure.
type EmployeeEntitlementSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEntitlementSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEntitlementSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmployeeEntitlementSearchOK creates a EmployeeEntitlementSearchOK with default headers values
func NewEmployeeEntitlementSearchOK() *EmployeeEntitlementSearchOK {
	return &EmployeeEntitlementSearchOK{}
}

/*EmployeeEntitlementSearchOK handles this case with default header values.

successful operation
*/
type EmployeeEntitlementSearchOK struct {
	Payload *models.ListResponseEntitlement
}

func (o *EmployeeEntitlementSearchOK) Error() string {
	return fmt.Sprintf("[GET /employee/entitlement][%d] employeeEntitlementSearchOK  %+v", 200, o.Payload)
}

func (o *EmployeeEntitlementSearchOK) GetPayload() *models.ListResponseEntitlement {
	return o.Payload
}

func (o *EmployeeEntitlementSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseEntitlement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
