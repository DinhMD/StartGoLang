package rest

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router *fiber.App) {
	//Authorize
	router.Post("/app/register", Register)
	router.Post("/app/login", Login)
	//Products
	router.Get("/app/products", GetProducts)
	router.Get("/app/products/:id", GetProductById)
	router.Post("/app/products", CreateProduct)
}
