// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// CheckNamespaceUserOKCode is the HTTP code returned for type CheckNamespaceUserOK
const CheckNamespaceUserOKCode int = 200

/*CheckNamespaceUserOK auth by namespace.

swagger:response checkNamespaceUserOK
*/
type CheckNamespaceUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewCheckNamespaceUserOK creates CheckNamespaceUserOK with default headers values
func NewCheckNamespaceUserOK() *CheckNamespaceUserOK {

	return &CheckNamespaceUserOK{}
}

// WithPayload adds the payload to the check namespace user o k response
func (o *CheckNamespaceUserOK) WithPayload(payload *models.Result) *CheckNamespaceUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace user o k response
func (o *CheckNamespaceUserOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckNamespaceUserUnauthorizedCode is the HTTP code returned for type CheckNamespaceUserUnauthorized
const CheckNamespaceUserUnauthorizedCode int = 401

/*CheckNamespaceUserUnauthorized Unauthorized

swagger:response checkNamespaceUserUnauthorized
*/
type CheckNamespaceUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckNamespaceUserUnauthorized creates CheckNamespaceUserUnauthorized with default headers values
func NewCheckNamespaceUserUnauthorized() *CheckNamespaceUserUnauthorized {

	return &CheckNamespaceUserUnauthorized{}
}

// WithPayload adds the payload to the check namespace user unauthorized response
func (o *CheckNamespaceUserUnauthorized) WithPayload(payload *models.Error) *CheckNamespaceUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace user unauthorized response
func (o *CheckNamespaceUserUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckNamespaceUserNotFoundCode is the HTTP code returned for type CheckNamespaceUserNotFound
const CheckNamespaceUserNotFoundCode int = 404

/*CheckNamespaceUserNotFound url to add namespace not found.

swagger:response checkNamespaceUserNotFound
*/
type CheckNamespaceUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckNamespaceUserNotFound creates CheckNamespaceUserNotFound with default headers values
func NewCheckNamespaceUserNotFound() *CheckNamespaceUserNotFound {

	return &CheckNamespaceUserNotFound{}
}

// WithPayload adds the payload to the check namespace user not found response
func (o *CheckNamespaceUserNotFound) WithPayload(payload *models.Error) *CheckNamespaceUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check namespace user not found response
func (o *CheckNamespaceUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckNamespaceUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
