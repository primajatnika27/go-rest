package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"status":  c.Response().StatusCode(),
				"method":  c.Method(),
				"path":    c.Path(),
				"message": err.Error(),
			}).Error("Failure in request")
			return err
		}

		if c.Response().StatusCode() == 404 {
			logrus.WithFields(logrus.Fields{
				"status": c.Response().StatusCode(),
				"method": c.Method(),
				"path":   c.Path(),
			}).Info("Request not found")
		}

		if c.Response().StatusCode() == 500 {
			logrus.WithFields(logrus.Fields{
				"status": c.Response().StatusCode(),
				"method": c.Method(),
				"path":   c.Path(),
			}).Error("Failure in request")
		}

		if c.Response().StatusCode() >= 200 && c.Response().StatusCode() <= 299 {
			logrus.WithFields(logrus.Fields{
				"status": c.Response().StatusCode(),
				"method": c.Method(),
				"path":   c.Path(),
			}).Info("Request success")
		}
		return nil
	}
}
