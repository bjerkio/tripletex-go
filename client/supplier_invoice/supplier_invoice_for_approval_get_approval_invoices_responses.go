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

// SupplierInvoiceForApprovalGetApprovalInvoicesReader is a Reader for the SupplierInvoiceForApprovalGetApprovalInvoices structure.
type SupplierInvoiceForApprovalGetApprovalInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplierInvoiceForApprovalGetApprovalInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplierInvoiceForApprovalGetApprovalInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSupplierInvoiceForApprovalGetApprovalInvoicesOK creates a SupplierInvoiceForApprovalGetApprovalInvoicesOK with default headers values
func NewSupplierInvoiceForApprovalGetApprovalInvoicesOK() *SupplierInvoiceForApprovalGetApprovalInvoicesOK {
	return &SupplierInvoiceForApprovalGetApprovalInvoicesOK{}
}

/*SupplierInvoiceForApprovalGetApprovalInvoicesOK handles this case with default header values.

successful operation
*/
type SupplierInvoiceForApprovalGetApprovalInvoicesOK struct {
	Payload *models.ListResponseSupplierInvoice
}

func (o *SupplierInvoiceForApprovalGetApprovalInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /supplierInvoice/forApproval][%d] supplierInvoiceForApprovalGetApprovalInvoicesOK  %+v", 200, o.Payload)
}

func (o *SupplierInvoiceForApprovalGetApprovalInvoicesOK) GetPayload() *models.ListResponseSupplierInvoice {
	return o.Payload
}

func (o *SupplierInvoiceForApprovalGetApprovalInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseSupplierInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
