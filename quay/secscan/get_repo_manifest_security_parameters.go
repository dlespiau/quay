// Code generated by go-swagger; DO NOT EDIT.

package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRepoManifestSecurityParams creates a new GetRepoManifestSecurityParams object
// with the default values initialized.
func NewGetRepoManifestSecurityParams() *GetRepoManifestSecurityParams {
	var ()
	return &GetRepoManifestSecurityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoManifestSecurityParamsWithTimeout creates a new GetRepoManifestSecurityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepoManifestSecurityParamsWithTimeout(timeout time.Duration) *GetRepoManifestSecurityParams {
	var ()
	return &GetRepoManifestSecurityParams{

		timeout: timeout,
	}
}

// NewGetRepoManifestSecurityParamsWithContext creates a new GetRepoManifestSecurityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepoManifestSecurityParamsWithContext(ctx context.Context) *GetRepoManifestSecurityParams {
	var ()
	return &GetRepoManifestSecurityParams{

		Context: ctx,
	}
}

// NewGetRepoManifestSecurityParamsWithHTTPClient creates a new GetRepoManifestSecurityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepoManifestSecurityParamsWithHTTPClient(client *http.Client) *GetRepoManifestSecurityParams {
	var ()
	return &GetRepoManifestSecurityParams{
		HTTPClient: client,
	}
}

/*GetRepoManifestSecurityParams contains all the parameters to send to the API endpoint
for the get repo manifest security operation typically these are written to a http.Request
*/
type GetRepoManifestSecurityParams struct {

	/*Manifestref
	  The digest of the manifest

	*/
	Manifestref string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Vulnerabilities
	  Include vulnerabilities informations

	*/
	Vulnerabilities *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithTimeout(timeout time.Duration) *GetRepoManifestSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithContext(ctx context.Context) *GetRepoManifestSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithHTTPClient(client *http.Client) *GetRepoManifestSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManifestref adds the manifestref to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithManifestref(manifestref string) *GetRepoManifestSecurityParams {
	o.SetManifestref(manifestref)
	return o
}

// SetManifestref adds the manifestref to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetManifestref(manifestref string) {
	o.Manifestref = manifestref
}

// WithRepository adds the repository to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithRepository(repository string) *GetRepoManifestSecurityParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithVulnerabilities adds the vulnerabilities to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) WithVulnerabilities(vulnerabilities *bool) *GetRepoManifestSecurityParams {
	o.SetVulnerabilities(vulnerabilities)
	return o
}

// SetVulnerabilities adds the vulnerabilities to the get repo manifest security params
func (o *GetRepoManifestSecurityParams) SetVulnerabilities(vulnerabilities *bool) {
	o.Vulnerabilities = vulnerabilities
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoManifestSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param manifestref
	if err := r.SetPathParam("manifestref", o.Manifestref); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if o.Vulnerabilities != nil {

		// query param vulnerabilities
		var qrVulnerabilities bool
		if o.Vulnerabilities != nil {
			qrVulnerabilities = *o.Vulnerabilities
		}
		qVulnerabilities := swag.FormatBool(qrVulnerabilities)
		if qVulnerabilities != "" {
			if err := r.SetQueryParam("vulnerabilities", qVulnerabilities); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
