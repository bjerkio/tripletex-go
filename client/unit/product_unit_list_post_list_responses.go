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

// ProductUnitListPostListReader is a Reader for the ProductUnitListPostList structure.
type ProductUnitListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductUnitListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProductUnitListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductUnitListPostListCreated creates a ProductUnitListPostListCreated with default headers values
func NewProductUnitListPostListCreated() *ProductUnitListPostListCreated {
	return &ProductUnitListPostListCreated{}
}

/*ProductUnitListPostListCreated handles this case with default header values.

successfully created
*/
type ProductUnitListPostListCreated struct {
	Payload *models.ListResponseProductUnit
}

func (o *ProductUnitListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /product/unit/list][%d] productUnitListPostListCreated  %+v", 201, o.Payload)
}

func (o *ProductUnitListPostListCreated) GetPayload() *models.ListResponseProductUnit {
	return o.Payload
}

func (o *ProductUnitListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProductUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
