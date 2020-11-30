// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportGroup report group
//
// swagger:model ReportGroup
type ReportGroup struct {

	// Subgroups that should be automatically generated
	// Required: true
	// Enum: [None Account Department PeriodMonths]
	AutoGroupType *string `json:"autoGroupType"`

	// Format string for cell (indentation, font size etc)
	// Max Length: 255
	CellFormat string `json:"cellFormat,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// Child groups
	Children []*ReportGroup `json:"children"`

	// Currently not in use
	Description string `json:"description,omitempty"`

	// Expression / formula according to Rule Engine Expression Language
	Expression string `json:"expression,omitempty"`

	// filter
	Filter *ReportGroupFilter `json:"filter,omitempty"`

	// Hide this group?
	HideSelf *bool `json:"hideSelf,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// The name to be shown for the column or row.
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// Used to select expression if both column and row expression is set. The bigger value wins.
	Precedence int32 `json:"precedence,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// Format string for value (how to print number, date etc)
	// Max Length: 100
	ValueFormat string `json:"valueFormat,omitempty"`

	// Variable name that can be used to reference this group
	// Max Length: 100
	VariableName string `json:"variableName,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this report group
func (m *ReportGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoGroupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCellFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportGroupTypeAutoGroupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Account","Department","PeriodMonths"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportGroupTypeAutoGroupTypePropEnum = append(reportGroupTypeAutoGroupTypePropEnum, v)
	}
}

const (

	// ReportGroupAutoGroupTypeNone captures enum value "None"
	ReportGroupAutoGroupTypeNone string = "None"

	// ReportGroupAutoGroupTypeAccount captures enum value "Account"
	ReportGroupAutoGroupTypeAccount string = "Account"

	// ReportGroupAutoGroupTypeDepartment captures enum value "Department"
	ReportGroupAutoGroupTypeDepartment string = "Department"

	// ReportGroupAutoGroupTypePeriodMonths captures enum value "PeriodMonths"
	ReportGroupAutoGroupTypePeriodMonths string = "PeriodMonths"
)

// prop value enum
func (m *ReportGroup) validateAutoGroupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportGroupTypeAutoGroupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportGroup) validateAutoGroupType(formats strfmt.Registry) error {

	if err := validate.Required("autoGroupType", "body", m.AutoGroupType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAutoGroupTypeEnum("autoGroupType", "body", *m.AutoGroupType); err != nil {
		return err
	}

	return nil
}

func (m *ReportGroup) validateCellFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.CellFormat) { // not required
		return nil
	}

	if err := validate.MaxLength("cellFormat", "body", string(m.CellFormat), 255); err != nil {
		return err
	}

	return nil
}

func (m *ReportGroup) validateChanges(formats strfmt.Registry) error {

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

func (m *ReportGroup) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportGroup) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *ReportGroup) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 255); err != nil {
		return err
	}

	return nil
}

func (m *ReportGroup) validateValueFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueFormat) { // not required
		return nil
	}

	if err := validate.MaxLength("valueFormat", "body", string(m.ValueFormat), 100); err != nil {
		return err
	}

	return nil
}

func (m *ReportGroup) validateVariableName(formats strfmt.Registry) error {

	if swag.IsZero(m.VariableName) { // not required
		return nil
	}

	if err := validate.MaxLength("variableName", "body", string(m.VariableName), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportGroup) UnmarshalBinary(b []byte) error {
	var res ReportGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
