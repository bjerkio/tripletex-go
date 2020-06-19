// Code generated by go-swagger; DO NOT EDIT.

package month

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetMonthReopenReopenReader is a Reader for the TimesheetMonthReopenReopen structure.
type TimesheetMonthReopenReopenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetMonthReopenReopenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetMonthReopenReopenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimesheetMonthReopenReopenOK creates a TimesheetMonthReopenReopenOK with default headers values
func NewTimesheetMonthReopenReopenOK() *TimesheetMonthReopenReopenOK {
	return &TimesheetMonthReopenReopenOK{}
}

/*TimesheetMonthReopenReopenOK handles this case with default header values.

successful operation
*/
type TimesheetMonthReopenReopenOK struct {
	Payload *models.ListResponseMonthlyStatus
}

func (o *TimesheetMonthReopenReopenOK) Error() string {
	return fmt.Sprintf("[PUT /timesheet/month/:reopen][%d] timesheetMonthReopenReopenOK  %+v", 200, o.Payload)
}

func (o *TimesheetMonthReopenReopenOK) GetPayload() *models.ListResponseMonthlyStatus {
	return o.Payload
}

func (o *TimesheetMonthReopenReopenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseMonthlyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
