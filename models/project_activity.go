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

// ProjectActivity project activity
//
// swagger:model ProjectActivity
type ProjectActivity struct {

	// Add existing project activity or create new project specific activity
	Activity *Activity `json:"activity,omitempty"`

	// Set budget fee
	BudgetFeeCurrency float64 `json:"budgetFeeCurrency,omitempty"`

	// Set budget hourly rate
	BudgetHourlyRateCurrency float64 `json:"budgetHourlyRateCurrency,omitempty"`

	// Set budget hours
	BudgetHours float64 `json:"budgetHours,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is closed
	IsClosed *bool `json:"isClosed,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this project activity
func (m *ProjectActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectActivity) validateActivity(formats strfmt.Registry) error {

	if swag.IsZero(m.Activity) { // not required
		return nil
	}

	if m.Activity != nil {
		if err := m.Activity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectActivity) validateChanges(formats strfmt.Registry) error {

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

func (m *ProjectActivity) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectActivity) UnmarshalBinary(b []byte) error {
	var res ProjectActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
