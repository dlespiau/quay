// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// GetOrganizationReader is a Reader for the GetOrganization structure.
type GetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationOK creates a GetOrganizationOK with default headers values
func NewGetOrganizationOK() *GetOrganizationOK {
	return &GetOrganizationOK{}
}

/*GetOrganizationOK handles this case with default header values.

Successful invocation
*/
type GetOrganizationOK struct {
}

func (o *GetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}][%d] getOrganizationOK ", 200)
}

func (o *GetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationBadRequest creates a GetOrganizationBadRequest with default headers values
func NewGetOrganizationBadRequest() *GetOrganizationBadRequest {
	return &GetOrganizationBadRequest{}
}

/*GetOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganizationBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}][%d] getOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationUnauthorized creates a GetOrganizationUnauthorized with default headers values
func NewGetOrganizationUnauthorized() *GetOrganizationUnauthorized {
	return &GetOrganizationUnauthorized{}
}

/*GetOrganizationUnauthorized handles this case with default header values.

Session required
*/
type GetOrganizationUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}][%d] getOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationForbidden creates a GetOrganizationForbidden with default headers values
func NewGetOrganizationForbidden() *GetOrganizationForbidden {
	return &GetOrganizationForbidden{}
}

/*GetOrganizationForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrganizationForbidden struct {
	Payload *models.APIError
}

func (o *GetOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}][%d] getOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationNotFound creates a GetOrganizationNotFound with default headers values
func NewGetOrganizationNotFound() *GetOrganizationNotFound {
	return &GetOrganizationNotFound{}
}

/*GetOrganizationNotFound handles this case with default header values.

Not found
*/
type GetOrganizationNotFound struct {
	Payload *models.APIError
}

func (o *GetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}][%d] getOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
