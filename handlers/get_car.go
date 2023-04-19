package handlers

import (
	runtime "Car-parking-management-systems"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/gen/models"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetCarByID handles request for retrieving a car by ID
func GetCarByID(rt *runtime.Runtime) operations.GetCarHandler {
	return &getCarByID{rt: rt}
}

type getCarByID struct {
	rt *runtime.Runtime
}

// Handle the get car by ID request
func (d *getCarByID) Handle(params operations.GetCarParams) middleware.Responder {
	car, err := d.rt.Service().GetCarByID(params.ID)
	if err != nil {
		switch apiErr := err.(*customErr.APIError); {
		case apiErr.IsError(customErr.NotFound):
			return operations.NewGetCarNotFound()
		default:
			return operations.NewGetCarInternalServerError()
		}
	}

	payload := &models.Car{
		ID:           car.ID,
		LicensePlate: car.LicensePlate,
		Model:        car.Model,
		Brand:        car.Brand,
		Color:        car.Color,
		PlateNumber:  car.PlateNumber,
	}

	return operations.NewGetCarOK().WithPayload(payload)
}
