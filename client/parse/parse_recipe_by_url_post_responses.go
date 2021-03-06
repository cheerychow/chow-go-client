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

// ParseRecipeByURLPostReader is a Reader for the ParseRecipeByURLPost structure.
type ParseRecipeByURLPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ParseRecipeByURLPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewParseRecipeByURLPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewParseRecipeByURLPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewParseRecipeByURLPostOK creates a ParseRecipeByURLPostOK with default headers values
func NewParseRecipeByURLPostOK() *ParseRecipeByURLPostOK {
	return &ParseRecipeByURLPostOK{}
}

/*ParseRecipeByURLPostOK handles this case with default header values.

No response was specified
*/
type ParseRecipeByURLPostOK struct {
	Payload *models.Recipe
}

func (o *ParseRecipeByURLPostOK) Error() string {
	return fmt.Sprintf("[POST /parse/via-url/recipe][%d] parseRecipeByUrlPostOK  %+v", 200, o.Payload)
}

func (o *ParseRecipeByURLPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewParseRecipeByURLPostBadRequest creates a ParseRecipeByURLPostBadRequest with default headers values
func NewParseRecipeByURLPostBadRequest() *ParseRecipeByURLPostBadRequest {
	return &ParseRecipeByURLPostBadRequest{}
}

/*ParseRecipeByURLPostBadRequest handles this case with default header values.

Description was not specified
*/
type ParseRecipeByURLPostBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *ParseRecipeByURLPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /parse/via-url/recipe][%d] parseRecipeByUrlPostBadRequest  %+v", 400, o.Payload)
}

func (o *ParseRecipeByURLPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
