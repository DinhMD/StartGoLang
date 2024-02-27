package main

import (
	"starter_go/configs"
	"starter_go/rest"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	router := fiber.New()

	router.Use(func(f *fiber.Ctx) error {
		if f.Path() != "/app/login" {
			return configs.Authenticate(f)
		}
		return f.Next()
	})

	rest.Routes(router)

	router.Listen(configs.ENV.Host + ":" + strconv.Itoa(configs.ENV.Port))
}
