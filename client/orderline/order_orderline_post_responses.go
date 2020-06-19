// Code generated by go-swagger; DO NOT EDIT.

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// OrderOrderlinePostReader is a Reader for the OrderOrderlinePost structure.
type OrderOrderlinePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderOrderlinePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrderOrderlinePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderOrderlinePostCreated creates a OrderOrderlinePostCreated with default headers values
func NewOrderOrderlinePostCreated() *OrderOrderlinePostCreated {
	return &OrderOrderlinePostCreated{}
}

/*OrderOrderlinePostCreated handles this case with default header values.

successfully created
*/
type OrderOrderlinePostCreated struct {
	Payload *models.ResponseWrapperOrderLine
}

func (o *OrderOrderlinePostCreated) Error() string {
	return fmt.Sprintf("[POST /order/orderline][%d] orderOrderlinePostCreated  %+v", 201, o.Payload)
}

func (o *OrderOrderlinePostCreated) GetPayload() *models.ResponseWrapperOrderLine {
	return o.Payload
}

func (o *OrderOrderlinePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperOrderLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
