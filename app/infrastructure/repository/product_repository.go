package repository

import (
	"starter_go/app/infrastructure/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) Create(product *models.Product) *gorm.DB {
	return repo.db.Create(product)
}

func (repo *ProductRepository) Update(product *models.Product) *gorm.DB {
	return repo.db.Updates(product)
}

func (repo *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *ProductRepository) FindById(id uint) (models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, id).Error
	return product, err
}
