package user

import "prc_hub_back/domain/model/user"

type GetUserListQuery user.GetUserListQueryParam

func GetList(q GetUserListQuery) (_ []user.User, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return user.GetList(repository, user.GetUserListQueryParam(q))
}
