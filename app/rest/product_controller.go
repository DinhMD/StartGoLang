package rest

import (
	"encoding/json"
	"net/http"
	"starter_go/app/common"
	"starter_go/app/infrastructure/services"
	rest_models "starter_go/app/rest/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	body := c.BodyRaw()
	if body == nil {
		c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		return nil
	}
	var product rest_models.ProductRequest
	if err := json.Unmarshal(body, &product); err != nil {
		return err
	}
	id, err := services.Create(&product)
	if err != nil || id == 0 {
		c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		return err
	}
	var response, _ = services.GetProductById(id)
	c.Status(http.StatusCreated).JSON(response)
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
	body := c.BodyRaw()
	reqId := c.Params("id")
	if body == nil || reqId == "" {
		c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		return nil
	}
	var product rest_models.ProductRequest
	if err := json.Unmarshal(body, &product); err != nil {
		return err
	}
	id, err := services.Update(common.ToInt(reqId), &product)
	if err != nil || id == 0 {
		c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		return err
	}
	var response, _ = services.GetProductById(id)
	c.Status(http.StatusCreated).JSON(response)
	return nil
}

func GetProducts(c *fiber.Ctx) error {
	var products, err = services.GetProducts()
	if err != nil {
		return err
	}
	c.JSON(products)
	return nil
}

func GetProductById(c *fiber.Ctx) error {
	var id = c.Params("id")
	var product, err = services.GetProductById(common.ToUInt(id))
	if err != nil {
		return err
	}
	c.JSON(product)
	return nil
}
