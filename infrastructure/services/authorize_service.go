package services

import (
	"net/http"
	"starter_go/common"
	"starter_go/configs"
	"starter_go/infrastructure/models"
	"starter_go/infrastructure/repository"
	dto "starter_go/rest/request_models"

	"github.com/gin-gonic/gin"
)

func CreateAnAccount(accountRequest dto.AccountRequest, c *gin.Context) error {
	validateAccount(accountRequest, c)
	passwordHashed, err := models.HashPassword(accountRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	account := models.Account{
		Username: accountRequest.Username,
		Password: passwordHashed,
	}
	repository := repository.NewAccountRepository(configs.DB)
	existedAccount, err := repository.FindByUsername(accountRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	if existedAccount.Username != "" {
		common.HandleFormError(common.StringPtr("username"), "Username is existed", c)
	}
	saveErr := repository.Create(&account)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr)
	}
	return nil
}

func ValidAccount(accountRequest dto.AccountRequest, c *gin.Context) error {

	repository := repository.NewAccountRepository(configs.DB)

	existedAccount, err := repository.FindByUsername(accountRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	if models.CheckPasswordHash(accountRequest.Password, existedAccount.Password) {
		c.String(http.StatusOK, "OK")
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
	}
	return nil
}

func validateAccount(account dto.AccountRequest, c *gin.Context) bool {
	if account.Username == "" || account.Password == "" {
		c.JSON(http.StatusBadRequest, "Username and password is required")
		common.HandleFormError(nil, "Username and password is required", c)
		return false
	}
	if len(account.Username) > 20 {
		common.HandleFormError(common.StringPtr("username"), "Username is too long. Max 20 characters", c)
		return false
	}
	if len(account.Username) < 6 {
		common.HandleFormError(common.StringPtr("username"), "Username is too short. Min 6 characters", c)
		return false
	}
	if len(account.Password) > 20 {
		common.HandleFormError(common.StringPtr("password"), "Password is too long. Max 20 characters", c)
		return false
	}
	if len(account.Password) < 8 {
		common.HandleFormError(common.StringPtr("password"), "Password is too short. Min 8 characters", c)
		return false
	}
	return true
}
