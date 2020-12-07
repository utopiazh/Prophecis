// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserGroupByUserIDAndGroupIDHandlerFunc turns a function with the right signature into a delete user group by user Id and group Id handler
type DeleteUserGroupByUserIDAndGroupIDHandlerFunc func(DeleteUserGroupByUserIDAndGroupIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserGroupByUserIDAndGroupIDHandlerFunc) Handle(params DeleteUserGroupByUserIDAndGroupIDParams) middleware.Responder {
	return fn(params)
}

// DeleteUserGroupByUserIDAndGroupIDHandler interface for that can handle valid delete user group by user Id and group Id params
type DeleteUserGroupByUserIDAndGroupIDHandler interface {
	Handle(DeleteUserGroupByUserIDAndGroupIDParams) middleware.Responder
}

// NewDeleteUserGroupByUserIDAndGroupID creates a new http.Handler for the delete user group by user Id and group Id operation
func NewDeleteUserGroupByUserIDAndGroupID(ctx *middleware.Context, handler DeleteUserGroupByUserIDAndGroupIDHandler) *DeleteUserGroupByUserIDAndGroupID {
	return &DeleteUserGroupByUserIDAndGroupID{Context: ctx, Handler: handler}
}

/*DeleteUserGroupByUserIDAndGroupID swagger:route DELETE /cc/v1/groups/userGroup/user/{userId}/group/{groupId} Groups deleteUserGroupByUserIdAndGroupId

Returns a userGroup.

Optional extended description in Markdown.

*/
type DeleteUserGroupByUserIDAndGroupID struct {
	Context *middleware.Context
	Handler DeleteUserGroupByUserIDAndGroupIDHandler
}

func (o *DeleteUserGroupByUserIDAndGroupID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserGroupByUserIDAndGroupIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
