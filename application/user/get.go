package user

import "prc_hub_back/domain/model/user"

func Get(id string) (_ user.User, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return user.Get(repository, id)
}

func GetByEmail(email string) (_ user.User, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return user.GetByEmail(repository, email)
}
