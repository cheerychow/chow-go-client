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

	models "github.com/cheerychow/chow-go-client/models"
)

// NewAddIngredientsToRecipeParams creates a new AddIngredientsToRecipeParams object
// with the default values initialized.
func NewAddIngredientsToRecipeParams() *AddIngredientsToRecipeParams {
	var ()
	return &AddIngredientsToRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientsToRecipeParamsWithTimeout creates a new AddIngredientsToRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddIngredientsToRecipeParamsWithTimeout(timeout time.Duration) *AddIngredientsToRecipeParams {
	var ()
	return &AddIngredientsToRecipeParams{

		timeout: timeout,
	}
}

// NewAddIngredientsToRecipeParamsWithContext creates a new AddIngredientsToRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddIngredientsToRecipeParamsWithContext(ctx context.Context) *AddIngredientsToRecipeParams {
	var ()
	return &AddIngredientsToRecipeParams{

		Context: ctx,
	}
}

// NewAddIngredientsToRecipeParamsWithHTTPClient creates a new AddIngredientsToRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddIngredientsToRecipeParamsWithHTTPClient(client *http.Client) *AddIngredientsToRecipeParams {
	var ()
	return &AddIngredientsToRecipeParams{
		HTTPClient: client,
	}
}

/*AddIngredientsToRecipeParams contains all the parameters to send to the API endpoint
for the add ingredients to recipe operation typically these are written to a http.Request
*/
type AddIngredientsToRecipeParams struct {

	/*RecipeID
	  The recipe ID

	*/
	RecipeID string
	/*RecipeIngredients
	  The recipe ID

	*/
	RecipeIngredients []*models.Ingredient

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) WithTimeout(timeout time.Duration) *AddIngredientsToRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) WithContext(ctx context.Context) *AddIngredientsToRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) WithHTTPClient(client *http.Client) *AddIngredientsToRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipeID adds the recipeID to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) WithRecipeID(recipeID string) *AddIngredientsToRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WithRecipeIngredients adds the recipeIngredients to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) WithRecipeIngredients(recipeIngredients []*models.Ingredient) *AddIngredientsToRecipeParams {
	o.SetRecipeIngredients(recipeIngredients)
	return o
}

// SetRecipeIngredients adds the recipeIngredients to the add ingredients to recipe params
func (o *AddIngredientsToRecipeParams) SetRecipeIngredients(recipeIngredients []*models.Ingredient) {
	o.RecipeIngredients = recipeIngredients
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientsToRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if o.RecipeIngredients != nil {
		if err := r.SetBodyParam(o.RecipeIngredients); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
