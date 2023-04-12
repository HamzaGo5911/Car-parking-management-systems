package models

import (
	"github.com/fatih/structs"
)

// Car holds information for a car
type Car struct {
	ID           string `json:"id,omitempty" db:"ID" default:""`
	LicensePlate string `json:"licensePlate" db:"LicensePlate"`
	Brand        string `json:"brand" db:"Brand"`
	Model        string `json:"model" db:"Model"`
	Color        string `json:"color" db:"Color"`
	PlateNumber  string `json:"plateNumber" db:"PlateNumber"`
}

// GetCarMap converts the Car struct to a map[string]interface{}.
// It returns the resulting map.
func (c *Car) GetCarMap() map[string]interface{} {
	return structs.Map(c)
}

// Names returns the field names of Car model
func (c *Car) Names() []string {
	fields := structs.Fields(c)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
