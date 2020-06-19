// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	InventoryDelete(params *InventoryDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	InventoryGet(params *InventoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryGetOK, error)

	InventoryPost(params *InventoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryPostCreated, error)

	InventoryPut(params *InventoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryPutOK, error)

	InventorySearch(params *InventorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*InventorySearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InventoryDelete bs e t a delete inventory
*/
func (a *Client) InventoryDelete(params *InventoryDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInventoryDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InventoryDelete",
		Method:             "DELETE",
		PathPattern:        "/inventory/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InventoryDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InventoryGet gets inventory by ID
*/
func (a *Client) InventoryGet(params *InventoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInventoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InventoryGet",
		Method:             "GET",
		PathPattern:        "/inventory/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InventoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InventoryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InventoryGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InventoryPost bs e t a create new inventory
*/
func (a *Client) InventoryPost(params *InventoryPostParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInventoryPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InventoryPost",
		Method:             "POST",
		PathPattern:        "/inventory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InventoryPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InventoryPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InventoryPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InventoryPut bs e t a update inventory
*/
func (a *Client) InventoryPut(params *InventoryPutParams, authInfo runtime.ClientAuthInfoWriter) (*InventoryPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInventoryPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InventoryPut",
		Method:             "PUT",
		PathPattern:        "/inventory/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InventoryPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InventoryPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InventoryPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InventorySearch finds inventory corresponding with sent data
*/
func (a *Client) InventorySearch(params *InventorySearchParams, authInfo runtime.ClientAuthInfoWriter) (*InventorySearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInventorySearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InventorySearch",
		Method:             "GET",
		PathPattern:        "/inventory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InventorySearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InventorySearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InventorySearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
