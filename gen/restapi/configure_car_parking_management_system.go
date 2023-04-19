// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Car-parking-management-systems/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name CarParkingManagementSystem --spec ../../swagger.yaml --principal interface{}

func configureFlags(api *operations.CarParkingManagementSystemAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CarParkingManagementSystemAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddCarHandler == nil {
		api.AddCarHandler = operations.AddCarHandlerFunc(func(params operations.AddCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddCar has not yet been implemented")
		})
	}
	if api.AddParkingHandler == nil {
		api.AddParkingHandler = operations.AddParkingHandlerFunc(func(params operations.AddParkingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddParking has not yet been implemented")
		})
	}
	if api.DeleteCarHandler == nil {
		api.DeleteCarHandler = operations.DeleteCarHandlerFunc(func(params operations.DeleteCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteCar has not yet been implemented")
		})
	}
	if api.DeleteParkingHandler == nil {
		api.DeleteParkingHandler = operations.DeleteParkingHandlerFunc(func(params operations.DeleteParkingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteParking has not yet been implemented")
		})
	}
	if api.GetCarHandler == nil {
		api.GetCarHandler = operations.GetCarHandlerFunc(func(params operations.GetCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCar has not yet been implemented")
		})
	}
	if api.GetParkingHandler == nil {
		api.GetParkingHandler = operations.GetParkingHandlerFunc(func(params operations.GetParkingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetParking has not yet been implemented")
		})
	}
	if api.GetParkingByDateHandler == nil {
		api.GetParkingByDateHandler = operations.GetParkingByDateHandlerFunc(func(params operations.GetParkingByDateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetParkingByDate has not yet been implemented")
		})
	}
	if api.UpdateCarHandler == nil {
		api.UpdateCarHandler = operations.UpdateCarHandlerFunc(func(params operations.UpdateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateCar has not yet been implemented")
		})
	}
	if api.UpdateParkingHandler == nil {
		api.UpdateParkingHandler = operations.UpdateParkingHandlerFunc(func(params operations.UpdateParkingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateParking has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
