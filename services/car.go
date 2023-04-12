package services

import (
	"github.com/Car-parking-management-systems/models"
)

// AddCar adds a new car to the database via the Service's database client
func (s *Service) AddCar(car *models.Car) (string, error) {
	return s.db.AddCar(car)
}

// GetCarByID retrieves a car from the database by its ID via the Service's database client
func (s *Service) GetCarByID(id string) (*models.Car, error) {
	return s.db.GetCarByID(id)
}

// DeleteCar deletes a car from the database by its ID via the Service's database client
func (s *Service) DeleteCar(id string) error {
	return s.db.DeleteCar(id)
}
