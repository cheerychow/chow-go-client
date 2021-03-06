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

// NewGetRecipeGroupsParams creates a new GetRecipeGroupsParams object
// with the default values initialized.
func NewGetRecipeGroupsParams() *GetRecipeGroupsParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &GetRecipeGroupsParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeGroupsParamsWithTimeout creates a new GetRecipeGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeGroupsParamsWithTimeout(timeout time.Duration) *GetRecipeGroupsParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &GetRecipeGroupsParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		timeout: timeout,
	}
}

// NewGetRecipeGroupsParamsWithContext creates a new GetRecipeGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeGroupsParamsWithContext(ctx context.Context) *GetRecipeGroupsParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &GetRecipeGroupsParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,

		Context: ctx,
	}
}

// NewGetRecipeGroupsParamsWithHTTPClient creates a new GetRecipeGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeGroupsParamsWithHTTPClient(client *http.Client) *GetRecipeGroupsParams {
	var (
		includeMenuItemsDefault = bool(false)
		includeRecipeDefault    = bool(false)
	)
	return &GetRecipeGroupsParams{
		IncludeMenuItems: &includeMenuItemsDefault,
		IncludeRecipe:    &includeRecipeDefault,
		HTTPClient:       client,
	}
}

/*GetRecipeGroupsParams contains all the parameters to send to the API endpoint
for the get recipe groups operation typically these are written to a http.Request
*/
type GetRecipeGroupsParams struct {

	/*Category
	  Find recipe groups by category.

	*/
	Category *string
	/*IncludeMenuItems
	  Should the recipe-group's menu items (1-to-many relationship) be included in the recipe-group object?

	*/
	IncludeMenuItems *bool
	/*IncludeRecipe
	  Should the menu-item's recipe object (1-to-1 relationship) be included in the menu-item object?

	*/
	IncludeRecipe *bool
	/*ReferenceID
	  Find recipe group by reference ID. When this parameter is included included in the query, a preceeding =, >=, <=, >, < or != must be included to indicate how the searching should be done.

	*/
	ReferenceID *string
	/*Slug
	  The slug to search for. You can include a comma seperated list of slugs to perform an OR search on each one.

	*/
	Slug *string
	/*UserID
	  Find recipe groups's g3r to the user. This will only provided publicly available lists. The API also needs a user id to search against, which can either be this value, or the logged in user's id (i.e. searching for their own).

	*/
	UserID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe groups params
func (o *GetRecipeGroupsParams) WithTimeout(timeout time.Duration) *GetRecipeGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe groups params
func (o *GetRecipeGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe groups params
func (o *GetRecipeGroupsParams) WithContext(ctx context.Context) *GetRecipeGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe groups params
func (o *GetRecipeGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe groups params
func (o *GetRecipeGroupsParams) WithHTTPClient(client *http.Client) *GetRecipeGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe groups params
func (o *GetRecipeGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get recipe groups params
func (o *GetRecipeGroupsParams) WithCategory(category *string) *GetRecipeGroupsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get recipe groups params
func (o *GetRecipeGroupsParams) SetCategory(category *string) {
	o.Category = category
}

// WithIncludeMenuItems adds the includeMenuItems to the get recipe groups params
func (o *GetRecipeGroupsParams) WithIncludeMenuItems(includeMenuItems *bool) *GetRecipeGroupsParams {
	o.SetIncludeMenuItems(includeMenuItems)
	return o
}

// SetIncludeMenuItems adds the includeMenuItems to the get recipe groups params
func (o *GetRecipeGroupsParams) SetIncludeMenuItems(includeMenuItems *bool) {
	o.IncludeMenuItems = includeMenuItems
}

// WithIncludeRecipe adds the includeRecipe to the get recipe groups params
func (o *GetRecipeGroupsParams) WithIncludeRecipe(includeRecipe *bool) *GetRecipeGroupsParams {
	o.SetIncludeRecipe(includeRecipe)
	return o
}

// SetIncludeRecipe adds the includeRecipe to the get recipe groups params
func (o *GetRecipeGroupsParams) SetIncludeRecipe(includeRecipe *bool) {
	o.IncludeRecipe = includeRecipe
}

// WithReferenceID adds the referenceID to the get recipe groups params
func (o *GetRecipeGroupsParams) WithReferenceID(referenceID *string) *GetRecipeGroupsParams {
	o.SetReferenceID(referenceID)
	return o
}

// SetReferenceID adds the referenceId to the get recipe groups params
func (o *GetRecipeGroupsParams) SetReferenceID(referenceID *string) {
	o.ReferenceID = referenceID
}

// WithSlug adds the slug to the get recipe groups params
func (o *GetRecipeGroupsParams) WithSlug(slug *string) *GetRecipeGroupsParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the get recipe groups params
func (o *GetRecipeGroupsParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithUserID adds the userID to the get recipe groups params
func (o *GetRecipeGroupsParams) WithUserID(userID *int64) *GetRecipeGroupsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get recipe groups params
func (o *GetRecipeGroupsParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.IncludeMenuItems != nil {

		// query param include-menu-items
		var qrIncludeMenuItems bool
		if o.IncludeMenuItems != nil {
			qrIncludeMenuItems = *o.IncludeMenuItems
		}
		qIncludeMenuItems := swag.FormatBool(qrIncludeMenuItems)
		if qIncludeMenuItems != "" {
			if err := r.SetQueryParam("include-menu-items", qIncludeMenuItems); err != nil {
				return err
			}
		}

	}

	if o.IncludeRecipe != nil {

		// query param include-recipe
		var qrIncludeRecipe bool
		if o.IncludeRecipe != nil {
			qrIncludeRecipe = *o.IncludeRecipe
		}
		qIncludeRecipe := swag.FormatBool(qrIncludeRecipe)
		if qIncludeRecipe != "" {
			if err := r.SetQueryParam("include-recipe", qIncludeRecipe); err != nil {
				return err
			}
		}

	}

	if o.ReferenceID != nil {

		// query param reference_id
		var qrReferenceID string
		if o.ReferenceID != nil {
			qrReferenceID = *o.ReferenceID
		}
		qReferenceID := qrReferenceID
		if qReferenceID != "" {
			if err := r.SetQueryParam("reference_id", qReferenceID); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param userId
		var qrUserID int64
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatInt64(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("userId", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
