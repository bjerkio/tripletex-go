// Code generated by go-swagger; DO NOT EDIT.

package unit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProductUnitGetReader is a Reader for the ProductUnitGet structure.
type ProductUnitGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductUnitGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductUnitGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProductUnitGetOK creates a ProductUnitGetOK with default headers values
func NewProductUnitGetOK() *ProductUnitGetOK {
	return &ProductUnitGetOK{}
}

/*ProductUnitGetOK handles this case with default header values.

successful operation
*/
type ProductUnitGetOK struct {
	Payload *models.ResponseWrapperProductUnit
}

func (o *ProductUnitGetOK) Error() string {
	return fmt.Sprintf("[GET /product/unit/{id}][%d] productUnitGetOK  %+v", 200, o.Payload)
}

func (o *ProductUnitGetOK) GetPayload() *models.ResponseWrapperProductUnit {
	return o.Payload
}

func (o *ProductUnitGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProductUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
