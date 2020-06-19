// Code generated by go-swagger; DO NOT EDIT.

package period

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProjectPeriodHourlistReportHourlistReportReader is a Reader for the ProjectPeriodHourlistReportHourlistReport structure.
type ProjectPeriodHourlistReportHourlistReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectPeriodHourlistReportHourlistReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectPeriodHourlistReportHourlistReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectPeriodHourlistReportHourlistReportOK creates a ProjectPeriodHourlistReportHourlistReportOK with default headers values
func NewProjectPeriodHourlistReportHourlistReportOK() *ProjectPeriodHourlistReportHourlistReportOK {
	return &ProjectPeriodHourlistReportHourlistReportOK{}
}

/*ProjectPeriodHourlistReportHourlistReportOK handles this case with default header values.

successful operation
*/
type ProjectPeriodHourlistReportHourlistReportOK struct {
	Payload *models.ResponseWrapperProjectPeriodHourlyReport
}

func (o *ProjectPeriodHourlistReportHourlistReportOK) Error() string {
	return fmt.Sprintf("[GET /project/{id}/period/hourlistReport][%d] projectPeriodHourlistReportHourlistReportOK  %+v", 200, o.Payload)
}

func (o *ProjectPeriodHourlistReportHourlistReportOK) GetPayload() *models.ResponseWrapperProjectPeriodHourlyReport {
	return o.Payload
}

func (o *ProjectPeriodHourlistReportHourlistReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectPeriodHourlyReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
