package event_mysql

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) Delete(id string) (err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `id`が一致する行を`events`テーブルから削除
	r2, err := d.Exec(
		`DELETE FROM events WHERE id = $1`,
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
		// `id`に一致する`event`が存在しない
		err = event.ErrEventNotFound
		return
	}

	return
}
