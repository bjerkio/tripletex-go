// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectPeriodMonthlyStatus project period monthly status
//
// swagger:model ProjectPeriodMonthlyStatus
type ProjectPeriodMonthlyStatus struct {

	// costs
	// Read Only: true
	Costs float64 `json:"costs,omitempty"`

	// date from
	// Read Only: true
	DateFrom string `json:"dateFrom,omitempty"`

	// date to
	// Read Only: true
	DateTo string `json:"dateTo,omitempty"`

	// income
	// Read Only: true
	Income float64 `json:"income,omitempty"`
}

// Validate validates this project period monthly status
func (m *ProjectPeriodMonthlyStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectPeriodMonthlyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectPeriodMonthlyStatus) UnmarshalBinary(b []byte) error {
	var res ProjectPeriodMonthlyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
