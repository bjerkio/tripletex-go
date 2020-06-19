// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// CustomerCategoryGetReader is a Reader for the CustomerCategoryGet structure.
type CustomerCategoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCategoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerCategoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerCategoryGetOK creates a CustomerCategoryGetOK with default headers values
func NewCustomerCategoryGetOK() *CustomerCategoryGetOK {
	return &CustomerCategoryGetOK{}
}

/*CustomerCategoryGetOK handles this case with default header values.

successful operation
*/
type CustomerCategoryGetOK struct {
	Payload *models.ResponseWrapperCustomerCategory
}

func (o *CustomerCategoryGetOK) Error() string {
	return fmt.Sprintf("[GET /customer/category/{id}][%d] customerCategoryGetOK  %+v", 200, o.Payload)
}

func (o *CustomerCategoryGetOK) GetPayload() *models.ResponseWrapperCustomerCategory {
	return o.Payload
}

func (o *CustomerCategoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCustomerCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
