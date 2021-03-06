// Code generated by go-swagger; DO NOT EDIT.

package vat_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vat type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vat type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LedgerVatTypeGet(params *LedgerVatTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerVatTypeGetOK, error)

	LedgerVatTypeSearch(params *LedgerVatTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerVatTypeSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LedgerVatTypeGet gets vat type by ID
*/
func (a *Client) LedgerVatTypeGet(params *LedgerVatTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerVatTypeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerVatTypeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerVatType_get",
		Method:             "GET",
		PathPattern:        "/ledger/vatType/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerVatTypeGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerVatTypeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerVatType_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerVatTypeSearch finds vat types corresponding with sent data
*/
func (a *Client) LedgerVatTypeSearch(params *LedgerVatTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerVatTypeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerVatTypeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerVatType_search",
		Method:             "GET",
		PathPattern:        "/ledger/vatType",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerVatTypeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerVatTypeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerVatType_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
