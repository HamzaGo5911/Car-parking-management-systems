package services

import (
	"Car-parking-management-systems/models"
)

// SaveParking adds a new parking to the database via the Service's database client
func (s *Service) SaveParking(parking *models.Parking) (string, error) {
	return s.db.SaveParking(parking)
}

// GetParkingByID retrieves a parking from the database by its ID via the Service's database client
func (s *Service) GetParkingByID(id string) (*models.Parking, error) {
	return s.db.GetParkingByID(id)
}

// DeleteParking deletes a parking from the database by its ID via the Service's database client
func (s *Service) DeleteParking(id string) error {
	return s.db.DeleteParking(id)
}
