// Code generated by go-swagger; DO NOT EDIT.

package supplier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// SupplierGetReader is a Reader for the SupplierGet structure.
type SupplierGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplierGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplierGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSupplierGetOK creates a SupplierGetOK with default headers values
func NewSupplierGetOK() *SupplierGetOK {
	return &SupplierGetOK{}
}

/*SupplierGetOK handles this case with default header values.

successful operation
*/
type SupplierGetOK struct {
	Payload *models.ResponseWrapperSupplier
}

func (o *SupplierGetOK) Error() string {
	return fmt.Sprintf("[GET /supplier/{id}][%d] supplierGetOK  %+v", 200, o.Payload)
}

func (o *SupplierGetOK) GetPayload() *models.ResponseWrapperSupplier {
	return o.Payload
}

func (o *SupplierGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSupplier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
