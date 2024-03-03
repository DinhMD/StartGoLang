package rest

import (
	"starter_go/oauth/configs"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)
	app.Post("/auth/verify", configs.Authenticate)
}
