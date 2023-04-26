package runtime

import (
	"Car-parking-management-systems/db"
	"Car-parking-management-systems/db/mongo"
	"Car-parking-management-systems/services"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	svc *services.Service
}

// NewRuntime creates a new database service
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: services.NewService(store)}, err
}

// Service returns pointer to our service variable
func (r Runtime) Service() *services.Service {
	return r.svc
}
