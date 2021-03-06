// Code generated by go-swagger; DO NOT EDIT.

package group_relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group relation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group relation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProductGroupRelationListDeleteByIds(params *ProductGroupRelationListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	ProductGroupRelationListPostList(params *ProductGroupRelationListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationListPostListCreated, error)

	ProductGroupRelationDelete(params *ProductGroupRelationDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	ProductGroupRelationGet(params *ProductGroupRelationGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationGetOK, error)

	ProductGroupRelationPost(params *ProductGroupRelationPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationPostCreated, error)

	ProductGroupRelationSearch(params *ProductGroupRelationSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProductGroupRelationListDeleteByIds bs e t a delete multiple product group relations
*/
func (a *Client) ProductGroupRelationListDeleteByIds(params *ProductGroupRelationListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelationList_deleteByIds",
		Method:             "DELETE",
		PathPattern:        "/product/groupRelation/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationListDeleteByIdsReader{formats: a.formats},
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
  ProductGroupRelationListPostList bs e t a add multiple products group relations
*/
func (a *Client) ProductGroupRelationListPostList(params *ProductGroupRelationListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelationList_postList",
		Method:             "POST",
		PathPattern:        "/product/groupRelation/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupRelationListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupRelationList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupRelationDelete bs e t a delete product group relation
*/
func (a *Client) ProductGroupRelationDelete(params *ProductGroupRelationDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelation_delete",
		Method:             "DELETE",
		PathPattern:        "/product/groupRelation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationDeleteReader{formats: a.formats},
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
  ProductGroupRelationGet bs e t a find product group relation by ID
*/
func (a *Client) ProductGroupRelationGet(params *ProductGroupRelationGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelation_get",
		Method:             "GET",
		PathPattern:        "/product/groupRelation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupRelationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupRelation_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupRelationPost bs e t a create new product group relation
*/
func (a *Client) ProductGroupRelationPost(params *ProductGroupRelationPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelation_post",
		Method:             "POST",
		PathPattern:        "/product/groupRelation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupRelationPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupRelation_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupRelationSearch bs e t a find product group relation with sent data
*/
func (a *Client) ProductGroupRelationSearch(params *ProductGroupRelationSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupRelationSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupRelationSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupRelation_search",
		Method:             "GET",
		PathPattern:        "/product/groupRelation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupRelationSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupRelationSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupRelation_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
