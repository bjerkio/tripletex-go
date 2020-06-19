// Code generated by go-swagger; DO NOT EDIT.

package supplier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new supplier API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for supplier API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SupplierGet(params *SupplierGetParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierGetOK, error)

	SupplierListPostList(params *SupplierListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierListPostListCreated, error)

	SupplierListPutList(params *SupplierListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierListPutListOK, error)

	SupplierPost(params *SupplierPostParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierPostCreated, error)

	SupplierPut(params *SupplierPutParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierPutOK, error)

	SupplierSearch(params *SupplierSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SupplierGet gets supplier by ID
*/
func (a *Client) SupplierGet(params *SupplierGetParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierGet",
		Method:             "GET",
		PathPattern:        "/supplier/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SupplierListPostList bs e t a create multiple suppliers related supplier addresses may also be created
*/
func (a *Client) SupplierListPostList(params *SupplierListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierListPostList",
		Method:             "POST",
		PathPattern:        "/supplier/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierListPostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SupplierListPutList bs e t a update multiple suppliers addresses can also be updated
*/
func (a *Client) SupplierListPutList(params *SupplierListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierListPutList",
		Method:             "PUT",
		PathPattern:        "/supplier/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierListPutList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SupplierPost creates supplier related supplier addresses may also be created
*/
func (a *Client) SupplierPost(params *SupplierPostParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierPost",
		Method:             "POST",
		PathPattern:        "/supplier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SupplierPut updates supplier
*/
func (a *Client) SupplierPut(params *SupplierPutParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierPut",
		Method:             "PUT",
		PathPattern:        "/supplier/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SupplierSearch finds suppliers corresponding with sent data
*/
func (a *Client) SupplierSearch(params *SupplierSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SupplierSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupplierSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SupplierSearch",
		Method:             "GET",
		PathPattern:        "/supplier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupplierSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupplierSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SupplierSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
