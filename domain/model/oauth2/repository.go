package oauth2

import "errors"

// Errors
var (
	ErrOAuth2TokenNotFound = errors.New("oauth2 token not found")
)

type OAuth2TokenRepository interface {
	Save(t OAuth2Token) (OAuth2Token, error)
	Get(userId string, provider string) (OAuth2Token, error)
	GetByOwnerId(provider string, ownerId string) (OAuth2Token, error)
	Delete(userId string, provider string) error
}
