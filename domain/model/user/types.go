package user

type User struct {
	Id                  int64   `json:"id" db:"id"`
	Name                string  `json:"name" db:"name"`
	Email               string  `json:"-" db:"email"`
	Password            string  `json:"-" db:"password"`
	PostEventAvailabled bool    `json:"post_event_availabled" db:"post_event_availabled"`
	Manage              bool    `json:"manage" db:"manage"`
	Admin               bool    `json:"admin" db:"admin"`
	TwitterId           *string `json:"twitter_id,omitempty" db:"twitter_id"`
	GithubUsername      *string `json:"github_username,omitempty" db:"github_username"`
}

type UserWithToken struct {
	User
	Token string `json:"token"`
}
