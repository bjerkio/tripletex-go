// Code generated by go-swagger; DO NOT EDIT.

package voucher_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new voucher message API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voucher message API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	VoucherMessagePost(params *VoucherMessagePostParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherMessagePostCreated, error)

	VoucherMessageSearch(params *VoucherMessageSearchParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherMessageSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VoucherMessagePost bs e t a post new voucher message
*/
func (a *Client) VoucherMessagePost(params *VoucherMessagePostParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherMessagePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherMessagePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherMessage_post",
		Method:             "POST",
		PathPattern:        "/voucherMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherMessagePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherMessagePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherMessage_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VoucherMessageSearch bs e t a find voucher message or a comment put on a voucher by inputting voucher ids
*/
func (a *Client) VoucherMessageSearch(params *VoucherMessageSearchParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherMessageSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherMessageSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherMessage_search",
		Method:             "GET",
		PathPattern:        "/voucherMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherMessageSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherMessageSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherMessage_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}