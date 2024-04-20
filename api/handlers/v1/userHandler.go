package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"primajatnika/api-server/api/dto"
	"primajatnika/api-server/api/presenter"
	"primajatnika/api-server/domain/services"
)

type UserHandler interface {
	Create() fiber.Handler
}

type userHandler struct {
	services services.UserService
}

func (u userHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users dto.UserRequestModel

		if err := c.BodyParser(&users); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Bad request")
			return c.Status(fiber.StatusBadRequest).
				JSON(presenter.FailedRes("Bad Request", err))
		}

		resp, err := u.services.Create(&users)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Internal Server Error")
			return c.Status(fiber.StatusInternalServerError).
				JSON(presenter.FailedRes("Internal Server Error", err))
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.SuccessRes(resp))
	}
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{
		services: service,
	}
}
