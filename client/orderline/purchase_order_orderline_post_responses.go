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

// PurchaseOrderOrderlinePostReader is a Reader for the PurchaseOrderOrderlinePost structure.
type PurchaseOrderOrderlinePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderOrderlinePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPurchaseOrderOrderlinePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurchaseOrderOrderlinePostCreated creates a PurchaseOrderOrderlinePostCreated with default headers values
func NewPurchaseOrderOrderlinePostCreated() *PurchaseOrderOrderlinePostCreated {
	return &PurchaseOrderOrderlinePostCreated{}
}

/*PurchaseOrderOrderlinePostCreated handles this case with default header values.

successfully created
*/
type PurchaseOrderOrderlinePostCreated struct {
	Payload *models.ResponseWrapperPurchaseOrderline
}

func (o *PurchaseOrderOrderlinePostCreated) Error() string {
	return fmt.Sprintf("[POST /purchaseOrder/orderline][%d] purchaseOrderOrderlinePostCreated  %+v", 201, o.Payload)
}

func (o *PurchaseOrderOrderlinePostCreated) GetPayload() *models.ResponseWrapperPurchaseOrderline {
	return o.Payload
}

func (o *PurchaseOrderOrderlinePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPurchaseOrderline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
