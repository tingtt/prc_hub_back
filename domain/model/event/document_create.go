package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
)

// Errors
var (
	ErrCannotCreateEventDocument = errors.New("sorry, you cannot create document to this event")
)

type CreateEventDocumentParam struct {
	EventId string `json:"event_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

func (p CreateEventDocumentParam) validate(requestUser user.User) error {
	// フィールドの検証
	err := validateDocumentName(p.Name)
	if err != nil {
		return err
	}
	err = validateUrl(p.Url)
	if err != nil {
		return err
	}
	err = validateEventId(p.EventId)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// Eventを取得
		e, err := GetEvent(p.EventId, GetEventQueryParam{}, requestUser)
		if err != nil {
			return err
		}

		if requestUser.Id != e.UserId {
			// `Admin`・`Manage`のいずれでもなく`Event.UserId`が自分ではない場合は追加不可
			return ErrCannotCreateEventDocument
		}
	}

	return nil
}

func CreateEventDocument(p CreateEventDocumentParam, requestUser user.User) (_ EventDocument, err error) {
	err = p.validate(requestUser)
	if err != nil {
		return
	}

	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	// TODO: UUID -> LastInsertedId()
	e := EventDocument{
		Id:      util.UUID(),
		EventId: p.EventId,
		Name:    p.Name,
		Url:     p.Url,
	}

	// `documents`テーブルに追加
	_, err = db.NamedExec(
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
