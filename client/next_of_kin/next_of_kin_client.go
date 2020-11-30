// Code generated by go-swagger; DO NOT EDIT.

package next_of_kin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new next of kin API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for next of kin API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	EmployeeNextOfKinGet(params *EmployeeNextOfKinGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinGetOK, error)

	EmployeeNextOfKinPost(params *EmployeeNextOfKinPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinPostCreated, error)

	EmployeeNextOfKinPut(params *EmployeeNextOfKinPutParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinPutOK, error)

	EmployeeNextOfKinSearch(params *EmployeeNextOfKinSearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EmployeeNextOfKinGet bs e t a find next of kin by ID
*/
func (a *Client) EmployeeNextOfKinGet(params *EmployeeNextOfKinGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeNextOfKinGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeNextOfKin_get",
		Method:             "GET",
		PathPattern:        "/employee/nextOfKin/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeNextOfKinGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeNextOfKinGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeNextOfKin_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeNextOfKinPost bs e t a create next of kin
*/
func (a *Client) EmployeeNextOfKinPost(params *EmployeeNextOfKinPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeNextOfKinPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeNextOfKin_post",
		Method:             "POST",
		PathPattern:        "/employee/nextOfKin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeNextOfKinPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeNextOfKinPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeNextOfKin_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeNextOfKinPut bs e t a update next of kin
*/
func (a *Client) EmployeeNextOfKinPut(params *EmployeeNextOfKinPutParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeNextOfKinPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeNextOfKin_put",
		Method:             "PUT",
		PathPattern:        "/employee/nextOfKin/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeNextOfKinPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeNextOfKinPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeNextOfKin_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EmployeeNextOfKinSearch finds all next of kin for employee
*/
func (a *Client) EmployeeNextOfKinSearch(params *EmployeeNextOfKinSearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeNextOfKinSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeNextOfKinSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeNextOfKin_search",
		Method:             "GET",
		PathPattern:        "/employee/nextOfKin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeNextOfKinSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeNextOfKinSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeNextOfKin_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
