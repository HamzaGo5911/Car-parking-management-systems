package handlers

import (
	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/gen/models"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// NewUpdateCar handles request for updating a car
func NewUpdateCar(rt *runtime.Runtime) operations.UpdateCarHandler {
	return &updateCar{rt: rt}
}

type updateCar struct {
	rt *runtime.Runtime
}

// Handle the update car request
func (d *updateCar) Handle(params operations.UpdateCarParams) middleware.Responder {
	// Retrieve the car to be updated from the database
	car, err := d.rt.Service().GetCarByID(params.ID)
	if err != nil {
		return operations.NewUpdateCarInternalServerError()
	}

	// Update the car with the request parameters
	car.LicensePlate = params.Car.LicensePlate
	car.Model = params.Car.Model
	car.Brand = params.Car.Brand
	car.Color = params.Car.Color
	car.PlateNumber = params.Car.PlateNumber

	// Save the updated car model into the database
	if _, err := d.rt.Service().SaveCar(car); err != nil {
		return operations.NewUpdateCarInternalServerError()
	}

	// Build the response payload with the updated car details
	payload := &models.Car{
		ID:           car.ID,
		LicensePlate: car.LicensePlate,
		Model:        car.Model,
		Brand:        car.Brand,
		Color:        car.Color,
		PlateNumber:  car.PlateNumber,
	}

	// Return a successful response with the updated car details
	return operations.NewUpdateCarOK().WithPayload(payload)
}
