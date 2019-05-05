// Code generated by go-swagger; DO NOT EDIT.

package scratch_recipe

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

// NewRemoveIngredientFromScratchPadRecipeParams creates a new RemoveIngredientFromScratchPadRecipeParams object
// with the default values initialized.
func NewRemoveIngredientFromScratchPadRecipeParams() *RemoveIngredientFromScratchPadRecipeParams {
	var ()
	return &RemoveIngredientFromScratchPadRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveIngredientFromScratchPadRecipeParamsWithTimeout creates a new RemoveIngredientFromScratchPadRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveIngredientFromScratchPadRecipeParamsWithTimeout(timeout time.Duration) *RemoveIngredientFromScratchPadRecipeParams {
	var ()
	return &RemoveIngredientFromScratchPadRecipeParams{

		timeout: timeout,
	}
}

// NewRemoveIngredientFromScratchPadRecipeParamsWithContext creates a new RemoveIngredientFromScratchPadRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveIngredientFromScratchPadRecipeParamsWithContext(ctx context.Context) *RemoveIngredientFromScratchPadRecipeParams {
	var ()
	return &RemoveIngredientFromScratchPadRecipeParams{

		Context: ctx,
	}
}

// NewRemoveIngredientFromScratchPadRecipeParamsWithHTTPClient creates a new RemoveIngredientFromScratchPadRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveIngredientFromScratchPadRecipeParamsWithHTTPClient(client *http.Client) *RemoveIngredientFromScratchPadRecipeParams {
	var ()
	return &RemoveIngredientFromScratchPadRecipeParams{
		HTTPClient: client,
	}
}

/*RemoveIngredientFromScratchPadRecipeParams contains all the parameters to send to the API endpoint
for the remove ingredient from scratch pad recipe operation typically these are written to a http.Request
*/
type RemoveIngredientFromScratchPadRecipeParams struct {

	/*IngredientID
	  The ID of the ingredient that you want to delete

	*/
	IngredientID int64
	/*RecipeID
	  The ID of the recipe that the ingredient belongs to

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) WithTimeout(timeout time.Duration) *RemoveIngredientFromScratchPadRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) WithContext(ctx context.Context) *RemoveIngredientFromScratchPadRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) WithHTTPClient(client *http.Client) *RemoveIngredientFromScratchPadRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientID adds the ingredientID to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) WithIngredientID(ingredientID int64) *RemoveIngredientFromScratchPadRecipeParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) SetIngredientID(ingredientID int64) {
	o.IngredientID = ingredientID
}

// WithRecipeID adds the recipeID to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) WithRecipeID(recipeID string) *RemoveIngredientFromScratchPadRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the remove ingredient from scratch pad recipe params
func (o *RemoveIngredientFromScratchPadRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveIngredientFromScratchPadRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ingredient-id
	if err := r.SetPathParam("ingredient-id", swag.FormatInt64(o.IngredientID)); err != nil {
		return err
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