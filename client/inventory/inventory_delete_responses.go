// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryDeleteReader is a Reader for the InventoryDelete structure.
type InventoryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewInventoryDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewInventoryDeleteDefault creates a InventoryDeleteDefault with default headers values
func NewInventoryDeleteDefault(code int) *InventoryDeleteDefault {
	return &InventoryDeleteDefault{
		_statusCode: code,
	}
}

/*InventoryDeleteDefault handles this case with default header values.

successful operation
*/
type InventoryDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the inventory delete default response
func (o *InventoryDeleteDefault) Code() int {
	return o._statusCode
}

func (o *InventoryDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /inventory/{id}][%d] Inventory_delete default ", o._statusCode)
}

func (o *InventoryDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
