package main

import (
	configs "starter_go/configs"
	"starter_go/rest"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig := configs.InitializeConfig()

	router := gin.Default()
	rest.Routes(router)
	router.Run(serverConfig.Host + ":" + strconv.Itoa(serverConfig.Port))
}
