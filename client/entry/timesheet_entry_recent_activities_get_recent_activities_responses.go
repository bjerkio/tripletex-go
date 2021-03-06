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

// TimesheetEntryRecentActivitiesGetRecentActivitiesReader is a Reader for the TimesheetEntryRecentActivitiesGetRecentActivities structure.
type TimesheetEntryRecentActivitiesGetRecentActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetEntryRecentActivitiesGetRecentActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetEntryRecentActivitiesGetRecentActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetEntryRecentActivitiesGetRecentActivitiesOK creates a TimesheetEntryRecentActivitiesGetRecentActivitiesOK with default headers values
func NewTimesheetEntryRecentActivitiesGetRecentActivitiesOK() *TimesheetEntryRecentActivitiesGetRecentActivitiesOK {
	return &TimesheetEntryRecentActivitiesGetRecentActivitiesOK{}
}

/*TimesheetEntryRecentActivitiesGetRecentActivitiesOK handles this case with default header values.

successful operation
*/
type TimesheetEntryRecentActivitiesGetRecentActivitiesOK struct {
	Payload *models.ListResponseActivity
}

func (o *TimesheetEntryRecentActivitiesGetRecentActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /timesheet/entry/>recentActivities][%d] timesheetEntryRecentActivitiesGetRecentActivitiesOK  %+v", 200, o.Payload)
}

func (o *TimesheetEntryRecentActivitiesGetRecentActivitiesOK) GetPayload() *models.ListResponseActivity {
	return o.Payload
}

func (o *TimesheetEntryRecentActivitiesGetRecentActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
