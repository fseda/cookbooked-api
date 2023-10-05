package services

import (
	"errors"
	"fmt"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository}
}

func (us *UserService) FindByID(id uint) (*models.User, error) {
	user, err := us.repository.FindOneById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User with ID %d not found", id)
		}

		return nil, err
	}
	return user, nil
}

func (us *UserService) Create(username, email, password string) (uint, error) {

	usernameExists, _ := us.repository.UserExists("username", username)
	if usernameExists {
		return 0, fmt.Errorf("username '%v' already exists", username)
	}

	emailExists, _ := us.repository.UserExists("email", email)
	if emailExists {
		return 0, fmt.Errorf("email '%v' already exists", email)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	id, err := us.repository.Create(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
