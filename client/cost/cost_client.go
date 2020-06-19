// Code generated by go-swagger; DO NOT EDIT.

package cost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cost API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cost API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TravelExpenseCostDelete(params *TravelExpenseCostDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseCostGet(params *TravelExpenseCostGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostGetOK, error)

	TravelExpenseCostPost(params *TravelExpenseCostPostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostPostCreated, error)

	TravelExpenseCostPut(params *TravelExpenseCostPutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostPutOK, error)

	TravelExpenseCostSearch(params *TravelExpenseCostSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TravelExpenseCostDelete bs e t a delete cost
*/
func (a *Client) TravelExpenseCostDelete(params *TravelExpenseCostDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCostDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCostDelete",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCostDeleteReader{formats: a.formats},
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
  TravelExpenseCostGet bs e t a get cost by ID
*/
func (a *Client) TravelExpenseCostGet(params *TravelExpenseCostGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCostGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCostGet",
		Method:             "GET",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCostGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCostGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCostGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseCostPost bs e t a create cost
*/
func (a *Client) TravelExpenseCostPost(params *TravelExpenseCostPostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCostPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCostPost",
		Method:             "POST",
		PathPattern:        "/travelExpense/cost",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCostPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCostPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCostPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseCostPut bs e t a update cost
*/
func (a *Client) TravelExpenseCostPut(params *TravelExpenseCostPutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCostPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCostPut",
		Method:             "PUT",
		PathPattern:        "/travelExpense/cost/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCostPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCostPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCostPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseCostSearch bs e t a find costs corresponding with sent data
*/
func (a *Client) TravelExpenseCostSearch(params *TravelExpenseCostSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCostSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCostSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCostSearch",
		Method:             "GET",
		PathPattern:        "/travelExpense/cost",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCostSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCostSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCostSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
