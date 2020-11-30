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

// DocumentArchiveCustomerCustomerPostReader is a Reader for the DocumentArchiveCustomerCustomerPost structure.
type DocumentArchiveCustomerCustomerPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentArchiveCustomerCustomerPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDocumentArchiveCustomerCustomerPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentArchiveCustomerCustomerPostCreated creates a DocumentArchiveCustomerCustomerPostCreated with default headers values
func NewDocumentArchiveCustomerCustomerPostCreated() *DocumentArchiveCustomerCustomerPostCreated {
	return &DocumentArchiveCustomerCustomerPostCreated{}
}

/*DocumentArchiveCustomerCustomerPostCreated handles this case with default header values.

successfully created
*/
type DocumentArchiveCustomerCustomerPostCreated struct {
	Payload *models.ResponseWrapperDocumentArchive
}

func (o *DocumentArchiveCustomerCustomerPostCreated) Error() string {
	return fmt.Sprintf("[POST /documentArchive/customer/{id}][%d] documentArchiveCustomerCustomerPostCreated  %+v", 201, o.Payload)
}

func (o *DocumentArchiveCustomerCustomerPostCreated) GetPayload() *models.ResponseWrapperDocumentArchive {
	return o.Payload
}

func (o *DocumentArchiveCustomerCustomerPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDocumentArchive)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
