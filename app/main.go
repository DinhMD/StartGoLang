package main

import (
	"starter_go/app/configs"
	"starter_go/app/rest"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitializeConfig()
	router := fiber.New()
	rest.Routes(router)
	router.Listen(configs.ENV.Host + ":" + strconv.Itoa(configs.ENV.Port))
}
