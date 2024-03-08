package rest

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router *fiber.App) {
	//Products
	router.Get("/app/products", GetProducts)
	router.Get("/app/products/:id", GetProductById)
	router.Put("/app/products/:id", UpdateProduct)
	router.Post("/app/products", CreateProduct)
}
