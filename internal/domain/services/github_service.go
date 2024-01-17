package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type GithubService struct{}

var once sync.Once

var instance *GithubService

func NewGithubService() *GithubService {
	once.Do(func() {
		instance = &GithubService{}
	})
	return instance
}

func (gh *GithubService) githubFlow(clientID, clientSecret, code string) (githubUser *GithubUser, githubAccessToken *GithubAccessToken, err error) {
	githubAccessTokenRequest, _ := json.Marshal(GithubAccessTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	})
	httpClient := &http.Client{}

	req, _ := http.NewRequest("POST", githubGetAccessTokenURL, bytes.NewBuffer(githubAccessTokenRequest))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	tokenResp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer tokenResp.Body.Close()
	accessTokenBody, err := io.ReadAll(tokenResp.Body)
	json.Unmarshal(accessTokenBody, &githubAccessToken)
	if githubAccessToken.Error != "" { // Github returns 200 even if the code is invalid
		return nil, nil, nil
	}

	githubUser, err = gh.getUserFromGithub(githubAccessToken)
	return githubUser, githubAccessToken, err
}

func (gh *GithubService) getUserFromGithub(githubAccessToken *GithubAccessToken) (githubUser *GithubUser, err error) {
	token := fmt.Sprintf("%s %s", githubAccessToken.TokenType, githubAccessToken.AccessToken)

	req, _ := http.NewRequest("GET", githubGetUserURL, nil)
	req.Header.Set("Authorization", token)
	httpClient := &http.Client{}
	userResp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if userResp.StatusCode != 200 {
		err = errors.New("unable to get user from github")
		return nil, err
	}

	defer userResp.Body.Close()
	userBody, err := io.ReadAll(userResp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(userBody, &githubUser)
	return githubUser, nil
}
