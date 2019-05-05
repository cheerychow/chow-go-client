// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MenuItem menu item
// swagger:model MenuItem
type MenuItem struct {

	// group id
	GroupID int64 `json:"group_id,omitempty"`

	// item currency
	ItemCurrency string `json:"item_currency,omitempty"`

	// item description
	ItemDescription string `json:"item_description,omitempty"`

	// item name
	ItemName string `json:"item_name,omitempty"`

	// item position
	ItemPosition int32 `json:"item_position,omitempty"`

	// item price
	ItemPrice float32 `json:"item_price,omitempty"`

	// menu item id
	MenuItemID int64 `json:"menu_item_id,omitempty"`

	// recipe
	Recipe *Recipe `json:"recipe,omitempty"`

	// recipe id
	RecipeID int64 `json:"recipe_id,omitempty"`
}

// Validate validates this menu item
func (m *MenuItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipe(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MenuItem) validateRecipe(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipe) { // not required
		return nil
	}

	if m.Recipe != nil {
		if err := m.Recipe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipe")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MenuItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MenuItem) UnmarshalBinary(b []byte) error {
	var res MenuItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}