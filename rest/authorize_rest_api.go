package rest

import (
	"encoding/json"
	"net/http"
	"starter_go/infrastructure/services"
	dto "starter_go/rest/request_models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body, err = c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	var accountRequest dto.AccountRequest
	if err = json.Unmarshal(body, &accountRequest); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	services.CreateAnAccount(accountRequest, c)
}

func Login(c *gin.Context) {
	var body, err = c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	var accountRequest dto.AccountRequest
	if err = json.Unmarshal(body, &accountRequest); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	services.ValidAccount(accountRequest, c)
}
