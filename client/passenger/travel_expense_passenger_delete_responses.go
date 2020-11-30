// Code generated by go-swagger; DO NOT EDIT.

package passenger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TravelExpensePassengerDeleteReader is a Reader for the TravelExpensePassengerDelete structure.
type TravelExpensePassengerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpensePassengerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewTravelExpensePassengerDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewTravelExpensePassengerDeleteDefault creates a TravelExpensePassengerDeleteDefault with default headers values
func NewTravelExpensePassengerDeleteDefault(code int) *TravelExpensePassengerDeleteDefault {
	return &TravelExpensePassengerDeleteDefault{
		_statusCode: code,
	}
}

/*TravelExpensePassengerDeleteDefault handles this case with default header values.

successful operation
*/
type TravelExpensePassengerDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the travel expense passenger delete default response
func (o *TravelExpensePassengerDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TravelExpensePassengerDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /travelExpense/passenger/{id}][%d] TravelExpensePassenger_delete default ", o._statusCode)
}

func (o *TravelExpensePassengerDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
