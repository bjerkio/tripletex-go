// Code generated by go-swagger; DO NOT EDIT.

package control_form

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectControlFormSearchReader is a Reader for the ProjectControlFormSearch structure.
type ProjectControlFormSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectControlFormSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectControlFormSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectControlFormSearchOK creates a ProjectControlFormSearchOK with default headers values
func NewProjectControlFormSearchOK() *ProjectControlFormSearchOK {
	return &ProjectControlFormSearchOK{}
}

/*ProjectControlFormSearchOK handles this case with default header values.

successful operation
*/
type ProjectControlFormSearchOK struct {
	Payload *models.ListResponseProjectControlForm
}

func (o *ProjectControlFormSearchOK) Error() string {
	return fmt.Sprintf("[GET /project/controlForm][%d] projectControlFormSearchOK  %+v", 200, o.Payload)
}

func (o *ProjectControlFormSearchOK) GetPayload() *models.ListResponseProjectControlForm {
	return o.Payload
}

func (o *ProjectControlFormSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProjectControlForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
