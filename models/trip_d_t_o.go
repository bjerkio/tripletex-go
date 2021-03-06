// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TripDTO trip d t o
//
// swagger:model TripDTO
type TripDTO struct {

	// distance
	Distance float64 `json:"distance,omitempty"`

	// sum toll n o k
	SumTollNOK float64 `json:"sumTollNOK,omitempty"`

	// travel time in minutes
	TravelTimeInMinutes float64 `json:"travelTimeInMinutes,omitempty"`
}

// Validate validates this trip d t o
func (m *TripDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TripDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripDTO) UnmarshalBinary(b []byte) error {
	var res TripDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
