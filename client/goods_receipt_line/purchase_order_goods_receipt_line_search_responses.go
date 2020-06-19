// Code generated by go-swagger; DO NOT EDIT.

package goods_receipt_line

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// PurchaseOrderGoodsReceiptLineSearchReader is a Reader for the PurchaseOrderGoodsReceiptLineSearch structure.
type PurchaseOrderGoodsReceiptLineSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderGoodsReceiptLineSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurchaseOrderGoodsReceiptLineSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurchaseOrderGoodsReceiptLineSearchOK creates a PurchaseOrderGoodsReceiptLineSearchOK with default headers values
func NewPurchaseOrderGoodsReceiptLineSearchOK() *PurchaseOrderGoodsReceiptLineSearchOK {
	return &PurchaseOrderGoodsReceiptLineSearchOK{}
}

/*PurchaseOrderGoodsReceiptLineSearchOK handles this case with default header values.

successful operation
*/
type PurchaseOrderGoodsReceiptLineSearchOK struct {
	Payload *models.ListResponseGoodsReceiptLine
}

func (o *PurchaseOrderGoodsReceiptLineSearchOK) Error() string {
	return fmt.Sprintf("[GET /purchaseOrder/goodsReceiptLine][%d] purchaseOrderGoodsReceiptLineSearchOK  %+v", 200, o.Payload)
}

func (o *PurchaseOrderGoodsReceiptLineSearchOK) GetPayload() *models.ListResponseGoodsReceiptLine {
	return o.Payload
}

func (o *PurchaseOrderGoodsReceiptLineSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseGoodsReceiptLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
