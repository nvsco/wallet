// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nvsco/wallet/pkg/scope"
)

// NewGetAppCodesParams creates a new GetAppCodesParams object
// no default values defined in spec.
func NewGetAppCodesParams() *GetAppCodesParams {

	return &GetAppCodesParams{}
}

// GetAppCodesParams contains all the bound params for the get app codes operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAppCodes
type GetAppCodesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	// Body
	RequestBody []byte

	// Scope
	Scope *scope.Scope
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAppCodesParams() beforehand.
func (o *GetAppCodesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}