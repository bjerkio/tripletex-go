// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpenseCreateVouchersCreateVouchersReader is a Reader for the TravelExpenseCreateVouchersCreateVouchers structure.
type TravelExpenseCreateVouchersCreateVouchersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseCreateVouchersCreateVouchersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseCreateVouchersCreateVouchersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpenseCreateVouchersCreateVouchersOK creates a TravelExpenseCreateVouchersCreateVouchersOK with default headers values
func NewTravelExpenseCreateVouchersCreateVouchersOK() *TravelExpenseCreateVouchersCreateVouchersOK {
	return &TravelExpenseCreateVouchersCreateVouchersOK{}
}

/*TravelExpenseCreateVouchersCreateVouchersOK handles this case with default header values.

successful operation
*/
type TravelExpenseCreateVouchersCreateVouchersOK struct {
	Payload *models.ListResponseTravelExpense
}

func (o *TravelExpenseCreateVouchersCreateVouchersOK) Error() string {
	return fmt.Sprintf("[PUT /travelExpense/:createVouchers][%d] travelExpenseCreateVouchersCreateVouchersOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseCreateVouchersCreateVouchersOK) GetPayload() *models.ListResponseTravelExpense {
	return o.Payload
}

func (o *TravelExpenseCreateVouchersCreateVouchersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseTravelExpense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
