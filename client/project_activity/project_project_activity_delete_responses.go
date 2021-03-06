// Code generated by go-swagger; DO NOT EDIT.

package project_activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectProjectActivityDeleteReader is a Reader for the ProjectProjectActivityDelete structure.
type ProjectProjectActivityDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectProjectActivityDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProjectProjectActivityDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProjectProjectActivityDeleteDefault creates a ProjectProjectActivityDeleteDefault with default headers values
func NewProjectProjectActivityDeleteDefault(code int) *ProjectProjectActivityDeleteDefault {
	return &ProjectProjectActivityDeleteDefault{
		_statusCode: code,
	}
}

/*ProjectProjectActivityDeleteDefault handles this case with default header values.

successful operation
*/
type ProjectProjectActivityDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the project project activity delete default response
func (o *ProjectProjectActivityDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProjectProjectActivityDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/projectActivity/{id}][%d] ProjectProjectActivity_delete default ", o._statusCode)
}

func (o *ProjectProjectActivityDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
