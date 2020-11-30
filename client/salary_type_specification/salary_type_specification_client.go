// Code generated by go-swagger; DO NOT EDIT.

package salary_type_specification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new salary type specification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for salary type specification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TimesheetSalaryTypeSpecificationDelete(params *TimesheetSalaryTypeSpecificationDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	TimesheetSalaryTypeSpecificationGet(params *TimesheetSalaryTypeSpecificationGetParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationGetOK, error)

	TimesheetSalaryTypeSpecificationPost(params *TimesheetSalaryTypeSpecificationPostParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationPostCreated, error)

	TimesheetSalaryTypeSpecificationPut(params *TimesheetSalaryTypeSpecificationPutParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationPutOK, error)

	TimesheetSalaryTypeSpecificationSearch(params *TimesheetSalaryTypeSpecificationSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TimesheetSalaryTypeSpecificationDelete bs e t a delete a timesheet salary type specification
*/
func (a *Client) TimesheetSalaryTypeSpecificationDelete(params *TimesheetSalaryTypeSpecificationDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetSalaryTypeSpecificationDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetSalaryTypeSpecification_delete",
		Method:             "DELETE",
		PathPattern:        "/timesheet/salaryTypeSpecification/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetSalaryTypeSpecificationDeleteReader{formats: a.formats},
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
  TimesheetSalaryTypeSpecificationGet bs e t a get timesheet salary type specification for a specific ID
*/
func (a *Client) TimesheetSalaryTypeSpecificationGet(params *TimesheetSalaryTypeSpecificationGetParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetSalaryTypeSpecificationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetSalaryTypeSpecification_get",
		Method:             "GET",
		PathPattern:        "/timesheet/salaryTypeSpecification/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetSalaryTypeSpecificationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetSalaryTypeSpecificationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetSalaryTypeSpecification_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetSalaryTypeSpecificationPost bs e t a create a timesheet salary type specification only one entry per employee date salary type
*/
func (a *Client) TimesheetSalaryTypeSpecificationPost(params *TimesheetSalaryTypeSpecificationPostParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetSalaryTypeSpecificationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetSalaryTypeSpecification_post",
		Method:             "POST",
		PathPattern:        "/timesheet/salaryTypeSpecification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetSalaryTypeSpecificationPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetSalaryTypeSpecificationPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetSalaryTypeSpecification_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetSalaryTypeSpecificationPut bs e t a update a timesheet salary type specification
*/
func (a *Client) TimesheetSalaryTypeSpecificationPut(params *TimesheetSalaryTypeSpecificationPutParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetSalaryTypeSpecificationPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetSalaryTypeSpecification_put",
		Method:             "PUT",
		PathPattern:        "/timesheet/salaryTypeSpecification/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetSalaryTypeSpecificationPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetSalaryTypeSpecificationPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetSalaryTypeSpecification_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TimesheetSalaryTypeSpecificationSearch bs e t a get list of timesheet salary type specifications
*/
func (a *Client) TimesheetSalaryTypeSpecificationSearch(params *TimesheetSalaryTypeSpecificationSearchParams, authInfo runtime.ClientAuthInfoWriter) (*TimesheetSalaryTypeSpecificationSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTimesheetSalaryTypeSpecificationSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TimesheetSalaryTypeSpecification_search",
		Method:             "GET",
		PathPattern:        "/timesheet/salaryTypeSpecification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TimesheetSalaryTypeSpecificationSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TimesheetSalaryTypeSpecificationSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TimesheetSalaryTypeSpecification_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
