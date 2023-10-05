package services

import (
	"errors"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository       *repositories.UserRepository
	CreateUserErrors *_createUserErrors
}

type _createUserErrors struct {
	EmailExists    error
	UsernameExists error
}

var createUserErrors = &_createUserErrors{
	EmailExists:    errors.New("email already exists"),
	UsernameExists: errors.New("username already exists"),
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository,
		createUserErrors,
	}
}

func (us *UserService) FindByID(id uint) (*models.User, error) {
	user, err := us.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Create(username, email, password string) (*models.User, error) {

	usernameExists, _ := us.repository.UserExists("username", username)
	if usernameExists {
		return nil, createUserErrors.UsernameExists
	}

	emailExists, _ := us.repository.UserExists("email", email)
	if emailExists {
		return nil, createUserErrors.EmailExists
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

func (us *UserService) Delete(id uint) (int64, error) {
	return us.repository.Delete(id)
}
