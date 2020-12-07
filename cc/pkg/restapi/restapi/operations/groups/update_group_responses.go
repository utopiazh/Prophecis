// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// UpdateGroupOKCode is the HTTP code returned for type UpdateGroupOK
const UpdateGroupOKCode int = 200

/*UpdateGroupOK Detailed group and group information.

swagger:response updateGroupOK
*/
type UpdateGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewUpdateGroupOK creates UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {

	return &UpdateGroupOK{}
}

// WithPayload adds the payload to the update group o k response
func (o *UpdateGroupOK) WithPayload(payload *models.Group) *UpdateGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group o k response
func (o *UpdateGroupOK) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateGroupUnauthorizedCode is the HTTP code returned for type UpdateGroupUnauthorized
const UpdateGroupUnauthorizedCode int = 401

/*UpdateGroupUnauthorized Unauthorized

swagger:response updateGroupUnauthorized
*/
type UpdateGroupUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateGroupUnauthorized creates UpdateGroupUnauthorized with default headers values
func NewUpdateGroupUnauthorized() *UpdateGroupUnauthorized {

	return &UpdateGroupUnauthorized{}
}

// WithPayload adds the payload to the update group unauthorized response
func (o *UpdateGroupUnauthorized) WithPayload(payload *models.Error) *UpdateGroupUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group unauthorized response
func (o *UpdateGroupUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateGroupNotFoundCode is the HTTP code returned for type UpdateGroupNotFound
const UpdateGroupNotFoundCode int = 404

/*UpdateGroupNotFound url to get allGroups not found.

swagger:response updateGroupNotFound
*/
type UpdateGroupNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateGroupNotFound creates UpdateGroupNotFound with default headers values
func NewUpdateGroupNotFound() *UpdateGroupNotFound {

	return &UpdateGroupNotFound{}
}

// WithPayload adds the payload to the update group not found response
func (o *UpdateGroupNotFound) WithPayload(payload *models.Error) *UpdateGroupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group not found response
func (o *UpdateGroupNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
