// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IngredientParser ingredient parser
// swagger:model IngredientParser
type IngredientParser struct {

	// amount s
	AmountS string `json:"amount_s,omitempty"`

	// confidence
	Confidence float32 `json:"confidence,omitempty"`

	// debug output
	DebugOutput string `json:"debug_output,omitempty"`

	// ingredient parser id
	IngredientParserID int64 `json:"ingredient_parser_id,omitempty"`

	// matcher id
	MatcherID int64 `json:"matcher_id,omitempty"`

	// multiple s
	MultipleS string `json:"multiple_s,omitempty"`

	// processed line
	ProcessedLine string `json:"processed_line,omitempty"`

	// regex id
	RegexID int64 `json:"regex_id,omitempty"`
}

// Validate validates this ingredient parser
func (m *IngredientParser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IngredientParser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientParser) UnmarshalBinary(b []byte) error {
	var res IngredientParser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
