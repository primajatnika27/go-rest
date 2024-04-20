package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() (*gorm.DB, error) {
	var (
		host     = viper.GetString("DATABASE_HOST")
		port     = viper.GetInt("DATABASE_PORT")
		user     = viper.GetString("DATABASE_USERNAME")
		password = viper.GetString("DATABASE_PASSWORD")
		db       = viper.GetString("DATABASE_NAME")
	)

	var err error
	dbConnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
	logrus.Info(dbConnect)

	gormDB, err := gorm.Open(postgres.Open(dbConnect), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Gorm open connection failed")
		return nil, nil
	}

	return gormDB, nil
}
