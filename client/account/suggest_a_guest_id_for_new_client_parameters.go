// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewSuggestAGuestIDForNewClientParams creates a new SuggestAGuestIDForNewClientParams object
// with the default values initialized.
func NewSuggestAGuestIDForNewClientParams() *SuggestAGuestIDForNewClientParams {
	var ()
	return &SuggestAGuestIDForNewClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSuggestAGuestIDForNewClientParamsWithTimeout creates a new SuggestAGuestIDForNewClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSuggestAGuestIDForNewClientParamsWithTimeout(timeout time.Duration) *SuggestAGuestIDForNewClientParams {
	var ()
	return &SuggestAGuestIDForNewClientParams{

		timeout: timeout,
	}
}

// NewSuggestAGuestIDForNewClientParamsWithContext creates a new SuggestAGuestIDForNewClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewSuggestAGuestIDForNewClientParamsWithContext(ctx context.Context) *SuggestAGuestIDForNewClientParams {
	var ()
	return &SuggestAGuestIDForNewClientParams{

		Context: ctx,
	}
}

// NewSuggestAGuestIDForNewClientParamsWithHTTPClient creates a new SuggestAGuestIDForNewClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSuggestAGuestIDForNewClientParamsWithHTTPClient(client *http.Client) *SuggestAGuestIDForNewClientParams {
	var ()
	return &SuggestAGuestIDForNewClientParams{
		HTTPClient: client,
	}
}

/*SuggestAGuestIDForNewClientParams contains all the parameters to send to the API endpoint
for the suggest a guest ID for new client operation typically these are written to a http.Request
*/
type SuggestAGuestIDForNewClientParams struct {

	/*JSONBody
	  A valid push note token.

	*/
	JSONBody *models.SuggestUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) WithTimeout(timeout time.Duration) *SuggestAGuestIDForNewClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) WithContext(ctx context.Context) *SuggestAGuestIDForNewClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) WithHTTPClient(client *http.Client) *SuggestAGuestIDForNewClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) WithJSONBody(jSONBody *models.SuggestUser) *SuggestAGuestIDForNewClientParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the suggest a guest ID for new client params
func (o *SuggestAGuestIDForNewClientParams) SetJSONBody(jSONBody *models.SuggestUser) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *SuggestAGuestIDForNewClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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