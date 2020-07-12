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
)

// ListResponseTravelExpenseRateCategoryGroup list response travel expense rate category group
//
// swagger:model ListResponseTravelExpenseRateCategoryGroup
type ListResponseTravelExpenseRateCategoryGroup struct {

	// count
	Count int32 `json:"count,omitempty"`

	// from
	From int32 `json:"from,omitempty"`

	// [DEPRECATED] Indicates whether there are more values available. Note: The value is not exact
	FullResultSize int32 `json:"fullResultSize,omitempty"`

	// values
	Values []*TravelExpenseRateCategoryGroup `json:"values"`

	// Used to know if the paginated list has changed.
	VersionDigest string `json:"versionDigest,omitempty"`
}

// Validate validates this list response travel expense rate category group
func (m *ListResponseTravelExpenseRateCategoryGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListResponseTravelExpenseRateCategoryGroup) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListResponseTravelExpenseRateCategoryGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListResponseTravelExpenseRateCategoryGroup) UnmarshalBinary(b []byte) error {
	var res ListResponseTravelExpenseRateCategoryGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
