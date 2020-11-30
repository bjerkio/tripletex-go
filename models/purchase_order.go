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

// PurchaseOrder purchase order
//
// swagger:model PurchaseOrder
type PurchaseOrder struct {

	// Attention
	Attention *Employee `json:"attention,omitempty"`

	// changes
	// Read Only: true
	Changes []*Change `json:"changes"`

	// Delivery information and invoice comments
	Comments string `json:"comments,omitempty"`

	// creation date
	CreationDate string `json:"creationDate,omitempty"`

	// Company currency
	Currency *Currency `json:"currency,omitempty"`

	// Delivery address
	DeliveryAddress *Address `json:"deliveryAddress,omitempty"`

	// delivery date
	// Required: true
	DeliveryDate *string `json:"deliveryDate"`

	// Department/order
	Department *Department `json:"department,omitempty"`

	// Discount Percentage
	Discount float64 `json:"discount,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is closed
	IsClosed *bool `json:"isClosed,omitempty"`

	// Purchase order number
	// Read Only: true
	// Max Length: 100
	Number string `json:"number,omitempty"`

	// Order lines tied to the purchase order
	// Read Only: true
	OrderLines []*PurchaseOrderline `json:"orderLines"`

	// our contact
	// Required: true
	OurContact *Employee `json:"ourContact"`

	// Message on packing note.Wholesaler specific.
	// Max Length: 50
	PackingNoteMessage string `json:"packingNoteMessage,omitempty"`

	// Project/order
	Project *Project `json:"project,omitempty"`

	// Email when purchase order is send by email.
	// Max Length: 100
	ReceiverEmail string `json:"receiverEmail,omitempty"`

	// restorder
	// Read Only: true
	Restorder *PurchaseOrder `json:"restorder,omitempty"`

	// status
	// Read Only: true
	// Enum: [STATUS_OPEN STATUS_SENT STATUS_RECEIVING STATUS_CONFIRMED_DEVIATION_DETECTED STATUS_DEVIATION_OPEN STATUS_DEVIATION_CONFIRMED STATUS_CONFIRMED]
	Status string `json:"status,omitempty"`

	// supplier
	// Required: true
	Supplier *Supplier `json:"supplier"`

	// Recipient when purchase order is send by email.
	SupplierContact *Employee `json:"supplierContact,omitempty"`

	// Transport type
	TransportType *TransportType `json:"transportType,omitempty"`

	// Message to transporter.Wholesaler specific.
	// Max Length: 255
	TransporterMessage string `json:"transporterMessage,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this purchase order
func (m *PurchaseOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOurContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackingNoteMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiverEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestorder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplierContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransporterMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseOrder) validateAttention(formats strfmt.Registry) error {

	if swag.IsZero(m.Attention) { // not required
		return nil
	}

	if m.Attention != nil {
		if err := m.Attention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attention")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrder) validateChanges(formats strfmt.Registry) error {

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

func (m *PurchaseOrder) validateCurrency(formats strfmt.Registry) error {

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

func (m *PurchaseOrder) validateDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveryAddress) { // not required
		return nil
	}

	if m.DeliveryAddress != nil {
		if err := m.DeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrder) validateDeliveryDate(formats strfmt.Registry) error {

	if err := validate.Required("deliveryDate", "body", m.DeliveryDate); err != nil {
		return err
	}

	return nil
}

func (m *PurchaseOrder) validateDepartment(formats strfmt.Registry) error {

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

func (m *PurchaseOrder) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaxLength("number", "body", string(m.Number), 100); err != nil {
		return err
	}

	return nil
}

func (m *PurchaseOrder) validateOrderLines(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderLines) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderLines); i++ {
		if swag.IsZero(m.OrderLines[i]) { // not required
			continue
		}

		if m.OrderLines[i] != nil {
			if err := m.OrderLines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PurchaseOrder) validateOurContact(formats strfmt.Registry) error {

	if err := validate.Required("ourContact", "body", m.OurContact); err != nil {
		return err
	}

	if m.OurContact != nil {
		if err := m.OurContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ourContact")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrder) validatePackingNoteMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.PackingNoteMessage) { // not required
		return nil
	}

	if err := validate.MaxLength("packingNoteMessage", "body", string(m.PackingNoteMessage), 50); err != nil {
		return err
	}

	return nil
}

func (m *PurchaseOrder) validateProject(formats strfmt.Registry) error {

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

func (m *PurchaseOrder) validateReceiverEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceiverEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("receiverEmail", "body", string(m.ReceiverEmail), 100); err != nil {
		return err
	}

	return nil
}

func (m *PurchaseOrder) validateRestorder(formats strfmt.Registry) error {

	if swag.IsZero(m.Restorder) { // not required
		return nil
	}

	if m.Restorder != nil {
		if err := m.Restorder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restorder")
			}
			return err
		}
	}

	return nil
}

var purchaseOrderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATUS_OPEN","STATUS_SENT","STATUS_RECEIVING","STATUS_CONFIRMED_DEVIATION_DETECTED","STATUS_DEVIATION_OPEN","STATUS_DEVIATION_CONFIRMED","STATUS_CONFIRMED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		purchaseOrderTypeStatusPropEnum = append(purchaseOrderTypeStatusPropEnum, v)
	}
}

