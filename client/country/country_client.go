// Code generated by go-swagger; DO NOT EDIT.

package country

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new country API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for country API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CountryGet(params *CountryGetParams, authInfo runtime.ClientAuthInfoWriter) (*CountryGetOK, error)

	CountrySearch(params *CountrySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CountrySearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CountryGet gets country by ID
*/
func (a *Client) CountryGet(params *CountryGetParams, authInfo runtime.ClientAuthInfoWriter) (*CountryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CountryGet",
		Method:             "GET",
		PathPattern:        "/country/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CountryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CountryGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CountrySearch finds countries corresponding with sent data
*/
func (a *Client) CountrySearch(params *CountrySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CountrySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountrySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CountrySearch",
		Method:             "GET",
		PathPattern:        "/country",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountrySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CountrySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CountrySearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
