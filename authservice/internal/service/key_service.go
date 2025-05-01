package service

import (
	"github.com/faruqii/soa/authservice/internal/models"
	"github.com/faruqii/soa/authservice/internal/repository"
)

type KeyService interface {
	CreateKey(apikey *models.APIKey) error
	ValidateKey(key string) (*models.APIKey, error)
}

type keyService struct {
	repo repository.KeyRepository
}

func NewKeyService(repo repository.KeyRepository) KeyService {
	return &keyService{repo: repo}
}

func (s *keyService) CreateKey(apikey *models.APIKey) error {
	if err := s.repo.CreateKey(apikey); err != nil {
		return err
	}
	return nil
}

func (s *keyService) ValidateKey(key string) (*models.APIKey, error) {
	apikey, err := s.repo.ValidateKey(key)
	if err != nil {
		return nil, err
	}
	return apikey, nil
}
