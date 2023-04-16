// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1APIKey v1 Api key
//
// swagger:model v1ApiKey
type V1APIKey struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// expiration
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last seen
	// Format: date-time
	LastSeen strfmt.DateTime `json:"lastSeen,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this v1 Api key
func (m *V1APIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1APIKey) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1APIKey) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1APIKey) validateLastSeen(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Api key based on context it is used
func (m *V1APIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1APIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1APIKey) UnmarshalBinary(b []byte) error {
	var res V1APIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
