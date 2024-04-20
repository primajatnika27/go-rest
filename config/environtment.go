package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupEnvironment() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.SetDefault("LOG_FILE_LOCATION", "./logs/application.log")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error reading config file", err)
	}
}
