// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpenseMileageAllowancePutReader is a Reader for the TravelExpenseMileageAllowancePut structure.
type TravelExpenseMileageAllowancePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseMileageAllowancePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseMileageAllowancePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseMileageAllowancePutOK creates a TravelExpenseMileageAllowancePutOK with default headers values
func NewTravelExpenseMileageAllowancePutOK() *TravelExpenseMileageAllowancePutOK {
	return &TravelExpenseMileageAllowancePutOK{}
}

/*TravelExpenseMileageAllowancePutOK handles this case with default header values.

successful operation
*/
type TravelExpenseMileageAllowancePutOK struct {
	Payload *models.ResponseWrapperMileageAllowance
}

func (o *TravelExpenseMileageAllowancePutOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/mileageAllowance/{id}][%d] travelExpenseMileageAllowancePutOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseMileageAllowancePutOK) GetPayload() *models.ResponseWrapperMileageAllowance {
	return o.Payload
}

func (o *TravelExpenseMileageAllowancePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperMileageAllowance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
