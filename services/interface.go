package services

import (
	"Car-parking-management-systems/db"
)

// Service represents the car parking management service
type Service struct {
	db db.DataStore
}

// NewService creates a new instance of the car service
// and returns a pointer to the instance
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
