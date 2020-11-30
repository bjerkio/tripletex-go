// Code generated by go-swagger; DO NOT EDIT.

package stocktaking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// InventoryStocktakingPostReader is a Reader for the InventoryStocktakingPost structure.
type InventoryStocktakingPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryStocktakingPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoryStocktakingPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryStocktakingPostCreated creates a InventoryStocktakingPostCreated with default headers values
func NewInventoryStocktakingPostCreated() *InventoryStocktakingPostCreated {
	return &InventoryStocktakingPostCreated{}
}

/*InventoryStocktakingPostCreated handles this case with default header values.

successfully created
*/
type InventoryStocktakingPostCreated struct {
	Payload *models.ResponseWrapperStocktaking
}

func (o *InventoryStocktakingPostCreated) Error() string {
	return fmt.Sprintf("[POST /inventory/stocktaking][%d] inventoryStocktakingPostCreated  %+v", 201, o.Payload)
}

func (o *InventoryStocktakingPostCreated) GetPayload() *models.ResponseWrapperStocktaking {
	return o.Payload
}

func (o *InventoryStocktakingPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperStocktaking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
