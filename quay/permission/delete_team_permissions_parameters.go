// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteTeamPermissionsParams creates a new DeleteTeamPermissionsParams object
// with the default values initialized.
func NewDeleteTeamPermissionsParams() *DeleteTeamPermissionsParams {
	var ()
	return &DeleteTeamPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamPermissionsParamsWithTimeout creates a new DeleteTeamPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTeamPermissionsParamsWithTimeout(timeout time.Duration) *DeleteTeamPermissionsParams {
	var ()
	return &DeleteTeamPermissionsParams{

		timeout: timeout,
	}
}

// NewDeleteTeamPermissionsParamsWithContext creates a new DeleteTeamPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTeamPermissionsParamsWithContext(ctx context.Context) *DeleteTeamPermissionsParams {
	var ()
	return &DeleteTeamPermissionsParams{

		Context: ctx,
	}
}

// NewDeleteTeamPermissionsParamsWithHTTPClient creates a new DeleteTeamPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTeamPermissionsParamsWithHTTPClient(client *http.Client) *DeleteTeamPermissionsParams {
	var ()
	return &DeleteTeamPermissionsParams{
		HTTPClient: client,
	}
}

/*DeleteTeamPermissionsParams contains all the parameters to send to the API endpoint
for the delete team permissions operation typically these are written to a http.Request
*/
type DeleteTeamPermissionsParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Teamname
	  The name of the team to which the permission applies

	*/
	Teamname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete team permissions params
func (o *DeleteTeamPermissionsParams) WithTimeout(timeout time.Duration) *DeleteTeamPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team permissions params
func (o *DeleteTeamPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team permissions params
func (o *DeleteTeamPermissionsParams) WithContext(ctx context.Context) *DeleteTeamPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team permissions params
func (o *DeleteTeamPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team permissions params
func (o *DeleteTeamPermissionsParams) WithHTTPClient(client *http.Client) *DeleteTeamPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team permissions params
func (o *DeleteTeamPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the delete team permissions params
func (o *DeleteTeamPermissionsParams) WithRepository(repository string) *DeleteTeamPermissionsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the delete team permissions params
func (o *DeleteTeamPermissionsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTeamname adds the teamname to the delete team permissions params
func (o *DeleteTeamPermissionsParams) WithTeamname(teamname string) *DeleteTeamPermissionsParams {
	o.SetTeamname(teamname)
	return o
}

// SetTeamname adds the teamname to the delete team permissions params
func (o *DeleteTeamPermissionsParams) SetTeamname(teamname string) {
	o.Teamname = teamname
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param teamname
	if err := r.SetPathParam("teamname", o.Teamname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
