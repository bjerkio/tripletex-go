// Code generated by go-swagger; DO NOT EDIT.

package division

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// DivisionListPutListReader is a Reader for the DivisionListPutList structure.
type DivisionListPutListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DivisionListPutListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDivisionListPutListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDivisionListPutListOK creates a DivisionListPutListOK with default headers values
func NewDivisionListPutListOK() *DivisionListPutListOK {
	return &DivisionListPutListOK{}
}

/*DivisionListPutListOK handles this case with default header values.

successful operation
*/
type DivisionListPutListOK struct {
	Payload *models.ListResponseDivision
}

func (o *DivisionListPutListOK) Error() string {
	return fmt.Sprintf("[PUT /division/list][%d] divisionListPutListOK  %+v", 200, o.Payload)
}

func (o *DivisionListPutListOK) GetPayload() *models.ListResponseDivision {
	return o.Payload
}

func (o *DivisionListPutListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseDivision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
