// Code generated by go-swagger; DO NOT EDIT.

package purchase_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// PurchaseOrderSendByEmailSendByEmailReader is a Reader for the PurchaseOrderSendByEmailSendByEmail structure.
type PurchaseOrderSendByEmailSendByEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderSendByEmailSendByEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurchaseOrderSendByEmailSendByEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurchaseOrderSendByEmailSendByEmailOK creates a PurchaseOrderSendByEmailSendByEmailOK with default headers values
func NewPurchaseOrderSendByEmailSendByEmailOK() *PurchaseOrderSendByEmailSendByEmailOK {
	return &PurchaseOrderSendByEmailSendByEmailOK{}
}

/*PurchaseOrderSendByEmailSendByEmailOK handles this case with default header values.

successful operation
*/
type PurchaseOrderSendByEmailSendByEmailOK struct {
	Payload *models.ResponseWrapperPurchaseOrder
}

func (o *PurchaseOrderSendByEmailSendByEmailOK) Error() string {
	return fmt.Sprintf("[PUT /purchaseOrder/{id}/:sendByEmail][%d] purchaseOrderSendByEmailSendByEmailOK  %+v", 200, o.Payload)
}

func (o *PurchaseOrderSendByEmailSendByEmailOK) GetPayload() *models.ResponseWrapperPurchaseOrder {
	return o.Payload
}

func (o *PurchaseOrderSendByEmailSendByEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPurchaseOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
