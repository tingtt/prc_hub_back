package user

import "strings"

type GetUserListQueryParam struct {
	Name                *string `query:"name"`
	NameContain         *string `query:"name_contain"`
	PostEventAvailabled *bool   `json:"post_event_availabled"`
	Manage              *bool   `json:"manage"`
	Admin               *bool   `json:"admin"`
}

func GetList(q GetUserListQueryParam) ([]User, error) {
	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return nil, err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// クエリを作成
	query := "SELECT * FROM users WHERE"
	queryParams := []interface{}{}
	if q.PostEventAvailabled != nil {
		// 権限で絞り込み
		query += " post_event_availabled = ? AND"
		queryParams = append(queryParams, *q.PostEventAvailabled)
	}
	if q.Manage != nil {
		// 権限で絞り込み
		query += " manage = ? AND"
		queryParams = append(queryParams, *q.Manage)
	}
	if q.Admin != nil {
		// 権限で絞り込み
		query += " admin = ? AND"
		queryParams = append(queryParams, *q.Admin)
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

	// `users`テーブルからを取得し、変数`users`に代入する
	// TODO: Scan
	var users []User
	err = d.Get(
		&users,
		query,
		queryParams...,
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}
