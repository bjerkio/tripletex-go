// Code generated by go-swagger; DO NOT EDIT.

package division

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// DivisionPutReader is a Reader for the DivisionPut structure.
type DivisionPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DivisionPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDivisionPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDivisionPutOK creates a DivisionPutOK with default headers values
func NewDivisionPutOK() *DivisionPutOK {
	return &DivisionPutOK{}
}

/*DivisionPutOK handles this case with default header values.

successful operation
*/
type DivisionPutOK struct {
	Payload *models.ResponseWrapperDivision
}

func (o *DivisionPutOK) Error() string {
	return fmt.Sprintf("[PUT /division/{id}][%d] divisionPutOK  %+v", 200, o.Payload)
}

func (o *DivisionPutOK) GetPayload() *models.ResponseWrapperDivision {
	return o.Payload
}

func (o *DivisionPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperDivision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
