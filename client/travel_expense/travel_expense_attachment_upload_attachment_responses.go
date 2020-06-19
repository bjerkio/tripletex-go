// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TravelExpenseAttachmentUploadAttachmentReader is a Reader for the TravelExpenseAttachmentUploadAttachment structure.
type TravelExpenseAttachmentUploadAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseAttachmentUploadAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewTravelExpenseAttachmentUploadAttachmentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewTravelExpenseAttachmentUploadAttachmentDefault creates a TravelExpenseAttachmentUploadAttachmentDefault with default headers values
func NewTravelExpenseAttachmentUploadAttachmentDefault(code int) *TravelExpenseAttachmentUploadAttachmentDefault {
	return &TravelExpenseAttachmentUploadAttachmentDefault{
		_statusCode: code,
	}
}

/*TravelExpenseAttachmentUploadAttachmentDefault handles this case with default header values.

successful operation
*/
type TravelExpenseAttachmentUploadAttachmentDefault struct {
	_statusCode int
}

// Code gets the status code for the travel expense attachment upload attachment default response
func (o *TravelExpenseAttachmentUploadAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *TravelExpenseAttachmentUploadAttachmentDefault) Error() string {
	return fmt.Sprintf("[POST /travelExpense/{travelExpenseId}/attachment][%d] TravelExpenseAttachmentUploadAttachment default ", o._statusCode)
}

func (o *TravelExpenseAttachmentUploadAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
