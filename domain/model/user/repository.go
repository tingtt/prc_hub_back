package user

import "errors"

// Errors
var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Add(u User) (User, error)
	Get(id string) (User, error)
	GetByEmail(email string) (User, error)
	GetList(q GetUserListQueryParam) ([]User, error)
	Update(id string, p UpdateUserParam) (User, error)
	Delete(userId string) error
}
