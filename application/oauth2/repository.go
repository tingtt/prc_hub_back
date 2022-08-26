package oauth2

import (
	"errors"
	"prc_hub_back/domain/model/oauth2"
)

var (
	initialized = false
	repository  oauth2.OAuth2TokenRepository
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(repo oauth2.OAuth2TokenRepository) {
	initialized = true
	repository = repo
}
