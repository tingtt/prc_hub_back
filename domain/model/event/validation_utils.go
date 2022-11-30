package event

import (
	"errors"
	"time"
)

// Errors
var (
	ErrValidateEventTitleCannotBeEmpty         = errors.New("Event title cannot be empty")
	ErrValidateDocumentNameCannotBeEmpty       = errors.New("Event document name cannot be empty")
	ErrValidateUrlCannotBeEmpty                = errors.New("Event document url cannot be empty")
	ErrValidateEventDatetimesCannotBeEmpty     = errors.New("Event datetime cannot be empty")
	ErrValidateEventDatetimeStartMustBeforeEnd = errors.New("Event start datetime must be before end datetime")
)

func validateTitle(title string) error {
	// 空文字チェック
	if title == "" {
		return ErrValidateEventTitleCannotBeEmpty
	}
	return nil
}

func validateDocumentName(name string) error {
	// 空文字チェック
	if name == "" {
		return ErrValidateDocumentNameCannotBeEmpty
	}
	return nil
}

func validateUrl(url string) error {
	// 空文字チェック
	if url == "" {
		return ErrValidateUrlCannotBeEmpty
	}
	return nil
}

func validateEventId(id string) error {
	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	// `documents`テーブルから`id`が一致する行を取得し、変数`tmpEd`に代入する
	var tmpEd EventDocument
	// TODO: 変数へのアサインをスキャンにする
	err = db.Get(
		&tmpEd,
		`SELECT * FROM events WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func validateEventDatetime(start time.Time, end time.Time) error {
	if !start.Before(end) {
		return ErrValidateEventDatetimeStartMustBeforeEnd
	}
	return nil
}
