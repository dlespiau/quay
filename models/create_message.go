// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateMessage Create a new message
// swagger:model CreateMessage
type CreateMessage struct {

	// message
	Message *CreateMessageMessage `json:"message,omitempty"`
}

// Validate validates this create message
func (m *CreateMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMessage) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {

		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMessage) UnmarshalBinary(b []byte) error {
	var res CreateMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
