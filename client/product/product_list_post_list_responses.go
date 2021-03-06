// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// ProductListPostListReader is a Reader for the ProductListPostList structure.
type ProductListPostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductListPostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProductListPostListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductListPostListCreated creates a ProductListPostListCreated with default headers values
func NewProductListPostListCreated() *ProductListPostListCreated {
	return &ProductListPostListCreated{}
}

/*ProductListPostListCreated handles this case with default header values.

successfully created
*/
type ProductListPostListCreated struct {
	Payload *models.ListResponseProduct
}

func (o *ProductListPostListCreated) Error() string {
	return fmt.Sprintf("[POST /product/list][%d] productListPostListCreated  %+v", 201, o.Payload)
}

func (o *ProductListPostListCreated) GetPayload() *models.ListResponseProduct {
	return o.Payload
}

func (o *ProductListPostListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProduct)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
