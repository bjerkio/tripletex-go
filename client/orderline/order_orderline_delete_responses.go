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

package orderline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrderOrderlineDeleteReader is a Reader for the OrderOrderlineDelete structure.
type OrderOrderlineDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderOrderlineDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewOrderOrderlineDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewOrderOrderlineDeleteDefault creates a OrderOrderlineDeleteDefault with default headers values
func NewOrderOrderlineDeleteDefault(code int) *OrderOrderlineDeleteDefault {
	return &OrderOrderlineDeleteDefault{
		_statusCode: code,
	}
}

/*OrderOrderlineDeleteDefault handles this case with default header values.

successful operation
*/
type OrderOrderlineDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the order orderline delete default response
func (o *OrderOrderlineDeleteDefault) Code() int {
	return o._statusCode
}

func (o *OrderOrderlineDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /order/orderline/{id}][%d] OrderOrderlineDelete default ", o._statusCode)
}

func (o *OrderOrderlineDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
