// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new external API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for external API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProductExternalGet(params *ProductExternalGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductExternalGetOK, error)

	ProductExternalSearch(params *ProductExternalSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductExternalSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProductExternalGet bs e t a get external product by ID
*/
func (a *Client) ProductExternalGet(params *ProductExternalGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductExternalGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductExternalGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductExternalGet",
		Method:             "GET",
		PathPattern:        "/product/external/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductExternalGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductExternalGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductExternalGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductExternalSearch bs e t a find external products corresponding with sent data the sorting field is not in use on this endpoint
*/
func (a *Client) ProductExternalSearch(params *ProductExternalSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductExternalSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductExternalSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductExternalSearch",
		Method:             "GET",
		PathPattern:        "/product/external",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductExternalSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductExternalSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductExternalSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
