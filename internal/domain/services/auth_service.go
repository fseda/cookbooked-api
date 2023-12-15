package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
	validationPkg "github.com/fseda/cookbooked-api/internal/domain/validation"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (token string, validation validationPkg.Validation, err error)
	Create(username, email, password string) (user *models.User, validation validationPkg.Validation, err error)
	GithubLogin(code string) (token string, err error)
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

	user, err = as.userRepository.FindOneForLogin(username)
	if err != nil {
		return
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
	err = as.userRepository.Create(&user)
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

		if usernameExists, _ := as.userRepository.UserExists("username", user.Username); usernameExists {
			validation.AddError("username", errors.New("username already in use"))
		}
	}

	if user.Email == "" {
		validation.AddError("email", errors.New("email is required"))
	} else {
		if !validationPkg.IsEmailLike(user.Email) {
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

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`      // "repo,gist"
	TokenType   string `json:"token_type"` // Bearer
}

type GithubUser struct {
	ID    uint `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

const githubGetAccessTokenURL = "https://github.com/login/oauth/access_token"
const githubGetUserURL = "https://api.github.com/user"

func (as *authService) GithubLogin(code string) (token string, err error) {
	accessTokenRequest, _ := json.Marshal(GithubAccessTokenRequest{
		ClientID:     as.env.Github.ClientID,
		ClientSecret: as.env.Github.ClientSecret,
		Code:         code,
	})
	client := &http.Client{}

	req, _ := http.NewRequest("POST", githubGetAccessTokenURL, bytes.NewBuffer(accessTokenRequest))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	tokenResp, err := client.Do(req)
	if err != nil {
		return
	}

	defer tokenResp.Body.Close()
	accessTokenBody, err := io.ReadAll(tokenResp.Body)
	var accessTokenResponse GithubAccessTokenResponse
	json.Unmarshal(accessTokenBody, &accessTokenResponse)
	tokenType := accessTokenResponse.TokenType
	accessToken := accessTokenResponse.AccessToken

	req, _ = http.NewRequest("GET", githubGetUserURL, nil)
	req.Header.Set("Authorization", tokenType+" "+accessToken)
	userResp, err := client.Do(req)
	if err != nil {
		return
	}

	defer userResp.Body.Close()
	userBody, err := io.ReadAll(userResp.Body)
	if err != nil {
		return
	}
	var githubUser GithubUser
	json.Unmarshal(userBody, &githubUser)

	user, err := as.userRepository.FindOneByGithubID(githubUser.ID)
	if err != nil {
		return
	}

	if user == nil {
		newUser := models.User{
			Username: githubUser.Login,
			Email:    githubUser.Email,
			GithubID: fmt.Sprint(githubUser.ID),
		}

		if err = as.userRepository.Create(&newUser); err != nil {
			return
		}

		user = &newUser
	}

	if err = as.authRepository.SaveGithubAccessToken(user.ID, accessTokenResponse.AccessToken); err != nil {
		return
	}

	token, err = jwtutil.GenerateToken(user, accessToken, as.env.Http.JWTSecretKey)
	if err != nil {
		return
	}

	return
}
