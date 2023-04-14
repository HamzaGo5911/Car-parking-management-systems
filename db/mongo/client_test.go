package mongo

import (
	"Car-parking-management-systems/db"
	"Car-parking-management-systems/models"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_client_SaveCar(t *testing.T) {
	// Set up environment variables for testing
	err := os.Setenv("DB_PORT", "27017")
	if err != nil {
		return
	}
	err = os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return
	}
	// Define input arguments and expected output for the test cases
	type args struct {
		car *models.Car
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add car in db",
			args:    args{car: &models.Car{Model: "Tesla", Color: "Red", PlateNumber: "ABC123"}},
			wantErr: false,
		},
		{
			name:    "success - update existing car in db",
			args:    args{car: &models.Car{ID: "1", Model: "Tesla", Color: "Red", PlateNumber: "DEF456"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid car in db",
			args:    args{car: &models.Car{ID: "1", Model: "", Color: "", PlateNumber: ""}},
			wantErr: true,
		},
	}

	// Loop through the test cases and run them one by one
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new instance of the MongoDB client
			c, _ := NewClient(db.Option{})
			// Call the SaveCar method with the input arguments
			_, err := c.SaveCar(tt.args.car)
			// Check if the output matches the expected output
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_GetCarById(t *testing.T) {
	err := os.Setenv("DB_PORT", "27017")
	if err != nil {
		return
	}
	err = os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return
	}

	c, _ := NewClient(db.Option{})
	car := &models.Car{Model: "Tesla", Color: "Red", PlateNumber: "ABC123"}
	_, _ = c.SaveCar(car)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Car
		wantErr bool
	}{
		{
			name:    "success - get car from db",
			args:    args{id: car.ID},
			want:    car,
			wantErr: false,
		},
		{
			name:    "fail - get invalid car from db",
			args:    args{id: "invalid-car-id"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetCarById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCar() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_client_DeleteCar(t *testing.T) {
	err := os.Setenv("DB_PORT", "27017")
	if err != nil {
		return
	}
	err = os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return
	}

	c, _ := NewClient(db.Option{})
	car := &models.Car{Model: "Tesla", Color: "Red", PlateNumber: "ABC123"}
	id, _ := c.SaveCar(car)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete car from db",
			args:    args{id: id},
			wantErr: false,
		},
		{
			name:    "fail - delete invalid car from db",
			args:    args{id: "invalid-car-id"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.DeleteCar(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestSaveParking(t *testing.T) {
	// Create a new Parking instance with sample values
	p := &models.Parking{
		ID:          "abc123",
		EntryTime:   time.Now(),
		ExitTime:    time.Now().Add(time.Hour),
		HourlyRate:  2.50,
		TotalAmount: 5.00,
	}

	// Save the parking in the database
	c, _ := NewClient(db.Option{})
	id, err := c.SaveParking(p)
	if err != nil {
		t.Errorf("Failed to save parking: %v", err)
	}

	// Test the ID field
	if id != "abc123" {
		t.Errorf("Expected ID to be 'abc123', got %s", id)
	}

	// Update the parking
	p.HourlyRate = 3.00
	p.TotalAmount = 6.00
	id, err = c.SaveParking(p)
	if err != nil {
		t.Errorf("Failed to update parking: %v", err)
	}

	// Test the updated HourlyRate and TotalAmount fields
	if p.HourlyRate != 3.00 {
		t.Errorf("Expected HourlyRate to be 3.00, got %f", p.HourlyRate)
	}
	if p.TotalAmount != 6.00 {
		t.Errorf("Expected TotalAmount to be 6.00, got %f", p.TotalAmount)
	}
}

func TestClient_DeleteParking(t *testing.T) {
	// Create a new Client instance with a sample Parking object
	c, err := NewClient(db.Option{})
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}

	p := &models.Parking{
		ID:          "abc123",
		EntryTime:   time.Now(),
		ExitTime:    time.Now().Add(time.Hour),
		HourlyRate:  2.50,
		TotalAmount: 5.00,
	}
	// Delete the parking object
	err = c.DeleteParking(p.ID)
	if err != nil {
		t.Fatalf("Failed to delete parking: %v", err)
	}
	// Try to retrieve the deleted parking object from the database
	deletedParking, err := c.GetParkingById(p.ID)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
	if deletedParking != nil {
		t.Errorf("Expected nil Parking object, but got a non-nil object")
	}
}
func TestClient_GetParkingById(t *testing.T) {
	// Create a new Client instance
	c, err := NewClient(db.Option{})
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}

	// Create a new Parking instance with sample values
	p := &models.Parking{
		ID:          "abc123",
		EntryTime:   time.Now(),
		ExitTime:    time.Now().Add(time.Hour),
		HourlyRate:  2.50,
		TotalAmount: 5.00,
	}

	// Test the GetParking() function with an existing parking ID
	existingParking, err := c.GetParkingById(p.ID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if existingParking == nil {
		t.Error("Expected non-nil Parking object, but got nil")
	} else if *existingParking != *p {
		t.Errorf("Expected %v, but got %v", p, existingParking)
	}

	// Test the GetParking() function with a non-existing parking ID
	nonExistingParking, err := c.GetParkingById("def456")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
	if nonExistingParking != nil {
		t.Error("Expected nil Parking object, but got a non-nil object")
	}
}
