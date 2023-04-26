package services

import "Car-parking-management-systems/models"

// SaveCar adds a new car to the database via the Service's database client
func (s *Service) SaveCar(car *models.Car) (string, error) {
	return s.db.SaveCar(car)
}

// GetCarByID retrieves a car from the database by its ID via the Service's database client
func (s *Service) GetCarByID(id string) (*models.Car, error) {
	return s.db.GetCarByID(id)

}

// DeleteCar deletes a car from the database by its ID via the Service's database client
func (s *Service) DeleteCar(id string) error {
	return s.db.DeleteCar(id)
}
