// Code generated by go-swagger; DO NOT EDIT.

package prospect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// CrmProspectSearchReader is a Reader for the CrmProspectSearch structure.
type CrmProspectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CrmProspectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCrmProspectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCrmProspectSearchOK creates a CrmProspectSearchOK with default headers values
func NewCrmProspectSearchOK() *CrmProspectSearchOK {
	return &CrmProspectSearchOK{}
}

/*CrmProspectSearchOK handles this case with default header values.

successful operation
*/
type CrmProspectSearchOK struct {
	Payload *models.ListResponseProspect
}

func (o *CrmProspectSearchOK) Error() string {
	return fmt.Sprintf("[GET /crm/prospect][%d] crmProspectSearchOK  %+v", 200, o.Payload)
}

func (o *CrmProspectSearchOK) GetPayload() *models.ListResponseProspect {
	return o.Payload
}

func (o *CrmProspectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProspect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
