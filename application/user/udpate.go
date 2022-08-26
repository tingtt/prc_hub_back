package user

import "prc_hub_back/domain/model/user"

type (
	UpdateUserParam user.UpdateUserParam
)

func Update(id string, p UpdateUserParam, requestUserId string) (_ user.UserWithToken, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := Get(id)
	if err != nil {
		return
	}

	return user.UpdateUser(
		repository,
		id,
		user.UpdateUserParam{
			Name:                p.Name,
			Email:               p.Email,
			Password:            p.Password,
			PostEventAvailabled: p.PostEventAvailabled,
			Manage:              p.Manage,
			TwitterId:           p.TwitterId,
			GithubUsername:      p.GithubUsername,
		},
		u,
	)
}
