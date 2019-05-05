// Code generated by go-swagger; DO NOT EDIT.

package nutrition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new nutrition API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nutrition API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNutritionTipsForNutrient gets a list of the nutrient tips for a particular nutrient
*/
func (a *Client) GetNutritionTipsForNutrient(params *GetNutritionTipsForNutrientParams) (*GetNutritionTipsForNutrientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNutritionTipsForNutrientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNutritionTipsForNutrient",
		Method:             "GET",
		PathPattern:        "/nutrition-tip/{nutrientId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNutritionTipsForNutrientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNutritionTipsForNutrientOK), nil

}

/*
ShowNutritionForAFood retrieves the nutrition present in a particular food present in amount grams of the food
*/
func (a *Client) ShowNutritionForAFood(params *ShowNutritionForAFoodParams) (*ShowNutritionForAFoodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowNutritionForAFoodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowNutritionForAFood",
		Method:             "GET",
		PathPattern:        "/food-nutrition/food/{foodId}/{amount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShowNutritionForAFoodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowNutritionForAFoodOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}