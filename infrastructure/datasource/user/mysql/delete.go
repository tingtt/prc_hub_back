package user_mysql

import (
	"prc_hub_back/application/user"
)

func (r Repository) Delete(id string) (err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `id`が一致する行を`users`テーブルから削除
	r2, err := d.Exec(
		`DELETE FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`usersが存在しない
		err = user.ErrUserNotFound
		return
	}

	return
}
