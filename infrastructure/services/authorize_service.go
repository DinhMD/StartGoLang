package services

import (
	"net/http"
	"starter_go/common"
	"starter_go/configs"
	"starter_go/infrastructure/models"
	"starter_go/infrastructure/repository"
	dto "starter_go/rest/request_models"

	"github.com/gofiber/fiber/v2"
)

func CreateAnAccount(accountRequest dto.AccountRequest, f *fiber.Ctx) *common.FormError {
	validateAccount(accountRequest, f)
	passwordHashed, err := models.HashPassword(accountRequest.Password)
	if err != nil {
		return &common.FormError{
			Value: "Internal Server Error",
		}
	}
	account := models.Account{
		Username: accountRequest.Username,
		Password: passwordHashed,
	}
	repository := repository.NewAccountRepository(configs.DB)
	existedAccount, err := repository.FindByUsername(accountRequest.Username)
	if err != nil {
		return &common.FormError{
			Value: "Internal Server Error",
		}

	}
	if existedAccount.Username != "" {
		return &common.FormError{
			Field: common.StringPtr("username"),
			Value: "Username is existed",
		}
	}
	saveErr := repository.Create(&account)
	if saveErr != nil {
		return &common.FormError{
			Value: "Internal Server Error",
		}
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

func validateAccount(account dto.AccountRequest, f *fiber.Ctx) bool {
	if account.Username == "" || account.Password == "" {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Value: "Username and password is required",
		})
		return false
	}
	if len(account.Username) > 20 {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Field: common.StringPtr("username"),
			Value: "Username is too long. Max 20 characters",
		})
		return false
	}
	if len(account.Username) < 6 {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Field: common.StringPtr("username"),
			Value: "Username is too short. Min 6 characters",
		})
		return false
	}
	if len(account.Password) > 20 {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Field: common.StringPtr("password"),
			Value: "Password is too long. Max 20 characters",
		})
		return false
	}
	if len(account.Password) < 8 {
		f.Status(http.StatusBadRequest).JSON(common.FormError{
			Field: common.StringPtr("password"),
			Value: "Password is too short. Min 8 characters",
		})
		return false
	}
	return true
}
