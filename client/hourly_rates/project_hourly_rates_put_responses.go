// Code generated by go-swagger; DO NOT EDIT.

package hourly_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProjectHourlyRatesPutReader is a Reader for the ProjectHourlyRatesPut structure.
type ProjectHourlyRatesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectHourlyRatesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectHourlyRatesPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectHourlyRatesPutOK creates a ProjectHourlyRatesPutOK with default headers values
func NewProjectHourlyRatesPutOK() *ProjectHourlyRatesPutOK {
	return &ProjectHourlyRatesPutOK{}
}

/*ProjectHourlyRatesPutOK handles this case with default header values.

successful operation
*/
type ProjectHourlyRatesPutOK struct {
	Payload *models.ResponseWrapperProjectHourlyRate
}

func (o *ProjectHourlyRatesPutOK) Error() string {
	return fmt.Sprintf("[PUT /project/hourlyRates/{id}][%d] projectHourlyRatesPutOK  %+v", 200, o.Payload)
}

func (o *ProjectHourlyRatesPutOK) GetPayload() *models.ResponseWrapperProjectHourlyRate {
	return o.Payload
}

func (o *ProjectHourlyRatesPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectHourlyRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
