// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewConductRepoSearchParams creates a new ConductRepoSearchParams object
// with the default values initialized.
func NewConductRepoSearchParams() *ConductRepoSearchParams {
	var ()
	return &ConductRepoSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConductRepoSearchParamsWithTimeout creates a new ConductRepoSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConductRepoSearchParamsWithTimeout(timeout time.Duration) *ConductRepoSearchParams {
	var ()
	return &ConductRepoSearchParams{

		timeout: timeout,
	}
}

// NewConductRepoSearchParamsWithContext creates a new ConductRepoSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewConductRepoSearchParamsWithContext(ctx context.Context) *ConductRepoSearchParams {
	var ()
	return &ConductRepoSearchParams{

		Context: ctx,
	}
}

// NewConductRepoSearchParamsWithHTTPClient creates a new ConductRepoSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConductRepoSearchParamsWithHTTPClient(client *http.Client) *ConductRepoSearchParams {
	var ()
	return &ConductRepoSearchParams{
		HTTPClient: client,
	}
}

/*ConductRepoSearchParams contains all the parameters to send to the API endpoint
for the conduct repo search operation typically these are written to a http.Request
*/
type ConductRepoSearchParams struct {

	/*Page
	  The page.

	*/
	Page *int64
	/*Query
	  The search query.

	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the conduct repo search params
func (o *ConductRepoSearchParams) WithTimeout(timeout time.Duration) *ConductRepoSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the conduct repo search params
func (o *ConductRepoSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the conduct repo search params
func (o *ConductRepoSearchParams) WithContext(ctx context.Context) *ConductRepoSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the conduct repo search params
func (o *ConductRepoSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the conduct repo search params
func (o *ConductRepoSearchParams) WithHTTPClient(client *http.Client) *ConductRepoSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the conduct repo search params
func (o *ConductRepoSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the conduct repo search params
func (o *ConductRepoSearchParams) WithPage(page *int64) *ConductRepoSearchParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the conduct repo search params
func (o *ConductRepoSearchParams) SetPage(page *int64) {
	o.Page = page
}

// WithQuery adds the query to the conduct repo search params
func (o *ConductRepoSearchParams) WithQuery(query *string) *ConductRepoSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the conduct repo search params
func (o *ConductRepoSearchParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ConductRepoSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}