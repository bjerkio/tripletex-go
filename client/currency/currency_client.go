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
	CurrencyGet(params *CurrencyGetParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyGetOK, error)

	CurrencyRateGetRate(params *CurrencyRateGetRateParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencyRateGetRateOK, error)

	CurrencySearch(params *CurrencySearchParams, authInfo runtime.ClientAuthInfoWriter) (*CurrencySearchOK, error)

	SetTransport(transport runtime.ClientTransport)
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
		ID:                 "CurrencyGet",
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
	msg := fmt.Sprintf("unexpected success response for CurrencyGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
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
		ID:                 "CurrencyRateGetRate",
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
	msg := fmt.Sprintf("unexpected success response for CurrencyRateGetRate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
		ID:                 "CurrencySearch",
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
	msg := fmt.Sprintf("unexpected success response for CurrencySearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
