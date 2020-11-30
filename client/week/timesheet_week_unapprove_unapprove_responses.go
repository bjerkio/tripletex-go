// Code generated by go-swagger; DO NOT EDIT.

package week

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetWeekUnapproveUnapproveReader is a Reader for the TimesheetWeekUnapproveUnapprove structure.
type TimesheetWeekUnapproveUnapproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetWeekUnapproveUnapproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetWeekUnapproveUnapproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetWeekUnapproveUnapproveOK creates a TimesheetWeekUnapproveUnapproveOK with default headers values
func NewTimesheetWeekUnapproveUnapproveOK() *TimesheetWeekUnapproveUnapproveOK {
	return &TimesheetWeekUnapproveUnapproveOK{}
}

/*TimesheetWeekUnapproveUnapproveOK handles this case with default header values.

successful operation
*/
type TimesheetWeekUnapproveUnapproveOK struct {
	Payload *models.ResponseWrapperWeek
}

func (o *TimesheetWeekUnapproveUnapproveOK) Error() string {
	return fmt.Sprintf("[PUT /timesheet/week/:unapprove][%d] timesheetWeekUnapproveUnapproveOK  %+v", 200, o.Payload)
}

func (o *TimesheetWeekUnapproveUnapproveOK) GetPayload() *models.ResponseWrapperWeek {
	return o.Payload
}

func (o *TimesheetWeekUnapproveUnapproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperWeek)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
