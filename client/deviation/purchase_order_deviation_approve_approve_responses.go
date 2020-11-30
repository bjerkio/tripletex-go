// Code generated by go-swagger; DO NOT EDIT.

package deviation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// PurchaseOrderDeviationApproveApproveReader is a Reader for the PurchaseOrderDeviationApproveApprove structure.
type PurchaseOrderDeviationApproveApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurchaseOrderDeviationApproveApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurchaseOrderDeviationApproveApproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurchaseOrderDeviationApproveApproveOK creates a PurchaseOrderDeviationApproveApproveOK with default headers values
func NewPurchaseOrderDeviationApproveApproveOK() *PurchaseOrderDeviationApproveApproveOK {
	return &PurchaseOrderDeviationApproveApproveOK{}
}

/*PurchaseOrderDeviationApproveApproveOK handles this case with default header values.

successful operation
*/
type PurchaseOrderDeviationApproveApproveOK struct {
	Payload *models.ResponseWrapperPurchaseOrder
}

func (o *PurchaseOrderDeviationApproveApproveOK) Error() string {
	return fmt.Sprintf("[PUT /purchaseOrder/deviation/{id}/:approve][%d] purchaseOrderDeviationApproveApproveOK  %+v", 200, o.Payload)
}

func (o *PurchaseOrderDeviationApproveApproveOK) GetPayload() *models.ResponseWrapperPurchaseOrder {
	return o.Payload
}

func (o *PurchaseOrderDeviationApproveApproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPurchaseOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
