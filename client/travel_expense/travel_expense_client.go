// Code generated by go-swagger; DO NOT EDIT.

package travel_expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new travel expense API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for travel expense API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TravelExpenseApproveApprove(params *TravelExpenseApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseApproveApproveOK, error)

	TravelExpenseAttachmentDeleteAttachment(params *TravelExpenseAttachmentDeleteAttachmentParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseAttachmentDownloadAttachment(params *TravelExpenseAttachmentDownloadAttachmentParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAttachmentDownloadAttachmentOK, error)

	TravelExpenseAttachmentListUploadAttachments(params *TravelExpenseAttachmentListUploadAttachmentsParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseAttachmentUploadAttachment(params *TravelExpenseAttachmentUploadAttachmentParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseCopyCopy(params *TravelExpenseCopyCopyParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCopyCopyOK, error)

	TravelExpenseCreateVouchersCreateVouchers(params *TravelExpenseCreateVouchersCreateVouchersParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCreateVouchersCreateVouchersOK, error)

	TravelExpenseDelete(params *TravelExpenseDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	TravelExpenseDeliverDeliver(params *TravelExpenseDeliverDeliverParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseDeliverDeliverOK, error)

	TravelExpenseGet(params *TravelExpenseGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseGetOK, error)

	TravelExpensePost(params *TravelExpensePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePostCreated, error)

	TravelExpensePut(params *TravelExpensePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePutOK, error)

	TravelExpenseSearch(params *TravelExpenseSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseSearchOK, error)

	TravelExpenseUnapproveUnapprove(params *TravelExpenseUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseUnapproveUnapproveOK, error)

	TravelExpenseUndeliverUndeliver(params *TravelExpenseUndeliverUndeliverParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseUndeliverUndeliverOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TravelExpenseApproveApprove bs e t a approve travel expenses
*/
func (a *Client) TravelExpenseApproveApprove(params *TravelExpenseApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseApproveApproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseApproveApproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseApproveApprove",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:approve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseApproveApproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseApproveApproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseApproveApprove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseAttachmentDeleteAttachment bs e t a delete attachment
*/
func (a *Client) TravelExpenseAttachmentDeleteAttachment(params *TravelExpenseAttachmentDeleteAttachmentParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAttachmentDeleteAttachmentParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAttachmentDeleteAttachment",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/{travelExpenseId}/attachment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAttachmentDeleteAttachmentReader{formats: a.formats},
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
  TravelExpenseAttachmentDownloadAttachment gets attachment by travel expense ID
*/
func (a *Client) TravelExpenseAttachmentDownloadAttachment(params *TravelExpenseAttachmentDownloadAttachmentParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseAttachmentDownloadAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAttachmentDownloadAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAttachmentDownloadAttachment",
		Method:             "GET",
		PathPattern:        "/travelExpense/{travelExpenseId}/attachment",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAttachmentDownloadAttachmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseAttachmentDownloadAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseAttachmentDownloadAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseAttachmentListUploadAttachments uploads multiple attachments to travel expense
*/
func (a *Client) TravelExpenseAttachmentListUploadAttachments(params *TravelExpenseAttachmentListUploadAttachmentsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAttachmentListUploadAttachmentsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAttachmentListUploadAttachments",
		Method:             "POST",
		PathPattern:        "/travelExpense/{travelExpenseId}/attachment/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAttachmentListUploadAttachmentsReader{formats: a.formats},
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
  TravelExpenseAttachmentUploadAttachment uploads attachment to travel expense
*/
func (a *Client) TravelExpenseAttachmentUploadAttachment(params *TravelExpenseAttachmentUploadAttachmentParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseAttachmentUploadAttachmentParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseAttachmentUploadAttachment",
		Method:             "POST",
		PathPattern:        "/travelExpense/{travelExpenseId}/attachment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseAttachmentUploadAttachmentReader{formats: a.formats},
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
  TravelExpenseCopyCopy bs e t a copy travel expense
*/
func (a *Client) TravelExpenseCopyCopy(params *TravelExpenseCopyCopyParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCopyCopyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCopyCopyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCopyCopy",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCopyCopyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCopyCopyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCopyCopy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseCreateVouchersCreateVouchers bs e t a create vouchers
*/
func (a *Client) TravelExpenseCreateVouchersCreateVouchers(params *TravelExpenseCreateVouchersCreateVouchersParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseCreateVouchersCreateVouchersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseCreateVouchersCreateVouchersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseCreateVouchersCreateVouchers",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:createVouchers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseCreateVouchersCreateVouchersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseCreateVouchersCreateVouchersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseCreateVouchersCreateVouchers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseDelete bs e t a delete travel expense
*/
func (a *Client) TravelExpenseDelete(params *TravelExpenseDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseDelete",
		Method:             "DELETE",
		PathPattern:        "/travelExpense/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseDeleteReader{formats: a.formats},
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
  TravelExpenseDeliverDeliver bs e t a deliver travel expenses
*/
func (a *Client) TravelExpenseDeliverDeliver(params *TravelExpenseDeliverDeliverParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseDeliverDeliverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseDeliverDeliverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseDeliverDeliver",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:deliver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseDeliverDeliverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseDeliverDeliverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseDeliverDeliver: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseGet bs e t a get travel expense by ID
*/
func (a *Client) TravelExpenseGet(params *TravelExpenseGetParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseGet",
		Method:             "GET",
		PathPattern:        "/travelExpense/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpensePost bs e t a create travel expense
*/
func (a *Client) TravelExpensePost(params *TravelExpensePostParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpensePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpensePost",
		Method:             "POST",
		PathPattern:        "/travelExpense",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpensePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpensePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpensePost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpensePut bs e t a update travel expense
*/
func (a *Client) TravelExpensePut(params *TravelExpensePutParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpensePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpensePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpensePut",
		Method:             "PUT",
		PathPattern:        "/travelExpense/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpensePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpensePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpensePut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseSearch bs e t a find travel expenses corresponding with sent data
*/
func (a *Client) TravelExpenseSearch(params *TravelExpenseSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseSearch",
		Method:             "GET",
		PathPattern:        "/travelExpense",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseUnapproveUnapprove bs e t a unapprove travel expenses
*/
func (a *Client) TravelExpenseUnapproveUnapprove(params *TravelExpenseUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseUnapproveUnapproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseUnapproveUnapproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseUnapproveUnapprove",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:unapprove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseUnapproveUnapproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseUnapproveUnapproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseUnapproveUnapprove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TravelExpenseUndeliverUndeliver bs e t a undeliver travel expenses
*/
func (a *Client) TravelExpenseUndeliverUndeliver(params *TravelExpenseUndeliverUndeliverParams, authInfo runtime.ClientAuthInfoWriter) (*TravelExpenseUndeliverUndeliverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTravelExpenseUndeliverUndeliverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TravelExpenseUndeliverUndeliver",
		Method:             "PUT",
		PathPattern:        "/travelExpense/:undeliver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TravelExpenseUndeliverUndeliverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TravelExpenseUndeliverUndeliverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TravelExpenseUndeliverUndeliver: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
