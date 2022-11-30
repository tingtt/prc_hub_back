package user

func Get(id string) (u User, err error) {
	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `users`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	// TODO: Scan
	err = d.Get(
		&u,
		`SELECT * FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return
	}
	return
}

func GetByEmail(email string) (u User, err error) {
	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `users`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	// TODO: Scan
	err = d.Get(
		&u,
		`SELECT * FROM users WHERE email = $1`,
		email,
	)
	if err != nil {
		return
	}
	return
}
