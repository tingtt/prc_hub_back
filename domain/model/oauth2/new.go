package oauth2

import "errors"

var (
	ErrUserTokenForProviderExists  = errors.New("an oauth2-token already exists for user/provider combination")
	ErrOwnerTokenForProviderExists = errors.New("an oauth2-token already exists for provider/owner combination")
)

func NewToken(repo OAuth2TokenRepository, userId string, provider string, c Client, token string) (_ OAuth2Token, err error) {
	// Check existence of token for user/provider combination
	_, err = repo.Get(userId, provider)
	if err == nil || err != ErrOAuth2TokenNotFound {
		err = ErrUserTokenForProviderExists
		return
	}

	// Check existence of token for provider/owner_id combination
	ownerId, err := c.GetOwnerId(token)
	if err != nil {
		return
	}
	_, err = repo.GetByOwnerId(provider, ownerId)
	if err == nil || err != ErrOAuth2TokenNotFound {
		err = ErrOwnerTokenForProviderExists
		return
	}

	ot := OAuth2Token{
		UserId:   userId,
		Provider: provider,
		Token:    token,
		OwnerId:  ownerId,
	}

	return repo.Save(ot)
}
