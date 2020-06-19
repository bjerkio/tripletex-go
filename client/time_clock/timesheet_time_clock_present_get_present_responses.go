// Code generated by go-swagger; DO NOT EDIT.

package time_clock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetTimeClockPresentGetPresentReader is a Reader for the TimesheetTimeClockPresentGetPresent structure.
type TimesheetTimeClockPresentGetPresentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetTimeClockPresentGetPresentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetTimeClockPresentGetPresentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimesheetTimeClockPresentGetPresentOK creates a TimesheetTimeClockPresentGetPresentOK with default headers values
func NewTimesheetTimeClockPresentGetPresentOK() *TimesheetTimeClockPresentGetPresentOK {
	return &TimesheetTimeClockPresentGetPresentOK{}
}

/*TimesheetTimeClockPresentGetPresentOK handles this case with default header values.

successful operation
*/
type TimesheetTimeClockPresentGetPresentOK struct {
	Payload *models.ResponseWrapperTimeClock
}

func (o *TimesheetTimeClockPresentGetPresentOK) Error() string {
	return fmt.Sprintf("[GET /timesheet/timeClock/present][%d] timesheetTimeClockPresentGetPresentOK  %+v", 200, o.Payload)
}

func (o *TimesheetTimeClockPresentGetPresentOK) GetPayload() *models.ResponseWrapperTimeClock {
	return o.Payload
}

func (o *TimesheetTimeClockPresentGetPresentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTimeClock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
