// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Cuisine cuisine
// swagger:model Cuisine
type Cuisine struct {

	// cuisine id
	CuisineID int64 `json:"cuisine_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this cuisine
func (m *Cuisine) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cuisine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cuisine) UnmarshalBinary(b []byte) error {
	var res Cuisine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}