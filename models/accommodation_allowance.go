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

// AccommodationAllowance accommodation allowance
//
// swagger:model AccommodationAllowance
type AccommodationAllowance struct {

	// address
	Address string `json:"address,omitempty"`

	// amount
	Amount float64 `json:"amount,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// count
	Count int32 `json:"count,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// location
	// Required: true
	// Max Length: 255
	Location *string `json:"location"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// rate category
	RateCategory *TravelExpenseRateCategory `json:"rateCategory,omitempty"`

	// rate type
	RateType *TravelExpenseRate `json:"rateType,omitempty"`

	// travel expense
	TravelExpense *TravelExpense `json:"travelExpense,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this accommodation allowance
func (m *AccommodationAllowance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateType(formats); err != nil {
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

func (m *AccommodationAllowance) validateChanges(formats strfmt.Registry) error {

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

func (m *AccommodationAllowance) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if err := validate.MaxLength("location", "body", string(*m.Location), 255); err != nil {
		return err
	}

	return nil
}

func (m *AccommodationAllowance) validateRateCategory(formats strfmt.Registry) error {

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

func (m *AccommodationAllowance) validateRateType(formats strfmt.Registry) error {

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

func (m *AccommodationAllowance) validateTravelExpense(formats strfmt.Registry) error {

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
func (m *AccommodationAllowance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccommodationAllowance) UnmarshalBinary(b []byte) error {
	var res AccommodationAllowance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
