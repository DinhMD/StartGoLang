package main

import (
	"starter_go/oauth/configs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	app := fiber.New()
	app.Listen(":" + strconv.Itoa(configs.ENV.Port))
}
