package db

import (
	"Car-parking-management-systems/models"
	"log"
	"time"
)

// DataStore is an interface for query ops
type DataStore interface {
	SaveCar(car *models.Car) (string, error)
	GetCarByID(id string) (*models.Car, error)
	DeleteCar(id string) error
	SaveParking(parking *models.Parking) (string, error)
	GetParkingByID(id string) (*models.Parking, error)
	GetParkingsOnDate(exitDate time.Time) ([]*models.Parking, error)
	DeleteParking(id string) error
	Close() error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
