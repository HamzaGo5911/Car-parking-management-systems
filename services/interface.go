package services

import "github.com/Car-parking-management-systems/db"

type Service struct {
	db db.DataStore
}

// Service creates a new instance of the car service

func NewCarService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
