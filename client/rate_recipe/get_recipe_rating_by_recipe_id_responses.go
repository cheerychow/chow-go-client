// Code generated by go-swagger; DO NOT EDIT.

package rate_recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// GetRecipeRatingByRecipeIDReader is a Reader for the GetRecipeRatingByRecipeID structure.
type GetRecipeRatingByRecipeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeRatingByRecipeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeRatingByRecipeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecipeRatingByRecipeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRecipeRatingByRecipeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeRatingByRecipeIDOK creates a GetRecipeRatingByRecipeIDOK with default headers values
func NewGetRecipeRatingByRecipeIDOK() *GetRecipeRatingByRecipeIDOK {
	return &GetRecipeRatingByRecipeIDOK{}
}

/*GetRecipeRatingByRecipeIDOK handles this case with default header values.

Description was not specified
*/
type GetRecipeRatingByRecipeIDOK struct {
	Payload *models.AvgTotal
}

func (o *GetRecipeRatingByRecipeIDOK) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-id/{recipeId}][%d] getRecipeRatingByRecipeIdOK  %+v", 200, o.Payload)
}

func (o *GetRecipeRatingByRecipeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvgTotal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeRatingByRecipeIDBadRequest creates a GetRecipeRatingByRecipeIDBadRequest with default headers values
func NewGetRecipeRatingByRecipeIDBadRequest() *GetRecipeRatingByRecipeIDBadRequest {
	return &GetRecipeRatingByRecipeIDBadRequest{}
}

/*GetRecipeRatingByRecipeIDBadRequest handles this case with default header values.

Description was not specified
*/
type GetRecipeRatingByRecipeIDBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeRatingByRecipeIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-id/{recipeId}][%d] getRecipeRatingByRecipeIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecipeRatingByRecipeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeRatingByRecipeIDNotFound creates a GetRecipeRatingByRecipeIDNotFound with default headers values
func NewGetRecipeRatingByRecipeIDNotFound() *GetRecipeRatingByRecipeIDNotFound {
	return &GetRecipeRatingByRecipeIDNotFound{}
}

/*GetRecipeRatingByRecipeIDNotFound handles this case with default header values.

Recipe not found
*/
type GetRecipeRatingByRecipeIDNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeRatingByRecipeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-id/{recipeId}][%d] getRecipeRatingByRecipeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRecipeRatingByRecipeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}