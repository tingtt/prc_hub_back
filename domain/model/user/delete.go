package user

import "errors"

// Errors
var (
	ErrAdminUserCannnotDelete = errors.New("cannot delete admin user")
)

func DeleteUesr(repo UserRepository, id string, requestUser User) error {
	// リポジトリから削除対象の`User`を取得
	u, err := Get(repo, id)
	if err != nil {
		return err
	}

	if !requestUser.Admin && requestUser.Id != id {
		// Admin権限なし 且つ IDが自分ではない場合は削除不可
		return ErrUserNotFound
	}

	if u.Admin {
		// Adminユーザーは削除不可
		return ErrAdminUserCannnotDelete
	}

	// リポジトリから`User`を削除
	return repo.Delete(id)
}
