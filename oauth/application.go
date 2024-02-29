package main

import (
	"oauth/configs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	app := fiber.New()
	app.Listen(":" + strconv.Itoa(configs.ENV.Port))
}
