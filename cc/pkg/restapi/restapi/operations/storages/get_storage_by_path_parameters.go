// Code generated by go-swagger; DO NOT EDIT.

package storages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStorageByPathParams creates a new GetStorageByPathParams object
// no default values defined in spec.
func NewGetStorageByPathParams() GetStorageByPathParams {

	return GetStorageByPathParams{}
}

// GetStorageByPathParams contains all the bound params for the get storage by path operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetStorageByPath
type GetStorageByPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the path of Storage
	  Required: true
	  In: query
	*/
	Path string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetStorageByPathParams() beforehand.
func (o *GetStorageByPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPath binds and validates parameter Path from query.
func (o *GetStorageByPathParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("path", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("path", "query", raw); err != nil {
		return err
	}

	o.Path = raw

	return nil
}
