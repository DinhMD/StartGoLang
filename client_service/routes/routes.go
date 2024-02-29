package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Routes(app *fiber.App) {
	app.Use("/app/*", func(c *fiber.Ctx) error {
		wildcardPath := c.Params("*")
		targetURL := "http://localhost:3001/app/" + wildcardPath
		return forwardRequest(c, targetURL)
	})
	app.Use("/auth/*", func(c *fiber.Ctx) error {
		wildcardPath := c.Params("*")
		targetURL := "http://localhost:3002/auth/" + wildcardPath
		return forwardRequest(c, targetURL)
	})
}

func forwardRequest(c *fiber.Ctx, targetURL string) error {
	forwardReq := &fasthttp.Request{}
	c.Request().CopyTo(forwardReq)

	forwardReq.SetRequestURI(targetURL)

	client := &fasthttp.Client{}

	err := client.Do(forwardReq, c.Response())

	return err
}
