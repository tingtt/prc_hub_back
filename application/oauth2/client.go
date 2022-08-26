package oauth2

import (
	"prc_hub_back/domain/model/oauth2"
)

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
