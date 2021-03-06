// Code generated by go-swagger; DO NOT EDIT.

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectOrderlineDeleteReader is a Reader for the ProjectOrderlineDelete structure.
type ProjectOrderlineDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectOrderlineDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProjectOrderlineDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProjectOrderlineDeleteDefault creates a ProjectOrderlineDeleteDefault with default headers values
func NewProjectOrderlineDeleteDefault(code int) *ProjectOrderlineDeleteDefault {
	return &ProjectOrderlineDeleteDefault{
		_statusCode: code,
	}
}

/*ProjectOrderlineDeleteDefault handles this case with default header values.

successful operation
*/
type ProjectOrderlineDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the project orderline delete default response
func (o *ProjectOrderlineDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProjectOrderlineDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/orderline/{id}][%d] ProjectOrderline_delete default ", o._statusCode)
}

func (o *ProjectOrderlineDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
