// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TimesheetEntryRecentProjectsGetRecentProjectsReader is a Reader for the TimesheetEntryRecentProjectsGetRecentProjects structure.
type TimesheetEntryRecentProjectsGetRecentProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetEntryRecentProjectsGetRecentProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetEntryRecentProjectsGetRecentProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetEntryRecentProjectsGetRecentProjectsOK creates a TimesheetEntryRecentProjectsGetRecentProjectsOK with default headers values
func NewTimesheetEntryRecentProjectsGetRecentProjectsOK() *TimesheetEntryRecentProjectsGetRecentProjectsOK {
	return &TimesheetEntryRecentProjectsGetRecentProjectsOK{}
}

/*TimesheetEntryRecentProjectsGetRecentProjectsOK handles this case with default header values.

successful operation
*/
type TimesheetEntryRecentProjectsGetRecentProjectsOK struct {
	Payload *models.ListResponseProject
}

func (o *TimesheetEntryRecentProjectsGetRecentProjectsOK) Error() string {
	return fmt.Sprintf("[GET /timesheet/entry/>recentProjects][%d] timesheetEntryRecentProjectsGetRecentProjectsOK  %+v", 200, o.Payload)
}

func (o *TimesheetEntryRecentProjectsGetRecentProjectsOK) GetPayload() *models.ListResponseProject {
	return o.Payload
}

func (o *TimesheetEntryRecentProjectsGetRecentProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
