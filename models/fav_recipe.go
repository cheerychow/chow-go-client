// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FavRecipe fav recipe
// swagger:model FavRecipe
type FavRecipe struct {

	// created
	Created int64 `json:"created,omitempty"`

	// fav recipe id
	FavRecipeID int64 `json:"fav_recipe_id,omitempty"`

	// owner id
	OwnerID int64 `json:"owner_id,omitempty"`

	// recipe id
	RecipeID int64 `json:"recipe_id,omitempty"`

	// recipe url
	RecipeURL string `json:"recipe_url,omitempty"`
}

// Validate validates this fav recipe
func (m *FavRecipe) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FavRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FavRecipe) UnmarshalBinary(b []byte) error {
	var res FavRecipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
