// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// CustomerCategoryPutReader is a Reader for the CustomerCategoryPut structure.
type CustomerCategoryPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCategoryPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerCategoryPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomerCategoryPutOK creates a CustomerCategoryPutOK with default headers values
func NewCustomerCategoryPutOK() *CustomerCategoryPutOK {
	return &CustomerCategoryPutOK{}
}

/*CustomerCategoryPutOK handles this case with default header values.

successful operation
*/
type CustomerCategoryPutOK struct {
	Payload *models.ResponseWrapperCustomerCategory
}

func (o *CustomerCategoryPutOK) Error() string {
	return fmt.Sprintf("[PUT /customer/category/{id}][%d] customerCategoryPutOK  %+v", 200, o.Payload)
}

func (o *CustomerCategoryPutOK) GetPayload() *models.ResponseWrapperCustomerCategory {
	return o.Payload
}

func (o *CustomerCategoryPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCustomerCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
