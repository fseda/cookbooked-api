package services

import (
	"errors"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	modelvalidation "github.com/fseda/cookbooked-api/internal/domain/models/validation"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	validationPkg "github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (token string, validation validationPkg.Validation, err error)
	Create(username, email, password string) (user *models.User, validation validationPkg.Validation, err error)
}

type authService struct {
	repository repositories.UserRepository
	env        *config.Config
}

func NewAuthService(repository repositories.UserRepository, env *config.Config) AuthService {
	return &authService{repository, env}
}

func (as *authService) Login(username, password string) (token string, validation validationPkg.Validation, err error) {
	var user *models.User

	validation = validationPkg.NewValidation()
	if username == "" {
		validation.AddError("username", errors.New("username is required"))
	}
	if password == "" {
		validation.AddError("password", errors.New("password is required"))
	}
	if validation.HasErrors() {
		return
	}

	user, err = as.repository.FindOneForLogin(username)
	if err != nil {
		return
	}

	noMatch := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if noMatch != nil || user == nil {
		validation.AddError("user", errors.New("invalid credentials"))
		return "", validation, nil
	}

	token, err = jwtutil.GenerateToken(user, as.env.Http.JWTSecretKey)
	if err != nil {
		return
	}

	return
}

func (as *authService) Create(username, email, password string) (*models.User, validationPkg.Validation, error) {
	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: password,
	}

	validation := as.validateUserRegistration(user)
	if validation.HasErrors() {
		return nil, validation, nil
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return nil, validation, err
	}

	user.PasswordHash = string(passwordHash)
	err = as.repository.Create(&user)
	if err != nil {
		return nil, validation, err
	}

	return &user, validation, nil
}

func (as *authService) validateUserRegistration(user models.User) (validation validationPkg.Validation) {
	validation = validationPkg.NewValidation()

	if user.Username == "" {
		validation.AddError("username", errors.New("username is required"))
	} else {
		if len(user.Username) < 3 {
			validation.AddError("username", errors.New("username must be at least 3 characters long"))
		}
	
		if len(user.Username) > 255 {
			validation.AddError("username", errors.New("username must be less than 255 characters long"))
		}

		if usernameExists, _ := as.repository.UserExists("username", user.Username); usernameExists {
			validation.AddError("username", errors.New("username already in use"))
		}
	}

	if user.Email == "" {
		validation.AddError("email", errors.New("email is required"))
	} else {
		if !modelvalidation.IsEmailLike(user.Email) {
			validation.AddError("email", errors.New("email is invalid"))
		}

		if emailExists, _ := as.repository.UserExists("email", user.Email); emailExists {
			validation.AddError("email", errors.New("email already in use"))
		}
	}

	if user.PasswordHash == "" {
		validation.AddError("password", errors.New("password is required"))
	} else {
		if len(user.PasswordHash) < 4 {
			validation.AddError("password", errors.New("password must be at least 4 characters long"))
		}
	
		if len(user.PasswordHash) > 72 {
			validation.AddError("password", errors.New("password must be less than 72 characters long"))
		}
	}

	return validation
}
