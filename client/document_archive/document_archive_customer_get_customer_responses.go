// Code generated by go-swagger; DO NOT EDIT.

package document_archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// DocumentArchiveCustomerGetCustomerReader is a Reader for the DocumentArchiveCustomerGetCustomer structure.
type DocumentArchiveCustomerGetCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentArchiveCustomerGetCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentArchiveCustomerGetCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentArchiveCustomerGetCustomerOK creates a DocumentArchiveCustomerGetCustomerOK with default headers values
func NewDocumentArchiveCustomerGetCustomerOK() *DocumentArchiveCustomerGetCustomerOK {
	return &DocumentArchiveCustomerGetCustomerOK{}
}

/*DocumentArchiveCustomerGetCustomerOK handles this case with default header values.

successful operation
*/
type DocumentArchiveCustomerGetCustomerOK struct {
	Payload *models.ListResponseDocumentArchive
}

func (o *DocumentArchiveCustomerGetCustomerOK) Error() string {
	return fmt.Sprintf("[GET /documentArchive/customer/{id}][%d] documentArchiveCustomerGetCustomerOK  %+v", 200, o.Payload)
}

func (o *DocumentArchiveCustomerGetCustomerOK) GetPayload() *models.ListResponseDocumentArchive {
	return o.Payload
}

func (o *DocumentArchiveCustomerGetCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDocumentArchive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
