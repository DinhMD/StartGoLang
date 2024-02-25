package rest

import (
	"encoding/json"
	"net/http"
	"starter_go/common"
	"starter_go/configs"
	"starter_go/infrastructure/models"
	"starter_go/infrastructure/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v9"
)

func CreateProduct(c *gin.Context) {
	repository := repository.NewProductRepository(configs.DB)
	body, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	var product models.Product
	if err = json.Unmarshal(body, &product); err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	product.CreatedOn = null.TimeFrom(time.Now())
	repository.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	var products, err = repository.NewProductRepository(configs.DB).FindAll()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	var id = c.Param("id")
	var product, err = repository.NewProductRepository(configs.DB).FindById(common.ToUInt(id))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, product)
}
