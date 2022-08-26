package user

import (
	"errors"
	"prc_hub_back/domain/model/user"
)

// Singleton fields
var (
	initialized = false
	repository  user.UserRepository
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(repo user.UserRepository) {
	initialized = true
	repository = repo
}
