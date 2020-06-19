// Code generated by go-swagger; DO NOT EDIT.

package mileage_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mileage allowance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mileage allowance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TravelExpenseMileageAllowanceDelete(params *TravelExpenseMileageAllowanceDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseMileageAllowanceGet(params *TravelExpenseMileageAllowanceGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowanceGetOK, error)

	TravelExpenseMileageAllowancePost(params *TravelExpenseMileageAllowancePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowancePostCreated, error)

	TravelExpenseMileageAllowancePut(params *TravelExpenseMileageAllowancePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowancePutOK, error)

	TravelExpenseMileageAllowanceSearch(params *TravelExpenseMileageAllowanceSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowanceSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TravelExpenseMileageAllowanceDelete bs e t a delete mileage allowance
*/
func (a *Client) TravelExpenseMileageAllowanceDelete(params *TravelExpenseMileageAllowanceDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseMileageAllowanceDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseMileageAllowanceDelete",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseMileageAllowanceDeleteReader{formats: a.formats},
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
  TravelExpenseMileageAllowanceGet bs e t a get mileage allowance by ID
*/
func (a *Client) TravelExpenseMileageAllowanceGet(params *TravelExpenseMileageAllowanceGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowanceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseMileageAllowanceGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseMileageAllowanceGet",
		Method:             "GET",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseMileageAllowanceGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseMileageAllowanceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseMileageAllowanceGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseMileageAllowancePost bs e t a create mileage allowance
*/
func (a *Client) TravelExpenseMileageAllowancePost(params *TravelExpenseMileageAllowancePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowancePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseMileageAllowancePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseMileageAllowancePost",
		Method:             "POST",
		PathPattern:        "/travelExpense/mileageAllowance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseMileageAllowancePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseMileageAllowancePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseMileageAllowancePost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseMileageAllowancePut bs e t a update mileage allowance
*/
func (a *Client) TravelExpenseMileageAllowancePut(params *TravelExpenseMileageAllowancePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowancePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseMileageAllowancePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseMileageAllowancePut",
		Method:             "PUT",
		PathPattern:        "/travelExpense/mileageAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseMileageAllowancePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseMileageAllowancePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseMileageAllowancePut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseMileageAllowanceSearch bs e t a find mileage allowances corresponding with sent data
*/
func (a *Client) TravelExpenseMileageAllowanceSearch(params *TravelExpenseMileageAllowanceSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseMileageAllowanceSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseMileageAllowanceSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseMileageAllowanceSearch",
		Method:             "GET",
		PathPattern:        "/travelExpense/mileageAllowance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseMileageAllowanceSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseMileageAllowanceSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseMileageAllowanceSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
