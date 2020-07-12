// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LedgerAccountDelete(params *LedgerAccountDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	LedgerAccountGet(params *LedgerAccountGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountGetOK, error)

	LedgerAccountListDeleteByIds(params *LedgerAccountListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	LedgerAccountListPostList(params *LedgerAccountListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountListPostListCreated, error)

	LedgerAccountListPutList(params *LedgerAccountListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountListPutListOK, error)

	LedgerAccountPost(params *LedgerAccountPostParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountPostCreated, error)

	LedgerAccountPut(params *LedgerAccountPutParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountPutOK, error)

	LedgerAccountSearch(params *LedgerAccountSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LedgerAccountDelete bs e t a delete account
*/
func (a *Client) LedgerAccountDelete(params *LedgerAccountDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountDelete",
		Method:             "DELETE",
		PathPattern:        "/ledger/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountDeleteReader{formats: a.formats},
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
  LedgerAccountGet gets account by ID
*/
func (a *Client) LedgerAccountGet(params *LedgerAccountGetParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountGet",
		Method:             "GET",
		PathPattern:        "/ledger/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerAccountListDeleteByIds bs e t a delete multiple accounts
*/
func (a *Client) LedgerAccountListDeleteByIds(params *LedgerAccountListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountListDeleteByIds",
		Method:             "DELETE",
		PathPattern:        "/ledger/account/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountListDeleteByIdsReader{formats: a.formats},
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
  LedgerAccountListPostList bs e t a create several accounts
*/
func (a *Client) LedgerAccountListPostList(params *LedgerAccountListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountListPostList",
		Method:             "POST",
		PathPattern:        "/ledger/account/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountListPostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerAccountListPutList bs e t a update multiple accounts
*/
func (a *Client) LedgerAccountListPutList(params *LedgerAccountListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountListPutList",
		Method:             "PUT",
		PathPattern:        "/ledger/account/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountListPutList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerAccountPost bs e t a create a new account
*/
func (a *Client) LedgerAccountPost(params *LedgerAccountPostParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountPost",
		Method:             "POST",
		PathPattern:        "/ledger/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerAccountPut bs e t a update account
*/
func (a *Client) LedgerAccountPut(params *LedgerAccountPutParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountPut",
		Method:             "PUT",
		PathPattern:        "/ledger/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LedgerAccountSearch finds accounts corresponding with sent data
*/
func (a *Client) LedgerAccountSearch(params *LedgerAccountSearchParams, authInfo runtime.ClientAuthInfoWriter) (*LedgerAccountSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLedgerAccountSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LedgerAccountSearch",
		Method:             "GET",
		PathPattern:        "/ledger/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LedgerAccountSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LedgerAccountSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LedgerAccountSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
