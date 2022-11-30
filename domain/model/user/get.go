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
	r, err := d.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return
	}
	if !r.Next() {
		// 1行もレコードが無い場合
		// not found
		err = ErrUserNotFound
		return
	}

	// 一時変数に割り当て
	var (
		id2                 string
		name                string
		email               string
		password            string
		postEventAvailabled bool
		manage              bool
		admin               bool
		twitterId           *string
		githubUsername      *string
	)
	err = r.Scan(
		&id2, &name, &email, &password, &postEventAvailabled,
		&manage, &admin, &twitterId, &githubUsername,
	)
	if err != nil {
		return
	}

	return User{
			id,
			name,
			email,
			password,
			postEventAvailabled,
			manage,
			admin,
			twitterId,
			githubUsername,
		},
		nil
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
	r, err := d.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return
	}
	if !r.Next() {
		// 1行もレコードが無い場合
		// not found
		err = ErrUserNotFound
		return
	}

	// 一時変数に割り当て
	var (
		id                  string
		name                string
		email2              string
		password            string
		postEventAvailabled bool
		manage              bool
		admin               bool
		twitterId           *string
		githubUsername      *string
	)
	err = r.Scan(
		&id, &name, &email2, &password, &postEventAvailabled,
		&manage, &admin, &twitterId, &githubUsername,
	)
	if err != nil {
		return
	}

	return User{
			id,
			name,
			email,
			password,
			postEventAvailabled,
			manage,
			admin,
			twitterId,
			githubUsername,
		},
		nil
}
