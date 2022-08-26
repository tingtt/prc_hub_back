package user_inmemory

import (
	"prc_hub_back/domain/model/user"
)

func (r Repository) Add(e user.User) (user.User, error) {
	data = append(data, e)
	return e, nil
}
