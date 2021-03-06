// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ContactPutReader is a Reader for the ContactPut structure.
type ContactPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContactPutOK creates a ContactPutOK with default headers values
func NewContactPutOK() *ContactPutOK {
	return &ContactPutOK{}
}

/*ContactPutOK handles this case with default header values.

successful operation
*/
type ContactPutOK struct {
	Payload *models.ResponseWrapperContact
}

func (o *ContactPutOK) Error() string {
	return fmt.Sprintf("[PUT /contact/{id}][%d] contactPutOK  %+v", 200, o.Payload)
}

func (o *ContactPutOK) GetPayload() *models.ResponseWrapperContact {
	return o.Payload
}

func (o *ContactPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
