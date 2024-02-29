package repository

import (
	"starter_go/app/infrastructure/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repo *AccountRepository) Create(account *models.Account) error {
	return repo.db.Create(account).Error
}

func (repo *AccountRepository) FindByUsername(username string) (models.Account, error) {
	var account models.Account
	err := repo.db.Where("username = ?", username).First(&account).Error
	return account, err
}
