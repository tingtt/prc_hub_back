package user

import (
	"errors"
	"prc_hub_back/domain/model/jwt"
	"prc_hub_back/domain/model/util"
	"strings"
)

// Errors
var (
	ErrPostEventAvailabledCannotUpdate = errors.New("sorry, you cannot update `post_event_availabled`")
	ErrManageCannotUpdate              = errors.New("sorry, you cannot update `manage`")
)

type UpdateUserParam struct {
	Name                *string                 `json:"name"`
	Email               *string                 `json:"email"`
	Password            *string                 `json:"password"`
	PostEventAvailabled *bool                   `json:"post_event_availabled"`
	Manage              *bool                   `json:"manage"`
	TwitterId           util.NullableJSONString `json:"twitter_id,omitempty"`
	GithubUsername      util.NullableJSONString `json:"github_username,omitempty"`
}

func (p UpdateUserParam) validate(id string, requestUser User) error {
	// フィールドの検証
	if p.Name != nil {
		err := validateName(*p.Name)
		if err != nil {
			return err
		}
	}
	if p.Email != nil {
		err := validateEmail(*p.Email)
		if err != nil {
			return err
		}
	}
	if p.Password != nil {
		err := validatePassword(*p.Password)
		if err != nil {
			return err
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// `Admin`でも`Manage`でもない場合
		if p.PostEventAvailabled != nil {
			// `User.PostEventAvailabled`は変更不可
			return ErrPostEventAvailabledCannotUpdate
		}
		if p.Manage != nil {
			// `User.Manage`は変更不可
			return ErrManageCannotUpdate
		}
	}
	if !requestUser.Admin && requestUser.Manage {
		// `Admin`ではないが`Manage`の場合
		if p.Manage != nil && !*p.Manage && id != requestUser.Id {
			// 自分以外の`User.Manage`を`false`に変更不可
			return ErrManageCannotUpdate
		}
	}

	return nil
}

func Update(id string, p UpdateUserParam, requestUser User) (_ UserWithToken, err error) {
	// 権限の検証
	if requestUser.Id != id && !requestUser.Admin {
		// Admin権限なし 且つ IDが自分ではない場合は削除不可
		err = ErrUserNotFound
		return
	}

	// リポジトリか更新対象の`User`を取得
	_, err = Get(id)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate(id, requestUser)
	if err != nil {
		return
	}

	// リポジトリ内の`User`を更新
	// MySQLサーバーに接続
	d, err := OpenMysql()
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
		err = ErrNoUpdates
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
		err = ErrUserNotFound
		return
	}

	// 更新後のデータを取得
	u, err := Get(id)
	if err != nil {
		return
	}

	// jwtを生成
	uwt := UserWithToken{User: u}
	uwt.Token, err = jwt.GenerateToken(jwt.GenerateTokenParam{Id: u.Id, Email: u.Email, Admin: u.Admin})
	if err != nil {
		return
	}

	return uwt, nil
}
