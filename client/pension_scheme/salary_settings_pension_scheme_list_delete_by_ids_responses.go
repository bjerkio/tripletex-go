// Code generated by go-swagger; DO NOT EDIT.

package pension_scheme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SalarySettingsPensionSchemeListDeleteByIdsReader is a Reader for the SalarySettingsPensionSchemeListDeleteByIds structure.
type SalarySettingsPensionSchemeListDeleteByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalarySettingsPensionSchemeListDeleteByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSalarySettingsPensionSchemeListDeleteByIdsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSalarySettingsPensionSchemeListDeleteByIdsDefault creates a SalarySettingsPensionSchemeListDeleteByIdsDefault with default headers values
func NewSalarySettingsPensionSchemeListDeleteByIdsDefault(code int) *SalarySettingsPensionSchemeListDeleteByIdsDefault {
	return &SalarySettingsPensionSchemeListDeleteByIdsDefault{
		_statusCode: code,
	}
}

/*SalarySettingsPensionSchemeListDeleteByIdsDefault handles this case with default header values.

successful operation
*/
type SalarySettingsPensionSchemeListDeleteByIdsDefault struct {
	_statusCode int
}

// Code gets the status code for the salary settings pension scheme list delete by ids default response
func (o *SalarySettingsPensionSchemeListDeleteByIdsDefault) Code() int {
	return o._statusCode
}

func (o *SalarySettingsPensionSchemeListDeleteByIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /salary/settings/pensionScheme/list][%d] SalarySettingsPensionSchemeList_deleteByIds default ", o._statusCode)
}

func (o *SalarySettingsPensionSchemeListDeleteByIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
