// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetParkingByDateParams creates a new GetParkingByDateParams object
//
// There are no default values defined in the spec.
func NewGetParkingByDateParams() GetParkingByDateParams {

	return GetParkingByDateParams{}
}

// GetParkingByDateParams contains all the bound params for the get parking by date operation
// typically these are obtained from a http.Request
//
// swagger:parameters getParkingByDate
type GetParkingByDateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Entry time of the parking
	  Required: true
	  In: path
	*/
	ExitTime string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetParkingByDateParams() beforehand.
func (o *GetParkingByDateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rExitTime, rhkExitTime, _ := route.Params.GetOK("exit_time")
	if err := o.bindExitTime(rExitTime, rhkExitTime, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExitTime binds and validates parameter ExitTime from path.
func (o *GetParkingByDateParams) bindExitTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ExitTime = raw

	return nil
}
