// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/v2/models"
)

// CustomerSearchReader is a Reader for the CustomerSearch structure.
type CustomerSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomerSearchOK creates a CustomerSearchOK with default headers values
func NewCustomerSearchOK() *CustomerSearchOK {
	return &CustomerSearchOK{}
}

/*CustomerSearchOK handles this case with default header values.

successful operation
*/
type CustomerSearchOK struct {
	Payload *models.ListResponseCustomer
}

func (o *CustomerSearchOK) Error() string {
	return fmt.Sprintf("[GET /customer][%d] customerSearchOK  %+v", 200, o.Payload)
}

func (o *CustomerSearchOK) GetPayload() *models.ListResponseCustomer {
	return o.Payload
}

func (o *CustomerSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseCustomer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
