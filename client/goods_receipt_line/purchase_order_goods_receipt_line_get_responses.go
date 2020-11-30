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

// PurchaseOrderGoodsReceiptLineGetReader is a Reader for the PurchaseOrderGoodsReceiptLineGet structure.
type PurchaseOrderGoodsReceiptLineGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderGoodsReceiptLineGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurchaseOrderGoodsReceiptLineGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurchaseOrderGoodsReceiptLineGetOK creates a PurchaseOrderGoodsReceiptLineGetOK with default headers values
func NewPurchaseOrderGoodsReceiptLineGetOK() *PurchaseOrderGoodsReceiptLineGetOK {
	return &PurchaseOrderGoodsReceiptLineGetOK{}
}

/*PurchaseOrderGoodsReceiptLineGetOK handles this case with default header values.

successful operation
*/
type PurchaseOrderGoodsReceiptLineGetOK struct {
	Payload *models.ResponseWrapperGoodsReceiptLine
}

func (o *PurchaseOrderGoodsReceiptLineGetOK) Error() string {
	return fmt.Sprintf("[GET /purchaseOrder/goodsReceiptLine/{id}][%d] purchaseOrderGoodsReceiptLineGetOK  %+v", 200, o.Payload)
}

func (o *PurchaseOrderGoodsReceiptLineGetOK) GetPayload() *models.ResponseWrapperGoodsReceiptLine {
	return o.Payload
}

func (o *PurchaseOrderGoodsReceiptLineGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperGoodsReceiptLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
