// Code generated by go-swagger; DO NOT EDIT.

package rda

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cheerychow/go-client/models"
)

// GetRDADataReader is a Reader for the GetRDAData structure.
type GetRDADataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRDADataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRDADataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRDADataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRDADataOK creates a GetRDADataOK with default headers values
func NewGetRDADataOK() *GetRDADataOK {
	return &GetRDADataOK{}
}

/*GetRDADataOK handles this case with default header values.

Description was not specified
*/
type GetRDADataOK struct {
	Payload *models.RDA
}

func (o *GetRDADataOK) Error() string {
	return fmt.Sprintf("[GET /rda/{country}][%d] getRDADataOK  %+v", 200, o.Payload)
}

func (o *GetRDADataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RDA)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRDADataBadRequest creates a GetRDADataBadRequest with default headers values
func NewGetRDADataBadRequest() *GetRDADataBadRequest {
	return &GetRDADataBadRequest{}
}

/*GetRDADataBadRequest handles this case with default header values.

Description was not specified
*/
type GetRDADataBadRequest struct {
	Payload *models.HTTPAPIError
}

func (o *GetRDADataBadRequest) Error() string {
	return fmt.Sprintf("[GET /rda/{country}][%d] getRDADataBadRequest  %+v", 400, o.Payload)
}

func (o *GetRDADataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}