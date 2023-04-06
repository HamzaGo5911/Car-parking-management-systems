package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	LicensePlate string    `json:"licensePlate"`
	Brand        string    `json:"brand"`
	CarModel     string    `json:"model"`
	Color        string    `json:"color"`
	PlateNumber  string    `json:"plateNumber"`
	Parking      []Parking `json:"parking,omitempty"`
}
