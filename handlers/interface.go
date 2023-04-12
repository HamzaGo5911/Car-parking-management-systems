package handlers

import (
	runtime "github.com/Car-parking-management-systems"
	"github.com/Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/loads"
)

// Handler replaces swagger handler
type Handler *operations.CarParkingAPIAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewCarParkingAPIAPI(spec)

	// car handlers
	handler.AddCarHandler = NewAddCar(rt)
	handler.GetCarHandler = GetCarByID(rt)
	handler.DeleteCarHandler = DeleteCarByID(rt)

	return handler
}
