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

// PurchaseOrderOrderlineGetReader is a Reader for the PurchaseOrderOrderlineGet structure.
type PurchaseOrderOrderlineGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderOrderlineGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurchaseOrderOrderlineGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurchaseOrderOrderlineGetOK creates a PurchaseOrderOrderlineGetOK with default headers values
func NewPurchaseOrderOrderlineGetOK() *PurchaseOrderOrderlineGetOK {
	return &PurchaseOrderOrderlineGetOK{}
}

/*PurchaseOrderOrderlineGetOK handles this case with default header values.

successful operation
*/
type PurchaseOrderOrderlineGetOK struct {
	Payload *models.ResponseWrapperPurchaseOrderline
}

func (o *PurchaseOrderOrderlineGetOK) Error() string {
	return fmt.Sprintf("[GET /purchaseOrder/orderline/{id}][%d] purchaseOrderOrderlineGetOK  %+v", 200, o.Payload)
}

func (o *PurchaseOrderOrderlineGetOK) GetPayload() *models.ResponseWrapperPurchaseOrderline {
	return o.Payload
}

func (o *PurchaseOrderOrderlineGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPurchaseOrderline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
