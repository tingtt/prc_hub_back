package user

import "prc_hub_back/domain/model/user"

func SaveAdmin(email string, password string) error {
	return user.SaveAdmin(repository, email, password)
}
