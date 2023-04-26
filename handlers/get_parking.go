package handlers

import (
	runtime "Car-parking-management-systems"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/gen/models"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetParkingByID handles request for retrieving a parking by ID
func NewGetParkingByID(rt *runtime.Runtime) operations.GetParkingHandler {
	return &getParkingByID{rt: rt}
}

type getParkingByID struct {
	rt *runtime.Runtime
}

// Handle the get parking by ID request
func (d *getParkingByID) Handle(params operations.GetParkingParams) middleware.Responder {
	parking, err := d.rt.Service().GetParkingByID(params.ID)
	if err != nil {
		switch apiErr := err.(*customErr.APIError); {
		case apiErr.IsError(customErr.NotFound):
			return operations.NewGetParkingNotFound()
		default:
			return operations.NewGetParkingInternalServerError()
		}
	}

	payload := &models.Parking{
		ID:          parking.ID,
		EntryTime:   strfmt.DateTime(parking.EntryTime),
		ExitTime:    strfmt.DateTime(parking.ExitTime),
		HourlyRate:  parking.HourlyRate,
		TotalAmount: parking.TotalAmount,
	}

	return operations.NewGetParkingOK().WithPayload(payload)
}
