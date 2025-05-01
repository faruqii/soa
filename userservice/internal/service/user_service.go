package service

import (
	"github.com/faruqii/soa/userservice/internal/models"
	"github.com/faruqii/soa/userservice/internal/repository"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	HashPassword(password string) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	if err := s.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	log.Infof("Fetching user with ID: %s", id)

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) HashPassword(password string) (string, error) {
	// Implement password hashing logic here
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
