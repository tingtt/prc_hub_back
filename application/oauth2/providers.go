package oauth2

import (
	"errors"
)

type Provider string

var (
	ProviderGithub Provider = "github"
)

func searchProvider(s string) (Provider, error) {
	switch s {
	case string(ProviderGithub):
		return ProviderGithub, nil
	default:
		return "", ErrProviderNotFound
	}
}

// Errors
var (
	ErrProviderNotFound = errors.New("provider not found")
)
