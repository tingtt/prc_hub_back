package oauth2_inmemory

import "prc_hub_back/domain/model/oauth2"

func (r Repository) Save(t oauth2.OAuth2Token) (oauth2.OAuth2Token, error) {
	data = append(data, t)
	return t, nil
}
