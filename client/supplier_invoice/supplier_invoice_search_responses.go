// Code generated by go-swagger; DO NOT EDIT.

package supplier_invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// SupplierInvoiceSearchReader is a Reader for the SupplierInvoiceSearch structure.
type SupplierInvoiceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplierInvoiceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplierInvoiceSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSupplierInvoiceSearchOK creates a SupplierInvoiceSearchOK with default headers values
func NewSupplierInvoiceSearchOK() *SupplierInvoiceSearchOK {
	return &SupplierInvoiceSearchOK{}
}

/*SupplierInvoiceSearchOK handles this case with default header values.

successful operation
*/
type SupplierInvoiceSearchOK struct {
	Payload *models.ListResponseSupplierInvoice
}

func (o *SupplierInvoiceSearchOK) Error() string {
	return fmt.Sprintf("[GET /supplierInvoice][%d] supplierInvoiceSearchOK  %+v", 200, o.Payload)
}

func (o *SupplierInvoiceSearchOK) GetPayload() *models.ListResponseSupplierInvoice {
	return o.Payload
}

func (o *SupplierInvoiceSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseSupplierInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
