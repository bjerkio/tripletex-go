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

package leave_of_absence_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new leave of absence type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for leave of absence type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	EmployeeEmploymentLeaveOfAbsenceTypeSearch(params *EmployeeEmploymentLeaveOfAbsenceTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeEmploymentLeaveOfAbsenceTypeSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EmployeeEmploymentLeaveOfAbsenceTypeSearch bs e t a find all leave of absence type i ds
*/
func (a *Client) EmployeeEmploymentLeaveOfAbsenceTypeSearch(params *EmployeeEmploymentLeaveOfAbsenceTypeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*EmployeeEmploymentLeaveOfAbsenceTypeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmployeeEmploymentLeaveOfAbsenceTypeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmployeeEmploymentLeaveOfAbsenceTypeSearch",
		Method:             "GET",
		PathPattern:        "/employee/employment/leaveOfAbsenceType",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmployeeEmploymentLeaveOfAbsenceTypeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmployeeEmploymentLeaveOfAbsenceTypeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EmployeeEmploymentLeaveOfAbsenceTypeSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
