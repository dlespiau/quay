// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewListUserLogsParams creates a new ListUserLogsParams object
// with the default values initialized.
func NewListUserLogsParams() *ListUserLogsParams {
	var ()
	return &ListUserLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUserLogsParamsWithTimeout creates a new ListUserLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUserLogsParamsWithTimeout(timeout time.Duration) *ListUserLogsParams {
	var ()
	return &ListUserLogsParams{

		timeout: timeout,
	}
}

// NewListUserLogsParamsWithContext creates a new ListUserLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUserLogsParamsWithContext(ctx context.Context) *ListUserLogsParams {
	var ()
	return &ListUserLogsParams{

		Context: ctx,
	}
}

// NewListUserLogsParamsWithHTTPClient creates a new ListUserLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUserLogsParamsWithHTTPClient(client *http.Client) *ListUserLogsParams {
	var ()
	return &ListUserLogsParams{
		HTTPClient: client,
	}
}

/*ListUserLogsParams contains all the parameters to send to the API endpoint
for the list user logs operation typically these are written to a http.Request
*/
type ListUserLogsParams struct {

	/*Endtime
	  Latest time to which to get logs. (%m/%d/%Y %Z)

	*/
	Endtime *string
	/*NextPage
	  The page token for the next page

	*/
	NextPage *string
	/*Performer
	  Username for which to filter logs.

	*/
	Performer *string
	/*Starttime
	  Earliest time from which to get logs. (%m/%d/%Y %Z)

	*/
	Starttime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list user logs params
func (o *ListUserLogsParams) WithTimeout(timeout time.Duration) *ListUserLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user logs params
func (o *ListUserLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user logs params
func (o *ListUserLogsParams) WithContext(ctx context.Context) *ListUserLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user logs params
func (o *ListUserLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user logs params
func (o *ListUserLogsParams) WithHTTPClient(client *http.Client) *ListUserLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user logs params
func (o *ListUserLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndtime adds the endtime to the list user logs params
func (o *ListUserLogsParams) WithEndtime(endtime *string) *ListUserLogsParams {
	o.SetEndtime(endtime)
	return o
}

// SetEndtime adds the endtime to the list user logs params
func (o *ListUserLogsParams) SetEndtime(endtime *string) {
	o.Endtime = endtime
}

// WithNextPage adds the nextPage to the list user logs params
func (o *ListUserLogsParams) WithNextPage(nextPage *string) *ListUserLogsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the list user logs params
func (o *ListUserLogsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPerformer adds the performer to the list user logs params
func (o *ListUserLogsParams) WithPerformer(performer *string) *ListUserLogsParams {
	o.SetPerformer(performer)
	return o
}

// SetPerformer adds the performer to the list user logs params
func (o *ListUserLogsParams) SetPerformer(performer *string) {
	o.Performer = performer
}

// WithStarttime adds the starttime to the list user logs params
func (o *ListUserLogsParams) WithStarttime(starttime *string) *ListUserLogsParams {
	o.SetStarttime(starttime)
	return o
}

// SetStarttime adds the starttime to the list user logs params
func (o *ListUserLogsParams) SetStarttime(starttime *string) {
	o.Starttime = starttime
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime string
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

	}

	if o.NextPage != nil {

		// query param next_page
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("next_page", qNextPage); err != nil {
				return err
			}
		}

	}

	if o.Performer != nil {

		// query param performer
		var qrPerformer string
		if o.Performer != nil {
			qrPerformer = *o.Performer
		}
		qPerformer := qrPerformer
		if qPerformer != "" {
			if err := r.SetQueryParam("performer", qPerformer); err != nil {
				return err
			}
		}

	}

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime string
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
