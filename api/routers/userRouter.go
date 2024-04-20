package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"primajatnika/api-server/adapters/repository"
	v1 "primajatnika/api-server/api/handlers/v1"
	"primajatnika/api-server/domain/services"
)

func SetupUserRouter(app *fiber.App, db *gorm.DB) {
	r := repository.NewUsersRepository(db)
	s := services.NewUserService(r)
	h := v1.NewUserHandler(s)

	apiV1 := app.Group("api-go/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})

	apiV1User := apiV1.Group("/user")
	apiV1User.Post("/save", h.Create())
}
