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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNutritionLabelHTMLStaticSnippetParams creates a new GetNutritionLabelHTMLStaticSnippetParams object
// with the default values initialized.
func NewGetNutritionLabelHTMLStaticSnippetParams() *GetNutritionLabelHTMLStaticSnippetParams {
	var (
		countryDefault = string("uk")
	)
	return &GetNutritionLabelHTMLStaticSnippetParams{
		Country: &countryDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNutritionLabelHTMLStaticSnippetParamsWithTimeout creates a new GetNutritionLabelHTMLStaticSnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNutritionLabelHTMLStaticSnippetParamsWithTimeout(timeout time.Duration) *GetNutritionLabelHTMLStaticSnippetParams {
	var (
		countryDefault = string("uk")
	)
	return &GetNutritionLabelHTMLStaticSnippetParams{
		Country: &countryDefault,

		timeout: timeout,
	}
}

// NewGetNutritionLabelHTMLStaticSnippetParamsWithContext creates a new GetNutritionLabelHTMLStaticSnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNutritionLabelHTMLStaticSnippetParamsWithContext(ctx context.Context) *GetNutritionLabelHTMLStaticSnippetParams {
	var (
		countryDefault = string("uk")
	)
	return &GetNutritionLabelHTMLStaticSnippetParams{
		Country: &countryDefault,

		Context: ctx,
	}
}

// NewGetNutritionLabelHTMLStaticSnippetParamsWithHTTPClient creates a new GetNutritionLabelHTMLStaticSnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNutritionLabelHTMLStaticSnippetParamsWithHTTPClient(client *http.Client) *GetNutritionLabelHTMLStaticSnippetParams {
	var (
		countryDefault = string("uk")
	)
	return &GetNutritionLabelHTMLStaticSnippetParams{
		Country:    &countryDefault,
		HTTPClient: client,
	}
}

/*GetNutritionLabelHTMLStaticSnippetParams contains all the parameters to send to the API endpoint
for the get nutrition label Html static snippet operation typically these are written to a http.Request
*/
type GetNutritionLabelHTMLStaticSnippetParams struct {

	/*Country
	  The country code which the RDA should reflect.

	*/
	Country *string
	/*PortionSize
	  How large do you want the portions to be? The recipe creators will often decide how many portions a recipe will produce, when following the instructions. However, some people have healthier appetites than others, so this parameter allows control of this element of the recipe. This value is in gramms.

	*/
	PortionSize *int32
	/*Portions
	  The number of portions the parsed ingredients will create.

	*/
	Portions *int32
	/*RecipeID
	  The recipe ID which the nutritional label should be calculated from. This can either be the recipe's integer recipe ID (the primary key) or it's string oid.

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithTimeout(timeout time.Duration) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithContext(ctx context.Context) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithHTTPClient(client *http.Client) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountry adds the country to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithCountry(country *string) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetCountry(country *string) {
	o.Country = country
}

// WithPortionSize adds the portionSize to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithPortionSize(portionSize *int32) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetPortionSize(portionSize)
	return o
}

// SetPortionSize adds the portionSize to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetPortionSize(portionSize *int32) {
	o.PortionSize = portionSize
}

// WithPortions adds the portions to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithPortions(portions *int32) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetPortions(portions)
	return o
}

// SetPortions adds the portions to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetPortions(portions *int32) {
	o.Portions = portions
}

// WithRecipeID adds the recipeID to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) WithRecipeID(recipeID string) *GetNutritionLabelHTMLStaticSnippetParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the get nutrition label Html static snippet params
func (o *GetNutritionLabelHTMLStaticSnippetParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNutritionLabelHTMLStaticSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.PortionSize != nil {

		// query param portionSize
		var qrPortionSize int32
		if o.PortionSize != nil {
			qrPortionSize = *o.PortionSize
		}
		qPortionSize := swag.FormatInt32(qrPortionSize)
		if qPortionSize != "" {
			if err := r.SetQueryParam("portionSize", qPortionSize); err != nil {
				return err
			}
		}

	}

	if o.Portions != nil {

		// query param portions
		var qrPortions int32
		if o.Portions != nil {
			qrPortions = *o.Portions
		}
		qPortions := swag.FormatInt32(qrPortions)
		if qPortions != "" {
			if err := r.SetQueryParam("portions", qPortions); err != nil {
				return err
			}
		}

	}

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}