// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EmployeeCategoryListDeleteByIdsReader is a Reader for the EmployeeCategoryListDeleteByIds structure.
type EmployeeCategoryListDeleteByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeCategoryListDeleteByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEmployeeCategoryListDeleteByIdsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEmployeeCategoryListDeleteByIdsDefault creates a EmployeeCategoryListDeleteByIdsDefault with default headers values
func NewEmployeeCategoryListDeleteByIdsDefault(code int) *EmployeeCategoryListDeleteByIdsDefault {
	return &EmployeeCategoryListDeleteByIdsDefault{
		_statusCode: code,
	}
}

/*EmployeeCategoryListDeleteByIdsDefault handles this case with default header values.

successful operation
*/
type EmployeeCategoryListDeleteByIdsDefault struct {
	_statusCode int
}

// Code gets the status code for the employee category list delete by ids default response
func (o *EmployeeCategoryListDeleteByIdsDefault) Code() int {
	return o._statusCode
}

func (o *EmployeeCategoryListDeleteByIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /employee/category/list][%d] EmployeeCategoryList_deleteByIds default ", o._statusCode)
}

func (o *EmployeeCategoryListDeleteByIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
