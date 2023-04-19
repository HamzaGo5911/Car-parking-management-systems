package handlers

import (
	runtime "Car-parking-management-systems"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/gen/models"
	"Car-parking-management-systems/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

// GetParkingsByDate handles request for retrieving parkings by exit date
func GetParkingsByDate(rt *runtime.Runtime) operations.GetParkingByDateHandler {
	return &getParkingsByDate{rt: rt}
}

type getParkingsByDate struct {
	rt *runtime.Runtime
}

func (d *getParkingsByDate) Handle(params operations.GetParkingByDateParams) middleware.Responder {
	exitTime, err := strfmt.ParseDateTime(params.ExitTime)
	if err != nil {
		return operations.NewGetParkingByDateNotFound()
	}

	parkings, err := d.rt.Service().GetParkingsOnDate(time.Time(exitTime))
	if err != nil {
		switch apiErr := err.(*customErr.APIError); {
		case apiErr.IsError(customErr.NotFound):
			return operations.NewGetParkingByDateNotFound()
		default:
			return operations.NewGetParkingByDateInternalServerError()
		}
	}

	var payload []*models.Parking
	for _, parking := range parkings {
		payload = append(payload, &models.Parking{
			ID:          parking.ID,
			EntryTime:   strfmt.DateTime(parking.EntryTime),
			ExitTime:    strfmt.DateTime(parking.ExitTime),
			HourlyRate:  parking.HourlyRate,
			TotalAmount: parking.TotalAmount,
		})
	}

	return operations.NewGetParkingByDateOK().WithPayload(payload)
}
