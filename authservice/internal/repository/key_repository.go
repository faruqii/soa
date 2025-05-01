package repository

import (
	"github.com/faruqii/soa/authservice/internal/models"
	"gorm.io/gorm"
)

type KeyRepository interface {
	CreateKey(apikey *models.APIKey) error
	ValidateKey(key string) (*models.APIKey, error)
}

type keyRepository struct {
	db *gorm.DB
}

func NewKeyRepository(db *gorm.DB) KeyRepository {
	return &keyRepository{db: db}
}

func (r *keyRepository) CreateKey(apikey *models.APIKey) error {
	if err := r.db.Create(&apikey).Error; err != nil {
		return err
	}
	return nil

}

func (r *keyRepository) ValidateKey(key string) (*models.APIKey, error) {
	var apikey models.APIKey
	if err := r.db.Where("key = ?", key).First(&apikey).Error; err != nil {
		return nil, err
	}
	return &apikey, nil
}
