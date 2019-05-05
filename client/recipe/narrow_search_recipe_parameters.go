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

// NewNarrowSearchRecipeParams creates a new NarrowSearchRecipeParams object
// with the default values initialized.
func NewNarrowSearchRecipeParams() *NarrowSearchRecipeParams {
	var (
		sortbyDefault  = string("name")
		sortdirDefault = string("desc")
	)
	return &NarrowSearchRecipeParams{
		Sortby:  &sortbyDefault,
		Sortdir: &sortdirDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNarrowSearchRecipeParamsWithTimeout creates a new NarrowSearchRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNarrowSearchRecipeParamsWithTimeout(timeout time.Duration) *NarrowSearchRecipeParams {
	var (
		sortbyDefault  = string("name")
		sortdirDefault = string("desc")
	)
	return &NarrowSearchRecipeParams{
		Sortby:  &sortbyDefault,
		Sortdir: &sortdirDefault,

		timeout: timeout,
	}
}

// NewNarrowSearchRecipeParamsWithContext creates a new NarrowSearchRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewNarrowSearchRecipeParamsWithContext(ctx context.Context) *NarrowSearchRecipeParams {
	var (
		sortbyDefault  = string("name")
		sortdirDefault = string("desc")
	)
	return &NarrowSearchRecipeParams{
		Sortby:  &sortbyDefault,
		Sortdir: &sortdirDefault,

		Context: ctx,
	}
}

// NewNarrowSearchRecipeParamsWithHTTPClient creates a new NarrowSearchRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNarrowSearchRecipeParamsWithHTTPClient(client *http.Client) *NarrowSearchRecipeParams {
	var (
		sortbyDefault  = string("name")
		sortdirDefault = string("desc")
	)
	return &NarrowSearchRecipeParams{
		Sortby:     &sortbyDefault,
		Sortdir:    &sortdirDefault,
		HTTPClient: client,
	}
}

/*NarrowSearchRecipeParams contains all the parameters to send to the API endpoint
for the narrow search recipe operation typically these are written to a http.Request
*/
type NarrowSearchRecipeParams struct {

	/*Fields
	  The recipe object fields to include in the response.

	*/
	Fields *string
	/*Limit
	  Limit the number of results returned.

	*/
	Limit *int32
	/*Offset
	  The offset into search results.

	*/
	Offset *int64
	/*OwnerHandle
	  The owner's unique handle.

	*/
	OwnerHandle *string
	/*OwnerID
	  The owner's primary key.

	*/
	OwnerID *int64
	/*RecipeName
	  Search the recipe name, used as a wild card on the name column like so; '%name%'

	*/
	RecipeName *string
	/*Sortby
	  The nutrition field to sort by.

	*/
	Sortby *string
	/*Sortdir
	  The sort direction for the results.

	*/
	Sortdir *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithTimeout(timeout time.Duration) *NarrowSearchRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithContext(ctx context.Context) *NarrowSearchRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithHTTPClient(client *http.Client) *NarrowSearchRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithFields(fields *string) *NarrowSearchRecipeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithLimit(limit *int32) *NarrowSearchRecipeParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithOffset(offset *int64) *NarrowSearchRecipeParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerHandle adds the ownerHandle to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithOwnerHandle(ownerHandle *string) *NarrowSearchRecipeParams {
	o.SetOwnerHandle(ownerHandle)
	return o
}

// SetOwnerHandle adds the ownerHandle to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetOwnerHandle(ownerHandle *string) {
	o.OwnerHandle = ownerHandle
}

// WithOwnerID adds the ownerID to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithOwnerID(ownerID *int64) *NarrowSearchRecipeParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetOwnerID(ownerID *int64) {
	o.OwnerID = ownerID
}

// WithRecipeName adds the recipeName to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithRecipeName(recipeName *string) *NarrowSearchRecipeParams {
	o.SetRecipeName(recipeName)
	return o
}

// SetRecipeName adds the recipeName to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetRecipeName(recipeName *string) {
	o.RecipeName = recipeName
}

// WithSortby adds the sortby to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithSortby(sortby *string) *NarrowSearchRecipeParams {
	o.SetSortby(sortby)
	return o
}

// SetSortby adds the sortby to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetSortby(sortby *string) {
	o.Sortby = sortby
}

// WithSortdir adds the sortdir to the narrow search recipe params
func (o *NarrowSearchRecipeParams) WithSortdir(sortdir *string) *NarrowSearchRecipeParams {
	o.SetSortdir(sortdir)
	return o
}

// SetSortdir adds the sortdir to the narrow search recipe params
func (o *NarrowSearchRecipeParams) SetSortdir(sortdir *string) {
	o.Sortdir = sortdir
}

// WriteToRequest writes these params to a swagger request
func (o *NarrowSearchRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.OwnerHandle != nil {

		// query param ownerHandle
		var qrOwnerHandle string
		if o.OwnerHandle != nil {
			qrOwnerHandle = *o.OwnerHandle
		}
		qOwnerHandle := qrOwnerHandle
		if qOwnerHandle != "" {
			if err := r.SetQueryParam("ownerHandle", qOwnerHandle); err != nil {
				return err
			}
		}

	}

	if o.OwnerID != nil {

		// query param ownerId
		var qrOwnerID int64
		if o.OwnerID != nil {
			qrOwnerID = *o.OwnerID
		}
		qOwnerID := swag.FormatInt64(qrOwnerID)
		if qOwnerID != "" {
			if err := r.SetQueryParam("ownerId", qOwnerID); err != nil {
				return err
			}
		}

	}

	if o.RecipeName != nil {

		// query param recipeName
		var qrRecipeName string
		if o.RecipeName != nil {
			qrRecipeName = *o.RecipeName
		}
		qRecipeName := qrRecipeName
		if qRecipeName != "" {
			if err := r.SetQueryParam("recipeName", qRecipeName); err != nil {
				return err
			}
		}

	}

	if o.Sortby != nil {

		// query param sortby
		var qrSortby string
		if o.Sortby != nil {
			qrSortby = *o.Sortby
		}
		qSortby := qrSortby
		if qSortby != "" {
			if err := r.SetQueryParam("sortby", qSortby); err != nil {
				return err
			}
		}

	}

	if o.Sortdir != nil {

		// query param sortdir
		var qrSortdir string
		if o.Sortdir != nil {
			qrSortdir = *o.Sortdir
		}
		qSortdir := qrSortdir
		if qSortdir != "" {
			if err := r.SetQueryParam("sortdir", qSortdir); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}