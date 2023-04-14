package config

import "github.com/spf13/viper"

const (
	DbName     = "db.name"
	DbHost     = "db.host"
	DbPort     = "db.port"
	DbUser     = "db.user"
	DbPass     = "db.pass"
	ServerHost = "server.host"
	ServerPort = "server.port"
)

func init() {
	// env var for db
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

	// env var for server
	err = viper.BindEnv(ServerHost, "SERVER_HOST")
	if err != nil {
		return
	}
	err = viper.BindEnv(ServerPort, "SERVER_PORT")
	if err != nil {
		return
	}

	// Defaults
	viper.SetDefault(DbName, "car_parking")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "")
	viper.SetDefault(DbPass, "")
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")

	// Read config file if it exists
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("config") // look for config file in the config directory
	err = viper.ReadInConfig()    // read config file
	if err != nil {
		// handle error
	}

}
