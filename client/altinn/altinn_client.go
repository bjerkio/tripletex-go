// Code generated by go-swagger; DO NOT EDIT.

package altinn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new altinn API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for altinn API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompanySettingsAltinnPut(params *CompanySettingsAltinnPutParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySettingsAltinnPutOK, error)

	CompanySettingsAltinnSearch(params *CompanySettingsAltinnSearchParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySettingsAltinnSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompanySettingsAltinnPut bs e t a update alt inn id and password
*/
func (a *Client) CompanySettingsAltinnPut(params *CompanySettingsAltinnPutParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySettingsAltinnPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanySettingsAltinnPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanySettingsAltinnPut",
		Method:             "PUT",
		PathPattern:        "/company/settings/altinn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanySettingsAltinnPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanySettingsAltinnPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanySettingsAltinnPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompanySettingsAltinnSearch bs e t a find altinn id for login in company
*/
func (a *Client) CompanySettingsAltinnSearch(params *CompanySettingsAltinnSearchParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySettingsAltinnSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanySettingsAltinnSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanySettingsAltinnSearch",
		Method:             "GET",
		PathPattern:        "/company/settings/altinn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanySettingsAltinnSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanySettingsAltinnSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanySettingsAltinnSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
