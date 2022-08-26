package oauth2

import (
	"prc_hub_back/domain/model/flag_with_env"
	"prc_hub_back/domain/model/oauth2"
	"prc_hub_back/domain/model/oauth2_github"
)

// コマンドライン引数 / 環境変数
var (
	githubClientId     = flag_with_env.String("github-clint-id", "OAUTH2_CLIENT_ID_GITHUB", "", "OAuth2 client-id for github,com")
	githubClientSecret = flag_with_env.String("github-clint-secret", "OAUTH2_CLIENT_SECRET_GITHUB", "", "OAuth2 client-secret for github,com")
)

var githubClient *oauth2_github.Client

func Init() {
	// コマンドライン引数 / 環境変数 の取得
	flag_with_env.Parse()

	githubClient = oauth2_github.NewClient(*githubClientId, *githubClientSecret)
}

func getClient(provider string) (oauth2.Client, error) {
	p, err := searchProvider(provider)
	if err != nil {
		return nil, err
	}

	switch p {
	case ProviderGithub:
		return githubClient, nil
	default:
		return nil, ErrProviderNotFound
	}
}
