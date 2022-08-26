package oauth2

import "prc_hub_back/domain/model/oauth2"

func Disconnect(userId string, provider string) error {
	if !initialized {
		return ErrRepositoryNotInitialized
	}

	return oauth2.Delete(repository, userId, provider)
}
