package user

import (
	"prc_hub_back/domain/model/util"

	"golang.org/x/crypto/bcrypt"
)

func SaveAdmin(email string, password string) error {
	tmpBool := true
	u, err := GetList(GetUserListQueryParam{Admin: &tmpBool})
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
		_, err := Update(u[0].Id, UpdateUserParam{
			Email:    newEmail,
			Password: newPassword,
		}, u[0])
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
		// TODO: UUID -> LastInsertedId()
		u := User{
			Id:                  util.UUID(),
			Name:                "admin",
			Email:               email,
			Password:            string(hashed),
			PostEventAvailabled: true,
			Manage:              true,
			Admin:               true,
			TwitterId:           nil,
			GithubUsername:      nil,
		}

		// リポジトリに追加
		// MySQLサーバーに接続
		d, err := OpenMysql()
		if err != nil {
			return err
		}
		// return時にMySQLサーバーとの接続を閉じる
		defer d.Close()

		// `users`テーブルに追加
		_, err = d.NamedExec(
			`INSERT INTO users
				(id, name, email, password, post_event_availabled, manage, admin, twitter_id, github_username)
			VALUES
				(:id, :name, :email, :password, :post_event_availabled, :manage, :admin, :twitter_id, :github_username)`,
			u,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
