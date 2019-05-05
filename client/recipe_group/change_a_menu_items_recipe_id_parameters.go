// Code generated by go-swagger; DO NOT EDIT.

package recipe_group

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

// NewChangeAMenuItemsRecipeIDParams creates a new ChangeAMenuItemsRecipeIDParams object
// with the default values initialized.
func NewChangeAMenuItemsRecipeIDParams() *ChangeAMenuItemsRecipeIDParams {
	var ()
	return &ChangeAMenuItemsRecipeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeAMenuItemsRecipeIDParamsWithTimeout creates a new ChangeAMenuItemsRecipeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeAMenuItemsRecipeIDParamsWithTimeout(timeout time.Duration) *ChangeAMenuItemsRecipeIDParams {
	var ()
	return &ChangeAMenuItemsRecipeIDParams{

		timeout: timeout,
	}
}

// NewChangeAMenuItemsRecipeIDParamsWithContext creates a new ChangeAMenuItemsRecipeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeAMenuItemsRecipeIDParamsWithContext(ctx context.Context) *ChangeAMenuItemsRecipeIDParams {
	var ()
	return &ChangeAMenuItemsRecipeIDParams{

		Context: ctx,
	}
}

// NewChangeAMenuItemsRecipeIDParamsWithHTTPClient creates a new ChangeAMenuItemsRecipeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeAMenuItemsRecipeIDParamsWithHTTPClient(client *http.Client) *ChangeAMenuItemsRecipeIDParams {
	var ()
	return &ChangeAMenuItemsRecipeIDParams{
		HTTPClient: client,
	}
}

/*ChangeAMenuItemsRecipeIDParams contains all the parameters to send to the API endpoint
for the change a menu items recipe Id operation typically these are written to a http.Request
*/
type ChangeAMenuItemsRecipeIDParams struct {

	/*MenuItemID
	  The id of the menu item to be change.

	*/
	MenuItemID int64
	/*RecipeID
	  The recipe id to be assigned.

	*/
	RecipeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) WithTimeout(timeout time.Duration) *ChangeAMenuItemsRecipeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) WithContext(ctx context.Context) *ChangeAMenuItemsRecipeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) WithHTTPClient(client *http.Client) *ChangeAMenuItemsRecipeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMenuItemID adds the menuItemID to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) WithMenuItemID(menuItemID int64) *ChangeAMenuItemsRecipeIDParams {
	o.SetMenuItemID(menuItemID)
	return o
}

// SetMenuItemID adds the menuItemId to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) SetMenuItemID(menuItemID int64) {
	o.MenuItemID = menuItemID
}

// WithRecipeID adds the recipeID to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) WithRecipeID(recipeID string) *ChangeAMenuItemsRecipeIDParams {
	o.SetRecipeID(recipeID)
	return o
}

// SetRecipeID adds the recipeId to the change a menu items recipe Id params
func (o *ChangeAMenuItemsRecipeIDParams) SetRecipeID(recipeID string) {
	o.RecipeID = recipeID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeAMenuItemsRecipeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param menuItemId
	if err := r.SetPathParam("menuItemId", swag.FormatInt64(o.MenuItemID)); err != nil {
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