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

// Subscription subscription
//
// swagger:model Subscription
type Subscription struct {

	// Custom authentication header name
	// Max Length: 255
	AuthHeaderName string `json:"authHeaderName,omitempty"`

	// Custom authentication header value (write only)
	// Max Length: 4096
	AuthHeaderValue string `json:"authHeaderValue,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// Event name (from /v2/event) you wish to subscribe to. Form should be: *subject.verb*.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Event *string `json:"event"`

	// The fields in the object delivered with the notification callback, nested as in other API calls.
	// Max Length: 255
	Fields string `json:"fields,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// The status of the subscription.
	// Read Only: true
	// Enum: [ACTIVE DISABLED DISABLED_TOO_MANY_ERRORS DISABLED_RATE_LIMIT_EXCEEDED DISABLED_MISUSE]
	Status string `json:"status,omitempty"`

	// The callback URL used for subscriptions (including authentication tokens). Must be absolute and use HTTPS. Basic authentication is supported.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	TargetURL *string `json:"targetUrl"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthHeaderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthHeaderValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateAuthHeaderName(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthHeaderName) { // not required
		return nil
	}

	if err := validate.MaxLength("authHeaderName", "body", string(m.AuthHeaderName), 255); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateAuthHeaderValue(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthHeaderValue) { // not required
		return nil
	}

	if err := validate.MaxLength("authHeaderValue", "body", string(m.AuthHeaderValue), 4096); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateChanges(formats strfmt.Registry) error {

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

func (m *Subscription) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	if err := validate.MinLength("event", "body", string(*m.Event), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("event", "body", string(*m.Event), 255); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if err := validate.MaxLength("fields", "body", string(m.Fields), 255); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","DISABLED","DISABLED_TOO_MANY_ERRORS","DISABLED_RATE_LIMIT_EXCEEDED","DISABLED_MISUSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeStatusPropEnum = append(subscriptionTypeStatusPropEnum, v)
	}
}

const (

	// SubscriptionStatusACTIVE captures enum value "ACTIVE"
	SubscriptionStatusACTIVE string = "ACTIVE"

	// SubscriptionStatusDISABLED captures enum value "DISABLED"
	SubscriptionStatusDISABLED string = "DISABLED"

	// SubscriptionStatusDISABLEDTOOMANYERRORS captures enum value "DISABLED_TOO_MANY_ERRORS"
	SubscriptionStatusDISABLEDTOOMANYERRORS string = "DISABLED_TOO_MANY_ERRORS"

	// SubscriptionStatusDISABLEDRATELIMITEXCEEDED captures enum value "DISABLED_RATE_LIMIT_EXCEEDED"
	SubscriptionStatusDISABLEDRATELIMITEXCEEDED string = "DISABLED_RATE_LIMIT_EXCEEDED"

	// SubscriptionStatusDISABLEDMISUSE captures enum value "DISABLED_MISUSE"
	SubscriptionStatusDISABLEDMISUSE string = "DISABLED_MISUSE"
)

// prop value enum
func (m *Subscription) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, subscriptionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateTargetURL(formats strfmt.Registry) error {

	if err := validate.Required("targetUrl", "body", m.TargetURL); err != nil {
		return err
	}

	if err := validate.MinLength("targetUrl", "body", string(*m.TargetURL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("targetUrl", "body", string(*m.TargetURL), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
