// Code generated by go-swagger; DO NOT EDIT.

package purchase_order_incoming_invoice_relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListReader is a Reader for the PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostList structure.
type PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated creates a PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated with default headers values
func NewPurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated() *PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated {
	return &PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated{}
}

/*PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated handles this case with default header values.

successfully created
*/
type PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated struct {
	Payload *models.ListResponsePurchaseOrderIncomingInvoiceRelation
}

func (o *PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /purchaseOrder/purchaseOrderIncomingInvoiceRelation/list][%d] purchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated  %+v", 201, o.Payload)
}

func (o *PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated) GetPayload() *models.ListResponsePurchaseOrderIncomingInvoiceRelation {
	return o.Payload
}

func (o *PurchaseOrderPurchaseOrderIncomingInvoiceRelationListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePurchaseOrderIncomingInvoiceRelation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
