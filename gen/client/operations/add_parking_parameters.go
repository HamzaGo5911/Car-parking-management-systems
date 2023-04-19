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

// NewAddParkingParams creates a new AddParkingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddParkingParams() *AddParkingParams {
	return &AddParkingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddParkingParamsWithTimeout creates a new AddParkingParams object
// with the ability to set a timeout on a request.
func NewAddParkingParamsWithTimeout(timeout time.Duration) *AddParkingParams {
	return &AddParkingParams{
		timeout: timeout,
	}
}

// NewAddParkingParamsWithContext creates a new AddParkingParams object
// with the ability to set a context for a request.
func NewAddParkingParamsWithContext(ctx context.Context) *AddParkingParams {
	return &AddParkingParams{
		Context: ctx,
	}
}

// NewAddParkingParamsWithHTTPClient creates a new AddParkingParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddParkingParamsWithHTTPClient(client *http.Client) *AddParkingParams {
	return &AddParkingParams{
		HTTPClient: client,
	}
}

/*
AddParkingParams contains all the parameters to send to the API endpoint

	for the add parking operation.

	Typically these are written to a http.Request.
*/
type AddParkingParams struct {

	/* Parking.

	   Parking's details
	*/
	Parking *models.Parking

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add parking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddParkingParams) WithDefaults() *AddParkingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add parking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddParkingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add parking params
func (o *AddParkingParams) WithTimeout(timeout time.Duration) *AddParkingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add parking params
func (o *AddParkingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add parking params
func (o *AddParkingParams) WithContext(ctx context.Context) *AddParkingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add parking params
func (o *AddParkingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add parking params
func (o *AddParkingParams) WithHTTPClient(client *http.Client) *AddParkingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add parking params
func (o *AddParkingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParking adds the parking to the add parking params
func (o *AddParkingParams) WithParking(parking *models.Parking) *AddParkingParams {
	o.SetParking(parking)
	return o
}

// SetParking adds the parking to the add parking params
func (o *AddParkingParams) SetParking(parking *models.Parking) {
	o.Parking = parking
}

// WriteToRequest writes these params to a swagger request
func (o *AddParkingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Parking != nil {
		if err := r.SetBodyParam(o.Parking); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}