const (

	// PurchaseOrderStatusSTATUSOPEN captures enum value "STATUS_OPEN"
	PurchaseOrderStatusSTATUSOPEN string = "STATUS_OPEN"

	// PurchaseOrderStatusSTATUSSENT captures enum value "STATUS_SENT"
	PurchaseOrderStatusSTATUSSENT string = "STATUS_SENT"

	// PurchaseOrderStatusSTATUSRECEIVING captures enum value "STATUS_RECEIVING"
	PurchaseOrderStatusSTATUSRECEIVING string = "STATUS_RECEIVING"

	// PurchaseOrderStatusSTATUSCONFIRMEDDEVIATIONDETECTED captures enum value "STATUS_CONFIRMED_DEVIATION_DETECTED"
	PurchaseOrderStatusSTATUSCONFIRMEDDEVIATIONDETECTED string = "STATUS_CONFIRMED_DEVIATION_DETECTED"

	// PurchaseOrderStatusSTATUSDEVIATIONOPEN captures enum value "STATUS_DEVIATION_OPEN"
	PurchaseOrderStatusSTATUSDEVIATIONOPEN string = "STATUS_DEVIATION_OPEN"

	// PurchaseOrderStatusSTATUSDEVIATIONCONFIRMED captures enum value "STATUS_DEVIATION_CONFIRMED"
	PurchaseOrderStatusSTATUSDEVIATIONCONFIRMED string = "STATUS_DEVIATION_CONFIRMED"

	// PurchaseOrderStatusSTATUSCONFIRMED captures enum value "STATUS_CONFIRMED"
	PurchaseOrderStatusSTATUSCONFIRMED string = "STATUS_CONFIRMED"
)

// prop value enum
func (m *PurchaseOrder) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, purchaseOrderTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PurchaseOrder) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PurchaseOrder) validateSupplier(formats strfmt.Registry) error {

	if err := validate.Required("supplier", "body", m.Supplier); err != nil {
		return err
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

func (m *PurchaseOrder) validateSupplierContact(formats strfmt.Registry) error {

	if swag.IsZero(m.SupplierContact) { // not required
		return nil
	}

	if m.SupplierContact != nil {
		if err := m.SupplierContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supplierContact")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrder) validateTransportType(formats strfmt.Registry) error {

	if swag.IsZero(m.TransportType) { // not required
		return nil
	}

	if m.TransportType != nil {
		if err := m.TransportType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transportType")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseOrder) validateTransporterMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.TransporterMessage) { // not required
		return nil
	}

	if err := validate.MaxLength("transporterMessage", "body", string(m.TransporterMessage), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurchaseOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchaseOrder) UnmarshalBinary(b []byte) error {
	var res PurchaseOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
