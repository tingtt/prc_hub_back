package oauth2

type Client interface {
	GetToken(code string) (token string, err error)
	GetOwnerId(token string) (string, error)
}
