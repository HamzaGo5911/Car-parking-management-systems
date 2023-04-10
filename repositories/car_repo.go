package repositories

import (
	"Car-parking-management-systems/models"
	"errors"

	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) CreateCar(car *models.Car) error {
	if car == nil {
		return errors.New("car cannot be nil")
	}

	return r.db.Create(car).Error
}

func (r *CarRepository) UpdateCar(carID string, car *models.Car) error {
	if carID == "" {
		return errors.New("car ID cannot be empty")
	}
	if car == nil {
		return errors.New("car cannot be nil")
	}

	err := r.db.Model(&models.Car{}).Where("id = ?", carID).Updates(car).Error
	if err != nil {
		return errors.New("car not found")
	}

	return nil
}

func (r *CarRepository) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *CarRepository) DeleteCar(carID string) error {
	if carID == "" {
		return errors.New("car ID cannot be empty")
	}

	err := r.db.Where("id = ?", carID).Delete(&models.Car{}).Error
	if err != nil {
		return errors.New("car not found")
	}

	return nil
}
