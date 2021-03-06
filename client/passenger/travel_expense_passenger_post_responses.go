// Code generated by go-swagger; DO NOT EDIT.

package passenger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TravelExpensePassengerPostReader is a Reader for the TravelExpensePassengerPost structure.
type TravelExpensePassengerPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpensePassengerPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTravelExpensePassengerPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpensePassengerPostCreated creates a TravelExpensePassengerPostCreated with default headers values
func NewTravelExpensePassengerPostCreated() *TravelExpensePassengerPostCreated {
	return &TravelExpensePassengerPostCreated{}
}

/*TravelExpensePassengerPostCreated handles this case with default header values.

successfully created
*/
type TravelExpensePassengerPostCreated struct {
	Payload *models.ResponseWrapperPassenger
}

func (o *TravelExpensePassengerPostCreated) Error() string {
	return fmt.Sprintf("[POST /travelExpense/passenger][%d] travelExpensePassengerPostCreated  %+v", 201, o.Payload)
}

func (o *TravelExpensePassengerPostCreated) GetPayload() *models.ResponseWrapperPassenger {
	return o.Payload
}

func (o *TravelExpensePassengerPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPassenger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
