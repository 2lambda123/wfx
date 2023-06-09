// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package northbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/siemens/wfx/generated/model"
)

// PostJobsCreatedCode is the HTTP code returned for type PostJobsCreated
const PostJobsCreatedCode int = 201

/*
PostJobsCreated Job was created

swagger:response postJobsCreated
*/
type PostJobsCreated struct {

	/*
	  In: Body
	*/
	Payload *model.Job `json:"body,omitempty"`
}

// NewPostJobsCreated creates PostJobsCreated with default headers values
func NewPostJobsCreated() *PostJobsCreated {

	return &PostJobsCreated{}
}

// WithPayload adds the payload to the post jobs created response
func (o *PostJobsCreated) WithPayload(payload *model.Job) *PostJobsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post jobs created response
func (o *PostJobsCreated) SetPayload(payload *model.Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostJobsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostJobsBadRequestCode is the HTTP code returned for type PostJobsBadRequest
const PostJobsBadRequestCode int = 400

/*
PostJobsBadRequest Bad Request

swagger:response postJobsBadRequest
*/
type PostJobsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewPostJobsBadRequest creates PostJobsBadRequest with default headers values
func NewPostJobsBadRequest() *PostJobsBadRequest {

	return &PostJobsBadRequest{}
}

// WithPayload adds the payload to the post jobs bad request response
func (o *PostJobsBadRequest) WithPayload(payload *model.ErrorResponse) *PostJobsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post jobs bad request response
func (o *PostJobsBadRequest) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostJobsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostJobsDefault Other error with any status code and response body format.

swagger:response postJobsDefault
*/
type PostJobsDefault struct {
	_statusCode int
}

// NewPostJobsDefault creates PostJobsDefault with default headers values
func NewPostJobsDefault(code int) *PostJobsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostJobsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post jobs default response
func (o *PostJobsDefault) WithStatusCode(code int) *PostJobsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post jobs default response
func (o *PostJobsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PostJobsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
