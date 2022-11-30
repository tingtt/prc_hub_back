package user

import "errors"

// Errors
var (
	ErrAdminUserCannnotDelete = errors.New("cannot delete admin user")
)

func DeleteUesr(id string, requestUser User) error {
	// リポジトリから削除対象の`User`を取得
	u, err := Get(id)
	if err != nil {
		return err
	}

	if !requestUser.Admin && requestUser.Id != id {
		// Admin権限なし 且つ IDが自分ではない場合は削除不可
		return ErrUserNotFound
	}

	if u.Admin {
		// Adminユーザーは削除不可
		return ErrAdminUserCannnotDelete
	}

	// リポジトリから`User`を削除
	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `id`が一致する行を`users`テーブルから削除
	r2, err := d.Exec(
		`DELETE FROM users WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return err
		}
		// `id`に一致する`usersが存在しない
		return ErrUserNotFound
	}
	return nil
}
