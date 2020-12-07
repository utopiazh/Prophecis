// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetAllGroupsOKCode is the HTTP code returned for type GetAllGroupsOK
const GetAllGroupsOKCode int = 200

/*GetAllGroupsOK Detailed Group and Group information.

swagger:response getAllGroupsOK
*/
type GetAllGroupsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PageGroupList `json:"body,omitempty"`
}

// NewGetAllGroupsOK creates GetAllGroupsOK with default headers values
func NewGetAllGroupsOK() *GetAllGroupsOK {

	return &GetAllGroupsOK{}
}

// WithPayload adds the payload to the get all groups o k response
func (o *GetAllGroupsOK) WithPayload(payload *models.PageGroupList) *GetAllGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all groups o k response
func (o *GetAllGroupsOK) SetPayload(payload *models.PageGroupList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllGroupsUnauthorizedCode is the HTTP code returned for type GetAllGroupsUnauthorized
const GetAllGroupsUnauthorizedCode int = 401

/*GetAllGroupsUnauthorized Unauthorized

swagger:response getAllGroupsUnauthorized
*/
type GetAllGroupsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllGroupsUnauthorized creates GetAllGroupsUnauthorized with default headers values
func NewGetAllGroupsUnauthorized() *GetAllGroupsUnauthorized {

	return &GetAllGroupsUnauthorized{}
}

// WithPayload adds the payload to the get all groups unauthorized response
func (o *GetAllGroupsUnauthorized) WithPayload(payload *models.Error) *GetAllGroupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all groups unauthorized response
func (o *GetAllGroupsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllGroupsNotFoundCode is the HTTP code returned for type GetAllGroupsNotFound
const GetAllGroupsNotFoundCode int = 404

/*GetAllGroupsNotFound url to get allGroups not found.

swagger:response getAllGroupsNotFound
*/
type GetAllGroupsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllGroupsNotFound creates GetAllGroupsNotFound with default headers values
func NewGetAllGroupsNotFound() *GetAllGroupsNotFound {

	return &GetAllGroupsNotFound{}
}

// WithPayload adds the payload to the get all groups not found response
func (o *GetAllGroupsNotFound) WithPayload(payload *models.Error) *GetAllGroupsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all groups not found response
func (o *GetAllGroupsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllGroupsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
