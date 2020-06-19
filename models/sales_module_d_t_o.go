// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SalesModuleDTO sales module d t o
//
// swagger:model SalesModuleDTO
type SalesModuleDTO struct {

	// cost start date
	CostStartDate string `json:"costStartDate,omitempty"`

	// name
	// Required: true
	// Enum: [MAMUT MAMUT_WITH_WAGE AGRO_LICENCE AGRO_CLIENT AGRO_DOCUMENT_CENTER AGRO_INVOICE AGRO_WAGE]
	Name *string `json:"name"`
}

// Validate validates this sales module d t o
func (m *SalesModuleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var salesModuleDTOTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MAMUT","MAMUT_WITH_WAGE","AGRO_LICENCE","AGRO_CLIENT","AGRO_DOCUMENT_CENTER","AGRO_INVOICE","AGRO_WAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		salesModuleDTOTypeNamePropEnum = append(salesModuleDTOTypeNamePropEnum, v)
	}
}

const (

	// SalesModuleDTONameMAMUT captures enum value "MAMUT"
	SalesModuleDTONameMAMUT string = "MAMUT"

	// SalesModuleDTONameMAMUTWITHWAGE captures enum value "MAMUT_WITH_WAGE"
	SalesModuleDTONameMAMUTWITHWAGE string = "MAMUT_WITH_WAGE"

	// SalesModuleDTONameAGROLICENCE captures enum value "AGRO_LICENCE"
	SalesModuleDTONameAGROLICENCE string = "AGRO_LICENCE"

	// SalesModuleDTONameAGROCLIENT captures enum value "AGRO_CLIENT"
	SalesModuleDTONameAGROCLIENT string = "AGRO_CLIENT"

	// SalesModuleDTONameAGRODOCUMENTCENTER captures enum value "AGRO_DOCUMENT_CENTER"
	SalesModuleDTONameAGRODOCUMENTCENTER string = "AGRO_DOCUMENT_CENTER"

	// SalesModuleDTONameAGROINVOICE captures enum value "AGRO_INVOICE"
	SalesModuleDTONameAGROINVOICE string = "AGRO_INVOICE"

	// SalesModuleDTONameAGROWAGE captures enum value "AGRO_WAGE"
	SalesModuleDTONameAGROWAGE string = "AGRO_WAGE"
)

// prop value enum
func (m *SalesModuleDTO) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, salesModuleDTOTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SalesModuleDTO) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SalesModuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesModuleDTO) UnmarshalBinary(b []byte) error {
	var res SalesModuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
