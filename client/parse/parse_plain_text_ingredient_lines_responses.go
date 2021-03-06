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

// ParsePlainTextIngredientLinesReader is a Reader for the ParsePlainTextIngredientLines structure.
type ParsePlainTextIngredientLinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ParsePlainTextIngredientLinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewParsePlainTextIngredientLinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewParsePlainTextIngredientLinesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewParsePlainTextIngredientLinesOK creates a ParsePlainTextIngredientLinesOK with default headers values
func NewParsePlainTextIngredientLinesOK() *ParsePlainTextIngredientLinesOK {
	return &ParsePlainTextIngredientLinesOK{}
}

/*ParsePlainTextIngredientLinesOK handles this case with default header values.

Retrieves a list of ingredients, including the nutritional content of the ingredient. The list of ingredients is contained within a standard RecipeWithFood object. WARNING- This assumes a number of portions=1. You'll need to divide up the values received in this object to find the actual nutrition per portion.
*/
type ParsePlainTextIngredientLinesOK struct {
	Payload *models.Recipe
}

func (o *ParsePlainTextIngredientLinesOK) Error() string {
	return fmt.Sprintf("[POST /parse/ingredients][%d] parsePlainTextIngredientLinesOK  %+v", 200, o.Payload)
}

func (o *ParsePlainTextIngredientLinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewParsePlainTextIngredientLinesBadRequest creates a ParsePlainTextIngredientLinesBadRequest with default headers values
func NewParsePlainTextIngredientLinesBadRequest() *ParsePlainTextIngredientLinesBadRequest {
	return &ParsePlainTextIngredientLinesBadRequest{}
}

/*ParsePlainTextIngredientLinesBadRequest handles this case with default header values.

Malformed JSON (this endpoint requires an array of strings)
*/
type ParsePlainTextIngredientLinesBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *ParsePlainTextIngredientLinesBadRequest) Error() string {
	return fmt.Sprintf("[POST /parse/ingredients][%d] parsePlainTextIngredientLinesBadRequest  %+v", 400, o.Payload)
}

func (o *ParsePlainTextIngredientLinesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
