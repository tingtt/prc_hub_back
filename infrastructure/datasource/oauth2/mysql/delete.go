package oauth2_mysql

import (
	"prc_hub_back/domain/model/oauth2"
)

func (r Repository) Delete(userId string, provider string) (err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `id`が一致する行を`oauth2_tokens`テーブルから削除
	r2, err := d.Exec(
		`DELETE FROM oauth2_tokens WHERE user_id = $1 AND provider = $2`,
		userId, provider,
	)
	if err != nil {
		return
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`event`が存在しない
		err = oauth2.ErrOAuth2TokenNotFound
		return
	}

	return
}
