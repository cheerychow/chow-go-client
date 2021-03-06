// Code generated by go-swagger; DO NOT EDIT.

package recipe

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

// NewGetNutritionLabelCSSParams creates a new GetNutritionLabelCSSParams object
// with the default values initialized.
func NewGetNutritionLabelCSSParams() *GetNutritionLabelCSSParams {

	return &GetNutritionLabelCSSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNutritionLabelCSSParamsWithTimeout creates a new GetNutritionLabelCSSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNutritionLabelCSSParamsWithTimeout(timeout time.Duration) *GetNutritionLabelCSSParams {

	return &GetNutritionLabelCSSParams{

		timeout: timeout,
	}
}

// NewGetNutritionLabelCSSParamsWithContext creates a new GetNutritionLabelCSSParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNutritionLabelCSSParamsWithContext(ctx context.Context) *GetNutritionLabelCSSParams {

	return &GetNutritionLabelCSSParams{

		Context: ctx,
	}
}

// NewGetNutritionLabelCSSParamsWithHTTPClient creates a new GetNutritionLabelCSSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNutritionLabelCSSParamsWithHTTPClient(client *http.Client) *GetNutritionLabelCSSParams {

	return &GetNutritionLabelCSSParams{
		HTTPClient: client,
	}
}

/*GetNutritionLabelCSSParams contains all the parameters to send to the API endpoint
for the get nutrition label Css operation typically these are written to a http.Request
*/
type GetNutritionLabelCSSParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) WithTimeout(timeout time.Duration) *GetNutritionLabelCSSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) WithContext(ctx context.Context) *GetNutritionLabelCSSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) WithHTTPClient(client *http.Client) *GetNutritionLabelCSSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nutrition label Css params
func (o *GetNutritionLabelCSSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNutritionLabelCSSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
