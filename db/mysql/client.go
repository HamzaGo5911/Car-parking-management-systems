package mysql

import (
	"database/sql"
	"fmt"
	"github.com/Car-parking-management-systems/config"
	"github.com/Car-parking-management-systems/db"
	"github.com/Car-parking-management-systems/models"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

const (
	carTableName = "car"
)

// Register the MySQL database connection
func init() {
	db.Register("mysql", NewClient)
}

// Client is a MySQL database client
type client struct {
	db *sqlx.DB
}

// formatDSN formats the DSN string for the MySQL database connection
func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a MySQL database connection
func NewClient(db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mysql://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Info("initializing mysql connection" + uri)
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

// AddCar adds a car record to the MySQL database
func (c *client) AddCar(car *models.Car) (string, error) {
	car.ID = uuid.NewString() // Generate a unique ID for the new car

	names := car.Names()
	placeholders := make([]string, len(names))
	for i, name := range names {
		placeholders[i] = ":" + name
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", carTableName, strings.Join(names, ","), strings.Join(placeholders, ","))

	result, err := c.db.NamedExec(query, car)
	if err != nil {
		return "", errors.Wrap(err, "failed to add car")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected != 1 {
		return "", errors.New("unexpected number of rows affected")
	}

	return car.ID, nil
}

// GetCarByID retrieves a car record by ID from the MySQL database
func (c *client) GetCarByID(id string) (*models.Car, error) {
	var car models.Car
	err := c.db.Get(&car, "SELECT * FROM "+carTableName+" WHERE ID = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("car with ID %s not found", id)
		}
		return nil, errors.Wrap(err, "failed to get car by ID")
	}

	return &car, nil
}

// DeleteCar deletes a car record from the MySQL database by ID
func (c *client) DeleteCar(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=:id", carTableName)
	_, err := c.db.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return errors.Wrap(err, "failed to delete car")
	}
	return nil
}
