// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetRepo Response body of a successful getRepo.
// swagger:model GetRepo
type GetRepo struct {

	// can admin
	CanAdmin bool `json:"can_admin,omitempty"`

	// can write
	CanWrite bool `json:"can_write,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// is organization
	IsOrganization bool `json:"is_organization,omitempty"`

	// is public
	IsPublic bool `json:"is_public,omitempty"`

	// is starred
	IsStarred bool `json:"is_starred,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// status token
	StatusToken string `json:"status_token,omitempty"`

	// tag expiration s
	TagExpirationS float64 `json:"tag_expiration_s,omitempty"`

	// tags
	Tags GetRepoTags `json:"tags,omitempty"`

	// trust enabled
	TrustEnabled bool `json:"trust_enabled,omitempty"`
}

// Validate validates this get repo
func (m *GetRepo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetRepo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRepo) UnmarshalBinary(b []byte) error {
	var res GetRepo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}