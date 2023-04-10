package services

import "Car-parking-management-systems/models"

type CarRepositoryInterface interface {
	CreateCar(car *models.Car) error
	UpdateCar(carID string, car *models.Car) error
	GetAllCars() ([]models.Car, error)
	DeleteCar(carID string) error
}
