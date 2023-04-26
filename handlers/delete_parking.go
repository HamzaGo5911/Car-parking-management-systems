package handlers

import (
	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// NewDeleteParking handles requests for deleting a parking by its ID
func NewDeleteParking(rt *runtime.Runtime) operations.DeleteParkingHandler {
	return &deleteParking{rt: rt}
}

type deleteParking struct {
	rt *runtime.Runtime
}

// Handle the delete parking request
func (d *deleteParking) Handle(params operations.DeleteParkingParams) middleware.Responder {
	// Delete the parking document from the database using the parking ID
	err := d.rt.Service().DeleteParking(params.ID)
	if err != nil {
		// If there is an error, return a bad request response
		log().Errorf("failed to delete parking: %s", err)
		return operations.NewDeleteCarNoContent()
	}

	// Return a successful response
	return operations.NewDeleteParkingOK()
}
