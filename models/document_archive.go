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

// DocumentArchive document archive
//
// swagger:model DocumentArchive
type DocumentArchive struct {

	// archive date
	ArchiveDate string `json:"archiveDate,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// The name of the document.
	// Required: true
	// Max Length: 255
	FileName *string `json:"fileName"`

	// id
	ID int32 `json:"id,omitempty"`

	// Type of the document
	// Required: true
	// Max Length: 100
	MimeType *string `json:"mimeType"`

	// The size of the document in readable format.
	// Read Only: true
	// Minimum: 0
	Size int32 `json:"size,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this document archive
func (m *DocumentArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentArchive) validateChanges(formats strfmt.Registry) error {

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

func (m *DocumentArchive) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("fileName", "body", m.FileName); err != nil {
		return err
	}

	if err := validate.MaxLength("fileName", "body", string(*m.FileName), 255); err != nil {
		return err
	}

	return nil
}

func (m *DocumentArchive) validateMimeType(formats strfmt.Registry) error {

	if err := validate.Required("mimeType", "body", m.MimeType); err != nil {
		return err
	}

	if err := validate.MaxLength("mimeType", "body", string(*m.MimeType), 100); err != nil {
		return err
	}

	return nil
}

func (m *DocumentArchive) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := validate.MinimumInt("size", "body", int64(m.Size), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentArchive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentArchive) UnmarshalBinary(b []byte) error {
	var res DocumentArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
