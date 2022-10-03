package user_mysql

import (
	"prc_hub_back/domain/model/user"
)

func (r Repository) Get(id string) (u user.User, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `users`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	err = d.Get(
		&u,
		`SELECT * FROM users WHERE id = $1`,
		id,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return
}

func (r Repository) GetByEmail(email string) (u user.User, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `users`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	err = d.Get(
		&u,
		`SELECT * FROM users WHERE email = $1`,
		email,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return
}
