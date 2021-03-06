// Code generated by go-swagger; DO NOT EDIT.

package productline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// InventoryStocktakingProductlinePostReader is a Reader for the InventoryStocktakingProductlinePost structure.
type InventoryStocktakingProductlinePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryStocktakingProductlinePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoryStocktakingProductlinePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryStocktakingProductlinePostCreated creates a InventoryStocktakingProductlinePostCreated with default headers values
func NewInventoryStocktakingProductlinePostCreated() *InventoryStocktakingProductlinePostCreated {
	return &InventoryStocktakingProductlinePostCreated{}
}

/*InventoryStocktakingProductlinePostCreated handles this case with default header values.

successfully created
*/
type InventoryStocktakingProductlinePostCreated struct {
	Payload *models.ResponseWrapperProductLine
}

func (o *InventoryStocktakingProductlinePostCreated) Error() string {
	return fmt.Sprintf("[POST /inventory/stocktaking/productline][%d] inventoryStocktakingProductlinePostCreated  %+v", 201, o.Payload)
}

func (o *InventoryStocktakingProductlinePostCreated) GetPayload() *models.ResponseWrapperProductLine {
	return o.Payload
}

func (o *InventoryStocktakingProductlinePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProductLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
