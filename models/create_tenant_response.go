// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateTenantResponse create tenant response
//
// swagger:model createTenantResponse
type CreateTenantResponse struct {

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// console
	Console *CreateTenantResponseConsole `json:"console,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`
}

// Validate validates this create tenant response
func (m *CreateTenantResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTenantResponse) validateConsole(formats strfmt.Registry) error {

	if swag.IsZero(m.Console) { // not required
		return nil
	}

	if m.Console != nil {
		if err := m.Console.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("console")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTenantResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTenantResponse) UnmarshalBinary(b []byte) error {
	var res CreateTenantResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateTenantResponseConsole create tenant response console
//
// swagger:model CreateTenantResponseConsole
type CreateTenantResponseConsole struct {

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`
}

// Validate validates this create tenant response console
func (m *CreateTenantResponseConsole) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTenantResponseConsole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTenantResponseConsole) UnmarshalBinary(b []byte) error {
	var res CreateTenantResponseConsole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
