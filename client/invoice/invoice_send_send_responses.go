// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InvoiceSendSendReader is a Reader for the InvoiceSendSend structure.
type InvoiceSendSendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoiceSendSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewInvoiceSendSendDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewInvoiceSendSendDefault creates a InvoiceSendSendDefault with default headers values
func NewInvoiceSendSendDefault(code int) *InvoiceSendSendDefault {
	return &InvoiceSendSendDefault{
		_statusCode: code,
	}
}

/*InvoiceSendSendDefault handles this case with default header values.

successful operation
*/
type InvoiceSendSendDefault struct {
	_statusCode int
}

// Code gets the status code for the invoice send send default response
func (o *InvoiceSendSendDefault) Code() int {
	return o._statusCode
}

func (o *InvoiceSendSendDefault) Error() string {
	return fmt.Sprintf("[PUT /invoice/{id}/:send][%d] InvoiceSendSend default ", o._statusCode)
}

func (o *InvoiceSendSendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
