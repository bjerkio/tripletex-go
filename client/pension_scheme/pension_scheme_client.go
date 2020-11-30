// Code generated by go-swagger; DO NOT EDIT.

package pension_scheme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pension scheme API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pension scheme API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SalarySettingsPensionSchemeListDeleteByIds(params *SalarySettingsPensionSchemeListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error

	SalarySettingsPensionSchemeListPostList(params *SalarySettingsPensionSchemeListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeListPostListCreated, error)

	SalarySettingsPensionSchemeListPutList(params *SalarySettingsPensionSchemeListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeListPutListOK, error)

	SalarySettingsPensionSchemeDelete(params *SalarySettingsPensionSchemeDeleteParams, authInfo runtime.ClientAuthInfoWriter) error

	SalarySettingsPensionSchemeGet(params *SalarySettingsPensionSchemeGetParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeGetOK, error)

	SalarySettingsPensionSchemePost(params *SalarySettingsPensionSchemePostParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemePostCreated, error)

	SalarySettingsPensionSchemePut(params *SalarySettingsPensionSchemePutParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemePutOK, error)

	SalarySettingsPensionSchemeSearch(params *SalarySettingsPensionSchemeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeSearchOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SalarySettingsPensionSchemeListDeleteByIds bs e t a delete multiple pension schemes
*/
func (a *Client) SalarySettingsPensionSchemeListDeleteByIds(params *SalarySettingsPensionSchemeListDeleteByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeListDeleteByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionSchemeList_deleteByIds",
		Method:             "DELETE",
		PathPattern:        "/salary/settings/pensionScheme/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeListDeleteByIdsReader{formats: a.formats},
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
  SalarySettingsPensionSchemeListPostList bs e t a create multiple pension schemes
*/
func (a *Client) SalarySettingsPensionSchemeListPostList(params *SalarySettingsPensionSchemeListPostListParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeListPostListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeListPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionSchemeList_postList",
		Method:             "POST",
		PathPattern:        "/salary/settings/pensionScheme/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeListPostListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemeListPostListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionSchemeList_postList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalarySettingsPensionSchemeListPutList bs e t a update multiple pension schemes
*/
func (a *Client) SalarySettingsPensionSchemeListPutList(params *SalarySettingsPensionSchemeListPutListParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeListPutListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeListPutListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionSchemeList_putList",
		Method:             "PUT",
		PathPattern:        "/salary/settings/pensionScheme/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeListPutListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemeListPutListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionSchemeList_putList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalarySettingsPensionSchemeDelete bs e t a delete a pension scheme
*/
func (a *Client) SalarySettingsPensionSchemeDelete(params *SalarySettingsPensionSchemeDeleteParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeDeleteParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionScheme_delete",
		Method:             "DELETE",
		PathPattern:        "/salary/settings/pensionScheme/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeDeleteReader{formats: a.formats},
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
  SalarySettingsPensionSchemeGet bs e t a get pension scheme for a specific ID
*/
func (a *Client) SalarySettingsPensionSchemeGet(params *SalarySettingsPensionSchemeGetParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionScheme_get",
		Method:             "GET",
		PathPattern:        "/salary/settings/pensionScheme/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionScheme_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalarySettingsPensionSchemePost bs e t a create a pension scheme
*/
func (a *Client) SalarySettingsPensionSchemePost(params *SalarySettingsPensionSchemePostParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionScheme_post",
		Method:             "POST",
		PathPattern:        "/salary/settings/pensionScheme",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemePostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionScheme_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalarySettingsPensionSchemePut bs e t a update a pension scheme
*/
func (a *Client) SalarySettingsPensionSchemePut(params *SalarySettingsPensionSchemePutParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionScheme_put",
		Method:             "PUT",
		PathPattern:        "/salary/settings/pensionScheme/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionScheme_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SalarySettingsPensionSchemeSearch bs e t a find pension schemes
*/
func (a *Client) SalarySettingsPensionSchemeSearch(params *SalarySettingsPensionSchemeSearchParams, authInfo runtime.ClientAuthInfoWriter) (*SalarySettingsPensionSchemeSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSalarySettingsPensionSchemeSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SalarySettingsPensionScheme_search",
		Method:             "GET",
		PathPattern:        "/salary/settings/pensionScheme",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SalarySettingsPensionSchemeSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SalarySettingsPensionSchemeSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SalarySettingsPensionScheme_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
