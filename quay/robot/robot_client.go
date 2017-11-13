// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new robot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for robot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateOrgRobot Create a new robot in the organization.
*/
func (a *Client) CreateOrgRobot(params *CreateOrgRobotParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrgRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrgRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrgRobot",
		Method:             "PUT",
		PathPattern:        "/api/v1/organization/{orgname}/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrgRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrgRobotOK), nil

}

/*
CreateUserRobot Create a new user robot with the specified name.
*/
func (a *Client) CreateUserRobot(params *CreateUserRobotParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUserRobot",
		Method:             "PUT",
		PathPattern:        "/api/v1/user/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserRobotOK), nil

}

/*
DeleteOrgRobot Delete an existing organization robot.
*/
func (a *Client) DeleteOrgRobot(params *DeleteOrgRobotParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrgRobotNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrgRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrgRobot",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organization/{orgname}/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrgRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrgRobotNoContent), nil

}

/*
DeleteUserRobot Delete an existing robot.
*/
func (a *Client) DeleteUserRobot(params *DeleteUserRobotParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserRobotNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserRobot",
		Method:             "DELETE",
		PathPattern:        "/api/v1/user/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserRobotNoContent), nil

}

/*
GetOrgRobot Returns the organization's robot with the specified name.
*/
func (a *Client) GetOrgRobot(params *GetOrgRobotParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrgRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrgRobot",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotOK), nil

}

/*
GetOrgRobotPermissions Returns the list of repository permissions for the org's robot.
*/
func (a *Client) GetOrgRobotPermissions(params *GetOrgRobotPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrgRobotPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgRobotPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrgRobotPermissions",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/robots/{robot_shortname}/permissions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgRobotPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotPermissionsOK), nil

}

/*
GetOrgRobots List the organization's robots.
*/
func (a *Client) GetOrgRobots(params *GetOrgRobotsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrgRobotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgRobotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrgRobots",
		Method:             "GET",
		PathPattern:        "/api/v1/organization/{orgname}/robots",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgRobotsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrgRobotsOK), nil

}

/*
GetUserRobot Returns the user's robot with the specified name.
*/
func (a *Client) GetUserRobot(params *GetUserRobotParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserRobotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserRobotParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserRobot",
		Method:             "GET",
		PathPattern:        "/api/v1/user/robots/{robot_shortname}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserRobotReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotOK), nil

}

/*
GetUserRobotPermissions Returns the list of repository permissions for the user's robot.
*/
func (a *Client) GetUserRobotPermissions(params *GetUserRobotPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserRobotPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserRobotPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserRobotPermissions",
		Method:             "GET",
		PathPattern:        "/api/v1/user/robots/{robot_shortname}/permissions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserRobotPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotPermissionsOK), nil

}

/*
GetUserRobots List the available robots for the user.
*/
func (a *Client) GetUserRobots(params *GetUserRobotsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserRobotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserRobotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserRobots",
		Method:             "GET",
		PathPattern:        "/api/v1/user/robots",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserRobotsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRobotsOK), nil

}

/*
RegenerateOrgRobotToken Regenerates the token for an organization robot.
*/
func (a *Client) RegenerateOrgRobotToken(params *RegenerateOrgRobotTokenParams, authInfo runtime.ClientAuthInfoWriter) (*RegenerateOrgRobotTokenCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenerateOrgRobotTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "regenerateOrgRobotToken",
		Method:             "POST",
		PathPattern:        "/api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegenerateOrgRobotTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegenerateOrgRobotTokenCreated), nil

}

/*
RegenerateUserRobotToken Regenerates the token for a user's robot.
*/
func (a *Client) RegenerateUserRobotToken(params *RegenerateUserRobotTokenParams, authInfo runtime.ClientAuthInfoWriter) (*RegenerateUserRobotTokenCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenerateUserRobotTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "regenerateUserRobotToken",
		Method:             "POST",
		PathPattern:        "/api/v1/user/robots/{robot_shortname}/regenerate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegenerateUserRobotTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegenerateUserRobotTokenCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
