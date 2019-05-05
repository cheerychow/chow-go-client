// Code generated by go-swagger; DO NOT EDIT.

package user

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

	models "github.com/cheerychow/go-client/models"
)

// NewSaveUserMetaInfoParams creates a new SaveUserMetaInfoParams object
// with the default values initialized.
func NewSaveUserMetaInfoParams() *SaveUserMetaInfoParams {
	var ()
	return &SaveUserMetaInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveUserMetaInfoParamsWithTimeout creates a new SaveUserMetaInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveUserMetaInfoParamsWithTimeout(timeout time.Duration) *SaveUserMetaInfoParams {
	var ()
	return &SaveUserMetaInfoParams{

		timeout: timeout,
	}
}

// NewSaveUserMetaInfoParamsWithContext creates a new SaveUserMetaInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveUserMetaInfoParamsWithContext(ctx context.Context) *SaveUserMetaInfoParams {
	var ()
	return &SaveUserMetaInfoParams{

		Context: ctx,
	}
}

// NewSaveUserMetaInfoParamsWithHTTPClient creates a new SaveUserMetaInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveUserMetaInfoParamsWithHTTPClient(client *http.Client) *SaveUserMetaInfoParams {
	var ()
	return &SaveUserMetaInfoParams{
		HTTPClient: client,
	}
}

/*SaveUserMetaInfoParams contains all the parameters to send to the API endpoint
for the save user meta info operation typically these are written to a http.Request
*/
type SaveUserMetaInfoParams struct {

	/*JSONBody
	  A user object containing the changes to be made.

	*/
	JSONBody []*models.UserMeta

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save user meta info params
func (o *SaveUserMetaInfoParams) WithTimeout(timeout time.Duration) *SaveUserMetaInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save user meta info params
func (o *SaveUserMetaInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save user meta info params
func (o *SaveUserMetaInfoParams) WithContext(ctx context.Context) *SaveUserMetaInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save user meta info params
func (o *SaveUserMetaInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save user meta info params
func (o *SaveUserMetaInfoParams) WithHTTPClient(client *http.Client) *SaveUserMetaInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save user meta info params
func (o *SaveUserMetaInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the save user meta info params
func (o *SaveUserMetaInfoParams) WithJSONBody(jSONBody []*models.UserMeta) *SaveUserMetaInfoParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the save user meta info params
func (o *SaveUserMetaInfoParams) SetJSONBody(jSONBody []*models.UserMeta) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *SaveUserMetaInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}