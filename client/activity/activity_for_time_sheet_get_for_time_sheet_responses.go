// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ActivityForTimeSheetGetForTimeSheetReader is a Reader for the ActivityForTimeSheetGetForTimeSheet structure.
type ActivityForTimeSheetGetForTimeSheetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivityForTimeSheetGetForTimeSheetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivityForTimeSheetGetForTimeSheetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActivityForTimeSheetGetForTimeSheetOK creates a ActivityForTimeSheetGetForTimeSheetOK with default headers values
func NewActivityForTimeSheetGetForTimeSheetOK() *ActivityForTimeSheetGetForTimeSheetOK {
	return &ActivityForTimeSheetGetForTimeSheetOK{}
}

/*ActivityForTimeSheetGetForTimeSheetOK handles this case with default header values.

successful operation
*/
type ActivityForTimeSheetGetForTimeSheetOK struct {
	Payload *models.ListResponseActivity
}

func (o *ActivityForTimeSheetGetForTimeSheetOK) Error() string {
	return fmt.Sprintf("[GET /activity/>forTimeSheet][%d] activityForTimeSheetGetForTimeSheetOK  %+v", 200, o.Payload)
}

func (o *ActivityForTimeSheetGetForTimeSheetOK) GetPayload() *models.ListResponseActivity {
	return o.Payload
}

func (o *ActivityForTimeSheetGetForTimeSheetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
