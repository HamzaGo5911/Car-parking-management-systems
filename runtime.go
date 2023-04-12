package runtime

import (
	"github.com/Car-parking-management-systems/db"
	"github.com/Car-parking-management-systems/db/mysql"
	"github.com/Car-parking-management-systems/services"
)

// Runtime contains the service layer for our application
type Runtime struct {
	svc *services.Service
}

// NewRuntime creates a new database service
func NewRuntime() (*Runtime, error) {
	store, err := mysql.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: services.NewCarService(store)}, nil
}

// Service returns pointer to our service variable
func (r *Runtime) Service() *services.Service {
	return r.svc
}
