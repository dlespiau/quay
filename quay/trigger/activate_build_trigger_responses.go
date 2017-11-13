// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// ActivateBuildTriggerReader is a Reader for the ActivateBuildTrigger structure.
type ActivateBuildTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivateBuildTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewActivateBuildTriggerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewActivateBuildTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewActivateBuildTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActivateBuildTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewActivateBuildTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActivateBuildTriggerCreated creates a ActivateBuildTriggerCreated with default headers values
func NewActivateBuildTriggerCreated() *ActivateBuildTriggerCreated {
	return &ActivateBuildTriggerCreated{}
}

/*ActivateBuildTriggerCreated handles this case with default header values.

Successful creation
*/
type ActivateBuildTriggerCreated struct {
}

func (o *ActivateBuildTriggerCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate][%d] activateBuildTriggerCreated ", 201)
}

func (o *ActivateBuildTriggerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActivateBuildTriggerBadRequest creates a ActivateBuildTriggerBadRequest with default headers values
func NewActivateBuildTriggerBadRequest() *ActivateBuildTriggerBadRequest {
	return &ActivateBuildTriggerBadRequest{}
}

/*ActivateBuildTriggerBadRequest handles this case with default header values.

Bad Request
*/
type ActivateBuildTriggerBadRequest struct {
	Payload *models.APIError
}

func (o *ActivateBuildTriggerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate][%d] activateBuildTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *ActivateBuildTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateBuildTriggerUnauthorized creates a ActivateBuildTriggerUnauthorized with default headers values
func NewActivateBuildTriggerUnauthorized() *ActivateBuildTriggerUnauthorized {
	return &ActivateBuildTriggerUnauthorized{}
}

/*ActivateBuildTriggerUnauthorized handles this case with default header values.

Session required
*/
type ActivateBuildTriggerUnauthorized struct {
	Payload *models.APIError
}

func (o *ActivateBuildTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate][%d] activateBuildTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *ActivateBuildTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateBuildTriggerForbidden creates a ActivateBuildTriggerForbidden with default headers values
func NewActivateBuildTriggerForbidden() *ActivateBuildTriggerForbidden {
	return &ActivateBuildTriggerForbidden{}
}

/*ActivateBuildTriggerForbidden handles this case with default header values.

Unauthorized access
*/
type ActivateBuildTriggerForbidden struct {
	Payload *models.APIError
}

func (o *ActivateBuildTriggerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate][%d] activateBuildTriggerForbidden  %+v", 403, o.Payload)
}

func (o *ActivateBuildTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateBuildTriggerNotFound creates a ActivateBuildTriggerNotFound with default headers values
func NewActivateBuildTriggerNotFound() *ActivateBuildTriggerNotFound {
	return &ActivateBuildTriggerNotFound{}
}

/*ActivateBuildTriggerNotFound handles this case with default header values.

Not found
*/
type ActivateBuildTriggerNotFound struct {
	Payload *models.APIError
}

func (o *ActivateBuildTriggerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate][%d] activateBuildTriggerNotFound  %+v", 404, o.Payload)
}

func (o *ActivateBuildTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
