package services

import (
	"net/http"
	"starter_go/app/common"
	"starter_go/oauth/configs"
	"starter_go/oauth/infrastructure/models"
	"starter_go/oauth/infrastructure/repository"
	dto "starter_go/oauth/rest/models"

	"github.com/gofiber/fiber/v2"
)

func CreateAnAccount(accountRequest dto.AccountRequest, f *fiber.Ctx) error {
	err := validateAccount(accountRequest, f)
	if err != nil {
		return err
	}
	passwordHashed, err := models.HashPassword(accountRequest.Password)
	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	account := models.Account{
		Username: accountRequest.Username,
		Password: passwordHashed,
	}
	repository := repository.NewAccountRepository(configs.DB)
	existedAccount, err := repository.FindByUsername(accountRequest.Username)
	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})

	}
	if existedAccount.Username != "" {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"field": "username", "error": "Username is existed"})
	}
	saveErr := repository.Create(&account)
	if saveErr != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return nil
}

func CheckAccount(accountRequest dto.AccountRequest, f *fiber.Ctx) *models.Account {

	repository := repository.NewAccountRepository(configs.DB)

	existedAccount, err := repository.FindByUsername(accountRequest.Username)
	if err != nil {
		return nil
	}
	if models.CheckPasswordHash(accountRequest.Password, existedAccount.Password) {
		return &existedAccount
	} else {
		return nil
	}
}

func validateAccount(account dto.AccountRequest, f *fiber.Ctx) error {
	if account.Username == "" || account.Password == "" {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username and password is required"})
	}
	if len(account.Username) > 20 {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Field: common.StringPtr("username"),
			Value: "Username is too long. Max 20 characters",
		})
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"field": "username", "error": "Username is too long. Max 20 characters"})
	}
	if len(account.Username) < 6 {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"field": "username", "error": "Username is too short. Min 6 characters"})
	}
	if len(account.Password) > 20 {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"field": "password", "error": "Password is too long. Max 20 characters"})
	}
	if len(account.Password) < 8 {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{"field": "password", "error": "Password is too short. Min 8 characters"})
	}
	return nil
}
