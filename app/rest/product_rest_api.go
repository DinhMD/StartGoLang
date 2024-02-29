package rest

import (
	"encoding/json"
	"net/http"
	"starter_go/app/common"
	"starter_go/app/configs"
	"starter_go/app/infrastructure/models"
	"starter_go/app/infrastructure/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v9"
)

func CreateProduct(c *fiber.Ctx) error {
	repository := repository.NewProductRepository(configs.DB)
	body := c.BodyRaw()
	if body == nil {
		c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		return nil
	}
	var product models.Product
	if err := json.Unmarshal(body, &product); err != nil {
		return err
	}
	product.CreatedOn = null.TimeFrom(time.Now())
	repository.Create(&product)
	c.Status(http.StatusCreated).JSON(product)
	return nil
}

func GetProducts(c *fiber.Ctx) error {
	var products, err = repository.NewProductRepository(configs.DB).FindAll()
	if err != nil {
		return err
	}
	c.JSON(products)
	return nil
}

func GetProductById(c *fiber.Ctx) error {
	var id = c.Params("id")
	var product, err = repository.NewProductRepository(configs.DB).FindById(common.ToUInt(id))
	if err != nil {
		return err
	}
	c.JSON(product)
	return nil
}
