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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmployeePeriod employee period
//
// swagger:model EmployeePeriod
type EmployeePeriod struct {

	// incoming vacation balance
	// Read Only: true
	IncomingVacationBalance float64 `json:"incomingVacationBalance,omitempty"`

	// outgoing vacation balance
	// Read Only: true
	OutgoingVacationBalance float64 `json:"outgoingVacationBalance,omitempty"`

	// vacation taken in period
	// Read Only: true
	VacationTakenInPeriod float64 `json:"vacationTakenInPeriod,omitempty"`

	// vacation taken this year
	// Read Only: true
	VacationTakenThisYear float64 `json:"vacationTakenThisYear,omitempty"`
}

// Validate validates this employee period
func (m *EmployeePeriod) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmployeePeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeePeriod) UnmarshalBinary(b []byte) error {
	var res EmployeePeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
