// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// ListReposReader is a Reader for the ListRepos structure.
type ListReposReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReposReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListReposBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListReposUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListReposForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListReposNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListReposOK creates a ListReposOK with default headers values
func NewListReposOK() *ListReposOK {
	return &ListReposOK{}
}

/*ListReposOK handles this case with default header values.

Successful invocation
*/
type ListReposOK struct {
	Payload *models.ListRepos
}

func (o *ListReposOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository][%d] listReposOK  %+v", 200, o.Payload)
}

func (o *ListReposOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRepos)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReposBadRequest creates a ListReposBadRequest with default headers values
func NewListReposBadRequest() *ListReposBadRequest {
	return &ListReposBadRequest{}
}

/*ListReposBadRequest handles this case with default header values.

Bad Request
*/
type ListReposBadRequest struct {
	Payload *models.APIError
}

func (o *ListReposBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository][%d] listReposBadRequest  %+v", 400, o.Payload)
}

func (o *ListReposBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReposUnauthorized creates a ListReposUnauthorized with default headers values
func NewListReposUnauthorized() *ListReposUnauthorized {
	return &ListReposUnauthorized{}
}

/*ListReposUnauthorized handles this case with default header values.

Session required
*/
type ListReposUnauthorized struct {
	Payload *models.APIError
}

func (o *ListReposUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository][%d] listReposUnauthorized  %+v", 401, o.Payload)
}

func (o *ListReposUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReposForbidden creates a ListReposForbidden with default headers values
func NewListReposForbidden() *ListReposForbidden {
	return &ListReposForbidden{}
}

/*ListReposForbidden handles this case with default header values.

Unauthorized access
*/
type ListReposForbidden struct {
	Payload *models.APIError
}

func (o *ListReposForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository][%d] listReposForbidden  %+v", 403, o.Payload)
}

func (o *ListReposForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReposNotFound creates a ListReposNotFound with default headers values
func NewListReposNotFound() *ListReposNotFound {
	return &ListReposNotFound{}
}

/*ListReposNotFound handles this case with default header values.

Not found
*/
type ListReposNotFound struct {
	Payload *models.APIError
}

func (o *ListReposNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository][%d] listReposNotFound  %+v", 404, o.Payload)
}

func (o *ListReposNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
