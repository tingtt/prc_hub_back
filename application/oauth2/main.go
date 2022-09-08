package oauth2

import (
	"errors"
	"prc_hub_back/domain/model/oauth2"
	"prc_hub_back/domain/model/oauth2_github"
)

// Singleton fields
var (
	initialized  = false
	repository   oauth2.OAuth2TokenRepository
	githubClient *oauth2_github.Client
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(repo oauth2.OAuth2TokenRepository, githubClientId string, githubClientSecret string) {
	githubClient = oauth2_github.NewClient(githubClientId, githubClientSecret)

	initialized = true
	repository = repo
}
