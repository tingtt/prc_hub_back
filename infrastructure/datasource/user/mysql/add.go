package user_mysql

import (
	"prc_hub_back/domain/model/user"
)

func (r Repository) Add(u user.User) (_ user.User, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
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
		return
	}

	return u, nil
}
