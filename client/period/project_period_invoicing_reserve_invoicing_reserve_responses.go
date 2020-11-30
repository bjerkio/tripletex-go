// Code generated by go-swagger; DO NOT EDIT.

package period

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProjectPeriodInvoicingReserveInvoicingReserveReader is a Reader for the ProjectPeriodInvoicingReserveInvoicingReserve structure.
type ProjectPeriodInvoicingReserveInvoicingReserveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectPeriodInvoicingReserveInvoicingReserveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectPeriodInvoicingReserveInvoicingReserveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectPeriodInvoicingReserveInvoicingReserveOK creates a ProjectPeriodInvoicingReserveInvoicingReserveOK with default headers values
func NewProjectPeriodInvoicingReserveInvoicingReserveOK() *ProjectPeriodInvoicingReserveInvoicingReserveOK {
	return &ProjectPeriodInvoicingReserveInvoicingReserveOK{}
}

/*ProjectPeriodInvoicingReserveInvoicingReserveOK handles this case with default header values.

successful operation
*/
type ProjectPeriodInvoicingReserveInvoicingReserveOK struct {
	Payload *models.ResponseWrapperProjectPeriodInvoicingReserve
}

func (o *ProjectPeriodInvoicingReserveInvoicingReserveOK) Error() string {
	return fmt.Sprintf("[GET /project/{id}/period/invoicingReserve][%d] projectPeriodInvoicingReserveInvoicingReserveOK  %+v", 200, o.Payload)
}

func (o *ProjectPeriodInvoicingReserveInvoicingReserveOK) GetPayload() *models.ResponseWrapperProjectPeriodInvoicingReserve {
	return o.Payload
}

func (o *ProjectPeriodInvoicingReserveInvoicingReserveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperProjectPeriodInvoicingReserve)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
