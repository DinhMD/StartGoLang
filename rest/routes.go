package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/app/products", GetProducts)
	router.GET("/app/products/:id", GetProductById)
	router.POST("/app/products", CreateProduct)
}
