package oauth2

type OAuth2Token struct {
	UserId   string `db:"user_id"`
	Provider string `db:"provider"`
	Token    string `db:"token"`
	OwnerId  string `db:"owner_id"`
}
