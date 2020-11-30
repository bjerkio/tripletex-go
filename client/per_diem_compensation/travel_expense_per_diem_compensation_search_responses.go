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

// TravelExpensePerDiemCompensationSearchReader is a Reader for the TravelExpensePerDiemCompensationSearch structure.
type TravelExpensePerDiemCompensationSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpensePerDiemCompensationSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpensePerDiemCompensationSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTravelExpensePerDiemCompensationSearchOK creates a TravelExpensePerDiemCompensationSearchOK with default headers values
func NewTravelExpensePerDiemCompensationSearchOK() *TravelExpensePerDiemCompensationSearchOK {
	return &TravelExpensePerDiemCompensationSearchOK{}
}

/*TravelExpensePerDiemCompensationSearchOK handles this case with default header values.

successful operation
*/
type TravelExpensePerDiemCompensationSearchOK struct {
	Payload *models.ListResponsePerDiemCompensation
}

func (o *TravelExpensePerDiemCompensationSearchOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/perDiemCompensation][%d] travelExpensePerDiemCompensationSearchOK  %+v", 200, o.Payload)
}

func (o *TravelExpensePerDiemCompensationSearchOK) GetPayload() *models.ListResponsePerDiemCompensation {
	return o.Payload
}

func (o *TravelExpensePerDiemCompensationSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePerDiemCompensation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
