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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MileageAllowance mileage allowance
//
// swagger:model MileageAllowance
type MileageAllowance struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// date
	// Required: true
	Date *string `json:"date"`

	// departure location
	// Required: true
	DepartureLocation *string `json:"departureLocation"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// id
	ID int32 `json:"id,omitempty"`

	// is company car
	IsCompanyCar *bool `json:"isCompanyCar,omitempty"`

	// km
	Km float64 `json:"km,omitempty"`

	// Passenger mileage allowance associated with this mileage allowance.
	PassengerSupplement *MileageAllowance `json:"passengerSupplement,omitempty"`

	// Link to individual passengers.
	// Read Only: true
	Passengers []*Passenger `json:"passengers"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// rate category
	RateCategory *TravelExpenseRateCategory `json:"rateCategory,omitempty"`

	// rate type
	RateType *TravelExpenseRate `json:"rateType,omitempty"`

	// Toll cost associated with this mileage allowance.
	TollCost *Cost `json:"tollCost,omitempty"`

	// travel expense
	TravelExpense *TravelExpense `json:"travelExpense,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this mileage allowance
func (m *MileageAllowance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassengerSupplement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassengers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTollCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTravelExpense(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MileageAllowance) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MileageAllowance) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *MileageAllowance) validateDepartureLocation(formats strfmt.Registry) error {

	if err := validate.Required("departureLocation", "body", m.DepartureLocation); err != nil {
		return err
	}

	return nil
}

func (m *MileageAllowance) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *MileageAllowance) validatePassengerSupplement(formats strfmt.Registry) error {

	if swag.IsZero(m.PassengerSupplement) { // not required
		return nil
	}

	if m.PassengerSupplement != nil {
		if err := m.PassengerSupplement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passengerSupplement")
			}
			return err
		}
	}

	return nil
}

func (m *MileageAllowance) validatePassengers(formats strfmt.Registry) error {

	if swag.IsZero(m.Passengers) { // not required
		return nil
	}

	for i := 0; i < len(m.Passengers); i++ {
		if swag.IsZero(m.Passengers[i]) { // not required
			continue
		}

		if m.Passengers[i] != nil {
			if err := m.Passengers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("passengers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MileageAllowance) validateRateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.RateCategory) { // not required
		return nil
	}

	if m.RateCategory != nil {
		if err := m.RateCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateCategory")
			}
			return err
		}
	}

	return nil
}

func (m *MileageAllowance) validateRateType(formats strfmt.Registry) error {

	if swag.IsZero(m.RateType) { // not required
		return nil
	}

	if m.RateType != nil {
		if err := m.RateType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateType")
			}
			return err
		}
	}

	return nil
}

func (m *MileageAllowance) validateTollCost(formats strfmt.Registry) error {

	if swag.IsZero(m.TollCost) { // not required
		return nil
	}

	if m.TollCost != nil {
		if err := m.TollCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tollCost")
			}
			return err
		}
	}

	return nil
}

func (m *MileageAllowance) validateTravelExpense(formats strfmt.Registry) error {

	if swag.IsZero(m.TravelExpense) { // not required
		return nil
	}

	if m.TravelExpense != nil {
		if err := m.TravelExpense.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("travelExpense")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MileageAllowance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MileageAllowance) UnmarshalBinary(b []byte) error {
	var res MileageAllowance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
