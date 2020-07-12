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

package project_specific_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/tripletex-go/models"
)

// ProjectHourlyRatesProjectSpecificRatesSearchReader is a Reader for the ProjectHourlyRatesProjectSpecificRatesSearch structure.
type ProjectHourlyRatesProjectSpecificRatesSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectHourlyRatesProjectSpecificRatesSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectHourlyRatesProjectSpecificRatesSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectHourlyRatesProjectSpecificRatesSearchOK creates a ProjectHourlyRatesProjectSpecificRatesSearchOK with default headers values
func NewProjectHourlyRatesProjectSpecificRatesSearchOK() *ProjectHourlyRatesProjectSpecificRatesSearchOK {
	return &ProjectHourlyRatesProjectSpecificRatesSearchOK{}
}

/*ProjectHourlyRatesProjectSpecificRatesSearchOK handles this case with default header values.

successful operation
*/
type ProjectHourlyRatesProjectSpecificRatesSearchOK struct {
	Payload *models.ListResponseProjectSpecificRate
}

func (o *ProjectHourlyRatesProjectSpecificRatesSearchOK) Error() string {
	return fmt.Sprintf("[GET /project/hourlyRates/projectSpecificRates][%d] projectHourlyRatesProjectSpecificRatesSearchOK  %+v", 200, o.Payload)
}

func (o *ProjectHourlyRatesProjectSpecificRatesSearchOK) GetPayload() *models.ListResponseProjectSpecificRate {
	return o.Payload
}

func (o *ProjectHourlyRatesProjectSpecificRatesSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListResponseProjectSpecificRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
