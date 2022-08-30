package user

import (
	"prc_hub_back/domain/model/util"
)

func SaveAdmin(repo UserRepository, email string, password string) error {
	tmpBool := true
	u, err := GetList(repo, GetUserListQueryParam{Admin: &tmpBool})
	if err != nil {
		return err
	}

	if len(u) == 1 {
		// `Admin`の`User`が登録済
		var newEmail *string = nil
		var newPassword *string = nil
		if u[0].Email != email {
			// `Email`が不一致
			newEmail = &email
		}
		if verify, err := u[0].Verify(password); err != nil {
			return err
		} else if !verify {
			// `Password`が不一致
			newPassword = &password
		}

		// `User`更新
		_, err := repo.Update(u[0].Id, UpdateUserParam{
			Email:    newEmail,
			Password: newPassword,
		})
		if err != nil {
			return err
		}
	} else {
		// `Admin`の`User`が未登録
		// `User`追加
		_, err := repo.Add(User{
			Id:                  util.UUID(),
			Name:                "admin",
			Email:               email,
			Password:            password,
			PostEventAvailabled: true,
			Manage:              true,
			Admin:               true,
			TwitterId:           nil,
			GithubUsername:      nil,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
