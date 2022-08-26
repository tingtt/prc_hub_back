package oauth2

func GetLoginUrl(provider string) (string, error) {
	p, err := searchProvider(provider)
	if err != nil {
		return "", err
	}
	switch p {
	case ProviderGithub:
		return githubClient.GetLoginUrl(), nil
	default:
		return "", ErrProviderNotFound
	}
}
