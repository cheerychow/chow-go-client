// Code generated by go-swagger; DO NOT EDIT.

package fav_recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// UnfavoriteARecipeByRecipeIDReader is a Reader for the UnfavoriteARecipeByRecipeID structure.
type UnfavoriteARecipeByRecipeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnfavoriteARecipeByRecipeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUnfavoriteARecipeByRecipeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUnfavoriteARecipeByRecipeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnfavoriteARecipeByRecipeIDOK creates a UnfavoriteARecipeByRecipeIDOK with default headers values
func NewUnfavoriteARecipeByRecipeIDOK() *UnfavoriteARecipeByRecipeIDOK {
	return &UnfavoriteARecipeByRecipeIDOK{}
}

/*UnfavoriteARecipeByRecipeIDOK handles this case with default header values.

Description was not specified
*/
type UnfavoriteARecipeByRecipeIDOK struct {
	Payload *models.HTTPAPIResponse
}

func (o *UnfavoriteARecipeByRecipeIDOK) Error() string {
	return fmt.Sprintf("[DELETE /fav-recipe/by-recipe-id/{recipeId}][%d] unfavoriteARecipeByRecipeIdOK  %+v", 200, o.Payload)
}

func (o *UnfavoriteARecipeByRecipeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnfavoriteARecipeByRecipeIDNotFound creates a UnfavoriteARecipeByRecipeIDNotFound with default headers values
func NewUnfavoriteARecipeByRecipeIDNotFound() *UnfavoriteARecipeByRecipeIDNotFound {
	return &UnfavoriteARecipeByRecipeIDNotFound{}
}

/*UnfavoriteARecipeByRecipeIDNotFound handles this case with default header values.

Fav recipe not found
*/
type UnfavoriteARecipeByRecipeIDNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *UnfavoriteARecipeByRecipeIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fav-recipe/by-recipe-id/{recipeId}][%d] unfavoriteARecipeByRecipeIdNotFound  %+v", 404, o.Payload)
}

func (o *UnfavoriteARecipeByRecipeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}