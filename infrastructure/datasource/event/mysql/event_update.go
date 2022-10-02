package event_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"prc_hub_back/domain/model/event"
	"strings"
)

func (r RepositoryEvent) Update(id string, p event.UpdateEventParam) (_ event.Event, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// トランザクション開始
	tx, err := d.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return
	}
	defer func() {
		// return時にトランザクションの後処理
		//* 17行目の`defer`より先に実行される
		if err != nil {
			// 失敗時はロールバック
			tx.Rollback()
		} else {
			// 成功時はコミット
			tx.Commit()
		}
	}()

	// `events`テーブル用のクエリを作成
	query1 := "UPDATE events SET"
	queryParams1 := []interface{}{}
	if p.Name != nil {
		// `name`を変更
		query1 += " name = ?,"
		queryParams1 = append(queryParams1, *p.Name)
	}
	if p.Description.KeyExists() {
		// `description`を変更
		query1 += " description = ?,"
		if p.Description.IsNull() {
			queryParams1 = append(queryParams1, nil)
		} else {
			queryParams1 = append(queryParams1, *p.Description.Value)
		}
	}
	if p.Location.KeyExists() {
		// `location`を変更
		query1 += " location = ?,"
		if p.Location.IsNull() {
			queryParams1 = append(queryParams1, nil)
		} else {
			queryParams1 = append(queryParams1, *p.Location.Value)
		}
	}
	if p.Published != nil {
		// `published`を変更
		query1 += " published = ?,"
		queryParams1 = append(queryParams1, *p.Published)
	}
	if p.Completed != nil {
		// `completed`を変更
		query1 += " completed = ?"
		queryParams1 = append(queryParams1, *p.Completed)
	}
	// 更新するフィールドがあるか確認
	if strings.HasSuffix(query1, "SET") {
		// 更新するフィールドが無いため中断
		err = event.ErrNoUpdates
		return
	}
	// 不要な末尾の句を切り取り
	query1 = strings.TrimSuffix(query1, ",")

	// `events`テーブルの`id`が一致する行を更新
	r2, err := tx.Exec(query1+" WHERE id = ?", append(queryParams1, id))
	if err != nil {
		return
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`event`が存在しない
		err = event.ErrEventDocumentNotFound
		return
	}

	if p.Datetimes != nil {
		// `event_datetimes`テーブルの更新

		// 既存のデータを削除
		_, err = tx.Exec(
			"DELETE FROM event_datetimes WHERE event_id = ?",
			id,
		)
		if err != nil {
			return
		}

		// 新規データの追加
		_, err = d.NamedExec(
			fmt.Sprintf(
				`INSERT INTO event_datetimes
					(event_id, start, end)
				VALUES
					(%s, :start, :end)`,
				id,
			),
			p.Datetimes,
		)
		if err != nil {
			return
		}
	}

	// TODO: 更新後のデータを取得
	return
}
