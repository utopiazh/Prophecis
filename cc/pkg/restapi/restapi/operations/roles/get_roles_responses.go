// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetRolesOKCode is the HTTP code returned for type GetRolesOK
const GetRolesOKCode int = 200

/*GetRolesOK Detailed Role and Role information.

swagger:response getRolesOK
*/
type GetRolesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Role `json:"body,omitempty"`
}

// NewGetRolesOK creates GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {

	return &GetRolesOK{}
}

// WithPayload adds the payload to the get roles o k response
func (o *GetRolesOK) WithPayload(payload *models.Role) *GetRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles o k response
func (o *GetRolesOK) SetPayload(payload *models.Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesUnauthorizedCode is the HTTP code returned for type GetRolesUnauthorized
const GetRolesUnauthorizedCode int = 401

/*GetRolesUnauthorized Unauthorized

swagger:response getRolesUnauthorized
*/
type GetRolesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRolesUnauthorized creates GetRolesUnauthorized with default headers values
func NewGetRolesUnauthorized() *GetRolesUnauthorized {

	return &GetRolesUnauthorized{}
}

// WithPayload adds the payload to the get roles unauthorized response
func (o *GetRolesUnauthorized) WithPayload(payload *models.Error) *GetRolesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles unauthorized response
func (o *GetRolesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesNotFoundCode is the HTTP code returned for type GetRolesNotFound
const GetRolesNotFoundCode int = 404

/*GetRolesNotFound url to add namespace not found.

swagger:response getRolesNotFound
*/
type GetRolesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRolesNotFound creates GetRolesNotFound with default headers values
func NewGetRolesNotFound() *GetRolesNotFound {

	return &GetRolesNotFound{}
}

// WithPayload adds the payload to the get roles not found response
func (o *GetRolesNotFound) WithPayload(payload *models.Error) *GetRolesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles not found response
func (o *GetRolesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
