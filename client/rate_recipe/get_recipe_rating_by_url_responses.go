// Code generated by go-swagger; DO NOT EDIT.

package rate_recipe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// GetRecipeRatingByURLReader is a Reader for the GetRecipeRatingByURL structure.
type GetRecipeRatingByURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeRatingByURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeRatingByURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecipeRatingByURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRecipeRatingByURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeRatingByURLOK creates a GetRecipeRatingByURLOK with default headers values
func NewGetRecipeRatingByURLOK() *GetRecipeRatingByURLOK {
	return &GetRecipeRatingByURLOK{}
}

/*GetRecipeRatingByURLOK handles this case with default header values.

Description was not specified
*/
type GetRecipeRatingByURLOK struct {
	Payload *models.AvgTotal
}

func (o *GetRecipeRatingByURLOK) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-url/{url}][%d] getRecipeRatingByUrlOK  %+v", 200, o.Payload)
}

func (o *GetRecipeRatingByURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvgTotal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeRatingByURLBadRequest creates a GetRecipeRatingByURLBadRequest with default headers values
func NewGetRecipeRatingByURLBadRequest() *GetRecipeRatingByURLBadRequest {
	return &GetRecipeRatingByURLBadRequest{}
}

/*GetRecipeRatingByURLBadRequest handles this case with default header values.

Description was not specified
*/
type GetRecipeRatingByURLBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeRatingByURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-url/{url}][%d] getRecipeRatingByUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecipeRatingByURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipeRatingByURLNotFound creates a GetRecipeRatingByURLNotFound with default headers values
func NewGetRecipeRatingByURLNotFound() *GetRecipeRatingByURLNotFound {
	return &GetRecipeRatingByURLNotFound{}
}

/*GetRecipeRatingByURLNotFound handles this case with default header values.

Recipe not found
*/
type GetRecipeRatingByURLNotFound struct {
	Payload *models.HTTPAPIError
}

func (o *GetRecipeRatingByURLNotFound) Error() string {
	return fmt.Sprintf("[GET /rate-recipe/by-url/{url}][%d] getRecipeRatingByUrlNotFound  %+v", 404, o.Payload)
}

func (o *GetRecipeRatingByURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
