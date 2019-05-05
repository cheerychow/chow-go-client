// Code generated by go-swagger; DO NOT EDIT.

package food

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

// NewShowSearchRecipeCategoriesParams creates a new ShowSearchRecipeCategoriesParams object
// with the default values initialized.
func NewShowSearchRecipeCategoriesParams() *ShowSearchRecipeCategoriesParams {

	return &ShowSearchRecipeCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowSearchRecipeCategoriesParamsWithTimeout creates a new ShowSearchRecipeCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowSearchRecipeCategoriesParamsWithTimeout(timeout time.Duration) *ShowSearchRecipeCategoriesParams {

	return &ShowSearchRecipeCategoriesParams{

		timeout: timeout,
	}
}

// NewShowSearchRecipeCategoriesParamsWithContext creates a new ShowSearchRecipeCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowSearchRecipeCategoriesParamsWithContext(ctx context.Context) *ShowSearchRecipeCategoriesParams {

	return &ShowSearchRecipeCategoriesParams{

		Context: ctx,
	}
}

// NewShowSearchRecipeCategoriesParamsWithHTTPClient creates a new ShowSearchRecipeCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowSearchRecipeCategoriesParamsWithHTTPClient(client *http.Client) *ShowSearchRecipeCategoriesParams {

	return &ShowSearchRecipeCategoriesParams{
		HTTPClient: client,
	}
}

/*ShowSearchRecipeCategoriesParams contains all the parameters to send to the API endpoint
for the show search recipe categories operation typically these are written to a http.Request
*/
type ShowSearchRecipeCategoriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) WithTimeout(timeout time.Duration) *ShowSearchRecipeCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) WithContext(ctx context.Context) *ShowSearchRecipeCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) WithHTTPClient(client *http.Client) *ShowSearchRecipeCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show search recipe categories params
func (o *ShowSearchRecipeCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ShowSearchRecipeCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}