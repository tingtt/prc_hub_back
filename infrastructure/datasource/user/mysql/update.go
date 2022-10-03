package user_mysql

import (
	"prc_hub_back/domain/model/user"
	"strings"
)

func (r Repository) Update(id string, p user.UpdateUserParam) (_ user.User, err error) {
	// MySQLサーバーに接続
	d, err := Open()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// クエリを作成
	query := "UPDATE users SET"
	queryParams := []interface{}{}

	if p.Name != nil {
		// `Name`を変更
		query += " name = ?,"
		queryParams = append(queryParams, *p.Name)
	}
	if p.Email != nil {
		// `Email`を変更
		query += " email = ?,"
		queryParams = append(queryParams, *p.Email)
	}
	if p.Password != nil {
		// `Password`を変更
		query += " password = ?,"
		queryParams = append(queryParams, *p.Password)
	}
	if p.PostEventAvailabled != nil {
		// `PostEventAvailabled`を変更
		query += " post_event_availabled = ?,"
		queryParams = append(queryParams, *p.PostEventAvailabled)
	}
	if p.Manage != nil {
		// `Manage`を変更
		query += " manage = ?,"
		queryParams = append(queryParams, *p.Manage)
	}
	if p.TwitterId.KeyExists() {
		// `TwitterId`を変更
		query += " twitter_id = ?,"
		if p.TwitterId.IsNull() {
			queryParams = append(queryParams, nil)
		} else {
			queryParams = append(queryParams, *p.TwitterId.Value)
		}
	}
	if p.GithubUsername.KeyExists() {
		// `GithubUsername`を変更
		query += " github_username = ?"
		if p.GithubUsername.IsNull() {
			queryParams = append(queryParams, nil)
		} else {
			queryParams = append(queryParams, *p.GithubUsername.Value)
		}
	}
	// 更新するフィールドがあるか確認
	if strings.HasSuffix(query, "SET") {
		// 更新するフィールドが無いため中断
		err = user.ErrNoUpdates
		return
	}
	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, ",")

	// `users`テーブルの`id`が一致する行を更新
	r2, err := d.Exec(query+" WHERE id = ?", append(queryParams, id))
	var a int64
	if a, err = r2.RowsAffected(); err != nil || a != 1 {
		if err != nil {
			return
		}
		// `id`に一致する`uesr`が存在しない
		err = user.ErrUserNotFound
		return
	}

	// TODO: 更新後のデータを取得
	return
}
