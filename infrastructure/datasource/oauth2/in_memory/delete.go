package oauth2_inmemory

import "prc_hub_back/domain/model/oauth2"

func (r Repository) Delete(userId string, provider string) (err error) {
	old := data
	data = nil
	err = oauth2.ErrOAuth2TokenNotFound
	for _, ot := range old {
		if ot.UserId == userId && ot.Provider == provider {
			err = nil
		} else {
			data = append(data, ot)
		}
	}
	return
}
