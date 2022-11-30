package user

import "prc_hub_back/domain/model/user"

type (
	UpdateUserParam user.UpdateUserParam
)

func Update(id string, p UpdateUserParam, requestUserId string) (_ user.UserWithToken, err error) {
	// リクエスト元のユーザーを取得
	u, err := Get(requestUserId)
	if err != nil {
		return
	}

	return user.Update(
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
