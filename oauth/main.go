package main

import (
	"starter_go/oauth/configs"
	"starter_go/oauth/rest"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	app := fiber.New()
	rest.Routes(app)
	app.Listen(":" + strconv.Itoa(configs.ENV.Port))
}
