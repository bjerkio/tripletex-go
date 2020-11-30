// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpenseSettingsGetReader is a Reader for the TravelExpenseSettingsGet structure.
type TravelExpenseSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseSettingsGetOK creates a TravelExpenseSettingsGetOK with default headers values
func NewTravelExpenseSettingsGetOK() *TravelExpenseSettingsGetOK {
	return &TravelExpenseSettingsGetOK{}
}

/*TravelExpenseSettingsGetOK handles this case with default header values.

successful operation
*/
type TravelExpenseSettingsGetOK struct {
	Payload *models.ResponseWrapperTravelExpenseSettings
}

func (o *TravelExpenseSettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/settings][%d] travelExpenseSettingsGetOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseSettingsGetOK) GetPayload() *models.ResponseWrapperTravelExpenseSettings {
	return o.Payload
}

func (o *TravelExpenseSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTravelExpenseSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
