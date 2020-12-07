// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetAllNamespacesOKCode is the HTTP code returned for type GetAllNamespacesOK
const GetAllNamespacesOKCode int = 200

/*GetAllNamespacesOK Detailed namespace and namespace information.

swagger:response getAllNamespacesOK
*/
type GetAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.PageNamespaceList `json:"body,omitempty"`
}

// NewGetAllNamespacesOK creates GetAllNamespacesOK with default headers values
func NewGetAllNamespacesOK() *GetAllNamespacesOK {

	return &GetAllNamespacesOK{}
}

// WithPayload adds the payload to the get all namespaces o k response
func (o *GetAllNamespacesOK) WithPayload(payload *models.PageNamespaceList) *GetAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all namespaces o k response
func (o *GetAllNamespacesOK) SetPayload(payload *models.PageNamespaceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllNamespacesUnauthorizedCode is the HTTP code returned for type GetAllNamespacesUnauthorized
const GetAllNamespacesUnauthorizedCode int = 401

/*GetAllNamespacesUnauthorized Unauthorized

swagger:response getAllNamespacesUnauthorized
*/
type GetAllNamespacesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllNamespacesUnauthorized creates GetAllNamespacesUnauthorized with default headers values
func NewGetAllNamespacesUnauthorized() *GetAllNamespacesUnauthorized {

	return &GetAllNamespacesUnauthorized{}
}

// WithPayload adds the payload to the get all namespaces unauthorized response
func (o *GetAllNamespacesUnauthorized) WithPayload(payload *models.Error) *GetAllNamespacesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all namespaces unauthorized response
func (o *GetAllNamespacesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllNamespacesNotFoundCode is the HTTP code returned for type GetAllNamespacesNotFound
const GetAllNamespacesNotFoundCode int = 404

/*GetAllNamespacesNotFound url to get allNamespaces not found.

swagger:response getAllNamespacesNotFound
*/
type GetAllNamespacesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllNamespacesNotFound creates GetAllNamespacesNotFound with default headers values
func NewGetAllNamespacesNotFound() *GetAllNamespacesNotFound {

	return &GetAllNamespacesNotFound{}
}

// WithPayload adds the payload to the get all namespaces not found response
func (o *GetAllNamespacesNotFound) WithPayload(payload *models.Error) *GetAllNamespacesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all namespaces not found response
func (o *GetAllNamespacesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllNamespacesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
