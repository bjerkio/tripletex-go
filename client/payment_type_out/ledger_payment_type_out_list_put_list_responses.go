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

package payment_type_out

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// LedgerPaymentTypeOutListPutListReader is a Reader for the LedgerPaymentTypeOutListPutList structure.
type LedgerPaymentTypeOutListPutListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerPaymentTypeOutListPutListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLedgerPaymentTypeOutListPutListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLedgerPaymentTypeOutListPutListOK creates a LedgerPaymentTypeOutListPutListOK with default headers values
func NewLedgerPaymentTypeOutListPutListOK() *LedgerPaymentTypeOutListPutListOK {
	return &LedgerPaymentTypeOutListPutListOK{}
}

/*LedgerPaymentTypeOutListPutListOK handles this case with default header values.

successful operation
*/
type LedgerPaymentTypeOutListPutListOK struct {
	Payload *models.ListResponsePaymentTypeOut
}

func (o *LedgerPaymentTypeOutListPutListOK) Error() string {
	return fmt.Sprintf("[PUT /ledger/paymentTypeOut/list][%d] ledgerPaymentTypeOutListPutListOK  %+v", 200, o.Payload)
}

func (o *LedgerPaymentTypeOutListPutListOK) GetPayload() *models.ListResponsePaymentTypeOut {
	return o.Payload
}

func (o *LedgerPaymentTypeOutListPutListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponsePaymentTypeOut)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
