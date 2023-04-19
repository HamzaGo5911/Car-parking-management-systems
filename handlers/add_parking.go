package handlers

import (
	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/gen/restapi/operations"
	"Car-parking-management-systems/models"
	"github.com/go-openapi/runtime/middleware"
	"time"
)

// NewSaveParking handles request for saving a parking
func NewSaveParking(rt *runtime.Runtime) operations.AddParkingHandler {
	return &createParking{rt: rt}
}

type createParking struct {
	rt *runtime.Runtime
}

// Handle the create parking request
func (d *createParking) Handle(params operations.AddParkingParams) middleware.Responder {
	// Set the entry time to the current time
	entryTime := time.Now().UTC()

	// Create a parking document to insert
	parkingDoc := &models.Parking{
		EntryTime:  entryTime,
		HourlyRate: params.Parking.HourlyRate,
	}

	// Insert the parking document into the database
	id, err := d.rt.Service().SaveParking(parkingDoc)
	if err != nil {
		// If there is an error, return a bad request response
		log().Errorf("failed to create parking: %s", err)
		return operations.NewAddParkingBadRequest()
	}

	// Set the parking ID from the insert result and return a successful response
	params.Parking.ID = id
	return operations.NewAddParkingCreated().WithPayload(params.Parking)
}
