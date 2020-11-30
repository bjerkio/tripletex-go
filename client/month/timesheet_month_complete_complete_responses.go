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

// TimesheetMonthCompleteCompleteReader is a Reader for the TimesheetMonthCompleteComplete structure.
type TimesheetMonthCompleteCompleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetMonthCompleteCompleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimesheetMonthCompleteCompleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTimesheetMonthCompleteCompleteOK creates a TimesheetMonthCompleteCompleteOK with default headers values
func NewTimesheetMonthCompleteCompleteOK() *TimesheetMonthCompleteCompleteOK {
	return &TimesheetMonthCompleteCompleteOK{}
}

/*TimesheetMonthCompleteCompleteOK handles this case with default header values.

successful operation
*/
type TimesheetMonthCompleteCompleteOK struct {
	Payload *models.ListResponseMonthlyStatus
}

func (o *TimesheetMonthCompleteCompleteOK) Error() string {
	return fmt.Sprintf("[PUT /timesheet/month/:complete][%d] timesheetMonthCompleteCompleteOK  %+v", 200, o.Payload)
}

func (o *TimesheetMonthCompleteCompleteOK) GetPayload() *models.ListResponseMonthlyStatus {
	return o.Payload
}

func (o *TimesheetMonthCompleteCompleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseMonthlyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
