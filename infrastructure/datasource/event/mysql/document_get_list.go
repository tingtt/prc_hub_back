package event_mysql

import (
	"prc_hub_back/domain/model/event"
	"strings"
)

func (r RepositoryEvent) GetDocumentList(q event.GetDocumentQueryParam) (documents []event.EventDocument, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// クエリを作成
	query := "SELECT * FROM documents WHERE"
	queryParams := []interface{}{}
	if q.EventId != nil {
		// イベントIDで絞り込み
		query += " event_id = ? AND"
		queryParams = append(queryParams, *q.EventId)
	}
	if q.Name != nil {
		// ドキュメント名の一致で絞り込み
		query += " name = ? AND"
		queryParams = append(queryParams, *q.Name)
	}
	if q.NameContain != nil {
		// ドキュメント名に文字列が含まれるかで絞り込み
		query += " name LIKE ?"
		queryParams = append(queryParams, "%"+*q.NameContain+"%")
	}
	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, " WHERE")
	query = strings.TrimSuffix(query, " AND")

	// `documents`テーブルからを取得し、変数`documents`に代入する
	err = d.Get(
		&documents,
		query,
		queryParams...,
	)
	// TODO: Return error if result set is empty
	if err != nil {
		return
	}

	return documents, nil
}
