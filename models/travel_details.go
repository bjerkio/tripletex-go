// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TravelDetails travel details
//
// swagger:model TravelDetails
type TravelDetails struct {

	// departure date
	DepartureDate string `json:"departureDate,omitempty"`

	// departure from
	// Max Length: 255
	DepartureFrom string `json:"departureFrom,omitempty"`

	// departure time
	// Max Length: 20
	DepartureTime string `json:"departureTime,omitempty"`

	// destination
	// Max Length: 255
	Destination string `json:"destination,omitempty"`

	// detailed journey description
	DetailedJourneyDescription string `json:"detailedJourneyDescription,omitempty"`

	// is compensation from rates
	IsCompensationFromRates *bool `json:"isCompensationFromRates,omitempty"`

	// is day trip
	IsDayTrip *bool `json:"isDayTrip,omitempty"`

	// is foreign travel
	IsForeignTravel *bool `json:"isForeignTravel,omitempty"`

	// purpose
	Purpose string `json:"purpose,omitempty"`

	// return date
	ReturnDate string `json:"returnDate,omitempty"`

	// return time
	// Max Length: 20
	ReturnTime string `json:"returnTime,omitempty"`
}

// Validate validates this travel details
func (m *TravelDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepartureFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TravelDetails) validateDepartureFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartureFrom) { // not required
		return nil
	}

	if err := validate.MaxLength("departureFrom", "body", string(m.DepartureFrom), 255); err != nil {
		return err
	}

	return nil
}

func (m *TravelDetails) validateDepartureTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartureTime) { // not required
		return nil
	}

	if err := validate.MaxLength("departureTime", "body", string(m.DepartureTime), 20); err != nil {
		return err
	}

	return nil
}

func (m *TravelDetails) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if err := validate.MaxLength("destination", "body", string(m.Destination), 255); err != nil {
		return err
	}

	return nil
}

func (m *TravelDetails) validateReturnTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnTime) { // not required
		return nil
	}

	if err := validate.MaxLength("returnTime", "body", string(m.ReturnTime), 20); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TravelDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TravelDetails) UnmarshalBinary(b []byte) error {
	var res TravelDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
