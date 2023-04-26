package handlers

import (
	"Car-parking-management-systems"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/gen/models"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewUpdateParking handles request for updating a car
func NewUpdateParking(rt *runtime.Runtime) operations.UpdateParkingHandler {
	return &updateParking{rt: rt}
}

type updateParking struct {
	rt *runtime.Runtime
}

func (d *updateParking) Handle(params operations.UpdateParkingParams) middleware.Responder {
	parking, err := d.rt.Service().GetParkingByID(params.ID)
	if err != nil {
		switch apiErr := err.(*customErr.APIError); {
		case apiErr.IsError(customErr.NotFound):
			return operations.NewUpdateParkingNotFound()
		default:
			return operations.NewUpdateParkingInternalServerError()
		}
	}

	// Update the parking document
	parking.HourlyRate = params.Parking.HourlyRate
	parking.TotalAmount = params.Parking.TotalAmount

	// Save the updated document to the database
	id, err := d.rt.Service().SaveParking(parking)
	if err != nil {
		log().Errorf("failed to update parking: %s", err)
		return operations.NewUpdateParkingBadRequest()
	}

	// Set the parking ID from the insert result and return a successful response
	params.Parking.ID = id
	return operations.NewUpdateParkingOK().WithPayload(&models.Parking{
		ID:          parking.ID,
		EntryTime:   strfmt.DateTime(parking.EntryTime),
		ExitTime:    strfmt.DateTime(parking.ExitTime),
		HourlyRate:  parking.HourlyRate,
		TotalAmount: parking.TotalAmount,
	})
}
