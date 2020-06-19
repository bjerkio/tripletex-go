// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// InvoiceCreateCreditNoteCreateCreditNoteReader is a Reader for the InvoiceCreateCreditNoteCreateCreditNote structure.
type InvoiceCreateCreditNoteCreateCreditNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoiceCreateCreditNoteCreateCreditNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoiceCreateCreditNoteCreateCreditNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvoiceCreateCreditNoteCreateCreditNoteOK creates a InvoiceCreateCreditNoteCreateCreditNoteOK with default headers values
func NewInvoiceCreateCreditNoteCreateCreditNoteOK() *InvoiceCreateCreditNoteCreateCreditNoteOK {
	return &InvoiceCreateCreditNoteCreateCreditNoteOK{}
}

/*InvoiceCreateCreditNoteCreateCreditNoteOK handles this case with default header values.

successful operation
*/
type InvoiceCreateCreditNoteCreateCreditNoteOK struct {
	Payload *models.ResponseWrapperInvoice
}

func (o *InvoiceCreateCreditNoteCreateCreditNoteOK) Error() string {
	return fmt.Sprintf("[PUT /invoice/{id}/:createCreditNote][%d] invoiceCreateCreditNoteCreateCreditNoteOK  %+v", 200, o.Payload)
}

func (o *InvoiceCreateCreditNoteCreateCreditNoteOK) GetPayload() *models.ResponseWrapperInvoice {
	return o.Payload
}

func (o *InvoiceCreateCreditNoteCreateCreditNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperInvoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
