package handlers

import (
	runtime "Car-parking-management-systems"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// DeleteCar handles requests for deleting a car by its ID
func DeleteCar(rt *runtime.Runtime) operations.DeleteCarHandler {
	return &deleteCar{rt: rt}
}

type deleteCar struct {
	rt *runtime.Runtime
}

func (d *deleteCar) Handle(params operations.DeleteCarParams) middleware.Responder {
	if err := d.rt.Service().DeleteCar(params.ID); err != nil {
		switch apiErr := err.(*customErr.APIError); {
		case apiErr.IsError(customErr.NotFound):
			return operations.NewDeleteCarNotFound()
		default:
			return operations.NewDeleteCarInternalServerError()
		}
	}
	return operations.NewDeleteCarNoContent()
}
