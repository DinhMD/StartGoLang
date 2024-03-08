package services

import (
	"starter_go/app/configs"
	"starter_go/app/infrastructure/models"
	"starter_go/app/infrastructure/repository"
	rest_models "starter_go/app/rest/models"
)

func Create(req *rest_models.ProductRequest) (uint, error) {
	repo := repository.NewProductRepository(configs.DB)
	product := requestToDomain(req)
	result := repo.Create(product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}
func Update(id int, req *rest_models.ProductRequest) (uint, error) {
	repo := repository.NewProductRepository(configs.DB)
	product, err := repo.FindById(uint(id))
	if err != nil {
		return 0, err
	}
	newProduct := updateToDomain(product, req)
	newProduct.ID = product.ID
	result := repo.Update(&newProduct)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

func GetProducts() ([]models.Product, error) {
	repo := repository.NewProductRepository(configs.DB)
	return repo.FindAll()
}

func GetProductById(id uint) (models.Product, error) {
	repo := repository.NewProductRepository(configs.DB)
	return repo.FindById(id)
}

func requestToDomain(product *rest_models.ProductRequest) *models.Product {
	return &models.Product{
		Name:        product.Name,
		Sku:         product.Sku,
		Price:       product.Price,
		Description: product.Description,
	}
}

func updateToDomain(product models.Product, req *rest_models.ProductRequest) models.Product {
	product.Name = req.Name
	product.Sku = req.Sku
	product.Price = req.Price
	product.Description = req.Description
	return product
}
