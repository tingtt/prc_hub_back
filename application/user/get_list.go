package user

import "prc_hub_back/domain/model/user"

type GetUserListQuery user.GetUserListQueryParam

func GetList(q GetUserListQuery) (users []user.User, err error) {
	return user.GetList(user.GetUserListQueryParam(q))
}
