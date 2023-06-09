// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"Car-parking-management-systems/gen/models"
)

// GetParkingByDateOKCode is the HTTP code returned for type GetParkingByDateOK
const GetParkingByDateOKCode int = 200

/*
GetParkingByDateOK Successful response

swagger:response getParkingByDateOK
*/
type GetParkingByDateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Parking `json:"body,omitempty"`
}

// NewGetParkingByDateOK creates GetParkingByDateOK with default headers values
func NewGetParkingByDateOK() *GetParkingByDateOK {

	return &GetParkingByDateOK{}
}

// WithPayload adds the payload to the get parking by date o k response
func (o *GetParkingByDateOK) WithPayload(payload []*models.Parking) *GetParkingByDateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get parking by date o k response
func (o *GetParkingByDateOK) SetPayload(payload []*models.Parking) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetParkingByDateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Parking, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetParkingByDateNotFoundCode is the HTTP code returned for type GetParkingByDateNotFound
const GetParkingByDateNotFoundCode int = 404

/*
GetParkingByDateNotFound Parking not found

swagger:response getParkingByDateNotFound
*/
type GetParkingByDateNotFound struct {
}

// NewGetParkingByDateNotFound creates GetParkingByDateNotFound with default headers values
func NewGetParkingByDateNotFound() *GetParkingByDateNotFound {

	return &GetParkingByDateNotFound{}
}

// WriteResponse to the client
func (o *GetParkingByDateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetParkingByDateInternalServerErrorCode is the HTTP code returned for type GetParkingByDateInternalServerError
const GetParkingByDateInternalServerErrorCode int = 500

/*
GetParkingByDateInternalServerError Internal server error

swagger:response getParkingByDateInternalServerError
*/
type GetParkingByDateInternalServerError struct {
}

// NewGetParkingByDateInternalServerError creates GetParkingByDateInternalServerError with default headers values
func NewGetParkingByDateInternalServerError() *GetParkingByDateInternalServerError {

	return &GetParkingByDateInternalServerError{}
}

// WriteResponse to the client
func (o *GetParkingByDateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
