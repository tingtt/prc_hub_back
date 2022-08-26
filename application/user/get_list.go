package user

import "prc_hub_back/domain/model/user"

func GetList() (_ []user.User, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return user.GetList(repository)
}
