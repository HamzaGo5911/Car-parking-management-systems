package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Car-parking-management-systems/config"
	"Car-parking-management-systems/db"
	customErr "Car-parking-management-systems/errors"
	"Car-parking-management-systems/models"
)

const (
	carCollection     = "car"
	parkingCollection = "parking"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn       *mongo.Client
	collection *mongo.Collection
}

// Close disconnects the client from the MongoDB server
func (c *client) Close() error {
	return c.conn.Disconnect(context.Background())
}

// NewClient initializes a mongodb database connection
func NewClient(db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err := cli.Ping(context.Background(), nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}

	return &client{conn: cli}, nil
}

// SaveCar saves a car in the database. If the car has an ID, it is updated, otherwise it is inserted as a new record.
func (c *client) SaveCar(car *models.Car) (string, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	if car.ID == "" { // If the car does not have an ID, generate a new one
		car.ID = uuid.NewV4().String()
		if _, err := collection.InsertOne(context.Background(), car); err != nil { // Insert the car as a new record
			return "", errors.Wrap(err, "failed to add car")
		}
	} else { // If the car has an ID, update the existing record
		filter := bson.M{"id": car.ID}
		update := bson.M{"$set": car}
		if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
			return "", errors.Wrap(err, "failed to update car")
		}
	}
	return car.ID, nil // Return the ID of the saved car
}

func (c *client) GetCarById(id string) (*models.Car, error) {
	var car *models.Car
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&car); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, customErr.NewAPIError(customErr.NotFound, fmt.Sprintf("car: %s not found", id))
		}
		return nil, errors.Wrap(err, "failed to get car")
	}
	return car, nil
}

func (c *client) DeleteCar(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	if _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete car")
	}

	return nil
}

func (c *client) SaveParking(parking *models.Parking) (string, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(parkingCollection)
	if parking.ID == "" {
		parking.ID = uuid.NewV4().String()
		if _, err := collection.InsertOne(context.Background(), parking); err != nil {
			return "", errors.Wrap(err, "failed to add parking")
		}
	} else {
		filter := bson.M{"id": parking.ID}
		update := bson.M{"$set": parking}
		if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
			return "", errors.Wrap(err, "failed to update parking")
		}
	}
	return parking.ID, nil
}

func (c *client) DeleteParking(parkingID string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(parkingCollection)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": parkingID})
	if err != nil {
		return errors.Wrap(err, "failed to delete parking")
	}
	return nil
}
func (c *client) GetParkingById(parkingID string) (*models.Parking, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(parkingCollection)
	filter := bson.M{"id": parkingID}
	var parking models.Parking
	err := collection.FindOne(context.Background(), filter).Decode(&parking)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get parking")
	}
	return &parking, nil
}
