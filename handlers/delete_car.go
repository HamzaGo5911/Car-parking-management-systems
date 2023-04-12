package handlers

import (
	runtime "github.com/Car-parking-management-systems"
	"github.com/Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// DeleteCarByID handles requests for deleting a car by its ID
func DeleteCarByID(rt *runtime.Runtime) operations.DeleteCarHandler {
	return &deleteCarByID{rt: rt}
}

type deleteCarByID struct {
	rt *runtime.Runtime
}

func (d *deleteCarByID) Handle(params operations.DeleteCarParams) middleware.Responder {
	logger := log()

	// Delete the car by ID
	if err := d.rt.Service().DeleteCar(params.ID); err != nil {
		logger.Errorf("Failed to delete car: %s", err)
		return operations.NewDeleteCarInternalServerError()
	}

	// Return a success response
	return operations.NewDeleteCarNotFound()
}
