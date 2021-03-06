// Code generated by go-swagger; DO NOT EDIT.

package document_archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// DocumentArchiveSupplierGetSupplierReader is a Reader for the DocumentArchiveSupplierGetSupplier structure.
type DocumentArchiveSupplierGetSupplierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentArchiveSupplierGetSupplierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentArchiveSupplierGetSupplierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentArchiveSupplierGetSupplierOK creates a DocumentArchiveSupplierGetSupplierOK with default headers values
func NewDocumentArchiveSupplierGetSupplierOK() *DocumentArchiveSupplierGetSupplierOK {
	return &DocumentArchiveSupplierGetSupplierOK{}
}

/*DocumentArchiveSupplierGetSupplierOK handles this case with default header values.

successful operation
*/
type DocumentArchiveSupplierGetSupplierOK struct {
	Payload *models.ListResponseDocumentArchive
}

func (o *DocumentArchiveSupplierGetSupplierOK) Error() string {
	return fmt.Sprintf("[GET /documentArchive/supplier/{id}][%d] documentArchiveSupplierGetSupplierOK  %+v", 200, o.Payload)
}

func (o *DocumentArchiveSupplierGetSupplierOK) GetPayload() *models.ListResponseDocumentArchive {
	return o.Payload
}

func (o *DocumentArchiveSupplierGetSupplierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDocumentArchive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
