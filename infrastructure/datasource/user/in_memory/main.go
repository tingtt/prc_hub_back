package user_inmemory

import (
	"prc_hub_back/domain/model/user"
)

type (
	Repository struct{}
)

var (
	data []user.User
)
