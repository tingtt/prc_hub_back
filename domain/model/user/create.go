package user

import (
	"prc_hub_back/domain/model/jwt"
	"prc_hub_back/domain/model/util"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserParam struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	TwitterId      *string `json:"twitter_id,omitempty"`
	GithubUsername *string `json:"github_username,omitempty"`
}

func (p CreateUserParam) validate() error {
	err := validateName(p.Name)
	if err != nil {
		return err
	}
	err = validateEmail(p.Email)
	if err != nil {
		return err
	}
	err = validatePassword(p.Password)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(p CreateUserParam) (_ UserWithToken, err error) {
	// バリデーション
	err = p.validate()
	if err != nil {
		return
	}

	// パスワードをハッシュ化
	hashed, err := bcrypt.GenerateFromPassword([]byte(p.Password), 10)
	if err != nil {
		return
	}

	// ""(空文字)を`null`に置き換え
	if p.TwitterId != nil && *p.TwitterId == "" {
		p.TwitterId = nil
	}
	if p.GithubUsername != nil && *p.GithubUsername == "" {
		p.GithubUsername = nil
	}

	// TODO: UUID -> LastInsertedId()
	u := User{
		Id:                  util.UUID(),
		Name:                p.Name,
		Email:               p.Email,
		Password:            string(hashed),
		PostEventAvailabled: false,
		Manage:              false,
		Admin:               false,
		TwitterId:           p.TwitterId,
		GithubUsername:      p.GithubUsername,
	}

	// リポジトリに追加
	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// `users`テーブルに追加
	_, err = d.NamedExec(
		`INSERT INTO users
			(id, name, email, password, post_event_availabled, manage, admin, twitter_id, github_username)
		VALUES
			(:id, :name, :email, :password, :post_event_availabled, :manage, :admin, :twitter_id, :github_username)`,
		u,
	)
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
