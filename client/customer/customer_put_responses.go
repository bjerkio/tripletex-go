// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// CustomerPutReader is a Reader for the CustomerPut structure.
type CustomerPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomerPutOK creates a CustomerPutOK with default headers values
func NewCustomerPutOK() *CustomerPutOK {
	return &CustomerPutOK{}
}

/*CustomerPutOK handles this case with default header values.

successful operation
*/
type CustomerPutOK struct {
	Payload *models.ResponseWrapperCustomer
}

func (o *CustomerPutOK) Error() string {
	return fmt.Sprintf("[PUT /customer/{id}][%d] customerPutOK  %+v", 200, o.Payload)
}

func (o *CustomerPutOK) GetPayload() *models.ResponseWrapperCustomer {
	return o.Payload
}

func (o *CustomerPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCustomer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
