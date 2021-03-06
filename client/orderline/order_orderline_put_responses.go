// Code generated by go-swagger; DO NOT EDIT.

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// OrderOrderlinePutReader is a Reader for the OrderOrderlinePut structure.
type OrderOrderlinePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderOrderlinePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderOrderlinePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrderOrderlinePutOK creates a OrderOrderlinePutOK with default headers values
func NewOrderOrderlinePutOK() *OrderOrderlinePutOK {
	return &OrderOrderlinePutOK{}
}

/*OrderOrderlinePutOK handles this case with default header values.

successful operation
*/
type OrderOrderlinePutOK struct {
	Payload *models.ResponseWrapperOrderLine
}

func (o *OrderOrderlinePutOK) Error() string {
	return fmt.Sprintf("[PUT /order/orderline/{id}][%d] orderOrderlinePutOK  %+v", 200, o.Payload)
}

func (o *OrderOrderlinePutOK) GetPayload() *models.ResponseWrapperOrderLine {
	return o.Payload
}

func (o *OrderOrderlinePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperOrderLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
