package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	logrus.WithFields(logrus.Fields{
		"status":  code,
		"method":  c.Method(),
		"path":    c.Path(),
		"message": err.Error(),
	}).Error("Failure in request")

	return c.Status(code).SendString(err.Error())
}
