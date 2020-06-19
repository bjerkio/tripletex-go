// Code generated by go-swagger; DO NOT EDIT.

package salesmodules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// CompanySalesmodulesPostReader is a Reader for the CompanySalesmodulesPost structure.
type CompanySalesmodulesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompanySalesmodulesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCompanySalesmodulesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompanySalesmodulesPostCreated creates a CompanySalesmodulesPostCreated with default headers values
func NewCompanySalesmodulesPostCreated() *CompanySalesmodulesPostCreated {
	return &CompanySalesmodulesPostCreated{}
}

/*CompanySalesmodulesPostCreated handles this case with default header values.

successfully created
*/
type CompanySalesmodulesPostCreated struct {
	Payload *models.ResponseWrapperSalesModuleDTO
}

func (o *CompanySalesmodulesPostCreated) Error() string {
	return fmt.Sprintf("[POST /company/salesmodules][%d] companySalesmodulesPostCreated  %+v", 201, o.Payload)
}

func (o *CompanySalesmodulesPostCreated) GetPayload() *models.ResponseWrapperSalesModuleDTO {
	return o.Payload
}

func (o *CompanySalesmodulesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperSalesModuleDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
