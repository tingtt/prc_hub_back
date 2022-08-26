package user

import (
	"errors"
	"prc_hub_back/domain/model/jwt"
	"prc_hub_back/domain/model/util"
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

func (p UpdateUserParam) validate(repo UserRepository, id string, requestUser User) error {
	// フィールドの検証
	if p.Name != nil {
		err := validateName(*p.Name)
		if err != nil {
			return err
		}
	}
	if p.Email != nil {
		err := validateEmail(repo, *p.Email)
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

func UpdateUser(repo UserRepository, id string, p UpdateUserParam, requestUser User) (_ UserWithToken, err error) {
	// 権限の検証
	if requestUser.Id != id && !requestUser.Admin {
		// Admin権限なし 且つ IDが自分ではない場合は削除不可
		err = ErrUserNotFound
		return
	}

	// リポジトリか更新対象の`User`を取得
	_, err = repo.Get(id)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate(repo, id, requestUser)
	if err != nil {
		return
	}

	// リポジトリ内の`User`を更新
	u, err := repo.Update(id, p)
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
