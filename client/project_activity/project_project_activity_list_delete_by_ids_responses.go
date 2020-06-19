// Code generated by go-swagger; DO NOT EDIT.

package project_activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectProjectActivityListDeleteByIdsReader is a Reader for the ProjectProjectActivityListDeleteByIds structure.
type ProjectProjectActivityListDeleteByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectProjectActivityListDeleteByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProjectProjectActivityListDeleteByIdsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProjectProjectActivityListDeleteByIdsDefault creates a ProjectProjectActivityListDeleteByIdsDefault with default headers values
func NewProjectProjectActivityListDeleteByIdsDefault(code int) *ProjectProjectActivityListDeleteByIdsDefault {
	return &ProjectProjectActivityListDeleteByIdsDefault{
		_statusCode: code,
	}
}

/*ProjectProjectActivityListDeleteByIdsDefault handles this case with default header values.

successful operation
*/
type ProjectProjectActivityListDeleteByIdsDefault struct {
	_statusCode int
}

// Code gets the status code for the project project activity list delete by ids default response
func (o *ProjectProjectActivityListDeleteByIdsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectProjectActivityListDeleteByIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/projectActivity/list][%d] ProjectProjectActivityListDeleteByIds default ", o._statusCode)
}

func (o *ProjectProjectActivityListDeleteByIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
