// Code generated by go-swagger; DO NOT EDIT.

package pension_scheme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// SalarySettingsPensionSchemeSearchReader is a Reader for the SalarySettingsPensionSchemeSearch structure.
type SalarySettingsPensionSchemeSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalarySettingsPensionSchemeSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalarySettingsPensionSchemeSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSalarySettingsPensionSchemeSearchOK creates a SalarySettingsPensionSchemeSearchOK with default headers values
func NewSalarySettingsPensionSchemeSearchOK() *SalarySettingsPensionSchemeSearchOK {
	return &SalarySettingsPensionSchemeSearchOK{}
}

/*SalarySettingsPensionSchemeSearchOK handles this case with default header values.

successful operation
*/
type SalarySettingsPensionSchemeSearchOK struct {
	Payload *models.ListResponsePensionScheme
}

func (o *SalarySettingsPensionSchemeSearchOK) Error() string {
	return fmt.Sprintf("[GET /salary/settings/pensionScheme][%d] salarySettingsPensionSchemeSearchOK  %+v", 200, o.Payload)
}

func (o *SalarySettingsPensionSchemeSearchOK) GetPayload() *models.ListResponsePensionScheme {
	return o.Payload
}

func (o *SalarySettingsPensionSchemeSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePensionScheme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
