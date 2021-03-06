// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TravelExpenseSettings travel expense settings
//
// swagger:model TravelExpenseSettings
type TravelExpenseSettings struct {

	// accommodation not compensated
	AccommodationNotCompensated *bool `json:"accommodationNotCompensated,omitempty"`

	// approval required
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// id
	ID int32 `json:"id,omitempty"`

	// mileage not compensated
	MileageNotCompensated *bool `json:"mileageNotCompensated,omitempty"`

	// per diem not compensated
	PerDiemNotCompensated *bool `json:"perDiemNotCompensated,omitempty"`

	// tax free mileage rates
	TaxFreeMileageRates *bool `json:"taxFreeMileageRates,omitempty"`

	// tax free per diem rates
	TaxFreePerDiemRates *bool `json:"taxFreePerDiemRates,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// use rates
	UseRates *bool `json:"useRates,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this travel expense settings
func (m *TravelExpenseSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TravelExpenseSettings) validateChanges(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *TravelExpenseSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TravelExpenseSettings) UnmarshalBinary(b []byte) error {
	var res TravelExpenseSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
