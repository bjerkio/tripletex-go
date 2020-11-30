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

// ProjectListPostListReader is a Reader for the ProjectListPostList structure.
type ProjectListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectListPostListCreated creates a ProjectListPostListCreated with default headers values
func NewProjectListPostListCreated() *ProjectListPostListCreated {
	return &ProjectListPostListCreated{}
}

/*ProjectListPostListCreated handles this case with default header values.

successfully created
*/
type ProjectListPostListCreated struct {
	Payload *models.ListResponseProject
}

func (o *ProjectListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /project/list][%d] projectListPostListCreated  %+v", 201, o.Payload)
}

func (o *ProjectListPostListCreated) GetPayload() *models.ListResponseProject {
	return o.Payload
}

func (o *ProjectListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
