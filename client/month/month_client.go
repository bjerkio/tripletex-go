// Code generated by go-swagger; DO NOT EDIT.

package month

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new month API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for month API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TimesheetMonthApproveApprove(params *TimesheetMonthApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthApproveApproveOK, error)

	TimesheetMonthByMonthNumberGetByMonthNumber(params *TimesheetMonthByMonthNumberGetByMonthNumberParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthByMonthNumberGetByMonthNumberOK, error)

	TimesheetMonthCompleteComplete(params *TimesheetMonthCompleteCompleteParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthCompleteCompleteOK, error)

	TimesheetMonthGet(params *TimesheetMonthGetParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthGetOK, error)

	TimesheetMonthReopenReopen(params *TimesheetMonthReopenReopenParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthReopenReopenOK, error)

	TimesheetMonthUnapproveUnapprove(params *TimesheetMonthUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthUnapproveUnapproveOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TimesheetMonthApproveApprove approves month s if id is provided the other args are ignored
*/
func (a *Client) TimesheetMonthApproveApprove(params *TimesheetMonthApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthApproveApproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthApproveApproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthApproveApprove",
		Method:             "PUT",
		PathPattern:        "/timesheet/month/:approve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthApproveApproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthApproveApproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthApproveApprove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetMonthByMonthNumberGetByMonthNumber finds monthly status for given month
*/
func (a *Client) TimesheetMonthByMonthNumberGetByMonthNumber(params *TimesheetMonthByMonthNumberGetByMonthNumberParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthByMonthNumberGetByMonthNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthByMonthNumberGetByMonthNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthByMonthNumberGetByMonthNumber",
		Method:             "GET",
		PathPattern:        "/timesheet/month/byMonthNumber",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthByMonthNumberGetByMonthNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthByMonthNumberGetByMonthNumberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthByMonthNumberGetByMonthNumber: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetMonthCompleteComplete completes month s if id is provided the other args are ignored
*/
func (a *Client) TimesheetMonthCompleteComplete(params *TimesheetMonthCompleteCompleteParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthCompleteCompleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthCompleteCompleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthCompleteComplete",
		Method:             "PUT",
		PathPattern:        "/timesheet/month/:complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthCompleteCompleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthCompleteCompleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthCompleteComplete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetMonthGet finds monthly status entry by ID
*/
func (a *Client) TimesheetMonthGet(params *TimesheetMonthGetParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthGet",
		Method:             "GET",
		PathPattern:        "/timesheet/month/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetMonthReopenReopen reopens month s if id is provided the other args are ignored
*/
func (a *Client) TimesheetMonthReopenReopen(params *TimesheetMonthReopenReopenParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthReopenReopenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthReopenReopenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthReopenReopen",
		Method:             "PUT",
		PathPattern:        "/timesheet/month/:reopen",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthReopenReopenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthReopenReopenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthReopenReopen: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetMonthUnapproveUnapprove unapproves month s if id is provided the other args are ignored
*/
func (a *Client) TimesheetMonthUnapproveUnapprove(params *TimesheetMonthUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetMonthUnapproveUnapproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetMonthUnapproveUnapproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetMonthUnapproveUnapprove",
		Method:             "PUT",
		PathPattern:        "/timesheet/month/:unapprove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetMonthUnapproveUnapproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetMonthUnapproveUnapproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetMonthUnapproveUnapprove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
