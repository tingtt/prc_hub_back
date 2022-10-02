package event_mysql

import (
	"prc_hub_back/domain/model/event"
)

func (r RepositoryEvent) AddDocument(e event.EventDocument) (_ event.EventDocument, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `documents`テーブルに追加
	_, err = d.NamedExec(
		`INSERT INTO documents
			(id, event_id, name, url)
		VALUES
			(:id, :event_id, :name, :url)`,
		e,
	)
	if err != nil {
		return
	}

	return e, nil
}
