// Code generated by go-swagger; DO NOT EDIT.

package company

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new company API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for company API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompanyDivisionsGetDivisions(params *CompanyDivisionsGetDivisionsParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyDivisionsGetDivisionsOK, error)

	CompanyGet(params *CompanyGetParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyGetOK, error)

	CompanyPut(params *CompanyPutParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyPutOK, error)

	CompanyWithLoginAccessGetWithLoginAccess(params *CompanyWithLoginAccessGetWithLoginAccessParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyWithLoginAccessGetWithLoginAccessOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompanyDivisionsGetDivisions ds e p r e c a t e d find divisions
*/
func (a *Client) CompanyDivisionsGetDivisions(params *CompanyDivisionsGetDivisionsParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyDivisionsGetDivisionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanyDivisionsGetDivisionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanyDivisionsGetDivisions",
		Method:             "GET",
		PathPattern:        "/company/divisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanyDivisionsGetDivisionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanyDivisionsGetDivisionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanyDivisionsGetDivisions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompanyGet finds company by ID
*/
func (a *Client) CompanyGet(params *CompanyGetParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanyGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanyGet",
		Method:             "GET",
		PathPattern:        "/company/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanyGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanyGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompanyPut updates company information
*/
func (a *Client) CompanyPut(params *CompanyPutParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanyPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanyPut",
		Method:             "PUT",
		PathPattern:        "/company",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanyPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanyPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanyPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompanyWithLoginAccessGetWithLoginAccess returns client customers with accountant auditor relation where the current user has login access proxy login
*/
func (a *Client) CompanyWithLoginAccessGetWithLoginAccess(params *CompanyWithLoginAccessGetWithLoginAccessParams, authInfo runtime.ClientAuthInfoWriter) (*CompanyWithLoginAccessGetWithLoginAccessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompanyWithLoginAccessGetWithLoginAccessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompanyWithLoginAccessGetWithLoginAccess",
		Method:             "GET",
		PathPattern:        "/company/>withLoginAccess",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompanyWithLoginAccessGetWithLoginAccessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompanyWithLoginAccessGetWithLoginAccessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompanyWithLoginAccessGetWithLoginAccess: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
