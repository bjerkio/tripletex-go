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

// OrderLine order line
//
// swagger:model OrderLine
type OrderLine struct {

	// Total amount on order line excluding VAT in the order's currency
	// Read Only: true
	AmountExcludingVatCurrency float64 `json:"amountExcludingVatCurrency,omitempty"`

	// Total amount on order line including VAT in the order's currency
	// Read Only: true
	AmountIncludingVatCurrency float64 `json:"amountIncludingVatCurrency,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// count
	Count float64 `json:"count,omitempty"`

	// The order line's currency. Determined by the order's currency.
	// Read Only: true
	Currency *Currency `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Discount given as a percentage (%)
	Discount float64 `json:"discount,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// inventory
	Inventory *Inventory `json:"inventory,omitempty"`

	// is subscription
	IsSubscription *bool `json:"isSubscription,omitempty"`

	// Markup given as a percentage (%)
	Markup float64 `json:"markup,omitempty"`

	// The related Order for this OrderLine. This field is required when creating a new OrderLine in the API.
	Order *Order `json:"order,omitempty"`

	// order group
	OrderGroup *OrderGroup `json:"orderGroup,omitempty"`

	// product
	Product *Product `json:"product,omitempty"`

	// subscription period end
	SubscriptionPeriodEnd string `json:"subscriptionPeriodEnd,omitempty"`

	// subscription period start
	SubscriptionPeriodStart string `json:"subscriptionPeriodStart,omitempty"`

	// Unit price purchase (cost) excluding VAT in the order's currency
	UnitCostCurrency float64 `json:"unitCostCurrency,omitempty"`

	// Unit price of purchase excluding VAT in the order's currency
	UnitPriceExcludingVatCurrency float64 `json:"unitPriceExcludingVatCurrency,omitempty"`

	// Unit price of purchase including VAT in the order's currency
	UnitPriceIncludingVatCurrency float64 `json:"unitPriceIncludingVatCurrency,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// vat type
	VatType *VatType `json:"vatType,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this order line
func (m *OrderLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInventory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderLine) validateChanges(formats strfmt.Registry) error {

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

func (m *OrderLine) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLine) validateInventory(formats strfmt.Registry) error {

	if swag.IsZero(m.Inventory) { // not required
		return nil
	}

	if m.Inventory != nil {
		if err := m.Inventory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLine) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLine) validateOrderGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderGroup) { // not required
		return nil
	}

	if m.OrderGroup != nil {
		if err := m.OrderGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderGroup")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLine) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLine) validateVatType(formats strfmt.Registry) error {

	if swag.IsZero(m.VatType) { // not required
		return nil
	}

	if m.VatType != nil {
		if err := m.VatType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vatType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderLine) UnmarshalBinary(b []byte) error {
	var res OrderLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
