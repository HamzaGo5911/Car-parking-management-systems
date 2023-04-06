package models

import "testing"

func TestCarModel(t *testing.T) {
	// create a new Car instance
	car := Car{
		LicensePlate: "ABC123",
		Brand:        "Ford",
		CarModel:     "Mustang",
		Color:        "Red",
		PlateNumber:  "123456",
	}

	// check if the car's brand is "Ford"
	if car.Brand != "Ford" {
		t.Errorf("expected brand to be 'Ford', got '%s'", car.Brand)
	}

	// check if the car's license plate is "ABC123"
	if car.LicensePlate != "ABC123" {
		t.Errorf("expected license plate to be 'ABC123', got '%s'", car.LicensePlate)
	}
}
