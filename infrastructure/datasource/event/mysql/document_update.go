package event_mysql

import (
	"prc_hub_back/domain/model/event"
	"strings"
)

func (r RepositoryEvent) UpdateDocument(id string, p event.UpdateEventDocumentParam) (_ event.EventDocument, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// クエリを作成
	query := "UPDATE documents SET"
	queryParams := []interface{}{}
	if p.Name != nil {
		// `name`を変更
		query += " name = ?,"
		queryParams = append(queryParams, *p.Name)
	}
	if p.Url != nil {
		// `url`を変更
		query += " url = ?"
		queryParams = append(queryParams, *p.Url)
	}
	// 更新するフィールドがあるか確認
	if strings.HasSuffix(query, "SET") {
		// 更新するフィールドが無いため中断
		err = event.ErrNoUpdates
		return
	}
	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, ",")

	// `documents`テーブルの`id`が一致する行を更新
	r2, err := d.Exec(query+" WHERE id = ?", append(queryParams, id))
	if err != nil {
		return
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`document`が存在しない
		err = event.ErrEventDocumentNotFound
		return
	}

	// TODO: 更新後のデータを取得
	return
}
