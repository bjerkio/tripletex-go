// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EmployeeCategoryDeleteReader is a Reader for the EmployeeCategoryDelete structure.
type EmployeeCategoryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeCategoryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEmployeeCategoryDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEmployeeCategoryDeleteDefault creates a EmployeeCategoryDeleteDefault with default headers values
func NewEmployeeCategoryDeleteDefault(code int) *EmployeeCategoryDeleteDefault {
	return &EmployeeCategoryDeleteDefault{
		_statusCode: code,
	}
}

/*EmployeeCategoryDeleteDefault handles this case with default header values.

successful operation
*/
type EmployeeCategoryDeleteDefault struct {
	_statusCode int
}

// Code gets the status code for the employee category delete default response
func (o *EmployeeCategoryDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EmployeeCategoryDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /employee/category/{id}][%d] EmployeeCategory_delete default ", o._statusCode)
}

func (o *EmployeeCategoryDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
