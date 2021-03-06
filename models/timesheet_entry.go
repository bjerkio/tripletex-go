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

// TimesheetEntry timesheet entry
//
// swagger:model TimesheetEntry
type TimesheetEntry struct {

	// activity
	// Required: true
	Activity *Activity `json:"activity"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// chargeable
	// Read Only: true
	Chargeable *bool `json:"chargeable,omitempty"`

	// chargeable hours
	// Read Only: true
	ChargeableHours float64 `json:"chargeableHours,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// date
	// Required: true
	Date *string `json:"date"`

	// employee
	// Required: true
	Employee *Employee `json:"employee"`

	// hourly cost
	// Read Only: true
	HourlyCost float64 `json:"hourlyCost,omitempty"`

	// hourly cost percentage
	// Read Only: true
	HourlyCostPercentage float64 `json:"hourlyCostPercentage,omitempty"`

	// hourly rate
	// Read Only: true
	HourlyRate float64 `json:"hourlyRate,omitempty"`

	// hours
	// Required: true
	Hours *float64 `json:"hours"`

	// id
	ID int32 `json:"id,omitempty"`

	// invoice
	// Read Only: true
	Invoice *Invoice `json:"invoice,omitempty"`

	// Indicates if the hour can be changed.
	// Read Only: true
	Locked *bool `json:"locked,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`

	// Link to stop watches on this hour.
	// Read Only: true
	TimeClocks []*TimeClock `json:"timeClocks"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this timesheet entry
func (m *TimesheetEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHours(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeClocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimesheetEntry) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", m.Activity); err != nil {
		return err
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

func (m *TimesheetEntry) validateChanges(formats strfmt.Registry) error {

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

func (m *TimesheetEntry) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *TimesheetEntry) validateEmployee(formats strfmt.Registry) error {

	if err := validate.Required("employee", "body", m.Employee); err != nil {
		return err
	}

	if m.Employee != nil {
		if err := m.Employee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee")
			}
			return err
		}
	}

	return nil
}

func (m *TimesheetEntry) validateHours(formats strfmt.Registry) error {

	if err := validate.Required("hours", "body", m.Hours); err != nil {
		return err
	}

	return nil
}

func (m *TimesheetEntry) validateInvoice(formats strfmt.Registry) error {

	if swag.IsZero(m.Invoice) { // not required
		return nil
	}

	if m.Invoice != nil {
		if err := m.Invoice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoice")
			}
			return err
		}
	}

	return nil
}

func (m *TimesheetEntry) validateProject(formats strfmt.Registry) error {

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

func (m *TimesheetEntry) validateTimeClocks(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeClocks) { // not required
		return nil
	}

	for i := 0; i < len(m.TimeClocks); i++ {
		if swag.IsZero(m.TimeClocks[i]) { // not required
			continue
		}

		if m.TimeClocks[i] != nil {
			if err := m.TimeClocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeClocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimesheetEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimesheetEntry) UnmarshalBinary(b []byte) error {
	var res TimesheetEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
