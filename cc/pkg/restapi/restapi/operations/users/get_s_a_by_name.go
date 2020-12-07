// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSAByNameHandlerFunc turns a function with the right signature into a get s a by name handler
type GetSAByNameHandlerFunc func(GetSAByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSAByNameHandlerFunc) Handle(params GetSAByNameParams) middleware.Responder {
	return fn(params)
}

// GetSAByNameHandler interface for that can handle valid get s a by name params
type GetSAByNameHandler interface {
	Handle(GetSAByNameParams) middleware.Responder
}

// NewGetSAByName creates a new http.Handler for the get s a by name operation
func NewGetSAByName(ctx *middleware.Context, handler GetSAByNameHandler) *GetSAByName {
	return &GetSAByName{Context: ctx, Handler: handler}
}

/*GetSAByName swagger:route GET /cc/v1/users/admin/{name} Users getSAByName

Returns a user.

Optional extended description in Markdown.

*/
type GetSAByName struct {
	Context *middleware.Context
	Handler GetSAByNameHandler
}

func (o *GetSAByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSAByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
