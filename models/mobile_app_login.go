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

// MobileAppLogin mobile app login
//
// swagger:model MobileAppLogin
type MobileAppLogin struct {

	// App secret
	// Required: true
	AppSecret *string `json:"appSecret"`

	// Optional employee ID. Default employee is used when null
	EmployeeID int32 `json:"employeeId,omitempty"`

	// Expiration date for the combined token
	// Required: true
	ExpirationDate *string `json:"expirationDate"`

	// Optional mfa auth code
	MfaCode int32 `json:"mfaCode,omitempty"`

	// Users password
	// Required: true
	Password *string `json:"password"`

	// Users username (email)
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this mobile app login
func (m *MobileAppLogin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileAppLogin) validateAppSecret(formats strfmt.Registry) error {

	if err := validate.Required("appSecret", "body", m.AppSecret); err != nil {
		return err
	}

	return nil
}

func (m *MobileAppLogin) validateExpirationDate(formats strfmt.Registry) error {

	if err := validate.Required("expirationDate", "body", m.ExpirationDate); err != nil {
		return err
	}

	return nil
}

func (m *MobileAppLogin) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *MobileAppLogin) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MobileAppLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileAppLogin) UnmarshalBinary(b []byte) error {
	var res MobileAppLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
