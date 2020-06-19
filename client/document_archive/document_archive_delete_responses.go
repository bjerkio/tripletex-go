// Code generated by go-swagger; DO NOT EDIT.

package document_archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DocumentArchiveDeleteReader is a Reader for the DocumentArchiveDelete structure.
type DocumentArchiveDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentArchiveDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDocumentArchiveDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDocumentArchiveDeleteDefault creates a DocumentArchiveDeleteDefault with default headers values
func NewDocumentArchiveDeleteDefault(code int) *DocumentArchiveDeleteDefault {
	return &DocumentArchiveDeleteDefault{
		_statusCode: code,
	}
}

/*DocumentArchiveDeleteDefault handles this case with default header values.

successful operation
*/
type DocumentArchiveDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the document archive delete default response
func (o *DocumentArchiveDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DocumentArchiveDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /documentArchive/{id}][%d] DocumentArchiveDelete default ", o._statusCode)
}

func (o *DocumentArchiveDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
