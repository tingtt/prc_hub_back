package oauth2_mysql

import "prc_hub_back/domain/model/oauth2"

func (r Repository) Get(userId string, provider string) (o oauth2.OAuth2Token, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `oauth2_tokens`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	err = d.Get(
		&o,
		`SELECT * FROM oauth2_tokens WHERE user_id = $1 AND provider = $2`,
		userId, provider,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return
}

func (r Repository) GetByOwnerId(provider string, ownerId string) (o oauth2.OAuth2Token, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `oauth2_tokens`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	err = d.Get(
		&o,
		`SELECT * FROM oauth2_tokens WHERE provider = $1 AND owner_id = $2`,
		provider, ownerId,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return
}
