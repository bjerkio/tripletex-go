// Code generated by go-swagger; DO NOT EDIT.

package voucher_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// VoucherMessageSearchReader is a Reader for the VoucherMessageSearch structure.
type VoucherMessageSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoucherMessageSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVoucherMessageSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVoucherMessageSearchOK creates a VoucherMessageSearchOK with default headers values
func NewVoucherMessageSearchOK() *VoucherMessageSearchOK {
	return &VoucherMessageSearchOK{}
}

/*VoucherMessageSearchOK handles this case with default header values.

successful operation
*/
type VoucherMessageSearchOK struct {
	Payload *models.ListResponseVoucherMessage
}

func (o *VoucherMessageSearchOK) Error() string {
	return fmt.Sprintf("[GET /voucherMessage][%d] voucherMessageSearchOK  %+v", 200, o.Payload)
}

func (o *VoucherMessageSearchOK) GetPayload() *models.ListResponseVoucherMessage {
	return o.Payload
}

func (o *VoucherMessageSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseVoucherMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
