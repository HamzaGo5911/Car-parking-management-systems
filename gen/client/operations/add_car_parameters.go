// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"Car-parking-management-systems/gen/models"
)

// NewAddCarParams creates a new AddCarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddCarParams() *AddCarParams {
	return &AddCarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddCarParamsWithTimeout creates a new AddCarParams object
// with the ability to set a timeout on a request.
func NewAddCarParamsWithTimeout(timeout time.Duration) *AddCarParams {
	return &AddCarParams{
		timeout: timeout,
	}
}

// NewAddCarParamsWithContext creates a new AddCarParams object
// with the ability to set a context for a request.
func NewAddCarParamsWithContext(ctx context.Context) *AddCarParams {
	return &AddCarParams{
		Context: ctx,
	}
}

// NewAddCarParamsWithHTTPClient creates a new AddCarParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddCarParamsWithHTTPClient(client *http.Client) *AddCarParams {
	return &AddCarParams{
		HTTPClient: client,
	}
}

/*
AddCarParams contains all the parameters to send to the API endpoint

	for the add car operation.

	Typically these are written to a http.Request.
*/
type AddCarParams struct {

	/* Car.

	   car's details
	*/
	Car *models.Car

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCarParams) WithDefaults() *AddCarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add car params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add car params
func (o *AddCarParams) WithTimeout(timeout time.Duration) *AddCarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add car params
func (o *AddCarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add car params
func (o *AddCarParams) WithContext(ctx context.Context) *AddCarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add car params
func (o *AddCarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add car params
func (o *AddCarParams) WithHTTPClient(client *http.Client) *AddCarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add car params
func (o *AddCarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCar adds the car to the add car params
func (o *AddCarParams) WithCar(car *models.Car) *AddCarParams {
	o.SetCar(car)
	return o
}

// SetCar adds the car to the add car params
func (o *AddCarParams) SetCar(car *models.Car) {
	o.Car = car
}

// WriteToRequest writes these params to a swagger request
func (o *AddCarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Car != nil {
		if err := r.SetBodyParam(o.Car); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
