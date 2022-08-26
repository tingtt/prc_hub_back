package oauth2_inmemory

import "prc_hub_back/domain/model/oauth2"

func (r Repository) Get(userId string, provider string) (oauth2.OAuth2Token, error) {
	for _, ot := range data {
		if ot.UserId == userId && ot.Provider == provider {
			return ot, nil
		}
	}
	return oauth2.OAuth2Token{}, oauth2.ErrOAuth2TokenNotFound
}

func (r Repository) GetByOwnerId(provider string, ownerId string) (oauth2.OAuth2Token, error) {
	for _, ot := range data {
		if ot.Provider == provider && ot.OwnerId == ownerId {
			return ot, nil
		}
	}
	return oauth2.OAuth2Token{}, oauth2.ErrOAuth2TokenNotFound
}
