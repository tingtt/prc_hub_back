package user_inmemory

import (
	"prc_hub_back/domain/model/user"
)

func (r Repository) Get(id string) (_ user.User, err error) {
	for _, u := range data {
		if u.Id == id {
			return u, nil
		}
	}
	err = user.ErrUserNotFound
	return
}

func (r Repository) GetByEmail(email string) (_ user.User, err error) {
	for _, u := range data {
		if u.Email == email {
			return u, nil
		}
	}
	err = user.ErrUserNotFound
	return
}
