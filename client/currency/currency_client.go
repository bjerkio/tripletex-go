// Code generated by go-swagger; DO NOT EDIT.

package currency

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new currency API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for currency API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CurrencyRateGetRate(params *CurrencyRateGetRateParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyRateGetRateOK, error)

	CurrencyGet(params *CurrencyGetParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyGetOK, error)

	CurrencySearch(params *CurrencySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencySearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CurrencyRateGetRate finds currency exchange rate corresponding with sent data
*/
func (a *Client) CurrencyRateGetRate(params *CurrencyRateGetRateParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyRateGetRateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCurrencyRateGetRateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CurrencyRate_getRate",
		Method:             "GET",
		PathPattern:        "/currency/{id}/rate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CurrencyRateGetRateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CurrencyRateGetRateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CurrencyRate_getRate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CurrencyGet gets currency by ID
*/
func (a *Client) CurrencyGet(params *CurrencyGetParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCurrencyGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Currency_get",
		Method:             "GET",
		PathPattern:        "/currency/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CurrencyGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CurrencyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Currency_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CurrencySearch finds currencies corresponding with sent data
*/
func (a *Client) CurrencySearch(params *CurrencySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCurrencySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Currency_search",
		Method:             "GET",
		PathPattern:        "/currency",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CurrencySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CurrencySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Currency_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
