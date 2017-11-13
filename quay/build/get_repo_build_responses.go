// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// GetRepoBuildReader is a Reader for the GetRepoBuild structure.
type GetRepoBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepoBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepoBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRepoBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRepoBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRepoBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRepoBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepoBuildOK creates a GetRepoBuildOK with default headers values
func NewGetRepoBuildOK() *GetRepoBuildOK {
	return &GetRepoBuildOK{}
}

/*GetRepoBuildOK handles this case with default header values.

Successful invocation
*/
type GetRepoBuildOK struct {
}

func (o *GetRepoBuildOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/{build_uuid}][%d] getRepoBuildOK ", 200)
}

func (o *GetRepoBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepoBuildBadRequest creates a GetRepoBuildBadRequest with default headers values
func NewGetRepoBuildBadRequest() *GetRepoBuildBadRequest {
	return &GetRepoBuildBadRequest{}
}

/*GetRepoBuildBadRequest handles this case with default header values.

Bad Request
*/
type GetRepoBuildBadRequest struct {
	Payload *models.APIError
}

func (o *GetRepoBuildBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/{build_uuid}][%d] getRepoBuildBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepoBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildUnauthorized creates a GetRepoBuildUnauthorized with default headers values
func NewGetRepoBuildUnauthorized() *GetRepoBuildUnauthorized {
	return &GetRepoBuildUnauthorized{}
}

/*GetRepoBuildUnauthorized handles this case with default header values.

Session required
*/
type GetRepoBuildUnauthorized struct {
	Payload *models.APIError
}

func (o *GetRepoBuildUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/{build_uuid}][%d] getRepoBuildUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepoBuildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildForbidden creates a GetRepoBuildForbidden with default headers values
func NewGetRepoBuildForbidden() *GetRepoBuildForbidden {
	return &GetRepoBuildForbidden{}
}

/*GetRepoBuildForbidden handles this case with default header values.

Unauthorized access
*/
type GetRepoBuildForbidden struct {
	Payload *models.APIError
}

func (o *GetRepoBuildForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/{build_uuid}][%d] getRepoBuildForbidden  %+v", 403, o.Payload)
}

func (o *GetRepoBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepoBuildNotFound creates a GetRepoBuildNotFound with default headers values
func NewGetRepoBuildNotFound() *GetRepoBuildNotFound {
	return &GetRepoBuildNotFound{}
}

/*GetRepoBuildNotFound handles this case with default header values.

Not found
*/
type GetRepoBuildNotFound struct {
	Payload *models.APIError
}

func (o *GetRepoBuildNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/build/{build_uuid}][%d] getRepoBuildNotFound  %+v", 404, o.Payload)
}

func (o *GetRepoBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}