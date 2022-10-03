package oauth2_mysql

import "prc_hub_back/domain/model/oauth2"

func (r Repository) Save(t oauth2.OAuth2Token) (_ oauth2.OAuth2Token, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `oauth2_tokens`テーブルに追加
	_, err = d.NamedExec(
		`INSERT INTO oauth2_tokens
			(user_id, provider, token, owner_id)
		VALUES
			(:user_id, :provider, :token, :owner_id)`,
		t,
	)
	if err != nil {
		return
	}

	return t, nil
}
