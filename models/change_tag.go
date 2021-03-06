// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ChangeTag Makes changes to a specific tag
// swagger:model ChangeTag
type ChangeTag struct {

	// (If specified) The expiration for the image
	Expiration float64 `json:"expiration,omitempty"`

	// (If specified) Image identifier to which the tag should point
	Image string `json:"image,omitempty"`
}

// Validate validates this change tag
func (m *ChangeTag) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeTag) UnmarshalBinary(b []byte) error {
	var res ChangeTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
