package handlers

import (
	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/loads"
)

// Handler replaces swagger handler
type Handler *operations.CarParkingManagementSystemAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewCarParkingManagementSystemAPI(spec)

	// car handlers
	handler.AddCarHandler = NewSaveCar(rt)
	handler.GetCarHandler = GetCarByID(rt)
	handler.UpdateCarHandler = NewUpdateCar(rt)
	handler.DeleteCarHandler = DeleteCar(rt)
	handler.AddParkingHandler = NewSaveParking(rt)
	handler.DeleteParkingHandler = NewDeleteParking(rt)
	handler.GetParkingHandler = NewGetParkingByID(rt)
	handler.UpdateParkingHandler = NewUpdateParking(rt)
	handler.GetParkingByDateHandler = GetParkingsByDate(rt)

	return handler
}
