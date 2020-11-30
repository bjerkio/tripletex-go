// Code generated by go-swagger; DO NOT EDIT.

package supplier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// SupplierListPutListReader is a Reader for the SupplierListPutList structure.
type SupplierListPutListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplierListPutListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplierListPutListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSupplierListPutListOK creates a SupplierListPutListOK with default headers values
func NewSupplierListPutListOK() *SupplierListPutListOK {
	return &SupplierListPutListOK{}
}

/*SupplierListPutListOK handles this case with default header values.

successful operation
*/
type SupplierListPutListOK struct {
	Payload *models.ListResponseSupplier
}

func (o *SupplierListPutListOK) Error() string {
	return fmt.Sprintf("[PUT /supplier/list][%d] supplierListPutListOK  %+v", 200, o.Payload)
}

func (o *SupplierListPutListOK) GetPayload() *models.ListResponseSupplier {
	return o.Payload
}

func (o *SupplierListPutListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseSupplier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
