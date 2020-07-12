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

package leave_of_absence_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// EmployeeEmploymentLeaveOfAbsenceTypeSearchReader is a Reader for the EmployeeEmploymentLeaveOfAbsenceTypeSearch structure.
type EmployeeEmploymentLeaveOfAbsenceTypeSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmployeeEmploymentLeaveOfAbsenceTypeSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmployeeEmploymentLeaveOfAbsenceTypeSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmployeeEmploymentLeaveOfAbsenceTypeSearchOK creates a EmployeeEmploymentLeaveOfAbsenceTypeSearchOK with default headers values
func NewEmployeeEmploymentLeaveOfAbsenceTypeSearchOK() *EmployeeEmploymentLeaveOfAbsenceTypeSearchOK {
	return &EmployeeEmploymentLeaveOfAbsenceTypeSearchOK{}
}

/*EmployeeEmploymentLeaveOfAbsenceTypeSearchOK handles this case with default header values.

successful operation
*/
type EmployeeEmploymentLeaveOfAbsenceTypeSearchOK struct {
	Payload *models.ListResponseLeaveOfAbsenceType
}

func (o *EmployeeEmploymentLeaveOfAbsenceTypeSearchOK) Error() string {
	return fmt.Sprintf("[GET /employee/employment/leaveOfAbsenceType][%d] employeeEmploymentLeaveOfAbsenceTypeSearchOK  %+v", 200, o.Payload)
}

func (o *EmployeeEmploymentLeaveOfAbsenceTypeSearchOK) GetPayload() *models.ListResponseLeaveOfAbsenceType {
	return o.Payload
}

func (o *EmployeeEmploymentLeaveOfAbsenceTypeSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseLeaveOfAbsenceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
