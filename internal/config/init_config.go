package config

import (
	"github.com/spf13/viper"
	"test_svc/internal/model"
)

func ReadConfigFromEnv() *model.Config {
	viper.AutomaticEnv()
	var config model.Config

	// Server
	config.Server.Host = viper.GetString("APP.SERVER.HOST")
	config.Server.Host = viper.GetString("APP.SERVER.PORT")

	// Database
	config.Database.Host = viper.GetString("APP.DB.HOST")
	config.Database.Port = viper.GetString("APP.DB.PORT")
	config.Database.DBName = viper.GetString("APP.DB.DB_NAME")
	config.Database.User = viper.GetString("APP.DB.USER")
	config.Database.Password = viper.GetString("APP.DB.PASSWORD")
	config.Database.SSLMode = viper.GetString("APP.DB.SSL_MODE")
	config.Database.TimeZone = viper.GetString("APP.DB.TIME_ZONE")
	config.Database.MaxCons = viper.GetInt("APP.DB.MAX_CONS")
	config.Database.MinCons = viper.GetInt("APP.DB.MIN_CONS")
	config.Database.MaxConnIdleTime = viper.GetInt("APP.DB.MAX_CONN_IDLE_TIME")

	return &config
}
