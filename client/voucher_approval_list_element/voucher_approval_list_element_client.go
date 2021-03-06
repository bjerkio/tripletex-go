// Code generated by go-swagger; DO NOT EDIT.

package voucher_approval_list_element

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new voucher approval list element API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voucher approval list element API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	VoucherApprovalListElementGet(params *VoucherApprovalListElementGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherApprovalListElementGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VoucherApprovalListElementGet bs e t a get by ID
*/
func (a *Client) VoucherApprovalListElementGet(params *VoucherApprovalListElementGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoucherApprovalListElementGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoucherApprovalListElementGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoucherApprovalListElement_get",
		Method:             "GET",
		PathPattern:        "/voucherApprovalListElement/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoucherApprovalListElementGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VoucherApprovalListElementGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VoucherApprovalListElement_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
