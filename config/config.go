package config

import "github.com/spf13/viper"

// Configuration keys
const (
	DbName     = "db.name"     // Key for the name of the database
	DbHost     = "db.host"     // Key for the hostname of the database server
	DbPort     = "db.port"     // Key for the port number of the database server
	DbUser     = "db.user"     // Key for the username used to authenticate to the database server
	DbPass     = "db.pass"     // Key for the password used to authenticate to the database server
	ServerHost = "server.host" // Key for the hostname of the server on which the API is running
	ServerPort = "server.port" // Key for the port number on which the API should listen for incoming connections
)

func init() {
	// Bind environment variables for database configuration
	err := viper.BindEnv(DbName, "MONGODB_NAME")
	if err != nil {
		return
	}
	err = viper.BindEnv(DbHost, "MONGODB_HOST")
	if err != nil {
		return
	}
	err = viper.BindEnv(DbPort, "MONGODB_PORT")
	if err != nil {
		return
	}
	err = viper.BindEnv(DbUser, "MONGODB_USER")
	if err != nil {
		return
	}
	err = viper.BindEnv(DbPass, "MONGODB_PASS")
	if err != nil {
		return
	}

	// Bind environment variables for server configuration
	err = viper.BindEnv(ServerHost, "SERVER_HOST")
	if err != nil {
		return
	}
	err = viper.BindEnv(ServerPort, "SERVER_PORT")
	if err != nil {
		return
	}

	// Set default values for configuration keys
	viper.SetDefault(DbName, "car_parking")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "")
	viper.SetDefault(DbPass, "")
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")

	// Read configuration file if it exists
	viper.SetConfigName("config") // name of the configuration file (without extension)
	viper.AddConfigPath("config") // look for the configuration file in the "config" directory
	err = viper.ReadInConfig()    // read the configuration file
	if err != nil {
		// handle error
	}
}
