package oauth2

func Delete(repo OAuth2TokenRepository, userId string, provider string) error {
	return repo.Delete(userId, provider)
}
