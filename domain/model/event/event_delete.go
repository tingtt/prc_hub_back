package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotDeleteEvent = errors.New("sorry, you cannot delete this event")
)

func DeleteEvent(id string, requestUser user.User) error {
	// Get event
	e, err := GetEvent(
		id,
		GetEventQueryParam{},
		requestUser,
	)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage && requestUser.Id != e.UserId {
		// Admin権限なし 且つ `Event.UserId`が自分ではない場合は削除不可
		return ErrCannotDeleteEvent
	}

	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	// `id`が一致する行を`events`テーブルから削除
	r2, err := db.Exec(
		`DELETE FROM events WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return err
		}
		// `id`に一致する`event`が存在しない
		err = ErrEventNotFound
		return err
	}

	return nil
}
