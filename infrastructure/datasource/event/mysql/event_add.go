package event_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"prc_hub_back/domain/model/event"
)

func (r RepositoryEvent) Add(e event.Event) (_ event.Event, err error) {
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

	// `events`テーブルに追加
	_, err = tx.NamedExec(
		`INSERT INTO events
			(id, name, description, location, published, completed, user_id)
		VALUES
			(:id, :name, :description, :location, :published, :completed, :user_id)`,
		e,
	)
	if err != nil {
		return
	}

	// `event_datetimes`テーブルに追加
	_, err = tx.NamedExec(
		fmt.Sprintf(
			`INSERT INTO event_datetimes
				(event_id, start, end)
			VALUES
				("%s", :start, :end)`,
			e.Id,
		),
		e.Datetimes,
	)
	if err != nil {
		return
	}

	return e, nil
}
