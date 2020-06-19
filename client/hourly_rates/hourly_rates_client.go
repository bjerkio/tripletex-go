// Code generated by go-swagger; DO NOT EDIT.

package hourly_rates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hourly rates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hourly rates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectHourlyRatesDelete(params *ProjectHourlyRatesDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	ProjectHourlyRatesGet(params *ProjectHourlyRatesGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesGetOK, error)

	ProjectHourlyRatesListDeleteByIds(params *ProjectHourlyRatesListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	ProjectHourlyRatesListPostList(params *ProjectHourlyRatesListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesListPostListCreated, error)

	ProjectHourlyRatesListPutList(params *ProjectHourlyRatesListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesListPutListOK, error)

	ProjectHourlyRatesPost(params *ProjectHourlyRatesPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesPostCreated, error)

	ProjectHourlyRatesPut(params *ProjectHourlyRatesPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesPutOK, error)

	ProjectHourlyRatesSearch(params *ProjectHourlyRatesSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProjectHourlyRatesDelete deletes project hourly rate
*/
func (a *Client) ProjectHourlyRatesDelete(params *ProjectHourlyRatesDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesDelete",
		Method:             "DELETE",
		PathPattern:        "/project/hourlyRates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesDeleteReader{formats: a.formats},
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
  ProjectHourlyRatesGet finds project hourly rate by ID
*/
func (a *Client) ProjectHourlyRatesGet(params *ProjectHourlyRatesGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesGet",
		Method:             "GET",
		PathPattern:        "/project/hourlyRates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectHourlyRatesListDeleteByIds deletes project hourly rates
*/
func (a *Client) ProjectHourlyRatesListDeleteByIds(params *ProjectHourlyRatesListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesListDeleteByIds",
		Method:             "DELETE",
		PathPattern:        "/project/hourlyRates/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesListDeleteByIdsReader{formats: a.formats},
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
  ProjectHourlyRatesListPostList creates multiple project hourly rates
*/
func (a *Client) ProjectHourlyRatesListPostList(params *ProjectHourlyRatesListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesListPostList",
		Method:             "POST",
		PathPattern:        "/project/hourlyRates/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesListPostList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectHourlyRatesListPutList updates multiple project hourly rates
*/
func (a *Client) ProjectHourlyRatesListPutList(params *ProjectHourlyRatesListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesListPutList",
		Method:             "PUT",
		PathPattern:        "/project/hourlyRates/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesListPutList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectHourlyRatesPost creates a project hourly rate
*/
func (a *Client) ProjectHourlyRatesPost(params *ProjectHourlyRatesPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesPost",
		Method:             "POST",
		PathPattern:        "/project/hourlyRates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectHourlyRatesPut updates a project hourly rate
*/
func (a *Client) ProjectHourlyRatesPut(params *ProjectHourlyRatesPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesPut",
		Method:             "PUT",
		PathPattern:        "/project/hourlyRates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectHourlyRatesSearch finds project hourly rates corresponding with sent data
*/
func (a *Client) ProjectHourlyRatesSearch(params *ProjectHourlyRatesSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectHourlyRatesSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectHourlyRatesSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectHourlyRatesSearch",
		Method:             "GET",
		PathPattern:        "/project/hourlyRates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectHourlyRatesSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectHourlyRatesSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectHourlyRatesSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
