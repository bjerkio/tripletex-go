// Code generated by go-swagger; DO NOT EDIT.

package delivery_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// DeliveryAddressGetReader is a Reader for the DeliveryAddressGet structure.
type DeliveryAddressGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeliveryAddressGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeliveryAddressGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeliveryAddressGetOK creates a DeliveryAddressGetOK with default headers values
func NewDeliveryAddressGetOK() *DeliveryAddressGetOK {
	return &DeliveryAddressGetOK{}
}

/*DeliveryAddressGetOK handles this case with default header values.

successful operation
*/
type DeliveryAddressGetOK struct {
	Payload *models.ResponseWrapperDeliveryAddress
}

func (o *DeliveryAddressGetOK) Error() string {
	return fmt.Sprintf("[GET /deliveryAddress/{id}][%d] deliveryAddressGetOK  %+v", 200, o.Payload)
}

func (o *DeliveryAddressGetOK) GetPayload() *models.ResponseWrapperDeliveryAddress {
	return o.Payload
}

func (o *DeliveryAddressGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDeliveryAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
