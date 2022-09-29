package event

import "prc_hub_back/domain/model/user"

func GetDocument(repo EventRepository, qs EventQueryService, id string, requestUser user.User) (ed EventDocument, err error) {
	// Get document
	tmpEd, err := repo.GetDocument(id)
	if err != nil {
		return
	}

	// Get event
	e, err := GetEvent(qs, tmpEd.EventId, GetEventQueryParam{}, requestUser)
	if err != nil {
		return
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage &&
		!e.Published && e.UserId != requestUser.Id {
		// `User`が`Admin`・`Manage`のいずれでもなく
		// `Published`でない 且つ 自分のものでない`Event`は取得不可
		err = ErrEventDocumentNotFound
		return
	}

	ed = tmpEd
	return
}
