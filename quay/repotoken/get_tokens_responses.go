// Code generated by go-swagger; DO NOT EDIT.

package repotoken

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// GetTokensReader is a Reader for the GetTokens structure.
type GetTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTokensNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTokensOK creates a GetTokensOK with default headers values
func NewGetTokensOK() *GetTokensOK {
	return &GetTokensOK{}
}

/*GetTokensOK handles this case with default header values.

Successful invocation
*/
type GetTokensOK struct {
}

func (o *GetTokensOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/{code}][%d] getTokensOK ", 200)
}

func (o *GetTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTokensBadRequest creates a GetTokensBadRequest with default headers values
func NewGetTokensBadRequest() *GetTokensBadRequest {
	return &GetTokensBadRequest{}
}

/*GetTokensBadRequest handles this case with default header values.

Bad Request
*/
type GetTokensBadRequest struct {
	Payload *models.APIError
}

func (o *GetTokensBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/{code}][%d] getTokensBadRequest  %+v", 400, o.Payload)
}

func (o *GetTokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensUnauthorized creates a GetTokensUnauthorized with default headers values
func NewGetTokensUnauthorized() *GetTokensUnauthorized {
	return &GetTokensUnauthorized{}
}

/*GetTokensUnauthorized handles this case with default header values.

Session required
*/
type GetTokensUnauthorized struct {
	Payload *models.APIError
}

func (o *GetTokensUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/{code}][%d] getTokensUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensForbidden creates a GetTokensForbidden with default headers values
func NewGetTokensForbidden() *GetTokensForbidden {
	return &GetTokensForbidden{}
}

/*GetTokensForbidden handles this case with default header values.

Unauthorized access
*/
type GetTokensForbidden struct {
	Payload *models.APIError
}

func (o *GetTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/{code}][%d] getTokensForbidden  %+v", 403, o.Payload)
}

func (o *GetTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensNotFound creates a GetTokensNotFound with default headers values
func NewGetTokensNotFound() *GetTokensNotFound {
	return &GetTokensNotFound{}
}

/*GetTokensNotFound handles this case with default header values.

Not found
*/
type GetTokensNotFound struct {
	Payload *models.APIError
}

func (o *GetTokensNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tokens/{code}][%d] getTokensNotFound  %+v", 404, o.Payload)
}

func (o *GetTokensNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
