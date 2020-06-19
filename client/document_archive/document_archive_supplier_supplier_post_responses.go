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

// DocumentArchiveSupplierSupplierPostReader is a Reader for the DocumentArchiveSupplierSupplierPost structure.
type DocumentArchiveSupplierSupplierPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentArchiveSupplierSupplierPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDocumentArchiveSupplierSupplierPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDocumentArchiveSupplierSupplierPostCreated creates a DocumentArchiveSupplierSupplierPostCreated with default headers values
func NewDocumentArchiveSupplierSupplierPostCreated() *DocumentArchiveSupplierSupplierPostCreated {
	return &DocumentArchiveSupplierSupplierPostCreated{}
}

/*DocumentArchiveSupplierSupplierPostCreated handles this case with default header values.

successfully created
*/
type DocumentArchiveSupplierSupplierPostCreated struct {
	Payload *models.ResponseWrapperDocumentArchive
}

func (o *DocumentArchiveSupplierSupplierPostCreated) Error() string {
	return fmt.Sprintf("[POST /documentArchive/supplier/{id}][%d] documentArchiveSupplierSupplierPostCreated  %+v", 201, o.Payload)
}

func (o *DocumentArchiveSupplierSupplierPostCreated) GetPayload() *models.ResponseWrapperDocumentArchive {
	return o.Payload
}

func (o *DocumentArchiveSupplierSupplierPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDocumentArchive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
