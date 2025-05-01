package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/faruqii/soa/authservice/internal/helper"
	"github.com/faruqii/soa/authservice/internal/models"
	"github.com/faruqii/soa/authservice/internal/repository"
	log "github.com/sirupsen/logrus"
)

type KeyService interface {
	CreateKey(userID string) (*models.APIKey, error)
	ValidateKey(key string) (*models.APIKey, error)
}

type keyService struct {
	repo           repository.KeyRepository
	userServiceURL string
}

func NewKeyService(repo repository.KeyRepository, userServiceURL string) KeyService {
	return &keyService{
		repo:           repo,
		userServiceURL: userServiceURL,
	}
}

func (s *keyService) CreateKey(userID string) (*models.APIKey, error) {
	if !s.isValidUserID(userID) {
		log.Warnf("Invalid user ID: %s", userID)
		return nil, fmt.Errorf("invalid user ID: %s", userID)
	}

	key := helper.GenerateSecureKey()
	apikey := &models.APIKey{
		Key:       key,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateKey(apikey); err != nil {
		log.Errorf("Failed to save API key: %v", err)
		return nil, err
	}

	log.Infof("API key %s created for user %s", apikey.Key, apikey.UserID)
	return apikey, nil
}

func (s *keyService) ValidateKey(key string) (*models.APIKey, error) {
	log.Infof("Validating API key: %s", key)
	// Validate the API key
	apikey, err := s.repo.ValidateKey(key)
	if err != nil {
		return nil, err
	}
	return apikey, nil
}

func (s *keyService) isValidUserID(userID string) bool {
	if userID == "" {
		log.Warn("User ID is empty")
		return false
	}

	log.Println(userID)

	url := fmt.Sprintf("%s/%s", s.userServiceURL, userID)
	log.Printf("Making request to user service: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Error making request to user service: %v", err)
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}
