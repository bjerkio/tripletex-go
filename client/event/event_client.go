// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new event API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for event API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	EventExample(params *EventExampleParams, authInfo runtime.ClientAuthInfoWriter) (*EventExampleOK, error)

	EventGet(params *EventGetParams) (*EventGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EventExample bs e t a get example webhook payload
*/
func (a *Client) EventExample(params *EventExampleParams, authInfo runtime.ClientAuthInfoWriter) (*EventExampleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEventExampleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Event_example",
		Method:             "GET",
		PathPattern:        "/event/{eventType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EventExampleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EventExampleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Event_example: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EventGet bs e t a get all web hook event keys
*/
func (a *Client) EventGet(params *EventGetParams) (*EventGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEventGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Event_get",
		Method:             "GET",
		PathPattern:        "/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EventGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EventGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Event_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
