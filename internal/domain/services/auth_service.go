package services

import (
	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Create(username, email, password string) (*models.User, error)
}

type authService struct {
	repository repositories.UserRepository
	env *config.Config
}

func NewAuthService(repository repositories.UserRepository, env *config.Config) AuthService {
	return &authService{repository, env}
}

func (as *authService) Login(username, password string) (string, error) {
	var user *models.User
	var err error

	user, err = as.repository.FindOneForLogin(username)
	if err != nil {
		log.Errorf("Error fetching user: %v", err)
		return "", globalerrors.GlobalInternalServerError
	}

	if user == nil {
		return "", globalerrors.AuthInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", globalerrors.AuthInvalidCredentials
	}

	token, err := jwtutil.GenerateToken(user, as.env.Http.JWTSecretKey)
	if err != nil {
		log.Errorf("Error generating token: %v", err)
		return "", globalerrors.GlobalInternalServerError
	}

	return token, nil
}


func (as *authService) Create(username, email, password string) (*models.User, error) {

	usernameExists, _ := as.repository.UserExists("username", username)
	if usernameExists {
		return nil, globalerrors.UserUsernameExists
	}

	emailExists, _ := as.repository.UserExists("email", email)
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
	err = as.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}