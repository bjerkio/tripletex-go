// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectForTimeSheetGetForTimeSheetReader is a Reader for the ProjectForTimeSheetGetForTimeSheet structure.
type ProjectForTimeSheetGetForTimeSheetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectForTimeSheetGetForTimeSheetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectForTimeSheetGetForTimeSheetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectForTimeSheetGetForTimeSheetOK creates a ProjectForTimeSheetGetForTimeSheetOK with default headers values
func NewProjectForTimeSheetGetForTimeSheetOK() *ProjectForTimeSheetGetForTimeSheetOK {
	return &ProjectForTimeSheetGetForTimeSheetOK{}
}

/*ProjectForTimeSheetGetForTimeSheetOK handles this case with default header values.

successful operation
*/
type ProjectForTimeSheetGetForTimeSheetOK struct {
	Payload *models.ListResponseProject
}

func (o *ProjectForTimeSheetGetForTimeSheetOK) Error() string {
	return fmt.Sprintf("[GET /project/>forTimeSheet][%d] projectForTimeSheetGetForTimeSheetOK  %+v", 200, o.Payload)
}

func (o *ProjectForTimeSheetGetForTimeSheetOK) GetPayload() *models.ListResponseProject {
	return o.Payload
}

func (o *ProjectForTimeSheetGetForTimeSheetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
