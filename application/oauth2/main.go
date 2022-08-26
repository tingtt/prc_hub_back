package oauth2

import (
	"errors"
	"prc_hub_back/domain/model/flag_with_env"
	"prc_hub_back/domain/model/oauth2"
	"prc_hub_back/domain/model/oauth2_github"
)

// コマンドライン引数 / 環境変数
var (
	githubClientId     = flag_with_env.String("github-clint-id", "OAUTH2_CLIENT_ID_GITHUB", "", "OAuth2 client-id for github,com")
	githubClientSecret = flag_with_env.String("github-clint-secret", "OAUTH2_CLIENT_SECRET_GITHUB", "", "OAuth2 client-secret for github,com")
)

// Singleton fields
var (
	initialized  = false
	repository   oauth2.OAuth2TokenRepository
	githubClient *oauth2_github.Client
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(repo oauth2.OAuth2TokenRepository) {
	// コマンドライン引数 / 環境変数 の取得
	flag_with_env.Parse()

	githubClient = oauth2_github.NewClient(*githubClientId, *githubClientSecret)

	initialized = true
	repository = repo
}
