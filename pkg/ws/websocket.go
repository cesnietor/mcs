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

// Package ws contains websocket utils for mcs project
package ws

import (
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/minio/mcs/pkg/auth"
)

// Authenticate validates websocket header and returns mcs jwt claims
//
// Authorization Header needs to be like "Authorization Bearer <jwt_token>"
func Authenticate(r *http.Request) (*auth.DecryptedClaims, error) {
	// Get Auth token
	reqToken := r.Header.Get("Authorization")
	// reqToken comes as "Bearer <token>"
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) <= 1 {
		return nil, errors.New(http.StatusBadRequest, "Authentication Header not valid")
	}

	reqToken = strings.TrimSpace(splitToken[1])

	// Perform authentication before upgrading to a Websocket Connection
	claims, err := auth.JWTAuthenticate(reqToken)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, err.Error())
	}
	return claims, nil
}
