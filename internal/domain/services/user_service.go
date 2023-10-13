package services

import (
	"github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(id uint) (*models.User, error)
	Create(username, email, password string) (*models.User, error)
	Delete(id uint) (int64, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (us *userService) FindByID(id uint) (*models.User, error) {
	user, err := us.repository.FindOneById(id)
	if err != nil {
		return nil, globalerrors.GlobalInternalServerError
	}
	if user == nil {
		return nil, globalerrors.UserNotFound
	}
	return user, nil
}

func (us *userService) Create(username, email, password string) (*models.User, error) {

	usernameExists, _ := us.repository.UserExists("username", username)
	if usernameExists {
		return nil, globalerrors.UserUsernameExists
	}

	emailExists, _ := us.repository.UserExists("email", email)
	if emailExists {
		return nil, globalerrors.UserEmailExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	err = us.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) Delete(id uint) (int64, error) {
	return us.repository.Delete(id)
}
