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

// Product product
//
// swagger:model Product
type Product struct {

	// account
	Account *Account `json:"account,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// Price purchase (cost) excluding VAT in the product's currency
	CostExcludingVatCurrency float64 `json:"costExcludingVatCurrency,omitempty"`

	// currency
	Currency *Currency `json:"currency,omitempty"`

	// department
	Department *Department `json:"department,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// discount price
	// Read Only: true
	DiscountPrice float64 `json:"discountPrice,omitempty"`

	// ean
	// Max Length: 14
	Ean string `json:"ean,omitempty"`

	// el number
	// Read Only: true
	// Max Length: 14
	ElNumber string `json:"elNumber,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is inactive
	IsInactive *bool `json:"isInactive,omitempty"`

	// is stock item
	IsStockItem *bool `json:"isStockItem,omitempty"`

	// name
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// nrf number
	// Read Only: true
	// Max Length: 14
	NrfNumber string `json:"nrfNumber,omitempty"`

	// number
	// Max Length: 100
	Number string `json:"number,omitempty"`

	// Price of purchase excluding VAT in the product's currency
	PriceExcludingVatCurrency float64 `json:"priceExcludingVatCurrency,omitempty"`

	// Price of purchase including VAT in the product's currency
	PriceIncludingVatCurrency float64 `json:"priceIncludingVatCurrency,omitempty"`

	// product unit
	ProductUnit *ProductUnit `json:"productUnit,omitempty"`

	// resale product
	ResaleProduct *Product `json:"resaleProduct,omitempty"`

	// stock of goods
	// Read Only: true
	StockOfGoods float64 `json:"stockOfGoods,omitempty"`

	// supplier
	Supplier *Supplier `json:"supplier,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// vat type
	VatType *VatType `json:"vatType,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNrfNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResaleProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplier(formats); err != nil {
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

func (m *Product) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateChanges(formats strfmt.Registry) error {

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

func (m *Product) validateCurrency(formats strfmt.Registry) error {

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

func (m *Product) validateDepartment(formats strfmt.Registry) error {

	if swag.IsZero(m.Department) { // not required
		return nil
	}

	if m.Department != nil {
		if err := m.Department.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("department")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateEan(formats strfmt.Registry) error {

	if swag.IsZero(m.Ean) { // not required
		return nil
	}

	if err := validate.MaxLength("ean", "body", string(m.Ean), 14); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateElNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.ElNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("elNumber", "body", string(m.ElNumber), 14); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 255); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateNrfNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.NrfNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("nrfNumber", "body", string(m.NrfNumber), 14); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaxLength("number", "body", string(m.Number), 100); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateProductUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductUnit) { // not required
		return nil
	}

	if m.ProductUnit != nil {
		if err := m.ProductUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productUnit")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateResaleProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.ResaleProduct) { // not required
		return nil
	}

	if m.ResaleProduct != nil {
		if err := m.ResaleProduct.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resaleProduct")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateSupplier(formats strfmt.Registry) error {

	if swag.IsZero(m.Supplier) { // not required
		return nil
	}

	if m.Supplier != nil {
		if err := m.Supplier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supplier")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateVatType(formats strfmt.Registry) error {

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
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
