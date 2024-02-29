package main

import (
	"fmt"
	"starter_go/oauth/configs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, World!")
		c.SendString("Hello, World!")
		return nil
	})
	app.Listen(":" + strconv.Itoa(configs.ENV.Port))
}
