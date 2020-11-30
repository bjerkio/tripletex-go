// Code generated by go-swagger; DO NOT EDIT.

package per_diem_compensation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpensePerDiemCompensationGetReader is a Reader for the TravelExpensePerDiemCompensationGet structure.
type TravelExpensePerDiemCompensationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpensePerDiemCompensationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpensePerDiemCompensationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpensePerDiemCompensationGetOK creates a TravelExpensePerDiemCompensationGetOK with default headers values
func NewTravelExpensePerDiemCompensationGetOK() *TravelExpensePerDiemCompensationGetOK {
	return &TravelExpensePerDiemCompensationGetOK{}
}

/*TravelExpensePerDiemCompensationGetOK handles this case with default header values.

successful operation
*/
type TravelExpensePerDiemCompensationGetOK struct {
	Payload *models.ResponseWrapperPerDiemCompensation
}

func (o *TravelExpensePerDiemCompensationGetOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/perDiemCompensation/{id}][%d] travelExpensePerDiemCompensationGetOK  %+v", 200, o.Payload)
}

func (o *TravelExpensePerDiemCompensationGetOK) GetPayload() *models.ResponseWrapperPerDiemCompensation {
	return o.Payload
}

func (o *TravelExpensePerDiemCompensationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperPerDiemCompensation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
