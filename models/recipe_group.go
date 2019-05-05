// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecipeGroup A recipe group is a collection of recipes, where each recipe in the group is represented by a MenuItem (note the recipe_id foreign key in the MenuItem object). A recipe group provides the flixibilty of creating a group of recipes for any purpose. For example, if you want to provide a recipe bookmarking service, you could create a particular recipe group, representing the collection of recipes. And recipes can be added to the collection, represented by menu-items objects.
// swagger:model RecipeGroup
type RecipeGroup struct {

	// category
	Category string `json:"category,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// group id
	GroupID int64 `json:"group_id,omitempty"`

	// menu items
	MenuItems []*MenuItem `json:"menu_items"`

	// name
	// Required: true
	// Min Length: 2
	Name *string `json:"name"`

	// owner id
	OwnerID int64 `json:"owner_id,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// public
	Public int64 `json:"public,omitempty"`

	// Optional integer field which can be used to identify this group using a number. This is useful when the group refers to a date, and this value can be the unixtime stamp, alllowing searching based on this value.
	ReferenceID int64 `json:"reference_id,omitempty"`

	// This is a unique identifier for a recipe group. It can be something like a recipe name, a day of the week (to collect recipes for that day).
	// Required: true
	// Min Length: 2
	Slug *string `json:"slug"`
}

// Validate validates this recipe group
func (m *RecipeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMenuItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeGroup) validateMenuItems(formats strfmt.Registry) error {

	if swag.IsZero(m.MenuItems) { // not required
		return nil
	}

	for i := 0; i < len(m.MenuItems); i++ {
		if swag.IsZero(m.MenuItems[i]) { // not required
			continue
		}

		if m.MenuItems[i] != nil {
			if err := m.MenuItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("menu_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecipeGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 2); err != nil {
		return err
	}

	return nil
}

func (m *RecipeGroup) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", string(*m.Slug), 2); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecipeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeGroup) UnmarshalBinary(b []byte) error {
	var res RecipeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}