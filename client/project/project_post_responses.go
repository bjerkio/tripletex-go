// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProjectPostReader is a Reader for the ProjectPost structure.
type ProjectPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectPostCreated creates a ProjectPostCreated with default headers values
func NewProjectPostCreated() *ProjectPostCreated {
	return &ProjectPostCreated{}
}

/*ProjectPostCreated handles this case with default header values.

successfully created
*/
type ProjectPostCreated struct {
	Payload *models.ResponseWrapperProject
}

func (o *ProjectPostCreated) Error() string {
	return fmt.Sprintf("[POST /project][%d] projectPostCreated  %+v", 201, o.Payload)
}

func (o *ProjectPostCreated) GetPayload() *models.ResponseWrapperProject {
	return o.Payload
}

func (o *ProjectPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
