package services

import (
	"Car-parking-management-systems/models"
	"errors"
)

type CarService struct {
	carRepository CarRepositoryInterface
}

func NewCarService(carRepository CarRepositoryInterface) *CarService {
	return &CarService{carRepository: carRepository}
}

func (s *CarService) CreateCar(car *models.Car) error {
	if car == nil {
		return errors.New("car cannot be nil")
	}
	return s.carRepository.CreateCar(car)
}

func (s *CarService) UpdateCar(carID string, car *models.Car) error {
	if carID == "" {
		return errors.New("car ID cannot be empty")
	}
	if car == nil {
		return errors.New("car cannot be nil")
	}
	return s.carRepository.UpdateCar(carID, car)
}

func (s *CarService) GetAllCars() ([]models.Car, error) {
	return s.carRepository.GetAllCars()
}

func (s *CarService) DeleteCar(carID string) error {
	if carID == "" {
		return errors.New("car ID cannot be empty")
	}
	return s.carRepository.DeleteCar(carID)
}
