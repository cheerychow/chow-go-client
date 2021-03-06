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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteScratchPadRecipeParams creates a new DeleteScratchPadRecipeParams object
// with the default values initialized.
func NewDeleteScratchPadRecipeParams() *DeleteScratchPadRecipeParams {
	var ()
	return &DeleteScratchPadRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScratchPadRecipeParamsWithTimeout creates a new DeleteScratchPadRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScratchPadRecipeParamsWithTimeout(timeout time.Duration) *DeleteScratchPadRecipeParams {
	var ()
	return &DeleteScratchPadRecipeParams{

		timeout: timeout,
	}
}

// NewDeleteScratchPadRecipeParamsWithContext creates a new DeleteScratchPadRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScratchPadRecipeParamsWithContext(ctx context.Context) *DeleteScratchPadRecipeParams {
	var ()
	return &DeleteScratchPadRecipeParams{

		Context: ctx,
	}
}

// NewDeleteScratchPadRecipeParamsWithHTTPClient creates a new DeleteScratchPadRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScratchPadRecipeParamsWithHTTPClient(client *http.Client) *DeleteScratchPadRecipeParams {
	var ()
	return &DeleteScratchPadRecipeParams{
		HTTPClient: client,
	}
}

/*DeleteScratchPadRecipeParams contains all the parameters to send to the API endpoint
for the delete scratch pad recipe operation typically these are written to a http.Request
*/
type DeleteScratchPadRecipeParams struct {

	/*RecipeID
	  The ID of the recipe you want to delete

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) WithTimeout(timeout time.Duration) *DeleteScratchPadRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) WithContext(ctx context.Context) *DeleteScratchPadRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) WithHTTPClient(client *http.Client) *DeleteScratchPadRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipeID adds the recipeID to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) WithRecipeID(recipeID string) *DeleteScratchPadRecipeParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the delete scratch pad recipe params
func (o *DeleteScratchPadRecipeParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScratchPadRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recipeId
	if err := r.SetPathParam("recipeId", o.RecipeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
