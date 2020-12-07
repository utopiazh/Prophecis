// Code generated by go-swagger; DO NOT EDIT.

package inters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserInterceptorHandlerFunc turns a function with the right signature into a user interceptor handler
type UserInterceptorHandlerFunc func(UserInterceptorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserInterceptorHandlerFunc) Handle(params UserInterceptorParams) middleware.Responder {
	return fn(params)
}

// UserInterceptorHandler interface for that can handle valid user interceptor params
type UserInterceptorHandler interface {
	Handle(UserInterceptorParams) middleware.Responder
}

// NewUserInterceptor creates a new http.Handler for the user interceptor operation
func NewUserInterceptor(ctx *middleware.Context, handler UserInterceptorHandler) *UserInterceptor {
	return &UserInterceptor{Context: ctx, Handler: handler}
}

/*UserInterceptor swagger:route GET /cc/v1/inter/user Inters userInterceptor

user interceptor .

Optional extended description in Markdown.

*/
type UserInterceptor struct {
	Context *middleware.Context
	Handler UserInterceptorHandler
}

func (o *UserInterceptor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserInterceptorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
