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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MaventaStatusDTO maventa status d t o
//
// swagger:model MaventaStatusDTO
type MaventaStatusDTO struct {

	// company id
	// Required: true
	// Read Only: true
	CompanyID string `json:"company_id"`

	// event
	// Required: true
	// Read Only: true
	Event string `json:"event"`

	// event data
	// Required: true
	// Read Only: true
	EventData *MaventaEventDataDTO `json:"event_data"`

	// event timestamp
	// Required: true
	// Read Only: true
	// Format: date-time
	EventTimestamp strfmt.DateTime `json:"event_timestamp"`
}

// Validate validates this maventa status d t o
func (m *MaventaStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaventaStatusDTO) validateCompanyID(formats strfmt.Registry) error {

	if err := validate.RequiredString("company_id", "body", string(m.CompanyID)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaStatusDTO) validateEvent(formats strfmt.Registry) error {

	if err := validate.RequiredString("event", "body", string(m.Event)); err != nil {
		return err
	}

	return nil
}

func (m *MaventaStatusDTO) validateEventData(formats strfmt.Registry) error {

	if err := validate.Required("event_data", "body", m.EventData); err != nil {
		return err
	}

	if m.EventData != nil {
		if err := m.EventData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_data")
			}
			return err
		}
	}

	return nil
}

func (m *MaventaStatusDTO) validateEventTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("event_timestamp", "body", strfmt.DateTime(m.EventTimestamp)); err != nil {
		return err
	}

	if err := validate.FormatOf("event_timestamp", "body", "date-time", m.EventTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaventaStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaventaStatusDTO) UnmarshalBinary(b []byte) error {
	var res MaventaStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
