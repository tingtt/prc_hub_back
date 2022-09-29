package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotDeleteEventDocument = errors.New("sorry, you cannot delete this document")
)

func DeleteEventDocument(repo EventRepository, qs EventQueryService, id string, requestUser user.User) error {
	// Get document
	ed, err := GetDocument(repo, qs, id, requestUser)
	if err != nil {
		return err
	}

	// Get event
	e, err := GetEvent(qs, ed.EventId, GetEventQueryParam{}, requestUser)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage &&
		requestUser.Id != e.UserId {
		// Admin権限なし 且つ `Event.UserId`が自分ではない場合は削除不可
		return ErrCannotDeleteEventDocument
	}

	return repo.DeleteDocument(id)
}
