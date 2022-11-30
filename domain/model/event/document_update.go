package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"strings"
)

// Errors
var (
	ErrCannotUpdateEventDocument = errors.New("sorry, you cannot update this document")
)

type UpdateEventDocumentParam struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

func (p UpdateEventDocumentParam) validate(id string, requestUser user.User) error {
	/**
	 * フィールドの検証
	**/
	if p.Name == nil && p.Url == nil {
		return ErrNoUpdates
	}
	// `Name`
	if p.Name != nil {
		err := validateDocumentName(*p.Name)
		if err != nil {
			return err
		}
	}
	// `Url`
	if p.Url != nil {
		err := validateUrl(*p.Url)
		if err != nil {
			return err
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		ed, err := GetDocument(id, requestUser)
		if err != nil {
			return err
		}

		// Eventを取得
		e, err := GetEvent(ed.EventId, GetEventQueryParam{}, requestUser)
		if err != nil {
			return err
		}

		if e.UserId != requestUser.Id {
			// `User`が`Admin`・`Manage`のいずれでもなく
			// `Published`でない 且つ 自分のものでない`Event`は変更不可
			return ErrCannotUpdateEventDocument
		}
	}

	return nil
}

func UpdateEventDocument(id string, p UpdateEventDocumentParam, requestUser user.User) (_ EventDocument, err error) {
	// 確認

	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `documents`テーブルから`id`が一致する行を確認
	_, err = GetDocument(id, requestUser)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate(id, requestUser)
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
		err = ErrNoUpdates
		return
	}
	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, ",")

	// `documents`テーブルの`id`が一致する行を更新
	r2, err := db.Exec(query+" WHERE id = ?", append(queryParams, id))
	if err != nil {
		return
	}
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`document`が存在しない
		err = ErrEventDocumentNotFound
		return
	}

	// 更新後のデータを取得
	return GetDocument(id, requestUser)
}
