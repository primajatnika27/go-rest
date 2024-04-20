package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"primajatnika/api-server/api/routers"
	"primajatnika/api-server/config"
	"primajatnika/api-server/pkg/logger"
	"primajatnika/api-server/pkg/middleware"
)

func main() {
	app := fiber.New()
	logger.SetupLogger()
	config.SetupEnvironment()

	gormDB, err := config.SetupDatabase()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Gorm open connection failed")
		return
	}

	app.Use(middleware.CorsConfig())
	app.Use(middleware.RequestLogger())

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return middleware.ErrorHandler(c, err)
		}
		return nil
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			logrus.WithFields(logrus.Fields{
				"error": e,
			}).Error("Panic occurred")
		},
	}))

	routers.SetupUserRouter(app, gormDB)

	appPort := viper.GetString("APP_PORT")
	logrus.Fatal(app.Listen(":" + appPort))
}
