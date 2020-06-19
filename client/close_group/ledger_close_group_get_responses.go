// Code generated by go-swagger; DO NOT EDIT.

package close_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// LedgerCloseGroupGetReader is a Reader for the LedgerCloseGroupGet structure.
type LedgerCloseGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerCloseGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerCloseGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLedgerCloseGroupGetOK creates a LedgerCloseGroupGetOK with default headers values
func NewLedgerCloseGroupGetOK() *LedgerCloseGroupGetOK {
	return &LedgerCloseGroupGetOK{}
}

/*LedgerCloseGroupGetOK handles this case with default header values.

successful operation
*/
type LedgerCloseGroupGetOK struct {
	Payload *models.ResponseWrapperCloseGroup
}

func (o *LedgerCloseGroupGetOK) Error() string {
	return fmt.Sprintf("[GET /ledger/closeGroup/{id}][%d] ledgerCloseGroupGetOK  %+v", 200, o.Payload)
}

func (o *LedgerCloseGroupGetOK) GetPayload() *models.ResponseWrapperCloseGroup {
	return o.Payload
}

func (o *LedgerCloseGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperCloseGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
