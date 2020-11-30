// Code generated by go-swagger; DO NOT EDIT.

package week

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new week API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for week API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TimesheetWeekApproveApprove(params *TimesheetWeekApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekApproveApproveOK, error)

	TimesheetWeekCompleteComplete(params *TimesheetWeekCompleteCompleteParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekCompleteCompleteOK, error)

	TimesheetWeekReopenReopen(params *TimesheetWeekReopenReopenParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekReopenReopenOK, error)

	TimesheetWeekUnapproveUnapprove(params *TimesheetWeekUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekUnapproveUnapproveOK, error)

	TimesheetWeekSearch(params *TimesheetWeekSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TimesheetWeekApproveApprove approves week by ID or i s o 8601 week and employee Id combination
*/
func (a *Client) TimesheetWeekApproveApprove(params *TimesheetWeekApproveApproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekApproveApproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetWeekApproveApproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetWeekApprove_approve",
		Method:             "PUT",
		PathPattern:        "/timesheet/week/:approve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetWeekApproveApproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetWeekApproveApproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetWeekApprove_approve: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetWeekCompleteComplete completes week by ID or i s o 8601 week and employee Id combination
*/
func (a *Client) TimesheetWeekCompleteComplete(params *TimesheetWeekCompleteCompleteParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekCompleteCompleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetWeekCompleteCompleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetWeekComplete_complete",
		Method:             "PUT",
		PathPattern:        "/timesheet/week/:complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetWeekCompleteCompleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetWeekCompleteCompleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetWeekComplete_complete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetWeekReopenReopen reopens week by ID or i s o 8601 week and employee Id combination
*/
func (a *Client) TimesheetWeekReopenReopen(params *TimesheetWeekReopenReopenParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekReopenReopenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetWeekReopenReopenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetWeekReopen_reopen",
		Method:             "PUT",
		PathPattern:        "/timesheet/week/:reopen",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetWeekReopenReopenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetWeekReopenReopenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetWeekReopen_reopen: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetWeekUnapproveUnapprove unapproves week by ID or i s o 8601 week and employee Id combination
*/
func (a *Client) TimesheetWeekUnapproveUnapprove(params *TimesheetWeekUnapproveUnapproveParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekUnapproveUnapproveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetWeekUnapproveUnapproveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetWeekUnapprove_unapprove",
		Method:             "PUT",
		PathPattern:        "/timesheet/week/:unapprove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetWeekUnapproveUnapproveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetWeekUnapproveUnapproveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetWeekUnapprove_unapprove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetWeekSearch finds weekly status by ID week year combination employee Id or an approver
*/
func (a *Client) TimesheetWeekSearch(params *TimesheetWeekSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetWeekSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetWeekSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetWeek_search",
		Method:             "GET",
		PathPattern:        "/timesheet/week",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetWeekSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetWeekSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetWeek_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
