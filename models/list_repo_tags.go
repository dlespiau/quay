// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListRepoTags Response body of a successful listRepoTags.
// swagger:model ListRepoTags
type ListRepoTags struct {

	// has additional
	HasAdditional bool `json:"has_additional,omitempty"`

	// page
	Page float64 `json:"page,omitempty"`

	// tags
	Tags ListRepoTagsTags `json:"tags"`
}

// Validate validates this list repo tags
func (m *ListRepoTags) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListRepoTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListRepoTags) UnmarshalBinary(b []byte) error {
	var res ListRepoTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
