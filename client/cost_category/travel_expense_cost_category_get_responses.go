// Code generated by go-swagger; DO NOT EDIT.

package cost_category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TravelExpenseCostCategoryGetReader is a Reader for the TravelExpenseCostCategoryGet structure.
type TravelExpenseCostCategoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelExpenseCostCategoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelExpenseCostCategoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTravelExpenseCostCategoryGetOK creates a TravelExpenseCostCategoryGetOK with default headers values
func NewTravelExpenseCostCategoryGetOK() *TravelExpenseCostCategoryGetOK {
	return &TravelExpenseCostCategoryGetOK{}
}

/*TravelExpenseCostCategoryGetOK handles this case with default header values.

successful operation
*/
type TravelExpenseCostCategoryGetOK struct {
	Payload *models.ResponseWrapperTravelCostCategory
}

func (o *TravelExpenseCostCategoryGetOK) Error() string {
	return fmt.Sprintf("[GET /travelExpense/costCategory/{id}][%d] travelExpenseCostCategoryGetOK  %+v", 200, o.Payload)
}

func (o *TravelExpenseCostCategoryGetOK) GetPayload() *models.ResponseWrapperTravelCostCategory {
	return o.Payload
}

func (o *TravelExpenseCostCategoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTravelCostCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
