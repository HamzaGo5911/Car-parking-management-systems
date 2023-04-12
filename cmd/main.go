package main

import (
	runtime "github.com/Car-parking-management-systems"
	"github.com/Car-parking-management-systems/config"
	"github.com/Car-parking-management-systems/gen/restapi"
	"github.com/Car-parking-management-systems/handlers"
	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// Load the Swagger spec
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	// Create a new runtime
	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	// Create a new API handler with the runtime and Swagger spec
	api := handlers.NewHandler(rt, swaggerSpec)

	// Create a new REST API server with the API handler
	server := restapi.NewServer(api)
	defer func(server *restapi.Server) {
		// Shut down the server when done
		err := server.Shutdown()
		if err != nil {
			// Handle any errors that occur during shutdown
		}
	}(server)

	// Configure the server's host and port from configuration
	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}

	// Configure the server's API
	server.ConfigureAPI()

	// Start serving requests
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
