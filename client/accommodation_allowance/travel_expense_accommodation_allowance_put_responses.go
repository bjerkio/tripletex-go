// Code generated by go-swagger; DO NOT EDIT.

package accommodation_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpenseAccommodationAllowancePutReader is a Reader for the TravelExpenseAccommodationAllowancePut structure.
type TravelExpenseAccommodationAllowancePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseAccommodationAllowancePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseAccommodationAllowancePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseAccommodationAllowancePutOK creates a TravelExpenseAccommodationAllowancePutOK with default headers values
func NewTravelExpenseAccommodationAllowancePutOK() *TravelExpenseAccommodationAllowancePutOK {
	return &TravelExpenseAccommodationAllowancePutOK{}
}

/*TravelExpenseAccommodationAllowancePutOK handles this case with default header values.

successful operation
*/
type TravelExpenseAccommodationAllowancePutOK struct {
	Payload *models.ResponseWrapperAccommodationAllowance
}

func (o *TravelExpenseAccommodationAllowancePutOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/accommodationAllowance/{id}][%d] travelExpenseAccommodationAllowancePutOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseAccommodationAllowancePutOK) GetPayload() *models.ResponseWrapperAccommodationAllowance {
	return o.Payload
}

func (o *TravelExpenseAccommodationAllowancePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperAccommodationAllowance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
