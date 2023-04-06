package models

import (
	"gorm.io/gorm"
	"time"
)

// Parking model
type Parking struct {
	gorm.Model
	CarID       uint      `json:"car_id"`
	Car         Car       `json:"car" gorm:"foreignKey:CarID"`
	EntryTime   time.Time `json:"entry_time"`
	ExitTime    time.Time `json:"exitTime"`
	HourlyRate  float64   `json:"hourlyRate"`
	TotalAmount float64   `json:"totalAmount"`
}

// CalculateTotalAmount calculates the total amount for the parking record.
// It takes into account the rate and the duration of the parking.
// It returns the total amount and any error encountered during the calculation.
func (p *Parking) CalculateTotalAmount() {
	duration := p.ExitTime.Sub(p.EntryTime)
	hours := duration.Hours()
	p.TotalAmount = hours * p.HourlyRate
}
