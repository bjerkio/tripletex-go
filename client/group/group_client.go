// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProductGroupListDeleteByIds(params *ProductGroupListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	ProductGroupListPostList(params *ProductGroupListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupListPostListCreated, error)

	ProductGroupListPutList(params *ProductGroupListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupListPutListOK, error)

	ProductGroupDelete(params *ProductGroupDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	ProductGroupGet(params *ProductGroupGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupGetOK, error)

	ProductGroupPost(params *ProductGroupPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupPostCreated, error)

	ProductGroupPut(params *ProductGroupPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupPutOK, error)

	ProductGroupSearch(params *ProductGroupSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProductGroupListDeleteByIds bs e t a delete multiple product groups
*/
func (a *Client) ProductGroupListDeleteByIds(params *ProductGroupListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupList_deleteByIds",
		Method:             "DELETE",
		PathPattern:        "/product/group/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupListDeleteByIdsReader{formats: a.formats},
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
  ProductGroupListPostList bs e t a add multiple products groups
*/
func (a *Client) ProductGroupListPostList(params *ProductGroupListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupList_postList",
		Method:             "POST",
		PathPattern:        "/product/group/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupListPutList bs e t a update a list of product groups
*/
func (a *Client) ProductGroupListPutList(params *ProductGroupListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroupList_putList",
		Method:             "PUT",
		PathPattern:        "/product/group/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroupList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupDelete bs e t a delete product group
*/
func (a *Client) ProductGroupDelete(params *ProductGroupDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroup_delete",
		Method:             "DELETE",
		PathPattern:        "/product/group/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupDeleteReader{formats: a.formats},
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
  ProductGroupGet bs e t a find product group by ID
*/
func (a *Client) ProductGroupGet(params *ProductGroupGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroup_get",
		Method:             "GET",
		PathPattern:        "/product/group/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroup_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupPost bs e t a create new product group
*/
func (a *Client) ProductGroupPost(params *ProductGroupPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroup_post",
		Method:             "POST",
		PathPattern:        "/product/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroup_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupPut bs e t a update product group
*/
func (a *Client) ProductGroupPut(params *ProductGroupPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroup_put",
		Method:             "PUT",
		PathPattern:        "/product/group/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroup_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductGroupSearch bs e t a find product group with sent data
*/
func (a *Client) ProductGroupSearch(params *ProductGroupSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProductGroupSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductGroupSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProductGroup_search",
		Method:             "GET",
		PathPattern:        "/product/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductGroupSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductGroupSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProductGroup_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
