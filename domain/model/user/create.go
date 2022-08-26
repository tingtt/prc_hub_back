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

func (p CreateUserParam) validate(repo UserRepository) error {
	err := validateName(p.Name)
	if err != nil {
		return err
	}
	err = validateEmail(repo, p.Email)
	if err != nil {
		return err
	}
	err = validatePassword(p.Password)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(repo UserRepository, p CreateUserParam) (_ UserWithToken, err error) {
	// バリデーション
	err = p.validate(repo)
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

	// リポジトリに追加
	u, err := repo.Add(User{
		Id:                  util.UUID(),
		Name:                p.Name,
		Email:               p.Email,
		Password:            string(hashed),
		PostEventAvailabled: false,
		Manage:              false,
		Admin:               false,
		TwitterId:           p.TwitterId,
		GithubUsername:      p.GithubUsername,
	})
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
