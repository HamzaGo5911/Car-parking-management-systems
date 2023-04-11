package models

import (
	"github.com/fatih/structs"
)

type Car struct {
	ID           uint   `json:"id"`
	LicensePlate string `json:"licensePlate"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	PlateNumber  string `json:"plateNumber"`
}

// GetCarMap converts the Car struct to a map[string]interface{}.
// It returns the resulting map.
func (c *Car) GetCarMap() map[string]interface{} {
	return structs.Map(c)
}
