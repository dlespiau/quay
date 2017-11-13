// Code generated by go-swagger; DO NOT EDIT.

package manifest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
)

// DeleteManifestLabelReader is a Reader for the DeleteManifestLabel structure.
type DeleteManifestLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteManifestLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteManifestLabelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteManifestLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteManifestLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteManifestLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteManifestLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteManifestLabelNoContent creates a DeleteManifestLabelNoContent with default headers values
func NewDeleteManifestLabelNoContent() *DeleteManifestLabelNoContent {
	return &DeleteManifestLabelNoContent{}
}

/*DeleteManifestLabelNoContent handles this case with default header values.

Deleted
*/
type DeleteManifestLabelNoContent struct {
}

func (o *DeleteManifestLabelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}][%d] deleteManifestLabelNoContent ", 204)
}

func (o *DeleteManifestLabelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteManifestLabelBadRequest creates a DeleteManifestLabelBadRequest with default headers values
func NewDeleteManifestLabelBadRequest() *DeleteManifestLabelBadRequest {
	return &DeleteManifestLabelBadRequest{}
}

/*DeleteManifestLabelBadRequest handles this case with default header values.

Bad Request
*/
type DeleteManifestLabelBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteManifestLabelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}][%d] deleteManifestLabelBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteManifestLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteManifestLabelUnauthorized creates a DeleteManifestLabelUnauthorized with default headers values
func NewDeleteManifestLabelUnauthorized() *DeleteManifestLabelUnauthorized {
	return &DeleteManifestLabelUnauthorized{}
}

/*DeleteManifestLabelUnauthorized handles this case with default header values.

Session required
*/
type DeleteManifestLabelUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteManifestLabelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}][%d] deleteManifestLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteManifestLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteManifestLabelForbidden creates a DeleteManifestLabelForbidden with default headers values
func NewDeleteManifestLabelForbidden() *DeleteManifestLabelForbidden {
	return &DeleteManifestLabelForbidden{}
}

/*DeleteManifestLabelForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteManifestLabelForbidden struct {
	Payload *models.APIError
}

func (o *DeleteManifestLabelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}][%d] deleteManifestLabelForbidden  %+v", 403, o.Payload)
}

func (o *DeleteManifestLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteManifestLabelNotFound creates a DeleteManifestLabelNotFound with default headers values
func NewDeleteManifestLabelNotFound() *DeleteManifestLabelNotFound {
	return &DeleteManifestLabelNotFound{}
}

/*DeleteManifestLabelNotFound handles this case with default header values.

Not found
*/
type DeleteManifestLabelNotFound struct {
	Payload *models.APIError
}

func (o *DeleteManifestLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}][%d] deleteManifestLabelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteManifestLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
