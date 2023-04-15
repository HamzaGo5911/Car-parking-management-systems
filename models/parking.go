package models

import (
	"github.com/fatih/structs"
	"time"
)

// Parking model
type Parking struct {
	ID          string    `json:"id" db:"ID"`
	EntryTime   time.Time `json:"entry_time" db:"EntryTime"`
	ExitTime    time.Time `json:"exit_time" db:"ExitTime"`
	HourlyRate  float64   `json:"hourly_rate" db:"HourlyRate"`
	TotalAmount float64   `json:"total_amount" db:"TotalAmount"`
}

// GetParkingMap converts the Parking struct to a map[string]interface{}.
// It returns the resulting map.
func (p *Parking) GetParkingMap() map[string]interface{} {
	return structs.Map(p)
}

// CalculateTotalAmount calculates the total amount for the parking record.
// It takes into account the rate and the duration of the parking.
// It returns the total amount and any error encountered during the calculation.
func (p *Parking) CalculateTotalAmount() {
	duration := p.ExitTime.Sub(p.EntryTime)
	hours := duration.Hours()
	p.TotalAmount = hours * p.HourlyRate
}
