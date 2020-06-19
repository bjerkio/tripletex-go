// Code generated by go-swagger; DO NOT EDIT.

package participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectParticipantListDeleteByIdsReader is a Reader for the ProjectParticipantListDeleteByIds structure.
type ProjectParticipantListDeleteByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectParticipantListDeleteByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProjectParticipantListDeleteByIdsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProjectParticipantListDeleteByIdsDefault creates a ProjectParticipantListDeleteByIdsDefault with default headers values
func NewProjectParticipantListDeleteByIdsDefault(code int) *ProjectParticipantListDeleteByIdsDefault {
	return &ProjectParticipantListDeleteByIdsDefault{
		_statusCode: code,
	}
}

/*ProjectParticipantListDeleteByIdsDefault handles this case with default header values.

successful operation
*/
type ProjectParticipantListDeleteByIdsDefault struct {
	_statusCode int
}

// Code gets the status code for the project participant list delete by ids default response
func (o *ProjectParticipantListDeleteByIdsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectParticipantListDeleteByIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/participant/list][%d] ProjectParticipantListDeleteByIds default ", o._statusCode)
}

func (o *ProjectParticipantListDeleteByIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
