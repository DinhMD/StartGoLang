package routes

import (
	"fmt"
	"net/http"
	ArrayUtils "starter_go/client_service/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type ServiceRoute struct {
	Prefix string `json:"verify_path"`
	Port   int    `json:"port"`
}

var IgnoreAuthRoutes = []string{"/auth/login", "/auth/register"}

var MicroServiceRoutes = []ServiceRoute{
	{
		Prefix: "/app",
		Port:   3001,
	},
	{
		Prefix: "/auth",
		Port:   3902,
	}}

func AuthMiddleware() fiber.Handler {
	// Assume we have an authentication service URL for token verification
	authServiceURL := "http://localhost:3902/auth/verify"

	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header
		path := c.Path()
		if ArrayUtils.IsContains(IgnoreAuthRoutes, path) {
			wildcardPath := c.Params("*")
			targetURL := fmt.Sprintf("http://localhost:3902/%s", wildcardPath)
			return forwardRequest(c, targetURL)
		}
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
		}

		// Split the Authorization header value to extract the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}
		token := parts[1]

		// Verify the token against the authentication service
		verifyReq, err := http.NewRequest("POST", authServiceURL, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify token"})
		}
		verifyReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		client := &http.Client{}
		resp, err := client.Do(verifyReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify token"})
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// If token is valid, forward request to microservice check in list MicroServiceRoutes
		for _, route := range MicroServiceRoutes {
			if strings.HasPrefix(path, route.Prefix) {
				wildcardPath := c.Params("*")
				targetURL := fmt.Sprintf("http://localhost:%d/%s", route.Port, wildcardPath)
				return forwardRequest(c, targetURL)
			}
		}
		return c.Next()
	}
}

func forwardRequest(c *fiber.Ctx, targetURL string) error {
	forwardReq := &fasthttp.Request{}
	c.Request().CopyTo(forwardReq)

	forwardReq.SetRequestURI(targetURL)

	client := &fasthttp.Client{}

	err := client.Do(forwardReq, c.Response())
	return err
}
