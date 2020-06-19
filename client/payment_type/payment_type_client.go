// Code generated by go-swagger; DO NOT EDIT.

package payment_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BankReconciliationPaymentTypeGet(params *BankReconciliationPaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationPaymentTypeGetOK, error)

	BankReconciliationPaymentTypeSearch(params *BankReconciliationPaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationPaymentTypeSearchOK, error)

	InvoicePaymentTypeGet(params *InvoicePaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*InvoicePaymentTypeGetOK, error)

	InvoicePaymentTypeSearch(params *InvoicePaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*InvoicePaymentTypeSearchOK, error)

	TravelExpensePaymentTypeGet(params *TravelExpensePaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePaymentTypeGetOK, error)

	TravelExpensePaymentTypeSearch(params *TravelExpensePaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePaymentTypeSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BankReconciliationPaymentTypeGet bs e t a get payment type by ID
*/
func (a *Client) BankReconciliationPaymentTypeGet(params *BankReconciliationPaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationPaymentTypeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationPaymentTypeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationPaymentTypeGet",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/paymentType/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationPaymentTypeGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationPaymentTypeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationPaymentTypeGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BankReconciliationPaymentTypeSearch bs e t a find payment type corresponding with sent data
*/
func (a *Client) BankReconciliationPaymentTypeSearch(params *BankReconciliationPaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*BankReconciliationPaymentTypeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBankReconciliationPaymentTypeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BankReconciliationPaymentTypeSearch",
		Method:             "GET",
		PathPattern:        "/bank/reconciliation/paymentType",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BankReconciliationPaymentTypeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BankReconciliationPaymentTypeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BankReconciliationPaymentTypeSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvoicePaymentTypeGet gets payment type by ID
*/
func (a *Client) InvoicePaymentTypeGet(params *InvoicePaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*InvoicePaymentTypeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvoicePaymentTypeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InvoicePaymentTypeGet",
		Method:             "GET",
		PathPattern:        "/invoice/paymentType/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InvoicePaymentTypeGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvoicePaymentTypeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InvoicePaymentTypeGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvoicePaymentTypeSearch finds payment type corresponding with sent data
*/
func (a *Client) InvoicePaymentTypeSearch(params *InvoicePaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*InvoicePaymentTypeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvoicePaymentTypeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InvoicePaymentTypeSearch",
		Method:             "GET",
		PathPattern:        "/invoice/paymentType",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InvoicePaymentTypeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvoicePaymentTypeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InvoicePaymentTypeSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpensePaymentTypeGet bs e t a get payment type by ID
*/
func (a *Client) TravelExpensePaymentTypeGet(params *TravelExpensePaymentTypeGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePaymentTypeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpensePaymentTypeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpensePaymentTypeGet",
		Method:             "GET",
		PathPattern:        "/travelExpense/paymentType/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpensePaymentTypeGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpensePaymentTypeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpensePaymentTypeGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpensePaymentTypeSearch bs e t a find payment type corresponding with sent data
*/
func (a *Client) TravelExpensePaymentTypeSearch(params *TravelExpensePaymentTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePaymentTypeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpensePaymentTypeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpensePaymentTypeSearch",
		Method:             "GET",
		PathPattern:        "/travelExpense/paymentType",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpensePaymentTypeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpensePaymentTypeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpensePaymentTypeSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
