// Code generated by go-swagger; DO NOT EDIT.

package project_specific_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectHourlyRatesProjectSpecificRatesDeleteReader is a Reader for the ProjectHourlyRatesProjectSpecificRatesDelete structure.
type ProjectHourlyRatesProjectSpecificRatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectHourlyRatesProjectSpecificRatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewProjectHourlyRatesProjectSpecificRatesDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewProjectHourlyRatesProjectSpecificRatesDeleteDefault creates a ProjectHourlyRatesProjectSpecificRatesDeleteDefault with default headers values
func NewProjectHourlyRatesProjectSpecificRatesDeleteDefault(code int) *ProjectHourlyRatesProjectSpecificRatesDeleteDefault {
	return &ProjectHourlyRatesProjectSpecificRatesDeleteDefault{
		_statusCode: code,
	}
}

/*ProjectHourlyRatesProjectSpecificRatesDeleteDefault handles this case with default header values.

successful operation
*/
type ProjectHourlyRatesProjectSpecificRatesDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the project hourly rates project specific rates delete default response
func (o *ProjectHourlyRatesProjectSpecificRatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProjectHourlyRatesProjectSpecificRatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /project/hourlyRates/projectSpecificRates/{id}][%d] ProjectHourlyRatesProjectSpecificRatesDelete default ", o._statusCode)
}

func (o *ProjectHourlyRatesProjectSpecificRatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
