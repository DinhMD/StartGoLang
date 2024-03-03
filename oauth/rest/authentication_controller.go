package rest

import (
	"encoding/json"
	"net/http"
	"starter_go/oauth/configs"
	"starter_go/oauth/infrastructure/services"
	dto "starter_go/oauth/rest/models"

	"github.com/gofiber/fiber/v2"
)

type OAuthServer struct {
	ClientStorage map[string]string
}

func Register(f *fiber.Ctx) error {
	var body = f.BodyRaw()
	if body == nil {
		f.Status(http.StatusBadRequest)
		return nil
	}
	var accountRequest dto.AccountRequest
	err := json.Unmarshal(body, &accountRequest)
	if err != nil {
		return err
	}
	saveError := services.CreateAnAccount(accountRequest, f)
	if saveError != nil {
		f.Status(http.StatusInternalServerError).JSON(saveError)
		return nil
	} else {
		f.SendString("OK")
		return nil
	}
}

func Login(f *fiber.Ctx) error {
	var body = f.BodyRaw()
	if body == nil {
		f.Status(fiber.StatusBadRequest)
		return nil
	}
	var accountRequest dto.AccountRequest
	if err := json.Unmarshal(body, &accountRequest); err != nil {
		f.Status(fiber.StatusBadRequest)
		return err
	}
	account := services.CheckAccount(accountRequest, f)
	if account == nil {
		f.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		return nil
	}
	tokenString, err := configs.CreateToken(account.Username)
	if err != nil {
		return err
	}
	f.JSON(configs.AuthResponse{
		Token: tokenString,
	})
	return nil
}
