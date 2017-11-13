// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repository API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repository API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeRepoTrust Change the visibility of a repository.
*/
func (a *Client) ChangeRepoTrust(params *ChangeRepoTrustParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeRepoTrustCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeRepoTrustParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeRepoTrust",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/changetrust",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeRepoTrustReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeRepoTrustCreated), nil

}

/*
ChangeRepoVisibility Change the visibility of a repository.
*/
func (a *Client) ChangeRepoVisibility(params *ChangeRepoVisibilityParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeRepoVisibilityCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeRepoVisibilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeRepoVisibility",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/changevisibility",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeRepoVisibilityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeRepoVisibilityCreated), nil

}

/*
CreateRepo Create a new repository.
*/
func (a *Client) CreateRepo(params *CreateRepoParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRepoCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRepo",
		Method:             "POST",
		PathPattern:        "/api/v1/repository",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRepoCreated), nil

}

/*
DeleteRepository Delete a repository.
*/
func (a *Client) DeleteRepository(params *DeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRepositoryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepositoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRepository",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRepositoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRepositoryNoContent), nil

}

/*
GetRepo Fetch the specified repository.
*/
func (a *Client) GetRepo(params *GetRepoParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepo",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoOK), nil

}

/*
ListRepos Fetch the list of repositories visible to the current user under a variety of situations.
*/
func (a *Client) ListRepos(params *ListReposParams, authInfo runtime.ClientAuthInfoWriter) (*ListReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepos",
		Method:             "GET",
		PathPattern:        "/api/v1/repository",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReposOK), nil

}

/*
UpdateRepo Update the description in the specified repository.
*/
func (a *Client) UpdateRepo(params *UpdateRepoParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRepo",
		Method:             "PUT",
		PathPattern:        "/api/v1/repository/{repository}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRepoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
