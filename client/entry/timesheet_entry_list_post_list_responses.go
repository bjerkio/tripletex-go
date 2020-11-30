// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetEntryListPostListReader is a Reader for the TimesheetEntryListPostList structure.
type TimesheetEntryListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetEntryListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTimesheetEntryListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetEntryListPostListCreated creates a TimesheetEntryListPostListCreated with default headers values
func NewTimesheetEntryListPostListCreated() *TimesheetEntryListPostListCreated {
	return &TimesheetEntryListPostListCreated{}
}

/*TimesheetEntryListPostListCreated handles this case with default header values.

successfully created
*/
type TimesheetEntryListPostListCreated struct {
	Payload *models.ListResponseTimesheetEntry
}

func (o *TimesheetEntryListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /timesheet/entry/list][%d] timesheetEntryListPostListCreated  %+v", 201, o.Payload)
}

func (o *TimesheetEntryListPostListCreated) GetPayload() *models.ListResponseTimesheetEntry {
	return o.Payload
}

func (o *TimesheetEntryListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseTimesheetEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
