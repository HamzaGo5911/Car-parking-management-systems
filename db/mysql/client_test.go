package mysql

import (
	"github.com/Car-parking-management-systems/db"
	"github.com/Car-parking-management-systems/models"
	"os"
	"testing"
)

func TestClient_AddCar(t *testing.T) {
	err := os.Setenv("DB_PORT", "3306")
	if err != nil {
		return
	}
	err = os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return
	}
	err = os.Setenv("DB_USER", "root")
	if err != nil {
		return
	}
	type args struct {
		car *models.Car
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add car to db",
			args: args{
				car: &models.Car{
					LicensePlate: "1002",
					Brand:        "Toyota",
					Model:        "Camry",
					Color:        "White",
					PlateNumber:  "ABC-123",
				},
			},
			wantErr: false,
		},
		{
			name: "fail - add invalid car to db",
			args: args{
				car: &models.Car{
					LicensePlate: "",
					Brand:        "Toyota",
					Model:        "Camry",
					Color:        "White",
					PlateNumber:  "ABC-123",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddCar(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
