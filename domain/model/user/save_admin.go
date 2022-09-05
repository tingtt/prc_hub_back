package user

import (
	"prc_hub_back/domain/model/util"

	"golang.org/x/crypto/bcrypt"
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
			// パスワードをハッシュ化
			hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
			if err != nil {
				return err
			}
			tmpPasswd := string(hashed)
			// 新規パスワード
			newPassword = &tmpPasswd
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

		// パスワードをハッシュ化
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			return err
		}

		// `User`追加
		_, err = repo.Add(User{
			Id:                  util.UUID(),
			Name:                "admin",
			Email:               email,
			Password:            string(hashed),
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
