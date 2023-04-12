package handlers

import (
	runtime "github.com/Car-parking-management-systems"
	"github.com/Car-parking-management-systems/gen/restapi/operations"
	"github.com/Car-parking-management-systems/models"
	"github.com/go-openapi/runtime/middleware"
)

// NewAddCar handles request for saving a car
func NewAddCar(rt *runtime.Runtime) operations.AddCarHandler {
	return &addCar{rt: rt}
}

type addCar struct {
	rt *runtime.Runtime
}

// Handle the add car request
func (d *addCar) Handle(params operations.AddCarParams) middleware.Responder {
	car := models.Car{
		LicensePlate: params.Car.LicensePlate,
		Model:        params.Car.Model,
		Brand:        params.Car.Brand,
		Color:        params.Car.Color,
		PlateNumber:  params.Car.PlateNumber,
	}

	id, err := d.rt.Service().AddCar(&car)
	if err != nil {
		log().Errorf("failed to add car: %s", err)
		return operations.NewAddCarBadRequest()
	}

	params.Car.ID = id
	return operations.NewAddCarCreated().WithPayload(params.Car)
}
