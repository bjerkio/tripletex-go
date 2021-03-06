// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// TravelExpenseMileageAllowanceSearchReader is a Reader for the TravelExpenseMileageAllowanceSearch structure.
type TravelExpenseMileageAllowanceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseMileageAllowanceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseMileageAllowanceSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseMileageAllowanceSearchOK creates a TravelExpenseMileageAllowanceSearchOK with default headers values
func NewTravelExpenseMileageAllowanceSearchOK() *TravelExpenseMileageAllowanceSearchOK {
	return &TravelExpenseMileageAllowanceSearchOK{}
}

/*TravelExpenseMileageAllowanceSearchOK handles this case with default header values.

successful operation
*/
type TravelExpenseMileageAllowanceSearchOK struct {
	Payload *models.ListResponseMileageAllowance
}

func (o *TravelExpenseMileageAllowanceSearchOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/mileageAllowance][%d] travelExpenseMileageAllowanceSearchOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseMileageAllowanceSearchOK) GetPayload() *models.ListResponseMileageAllowance {
	return o.Payload
}

func (o *TravelExpenseMileageAllowanceSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseMileageAllowance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
