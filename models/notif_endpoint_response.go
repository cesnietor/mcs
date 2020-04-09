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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotifEndpointResponse notif endpoint response
//
// swagger:model notifEndpointResponse
type NotifEndpointResponse struct {

	// notification endpoints
	NotificationEndpoints []*NotificationEndpointItem `json:"notification_endpoints"`
}

// Validate validates this notif endpoint response
func (m *NotifEndpointResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifEndpointResponse) validateNotificationEndpoints(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationEndpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.NotificationEndpoints); i++ {
		if swag.IsZero(m.NotificationEndpoints[i]) { // not required
			continue
		}

		if m.NotificationEndpoints[i] != nil {
			if err := m.NotificationEndpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notification_endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotifEndpointResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifEndpointResponse) UnmarshalBinary(b []byte) error {
	var res NotifEndpointResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
