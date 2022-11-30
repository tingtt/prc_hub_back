package user

import "prc_hub_back/domain/model/user"

func Delete(id string, requestUserId string) (err error) {
	// リクエスト元のユーザーを取得
	u, err := Get(id)
	if err != nil {
		return
	}

	return user.DeleteUesr(id, u)
}
