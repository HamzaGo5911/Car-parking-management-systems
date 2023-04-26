package services

import (
	"Car-parking-management-systems/models"
	"time"
)

// GetParkingsOnDate retrieves a parking from the database by date
func (s *Service) GetParkingsOnDate(exitTime time.Time) ([]*models.Parking, error) {
	return s.db.GetParkingsOnDate(exitTime)
}
