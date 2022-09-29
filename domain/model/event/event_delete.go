package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotDeleteEvent = errors.New("sorry, you cannot delete this event")
)

func DeleteEvent(repo EventRepository, qs EventQueryService, id string, requestUser user.User) error {
	// Get event
	e, err := GetEvent(
		qs,
		id,
		GetEventQueryParam{},
		requestUser,
	)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage && requestUser.Id != e.UserId {
		// Admin権限なし 且つ `Event.UserId`が自分ではない場合は削除不可
		return ErrCannotDeleteEvent
	}

	return repo.Delete(id)
}
