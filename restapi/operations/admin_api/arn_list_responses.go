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

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/mcs/models"
)

// ArnListOKCode is the HTTP code returned for type ArnListOK
const ArnListOKCode int = 200

/*ArnListOK A successful response.

swagger:response arnListOK
*/
type ArnListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ArnsResponse `json:"body,omitempty"`
}

// NewArnListOK creates ArnListOK with default headers values
func NewArnListOK() *ArnListOK {

	return &ArnListOK{}
}

// WithPayload adds the payload to the arn list o k response
func (o *ArnListOK) WithPayload(payload *models.ArnsResponse) *ArnListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the arn list o k response
func (o *ArnListOK) SetPayload(payload *models.ArnsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ArnListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ArnListDefault Generic error response.

swagger:response arnListDefault
*/
type ArnListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewArnListDefault creates ArnListDefault with default headers values
func NewArnListDefault(code int) *ArnListDefault {
	if code <= 0 {
		code = 500
	}

	return &ArnListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the arn list default response
func (o *ArnListDefault) WithStatusCode(code int) *ArnListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the arn list default response
func (o *ArnListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the arn list default response
func (o *ArnListDefault) WithPayload(payload *models.Error) *ArnListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the arn list default response
func (o *ArnListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ArnListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}