// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectForTimeSheetGetForTimeSheet(params *ProjectForTimeSheetGetForTimeSheetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectForTimeSheetGetForTimeSheetOK, error)

	ProjectListDeleteByIds(params *ProjectListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	ProjectListPostList(params *ProjectListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectListPostListCreated, error)

	ProjectListPutList(params *ProjectListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectListPutListOK, error)

	ProjectDelete(params *ProjectDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	ProjectDeleteList(params *ProjectDeleteListParams, authInfo runtime.ClientAuthInfoWriter) error

	ProjectGet(params *ProjectGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetOK, error)

	ProjectPost(params *ProjectPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectPostCreated, error)

	ProjectPut(params *ProjectPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectPutOK, error)

	ProjectSearch(params *ProjectSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProjectForTimeSheetGetForTimeSheet finds projects applicable for time sheet registration on a specific day
*/
func (a *Client) ProjectForTimeSheetGetForTimeSheet(params *ProjectForTimeSheetGetForTimeSheetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectForTimeSheetGetForTimeSheetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectForTimeSheetGetForTimeSheetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectForTimeSheet_getForTimeSheet",
		Method:             "GET",
		PathPattern:        "/project/>forTimeSheet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectForTimeSheetGetForTimeSheetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectForTimeSheetGetForTimeSheetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectForTimeSheet_getForTimeSheet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectListDeleteByIds bs e t a delete projects
*/
func (a *Client) ProjectListDeleteByIds(params *ProjectListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectList_deleteByIds",
		Method:             "DELETE",
		PathPattern:        "/project/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectListDeleteByIdsReader{formats: a.formats},
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
  ProjectListPostList bs e t a register new projects multiple projects for different users can be sent in the same request
*/
func (a *Client) ProjectListPostList(params *ProjectListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectList_postList",
		Method:             "POST",
		PathPattern:        "/project/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectListPutList bs e t a update multiple projects
*/
func (a *Client) ProjectListPutList(params *ProjectListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectList_putList",
		Method:             "PUT",
		PathPattern:        "/project/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectDelete bs e t a delete project
*/
func (a *Client) ProjectDelete(params *ProjectDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_delete",
		Method:             "DELETE",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectDeleteReader{formats: a.formats},
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
  ProjectDeleteList bs e t a delete multiple projects
*/
func (a *Client) ProjectDeleteList(params *ProjectDeleteListParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDeleteListParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_deleteList",
		Method:             "DELETE",
		PathPattern:        "/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectDeleteListReader{formats: a.formats},
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
  ProjectGet finds project by ID
*/
func (a *Client) ProjectGet(params *ProjectGetParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_get",
		Method:             "GET",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Project_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectPost bs e t a add new project
*/
func (a *Client) ProjectPost(params *ProjectPostParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_post",
		Method:             "POST",
		PathPattern:        "/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Project_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectPut bs e t a update project
*/
func (a *Client) ProjectPut(params *ProjectPutParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_put",
		Method:             "PUT",
		PathPattern:        "/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Project_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectSearch finds projects corresponding with sent data
*/
func (a *Client) ProjectSearch(params *ProjectSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Project_search",
		Method:             "GET",
		PathPattern:        "/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Project_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
