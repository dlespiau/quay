// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetRepoTagsAdditionalProperties get repo tags additional properties
// swagger:model getRepoTagsAdditionalProperties
type GetRepoTagsAdditionalProperties struct {

	// image id
	ImageID string `json:"image_id,omitempty"`

	// last modififed
	LastModififed string `json:"last_modififed,omitempty"`

	// manifest digest
	ManifestDigest string `json:"manifest_digest,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size float64 `json:"size,omitempty"`
}

// Validate validates this get repo tags additional properties
func (m *GetRepoTagsAdditionalProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetRepoTagsAdditionalProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRepoTagsAdditionalProperties) UnmarshalBinary(b []byte) error {
	var res GetRepoTagsAdditionalProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
