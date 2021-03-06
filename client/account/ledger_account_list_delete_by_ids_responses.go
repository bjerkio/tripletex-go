// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LedgerAccountListDeleteByIdsReader is a Reader for the LedgerAccountListDeleteByIds structure.
type LedgerAccountListDeleteByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LedgerAccountListDeleteByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLedgerAccountListDeleteByIdsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLedgerAccountListDeleteByIdsDefault creates a LedgerAccountListDeleteByIdsDefault with default headers values
func NewLedgerAccountListDeleteByIdsDefault(code int) *LedgerAccountListDeleteByIdsDefault {
	return &LedgerAccountListDeleteByIdsDefault{
		_statusCode: code,
	}
}

/*LedgerAccountListDeleteByIdsDefault handles this case with default header values.

successful operation
*/
type LedgerAccountListDeleteByIdsDefault struct {
	_statusCode int
}

// Code gets the status code for the ledger account list delete by ids default response
func (o *LedgerAccountListDeleteByIdsDefault) Code() int {
	return o._statusCode
}

func (o *LedgerAccountListDeleteByIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /ledger/account/list][%d] LedgerAccountList_deleteByIds default ", o._statusCode)
}

func (o *LedgerAccountListDeleteByIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
