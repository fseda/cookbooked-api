package services

import (
	"errors"
	"fmt"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	"github.com/fseda/cookbooked-api/internal/domain/validator"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (token string, validation validator.Validation, err error)
	Create(username, email, password string) (user *models.User, validation validator.Validation, err error)
	GithubLogin(code string) (token string, validation validator.Validation, err error)
}

type authService struct {
	authRepository repositories.AuthRepository
	userRepository repositories.UserRepository
	env            *config.Config
}

func NewAuthService(
	authRepository repositories.AuthRepository,
	userRepository repositories.UserRepository,
	env *config.Config,
) AuthService {
	return &authService{
		authRepository,
		userRepository,
		env,
	}
}

func (as *authService) Login(username, password string) (token string, validation validator.Validation, err error) {
	var user *models.User

	validation = validator.NewValidation()
	if username == "" {
		validation.AddError("username", errors.New("username is required"))
	}
	if password == "" {
		validation.AddError("password", errors.New("password is required"))
	}
	if validation.HasErrors() {
		return
	}

	user, err = as.userRepository.FindOneForLogin(username)
	if err != nil {
		return
	}
	if user == nil {
		validation.AddError("user", errors.New("invalid credentials"))
		return "", validation, nil
	}

	noMatch := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if noMatch != nil || user == nil {
		validation.AddError("user", errors.New("invalid credentials"))
		return "", validation, nil
	}

	token, err = jwtutil.GenerateToken(user, "", as.env.Http.JWTSecretKey)
	if err != nil {
		return
	}

	return
}

func (as *authService) Create(username, email, password string) (*models.User, validator.Validation, error) {
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
	err = as.userRepository.Create(&user)
	if err != nil {
		return nil, validation, err
	}

	return &user, validation, nil
}

func (as *authService) validateUserRegistration(user models.User) (validation validator.Validation) {
	validation = validator.NewValidation()

	if user.Username == "" {
		validation.AddError("username", errors.New("username is required"))
	} else {
		if len(user.Username) < 3 {
			validation.AddError("username", errors.New("username must be at least 3 characters long"))
		}

		if len(user.Username) > 255 {
			validation.AddError("username", errors.New("username must be less than 255 characters long"))
		}

		if usernameExists, _ := as.userRepository.UserExists("username", user.Username); usernameExists {
			validation.AddError("username", errors.New("username already in use"))
		}
	}

	if user.Email == "" {
		validation.AddError("email", errors.New("email is required"))
	} else {
		if !validator.IsEmailLike(user.Email) {
			validation.AddError("email", errors.New("email is invalid"))
		}

		if emailExists, _ := as.userRepository.UserExists("email", user.Email); emailExists {
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

type GithubAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type GithubAccessToken struct {
	ErrorResponse
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`      // "repo,gist"
	TokenType   string `json:"token_type"` // Bearer
}

type ErrorResponse struct {
	Message string `json:"error_description"`
	Error   string `json:"error"`
}

type GithubUser struct {
	ID       uint   `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar_url"`
	Location string `json:"location"`
}

const githubGetAccessTokenURL = "https://github.com/login/oauth/access_token"
const githubGetUserURL = "https://api.github.com/user"

func (as *authService) GithubLogin(code string) (token string, validation validator.Validation, err error) {
	githubService := NewGithubService()
	validation = validator.NewValidation()

	if code == "" {
		validation.AddError("code", errors.New("code is required"))
		return
	}

	githubUser, githubAccessToken, err := githubService.githubFlow(as.env.Github.ClientID, as.env.Github.ClientSecret, code)
	if err != nil {
		return "", validation, err
	}
	if githubAccessToken.Error != "" {
		validation.AddError("code", errors.New(githubAccessToken.Message))
		return
	}

	user, err := as.userRepository.FindOneByGithubID(githubUser.ID)
	if err != nil {
		return
	}

	if user == nil {
		newUser := models.User{
			Username: githubUser.Login,
			Email:    githubUser.Email,
			Name:     githubUser.Name,
			Bio:      githubUser.Bio,
			Avatar:   githubUser.Avatar,
			Location: githubUser.Location,
			GithubID: fmt.Sprint(githubUser.ID),
		}

		if err = as.userRepository.Create(&newUser); err != nil {
			return
		}

		user = &newUser
	}

	if err = as.authRepository.SaveGithubAccessToken(user.ID, githubAccessToken.AccessToken); err != nil {
		return
	}

	token, err = jwtutil.GenerateToken(user, githubAccessToken.AccessToken, as.env.Http.JWTSecretKey)
	if err != nil {
		return
	}

	return
}
