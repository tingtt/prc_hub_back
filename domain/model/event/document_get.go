package event

import "prc_hub_back/domain/model/user"

func GetDocument(id string, requestUser user.User) (ed EventDocument, err error) {
	// Get document

	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	// `documents`テーブルから`id`が一致する行を取得し、変数`tmpEd`に代入する
	var tmpEd EventDocument
	r, err := db.Query("SELECT * FROM documents WHERE id = ?", id)
	if err != nil {
		return
	}
	defer r.Close()
	if !r.Next() {
		// 1行もレコードが無い場合
		// not found
		err = ErrEventDocumentNotFound
		return
	}
	err = r.Scan(&tmpEd.Id, &tmpEd.EventId, &tmpEd.Name, &tmpEd.Url)
	if err != nil {
		return
	}

	// Get event
	e, err := GetEvent(tmpEd.EventId, GetEventQueryParam{}, requestUser)
	if err != nil {
		return
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage &&
		!e.Published && e.UserId != requestUser.Id {
		// `User`が`Admin`・`Manage`のいずれでもなく
		// `Published`でない 且つ 自分のものでない`Event`は取得不可
		err = ErrEventDocumentNotFound
		return
	}

	ed = tmpEd
	return
}
