// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EventSubscriptionDeleteReader is a Reader for the EventSubscriptionDelete structure.
type EventSubscriptionDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventSubscriptionDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEventSubscriptionDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEventSubscriptionDeleteDefault creates a EventSubscriptionDeleteDefault with default headers values
func NewEventSubscriptionDeleteDefault(code int) *EventSubscriptionDeleteDefault {
	return &EventSubscriptionDeleteDefault{
		_statusCode: code,
	}
}

/*EventSubscriptionDeleteDefault handles this case with default header values.

successful operation
*/
type EventSubscriptionDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the event subscription delete default response
func (o *EventSubscriptionDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EventSubscriptionDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /event/subscription/{id}][%d] EventSubscriptionDelete default ", o._statusCode)
}

func (o *EventSubscriptionDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
