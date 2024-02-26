package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	//Authorize
	router.POST("/app/register", Register)
	router.POST("/app/login", Login)
	//Products
	router.GET("/app/products", GetProducts)
	router.GET("/app/products/:id", GetProductById)
	router.POST("/app/products", CreateProduct)
}
