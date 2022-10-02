package event_mysql

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) GetDocument(id string) (e event.EventDocument, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `events`テーブルから`id`が一致する行を取得し、変数`e`に代入する
	err = d.Get(
		&e,
		`SELECT * FROM events WHERE id = $1`,
		id,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return e, nil
}
