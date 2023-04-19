package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/gen/restapi/operations"
	"Car-parking-management-systems/models"
)

// NewSaveCar handles request for saving a car
func NewSaveCar(rt *runtime.Runtime) operations.AddCarHandler {
	return &saveCar{rt: rt}
}

type saveCar struct {
	rt *runtime.Runtime
}

// Handle the add car request
func (d *saveCar) Handle(params operations.AddCarParams) middleware.Responder {
	// Create a car model from the request parameters
	car := models.Car{
		LicensePlate: params.Car.LicensePlate,
		Model:        params.Car.Model,
		Brand:        params.Car.Brand,
		Color:        params.Car.Color,
		PlateNumber:  params.Car.PlateNumber,
	}

	// Save the car model into the database
	id, err := d.rt.Service().SaveCar(&car)
	if err != nil {
		// If there is an error, return a bad request response
		log().Errorf("failed to add car: %s", err)
		return operations.NewAddCarBadRequest()
	}

	// Set the car ID from the database and return a successful response
	params.Car.ID = id
	return operations.NewAddCarCreated().WithPayload(params.Car)
}
