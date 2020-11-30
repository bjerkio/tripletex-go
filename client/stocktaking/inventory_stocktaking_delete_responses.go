// Code generated by go-swagger; DO NOT EDIT.

package stocktaking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryStocktakingDeleteReader is a Reader for the InventoryStocktakingDelete structure.
type InventoryStocktakingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryStocktakingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewInventoryStocktakingDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewInventoryStocktakingDeleteDefault creates a InventoryStocktakingDeleteDefault with default headers values
func NewInventoryStocktakingDeleteDefault(code int) *InventoryStocktakingDeleteDefault {
	return &InventoryStocktakingDeleteDefault{
		_statusCode: code,
	}
}

/*InventoryStocktakingDeleteDefault handles this case with default header values.

successful operation
*/
type InventoryStocktakingDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the inventory stocktaking delete default response
func (o *InventoryStocktakingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *InventoryStocktakingDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /inventory/stocktaking/{id}][%d] InventoryStocktaking_delete default ", o._statusCode)
}

func (o *InventoryStocktakingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
