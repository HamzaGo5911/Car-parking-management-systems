package models

import "gorm.io/gorm"

// Car models
type Car struct {
	gorm.Model
	LicensePlate string    `json:"licensePlate"`
	Brand        string    `json:"brand"`
	Color        string    `json:"color"`
	PlateNumber  string    `json:"plateNumber"`
	Parking      []Parking `json:"parking,omitempty"`
}
