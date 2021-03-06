// Code generated by go-swagger; DO NOT EDIT.

package parse

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/chow-go-client/models"
)

// ParseRecipeViaPageContentReader is a Reader for the ParseRecipeViaPageContent structure.
type ParseRecipeViaPageContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ParseRecipeViaPageContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewParseRecipeViaPageContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewParseRecipeViaPageContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewParseRecipeViaPageContentOK creates a ParseRecipeViaPageContentOK with default headers values
func NewParseRecipeViaPageContentOK() *ParseRecipeViaPageContentOK {
	return &ParseRecipeViaPageContentOK{}
}

/*ParseRecipeViaPageContentOK handles this case with default header values.

New recipe object parsed.
*/
type ParseRecipeViaPageContentOK struct {
	Payload *models.Recipe
}

func (o *ParseRecipeViaPageContentOK) Error() string {
	return fmt.Sprintf("[POST /parse/via-page-content/recipe][%d] parseRecipeViaPageContentOK  %+v", 200, o.Payload)
}

func (o *ParseRecipeViaPageContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewParseRecipeViaPageContentBadRequest creates a ParseRecipeViaPageContentBadRequest with default headers values
func NewParseRecipeViaPageContentBadRequest() *ParseRecipeViaPageContentBadRequest {
	return &ParseRecipeViaPageContentBadRequest{}
}

/*ParseRecipeViaPageContentBadRequest handles this case with default header values.

Description was not specified
*/
type ParseRecipeViaPageContentBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *ParseRecipeViaPageContentBadRequest) Error() string {
	return fmt.Sprintf("[POST /parse/via-page-content/recipe][%d] parseRecipeViaPageContentBadRequest  %+v", 400, o.Payload)
}

func (o *ParseRecipeViaPageContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
