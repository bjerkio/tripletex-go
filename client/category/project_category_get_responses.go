// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProjectCategoryGetReader is a Reader for the ProjectCategoryGet structure.
type ProjectCategoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectCategoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectCategoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectCategoryGetOK creates a ProjectCategoryGetOK with default headers values
func NewProjectCategoryGetOK() *ProjectCategoryGetOK {
	return &ProjectCategoryGetOK{}
}

/*ProjectCategoryGetOK handles this case with default header values.

successful operation
*/
type ProjectCategoryGetOK struct {
	Payload *models.ResponseWrapperProjectCategory
}

func (o *ProjectCategoryGetOK) Error() string {
	return fmt.Sprintf("[GET /project/category/{id}][%d] projectCategoryGetOK  %+v", 200, o.Payload)
}

func (o *ProjectCategoryGetOK) GetPayload() *models.ResponseWrapperProjectCategory {
	return o.Payload
}

func (o *ProjectCategoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
