package main

import (
	"Car-parking-management-systems/db"
	"Car-parking-management-systems/db/mongo"
	"github.com/subosito/gotenv"
	"log"
)

func main() {
	// Load environment variables from .env file
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Initialize the database client
	dbClient, err := mongo.NewClient(db.Option{})
	if err != nil {
		log.Fatalf("failed to initialize database client: %v", err)
	}
	defer func() {
		if err := dbClient.Close(); err != nil {
			log.Fatalf("failed to close database client: %v", err)
		}
	}()
}
