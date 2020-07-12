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

package voucher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LedgerVoucherAttachmentUploadAttachmentReader is a Reader for the LedgerVoucherAttachmentUploadAttachment structure.
type LedgerVoucherAttachmentUploadAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerVoucherAttachmentUploadAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLedgerVoucherAttachmentUploadAttachmentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLedgerVoucherAttachmentUploadAttachmentDefault creates a LedgerVoucherAttachmentUploadAttachmentDefault with default headers values
func NewLedgerVoucherAttachmentUploadAttachmentDefault(code int) *LedgerVoucherAttachmentUploadAttachmentDefault {
	return &LedgerVoucherAttachmentUploadAttachmentDefault{
		_statusCode: code,
	}
}

/*LedgerVoucherAttachmentUploadAttachmentDefault handles this case with default header values.

successful operation
*/
type LedgerVoucherAttachmentUploadAttachmentDefault struct {
	_statusCode int
}

// Code gets the status code for the ledger voucher attachment upload attachment default response
func (o *LedgerVoucherAttachmentUploadAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *LedgerVoucherAttachmentUploadAttachmentDefault) Error() string {
	return fmt.Sprintf("[POST /ledger/voucher/{voucherId}/attachment][%d] LedgerVoucherAttachmentUploadAttachment default ", o._statusCode)
}

func (o *LedgerVoucherAttachmentUploadAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
