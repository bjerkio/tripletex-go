// Code generated by go-swagger; DO NOT EDIT.

package salesmodules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new salesmodules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for salesmodules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompanySalesmodulesGet(params *CompanySalesmodulesGetParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySalesmodulesGetOK, error)

	CompanySalesmodulesPost(params *CompanySalesmodulesPostParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySalesmodulesPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompanySalesmodulesGet bs e t a get active sales modules
*/
func (a *Client) CompanySalesmodulesGet(params *CompanySalesmodulesGetParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySalesmodulesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanySalesmodulesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanySalesmodules_get",
		Method:             "GET",
		PathPattern:        "/company/salesmodules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanySalesmodulesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanySalesmodulesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanySalesmodules_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompanySalesmodulesPost bs e t a add activate a new sales module
*/
func (a *Client) CompanySalesmodulesPost(params *CompanySalesmodulesPostParams, authInfo runtime.ClientAuthInfoWriter) (*CompanySalesmodulesPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanySalesmodulesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanySalesmodules_post",
		Method:             "POST",
		PathPattern:        "/company/salesmodules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanySalesmodulesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanySalesmodulesPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanySalesmodules_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
