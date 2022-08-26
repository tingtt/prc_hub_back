package oauth2

import (
	"prc_hub_back/domain/model/oauth2"
)

func Connect(userId string, provider string, code string) (err error) {
	if !initialized {
		return ErrRepositoryNotInitialized
	}

	c, err := getClient(provider)
	if err != nil {
		return
	}

	t, err := c.GetToken(code)
	if err != nil {
		return
	}

	_, err = oauth2.NewToken(repository, userId, provider, c, t)
	if err != nil {
		return
	}
	return
}
