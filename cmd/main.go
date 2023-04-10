package main

import (
	"Car-parking-management-systems/config"
	"Car-parking-management-systems/handlers"
	"Car-parking-management-systems/repositories"
	"Car-parking-management-systems/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Create a new instance of the CarRepository
	carRepo := repositories.NewCarRepository(db)

	// Create a new instance of the CarService
	carService := services.NewCarService(carRepo)

	// Create a new router instance
	r := mux.NewRouter()

	// Register the carHandler routes
	carHandler := handlers.NewCarHandler(carService)
	r.HandleFunc("/api/cars", carHandler.CreateCar).Methods(http.MethodPost)
	r.HandleFunc("/api/cars/{id}", carHandler.UpdateCar).Methods(http.MethodPut)
	r.HandleFunc("/api/get-cars", carHandler.GetAllCars).Methods(http.MethodGet)
	r.HandleFunc("/api/cars/{id}", carHandler.DeleteCar).Methods(http.MethodDelete)
	fmt.Println("Server is starting")
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
