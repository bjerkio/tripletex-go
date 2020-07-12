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

package salary_type_specification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// TimesheetSalaryTypeSpecificationPostReader is a Reader for the TimesheetSalaryTypeSpecificationPost structure.
type TimesheetSalaryTypeSpecificationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetSalaryTypeSpecificationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTimesheetSalaryTypeSpecificationPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimesheetSalaryTypeSpecificationPostCreated creates a TimesheetSalaryTypeSpecificationPostCreated with default headers values
func NewTimesheetSalaryTypeSpecificationPostCreated() *TimesheetSalaryTypeSpecificationPostCreated {
	return &TimesheetSalaryTypeSpecificationPostCreated{}
}

/*TimesheetSalaryTypeSpecificationPostCreated handles this case with default header values.

successfully created
*/
type TimesheetSalaryTypeSpecificationPostCreated struct {
	Payload *models.ResponseWrapperTimesheetSalaryTypeSpecification
}

func (o *TimesheetSalaryTypeSpecificationPostCreated) Error() string {
	return fmt.Sprintf("[POST /timesheet/salaryTypeSpecification][%d] timesheetSalaryTypeSpecificationPostCreated  %+v", 201, o.Payload)
}

func (o *TimesheetSalaryTypeSpecificationPostCreated) GetPayload() *models.ResponseWrapperTimesheetSalaryTypeSpecification {
	return o.Payload
}

func (o *TimesheetSalaryTypeSpecificationPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperTimesheetSalaryTypeSpecification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
