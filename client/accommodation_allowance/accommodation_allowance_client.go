// Code generated by go-swagger; DO NOT EDIT.

package accommodation_allowance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accommodation allowance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accommodation allowance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TravelExpenseAccommodationAllowanceDelete(params *TravelExpenseAccommodationAllowanceDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseAccommodationAllowanceGet(params *TravelExpenseAccommodationAllowanceGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowanceGetOK, error)

	TravelExpenseAccommodationAllowancePost(params *TravelExpenseAccommodationAllowancePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowancePostCreated, error)

	TravelExpenseAccommodationAllowancePut(params *TravelExpenseAccommodationAllowancePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowancePutOK, error)

	TravelExpenseAccommodationAllowanceSearch(params *TravelExpenseAccommodationAllowanceSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowanceSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TravelExpenseAccommodationAllowanceDelete bs e t a delete accommodation allowance
*/
func (a *Client) TravelExpenseAccommodationAllowanceDelete(params *TravelExpenseAccommodationAllowanceDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAccommodationAllowanceDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAccommodationAllowanceDelete",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/accommodationAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAccommodationAllowanceDeleteReader{formats: a.formats},
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
  TravelExpenseAccommodationAllowanceGet bs e t a get travel accommodation allowance by ID
*/
func (a *Client) TravelExpenseAccommodationAllowanceGet(params *TravelExpenseAccommodationAllowanceGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowanceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAccommodationAllowanceGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAccommodationAllowanceGet",
		Method:             "GET",
		PathPattern:        "/travelExpense/accommodationAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAccommodationAllowanceGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseAccommodationAllowanceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseAccommodationAllowanceGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseAccommodationAllowancePost bs e t a create accommodation allowance
*/
func (a *Client) TravelExpenseAccommodationAllowancePost(params *TravelExpenseAccommodationAllowancePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowancePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAccommodationAllowancePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAccommodationAllowancePost",
		Method:             "POST",
		PathPattern:        "/travelExpense/accommodationAllowance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAccommodationAllowancePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseAccommodationAllowancePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseAccommodationAllowancePost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseAccommodationAllowancePut bs e t a update accommodation allowance
*/
func (a *Client) TravelExpenseAccommodationAllowancePut(params *TravelExpenseAccommodationAllowancePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowancePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAccommodationAllowancePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAccommodationAllowancePut",
		Method:             "PUT",
		PathPattern:        "/travelExpense/accommodationAllowance/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAccommodationAllowancePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseAccommodationAllowancePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseAccommodationAllowancePut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseAccommodationAllowanceSearch bs e t a find accommodation allowances corresponding with sent data
*/
func (a *Client) TravelExpenseAccommodationAllowanceSearch(params *TravelExpenseAccommodationAllowanceSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAccommodationAllowanceSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAccommodationAllowanceSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAccommodationAllowanceSearch",
		Method:             "GET",
		PathPattern:        "/travelExpense/accommodationAllowance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAccommodationAllowanceSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseAccommodationAllowanceSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseAccommodationAllowanceSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
