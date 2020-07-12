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

package goods_receipt_line

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new goods receipt line API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for goods receipt line API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PurchaseOrderGoodsReceiptLineDelete(params *PurchaseOrderGoodsReceiptLineDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	PurchaseOrderGoodsReceiptLineGet(params *PurchaseOrderGoodsReceiptLineGetParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineGetOK, error)

	PurchaseOrderGoodsReceiptLineListPostList(params *PurchaseOrderGoodsReceiptLineListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineListPostListCreated, error)

	PurchaseOrderGoodsReceiptLinePost(params *PurchaseOrderGoodsReceiptLinePostParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLinePostCreated, error)

	PurchaseOrderGoodsReceiptLinePut(params *PurchaseOrderGoodsReceiptLinePutParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLinePutOK, error)

	PurchaseOrderGoodsReceiptLineSearch(params *PurchaseOrderGoodsReceiptLineSearchParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PurchaseOrderGoodsReceiptLineDelete bs e t a delete goods receipt line by ID only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLineDelete(params *PurchaseOrderGoodsReceiptLineDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLineDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLineDelete",
		Method:             "DELETE",
		PathPattern:        "/purchaseOrder/goodsReceiptLine/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLineDeleteReader{formats: a.formats},
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
  PurchaseOrderGoodsReceiptLineGet bs e t a get goods receipt line by purchase order line ID only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLineGet(params *PurchaseOrderGoodsReceiptLineGetParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLineGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLineGet",
		Method:             "GET",
		PathPattern:        "/purchaseOrder/goodsReceiptLine/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLineGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurchaseOrderGoodsReceiptLineGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PurchaseOrderGoodsReceiptLineGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseOrderGoodsReceiptLineListPostList bs e t a register multiple new goods receipt on an existing purchase order only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLineListPostList(params *PurchaseOrderGoodsReceiptLineListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLineListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLineListPostList",
		Method:             "POST",
		PathPattern:        "/purchaseOrder/goodsReceiptLine/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLineListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurchaseOrderGoodsReceiptLineListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PurchaseOrderGoodsReceiptLineListPostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseOrderGoodsReceiptLinePost bs e t a register new goods receipt new product on an existing purchase order when registration of several goods receipt use list for better performance only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLinePost(params *PurchaseOrderGoodsReceiptLinePostParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLinePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLinePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLinePost",
		Method:             "POST",
		PathPattern:        "/purchaseOrder/goodsReceiptLine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLinePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurchaseOrderGoodsReceiptLinePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PurchaseOrderGoodsReceiptLinePost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseOrderGoodsReceiptLinePut bs e t a enter goods receipt on purchase order line only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLinePut(params *PurchaseOrderGoodsReceiptLinePutParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLinePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLinePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLinePut",
		Method:             "PUT",
		PathPattern:        "/purchaseOrder/goodsReceiptLine/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLinePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurchaseOrderGoodsReceiptLinePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PurchaseOrderGoodsReceiptLinePut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseOrderGoodsReceiptLineSearch bs e t a find goods receipt lines for purchase order only available for users that have activated the logistics plus beta program in our customer account
*/
func (a *Client) PurchaseOrderGoodsReceiptLineSearch(params *PurchaseOrderGoodsReceiptLineSearchParams, authInfo runtime.ClientAuthInfoWriter) (*PurchaseOrderGoodsReceiptLineSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseOrderGoodsReceiptLineSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurchaseOrderGoodsReceiptLineSearch",
		Method:             "GET",
		PathPattern:        "/purchaseOrder/goodsReceiptLine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseOrderGoodsReceiptLineSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurchaseOrderGoodsReceiptLineSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PurchaseOrderGoodsReceiptLineSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
