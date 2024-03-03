package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Use("/*", AuthMiddleware())
	// app.Use("/app/*", func(c *fiber.Ctx) error {
	// 	wildcardPath := c.Params("*")
	// 	targetURL := "http://localhost:3001/app/" + wildcardPath
	// 	return forwardRequest(c, targetURL)
	// })
	// app.Use("/auth/*", func(c *fiber.Ctx) error {
	// 	wildcardPath := c.Params("*")
	// 	targetURL := "http://localhost:3002/auth/" + wildcardPath
	// 	return forwardRequest(c, targetURL)
	// })
}
