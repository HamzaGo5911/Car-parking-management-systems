package main

import (
	runtime "Car-parking-management-systems"
	"Car-parking-management-systems/config"
	"Car-parking-management-systems/gen/restapi"
	"Car-parking-management-systems/handlers"
	"github.com/go-openapi/loads"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"strconv"
)

func main() {
	// Load environment variables from .env file
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer func(server *restapi.Server) {
		err := server.Shutdown()
		if err != nil {

		}
	}(server)

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
